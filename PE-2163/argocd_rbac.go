package app

import (
	"bytes"
	"cmp"
	"compress/gzip"
	"context"
	_ "embed"
	"encoding/base64"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"text/template"

	argocdtypes "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	argorbacutil "github.com/argoproj/argo-cd/v2/util/rbac"
	set "github.com/deckarep/golang-set/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyconfigurationsv1 "k8s.io/client-go/applyconfigurations/core/v1"

	c "gitlab.localhost.com/platform/kubernetes/controllers/argocd-coordination-controller/pkg/constants"
)

//go:embed templates/argocd_rbac.gotpl
var argocdRBACTemplateUnparsed string

// this is separated into two methods because the fake clientset does not
// support server-side apply at present, so there is a separate method that
// generates the apply configuration object which we can inspect as part of
// the unit tests
func (arm *ArgoCDResourceManager) generateArgoCDRBAC(ctx context.Context, gitlabClient *GitlabClient, appProjects *[]*argocdtypes.AppProject, additionalRBACRules string, privilegedGroupConfiguration PrivilegedGroupConfiguration) (warnings []error, err error) {
	var rbacCM *applyconfigurationsv1.ConfigMapApplyConfiguration

	// this will be used to build a full tree that consists of parent (parent group) and child nodes (subgroups and projects)
	// to represent their relationship and retrieve the correct LDAP group links
	GroupMap := make(map[string]*PostTraverse)
	group, err := gitlabClient.GetGitlabGroupBaseID(arm.GitlabGroupBase)
	if err != nil {
		arm.Logger.Error(err, "failed to get gitlab group base id")
		return nil, err
	}
	tree, err := BuildTreeBase(gitlabClient, int64(group.ID))
	if err != nil {
		arm.Logger.Error(err, "failed to build tree base for gitlab projects")
		return nil, err
	}

	BuildFullTree(gitlabClient, tree)
	emptyLdapGroupLinks := []GitlabLDAPGroupLink{}
	TraverseTree(tree, emptyLdapGroupLinks, GroupMap)

	ldapToProjectMap := CreateLookUpTable(GroupMap)
	rbacCM, warnings, err = arm.getArgoCDRBACApplyConfiguration(appProjects, additionalRBACRules, privilegedGroupConfiguration, ldapToProjectMap)
	if err != nil {
		return warnings, err
	}

	_, err = arm.KubernetesClientset.CoreV1().ConfigMaps(arm.ArgoCDControllerNamespace).Apply(ctx, rbacCM, metav1.ApplyOptions{FieldManager: c.CONTROLLER_NAME, Force: true})
	if err != nil {
		return warnings, err
	}

	arm.Logger.Info("updated RBAC configmap", "ConfigMap", arm.ArgoCDRBACConfiguration.ConfigMapName)

	return warnings, nil
}

func (arm *ArgoCDResourceManager) getArgoCDRBACApplyConfiguration(appProjects *[]*argocdtypes.AppProject, additionalRBACRules string, privilegedGroupConfiguration PrivilegedGroupConfiguration, ldapToProjectMap map[string][]GitlabLDAPGroupLink) (*applyconfigurationsv1.ConfigMapApplyConfiguration, []error, error) {
	warnings := []error{}

	argocdRBACTemplate, err := template.New("argocdRBACTemplate").Parse(argocdRBACTemplateUnparsed)
	if err != nil {
		return nil, warnings, err
	}

	// when trying to figure out the relationship between an appProject and gitlab project
	// the key must contain the name of the project and the namespace

	unprivilegedEnvironmentsSet := set.NewSet[string](arm.ArgoCDRBACConfiguration.UnprivilegedEnvironments...)
	privilegedGroupAllowedActions := set.NewSet[string](c.ALLOWED_PRIVILEGED_GROUP_ACTIONS...)

	adminLdapGroupsToProjects := make(map[string][]*argocdtypes.AppProject)
	readOnlyLdapGroupsToProjects := make(map[string][]*argocdtypes.AppProject)
	readWriteLdapGroupsToProjects := make(map[string][]*argocdtypes.AppProject)

	privilegedGroupToProjectsMap := make(map[string]map[string][]string)
	var mapToUpdate *map[string][]*argocdtypes.AppProject

	adminGroupCoreMap := make(map[string]set.Set[*argocdtypes.AppProject])

	// group names are expected to be in one of two forms -
	// 1) APP_AWS_lob-scope-env_privtype (e.g. APP_AWS_platform-core-dev_admin)
	// 2) APP_k8s_(onprem|)_lob (e.g. APP_k8s_ifi or APP_k8s_onprem_platform) <- these are Rafay groups

	// TODO there must be some way to get this inner for loop outside
	// this code is all insanely unoptimized and needs to be refactored
	for _, project := range *appProjects {
		gitSourceRepo := ""
		for _, repo := range project.Spec.SourceRepos {
			if strings.HasPrefix(repo, "https://gitlab.localhost.com/") {
				gitSourceRepo = strings.TrimPrefix(repo, "https://gitlab.localhost.com/")
				gitSourceRepo = strings.TrimSuffix(gitSourceRepo, ".git")
				break
			} else {
				continue
			}
		}
		if gitSourceRepo == "" {
			arm.Logger.Error(nil, "project source repo not found, skipping", "project", project.Name)
			continue
		}

		if _, ok := project.Labels[c.TW_ACCOUNT_NAME_KEY]; ok {
			projectEnv, penvok := project.Labels[c.TW_ENVIRONMENT_KEY]
			if !penvok {
				arm.Logger.Error(nil, "project environment label not found, skipping", "project", project.Name)
				continue
			}
			// if the project's environment key is in the list of unprivileged
			// environments, proceed as normal; otherwise associate it with
			// any relevant privileged groups and ignore the associations determined above
			lobScope, lsok := project.Labels[c.TW_LOB_SCOPE_KEY]

			// check if the app project is in present in gitlab projects to associate the ldap groups with the correct roles in argoCD
			if ldapLinks, ok := ldapToProjectMap[gitSourceRepo]; ok {

				for _, ldapGroup := range ldapLinks {

					if ldapGroup.GroupAccess == 40 || ldapGroup.GroupAccess == 50 { // 40 maintainer, 40 owner
						mapToUpdate = &adminLdapGroupsToProjects
					} else if ldapGroup.GroupAccess <= 20 { // 5 minimal access, 10 Guest, 20 Reporter
						mapToUpdate = &readOnlyLdapGroupsToProjects
					} else if ldapGroup.GroupAccess == 30 { // 30 developer
						mapToUpdate = &readWriteLdapGroupsToProjects
					} else {
						warnings = append(warnings, fmt.Errorf("LDAP group %s did not have a recognized prefix or suffix, skipping", ldapGroup.CN))
						continue
					}
					if unprivilegedEnvironmentsSet.ContainsOne(projectEnv) {
						(*mapToUpdate)[ldapGroup.CN] = append((*mapToUpdate)[ldapGroup.CN], project)
						// CORE groups should also be granted access to projects in unprivileged envs where the LOB scope of the project matches an LOB scope on the privileged group
						for _, pg := range privilegedGroupConfiguration.PrivilegedGroups {
							lobscopes := set.NewSet[string](pg.LOBScopes...)
							if lobscopes.ContainsOne(lobScope) {
								if _, ok := adminGroupCoreMap[pg.Group]; !ok {
									adminGroupCoreMap[pg.Group] = set.NewSet[*argocdtypes.AppProject]()
								}
								adminGroupCoreMap[pg.Group].Add(project)
							}
						}

					}
				}
			}
			if !lsok {
				arm.Logger.Error(nil, "project is in privileged environment but has no LOB scope label, skipping", "project", project.Name)
				continue
			}
			for _, pg := range privilegedGroupConfiguration.PrivilegedGroups {
				if !privilegedGroupAllowedActions.Contains(pg.Actions...) {
					arm.Logger.Error(nil, "privileged group contains unrecognized action", "group", pg.Group, "actions", pg.Actions, "allowedActions", c.ALLOWED_PRIVILEGED_GROUP_ACTIONS)
					continue
				}

				envs := set.NewSet[string](pg.Environments...)
				if !envs.ContainsOne(projectEnv) {
					continue
				}

				lobscopes := set.NewSet[string](pg.LOBScopes...)
				if !lobscopes.ContainsOne(lobScope) {
					continue
				}

				if pg.NameFilter != "" {
					nameFilterRegex, regexerr := regexp.Compile(pg.NameFilter)
					if regexerr != nil {
						arm.Logger.Error(err, "error compiling regular expression", "regex", nameFilterRegex)
						continue
					}
					if !nameFilterRegex.MatchString(project.Name) {
						continue
					}
				}
				// passed all the checks, add an entry to associate this group and its actions with the project
				pgmap, ok := privilegedGroupToProjectsMap[pg.Group]
				if !ok {
					pgmap = make(map[string][]string)
					privilegedGroupToProjectsMap[pg.Group] = pgmap
				}
				pgmap[project.Name] = pg.Actions
			}

		}
	}

	// add the groups from core to admin map
	for k, v := range adminGroupCoreMap {
		adminLdapGroupsToProjects[k] = v.ToSlice()
	}

	// sort maps for ease of comparison and testing
	for _, projects := range adminLdapGroupsToProjects {
		slices.SortStableFunc(projects, func(a, b *argocdtypes.AppProject) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}

	for _, projects := range readOnlyLdapGroupsToProjects {
		slices.SortStableFunc(projects, func(a, b *argocdtypes.AppProject) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}

	for _, projects := range readWriteLdapGroupsToProjects {
		slices.SortStableFunc(projects, func(a, b *argocdtypes.AppProject) int {
			return cmp.Compare(a.Name, b.Name)
		})
	}

	var argocdRBACTemplateOutput bytes.Buffer
	err = argocdRBACTemplate.Execute(&argocdRBACTemplateOutput, map[string]any{
		"additionalRBACRules":           additionalRBACRules,
		"adminLdapGroupsToProjects":     adminLdapGroupsToProjects,
		"privilegedGroups":              privilegedGroupToProjectsMap,
		"readOnlyLdapGroupsToProjects":  readOnlyLdapGroupsToProjects,
		"readWriteLdapGroupsToProjects": readWriteLdapGroupsToProjects,
		"superadminGroups":              arm.ArgoCDRBACConfiguration.SuperadminGroups,
	})
	if err != nil {
		return nil, warnings, fmt.Errorf("error rendering Argo CD RBAC template: %w", err)
	}

	rbacPolicy := strings.TrimSpace(argocdRBACTemplateOutput.String())
	err = argorbacutil.ValidatePolicy(rbacPolicy)
	if err != nil {
		return nil, warnings, fmt.Errorf("error validating policy: %w", err)
	}

	var policyCsvKey string
	var rbacPolicyFinal string
	if arm.RBACGzip {
		policyCsvKey = fmt.Sprintf("%s.gz", c.RBAC_POLICY_CSV_KEY)
		buf := bytes.Buffer{}
		writer := gzip.NewWriter(&buf)
		_, err = writer.Write([]byte(rbacPolicy))
		if err != nil {
			return nil, warnings, fmt.Errorf("error compressing policy: %w", err)
		}
		err = writer.Close()
		if err != nil {
			return nil, warnings, fmt.Errorf("error closing writer: %w", err)
		}
		rbacPolicyFinal = base64.StdEncoding.EncodeToString(buf.Bytes())
	} else {
		policyCsvKey = c.RBAC_POLICY_CSV_KEY
		rbacPolicyFinal = rbacPolicy
	}

	cmData := map[string]string{
		policyCsvKey: rbacPolicyFinal,
	}
	if arm.ArgoCDRBACConfiguration.DefaultRole != c.RBAC_SPECIAL_NONE_STRING {
		cmData[c.RBAC_POLICY_DEFAULT_KEY] = fmt.Sprintf("role:%s", arm.ArgoCDRBACConfiguration.DefaultRole)
	}

	rbacCM := applyconfigurationsv1.ConfigMap(arm.ArgoCDRBACConfiguration.ConfigMapName, arm.ArgoCDControllerNamespace).
		WithLabels(
			map[string]string{
				c.TW_MANAGED_RESOURCE_KEY: c.CONTROLLER_NAME,
			},
		).WithData(cmData)

	return rbacCM, warnings, nil
}
