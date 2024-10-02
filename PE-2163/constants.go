package constants

import (
	"fmt"
)

const (
	ARGOCD_COORDINATOR_METADATA_CONFIGMAP = "argocd-coordinator-metadata-cm" // this is the name of the CM in kube-system that the coordinator tries to read from remote clusters to discover extra labels/annotations
	ACM_CM_ANNOTATIONS_KEY                = "annotations.yaml"
	ACM_CM_LABELS_KEY                     = "labels.yaml"
	ALL_ENVS                              = "allenvs"
	ARGOCD_DEPLOYER                       = "argocd-deployer"
	ARGOCD_DEFAULT_NAMESPACE              = "argocd"
	ARGOCD_DEFAULT_PROJECT_NAME           = "default"
	ARGOCD_SECRET_TYPE_LABEL              = "argocd.argoproj.io/secret-type"
	ARGOCD_SECRET_TYPE_CLUSTER_VALUE      = "cluster"
	ARGOCD_RESOURCES_FINALIZER            = "resources-finalizer.argocd.argoproj.io"
	CONTROLLER_NAME                       = "argocd-coordinator-controller"
	DISABLED                              = "disabled"
	ENABLED                               = "enabled"
	KUBERNETES                            = "kubernetes"
	KUBE_SYSTEM                           = "kube-system"
	KUBERNETES_APP_DOMAIN_PREFIX          = "app.kubernetes.io"
	KUBERNETES_APP_NAME_LABEL             = "app.kubernetes.io/name"
	POD_READINESS_GATE_LABEL              = "elbv2.k8s.aws/pod-readiness-gate-inject"
	RAFAY_HAPROXY                         = "rafay-haproxy"
	localhost_DOMAIN_PREFIX               = "localhost.cloud"

	DEFAULT_CONFIG_PATH                             = "/mnt/configs"
	APP_PROJECT_OVERRIDE_DEFAULT_FILE_NAME          = "app-project-override.yaml"
	APP_PROJECT_METADATA_OVERRIDE_DEFAULT_FILE_NAME = "app-project-metadata-override.yaml"
	PRIVILEGED_GROUP_CONFIG_DEFAULT_FILE_NAME       = "privileged-group-configuration.yaml"
	RBAC_OVERRIDE_CONFIG_DEFAULT_FILE_NAME          = "rbac-override-config.yaml"

	COORDINATOR_MODE_APPLICATION    = "application"
	COORDINATOR_MODE_INFRASTRUCTURE = "infrastructure"

	CLUSTER_TYPE_EKS    = "eks"
	CLUSTER_TYPE_RAFAY  = "rafay"
	CLUSTER_TYPE_ONPREM = "onprem"

	CLUSTER_SECRET_CONFIG_KEY = "config"
	CLUSTER_SECRET_NAME_KEY   = "name"
	CLUSTER_SECRET_SERVER_KEY = "server"

	SYNCED_MAP_CLUSTER_SECRETS_KEY = "clusterSecrets"
	SYNCED_MAP_APPPROJECTS_KEY     = "appProjects"
	SYNCED_MAP_RAFAY_CLUSTERS_KEY  = "rafayClusters"
	SYNCED_MAP_ONPREM_CLUSTERS_KEY = "onpremClusters"

	LDAP_GROUP_NAME_ADMIN_SUFFIX     = "_admin"
	LDAP_GROUP_NAME_PREFIX           = "APP_AWS_"
	LDAP_GROUP_NAME_READONLY_SUFFIX  = "_read-only"
	LDAP_GROUP_NAME_READWRITE_SUFFIX = "_read-write"
	RAFAY_AD_GROUP_PREFIX            = "APP_k8s_"
	RBAC_POLICY_DEFAULT_KEY          = "policy.default"
	RBAC_POLICY_CSV_KEY              = "policy.csv"
	RBAC_SPECIAL_NONE_STRING         = "NONE"

	METRIC_LABEL_APPPROJECT_NAME               = "appproject_name"
	METRIC_LABEL_APPPROJECT_STATUS             = "appproject_status"
	METRIC_LABEL_DISCOVERED_CLUSTER_IDENTIFIER = "discovered_cluster_identifier"
	METRIC_LABEL_DISCOVERED_CLUSTER_STATUS     = "discovery_status"
	METRIC_LABEL_DISCOVERED_CLUSTER_TYPE       = "discovered_cluster_type"
	METRIC_LABEL_EVALUATED_NAMESPACE           = "evaluated_namespace"
	METRIC_LABEL_NAMESPACE_GENERATION_STATUS   = "namespace_generation_status"

	METRIC_STATUS_DESCRIBE_FAILED            = "describe_failed"
	METRIC_STATUS_DISCOVERY_FAILED           = "discovery_failed"
	METRIC_STATUS_KUBERNETES_CLIENT_ERROR    = "kubernetes_client_error"
	METRIC_STATUS_NAMESPACE_RECONCILE_ERROR  = "namespace_reconcile_error"
	METRIC_STATUS_SCRAPE_SA_MISCONFIGURATION = "scrape_serviceaccount_misconfiguration"
	METRIC_STATUS_SCRAPE_SA_RBAC_ERROR       = "scrape_serviceaccount_rbac_error"
	METRIC_STATUS_SCRAPE_SA_SECRET_ERROR     = "scrape_serviceaccount_secret_error"
	METRIC_STATUS_SCRAPE_SA_UNAVAILABLE      = "scrape_serviceaccount_unavailable"
	METRIC_STATUS_SETUP_FAILURE              = "setup_failure"
	METRIC_STATUS_SUCCEEDED                  = "succeeded"

	METRIC_VALUE_OK   float64 = 0
	METRIC_VALUE_FAIL float64 = 1
)

var (
	ALLOWED_PRIVILEGED_GROUP_ACTIONS = []string{"delete", "exec", "sync"}
	COORDINATOR_MODES                = []string{COORDINATOR_MODE_APPLICATION, COORDINATOR_MODE_INFRASTRUCTURE}

	KUBE_APP_COMPONENT_KEY = fmt.Sprintf("%s/component", KUBERNETES_APP_DOMAIN_PREFIX)
	KUBE_APP_INSTANCE_KEY  = fmt.Sprintf("%s/instance", KUBERNETES_APP_DOMAIN_PREFIX)
	KUBE_APP_NAME_KEY      = fmt.Sprintf("%s/name", KUBERNETES_APP_DOMAIN_PREFIX)
	KUBE_APP_PART_OF_KEY   = fmt.Sprintf("%s/part-of", KUBERNETES_APP_DOMAIN_PREFIX)
	KUBE_APP_VERSION_KEY   = fmt.Sprintf("%s/version", KUBERNETES_APP_DOMAIN_PREFIX)

	// this is for removing Gitlab stuff that we should not create ArgoCD resources for (like template repos)
	FORBIDDEN_GITLAB_PROJECT_SUBSTRINGS = []string{"template"}
	FORBIDDEN_APP_KUBERNETES_NAMESPACES = []string{
		"aws-observability",
		"cert-manager",
		"default",
		"istio-system",
		"karpenter",
		"kube-node-lease",
		"kube-public",
		"kube-system",
		"logging",
		"monitoring",
		"prometheus-scraper",
		"rafay-infra",
		"rafay-system",
	}

	// these are here as a guard against mistakes (so that we don't accidentally deploy infra stuff to default)
	FORBIDDEN_INFRA_KUBERNETES_NAMESPACES = []string{
		"default",
		"kube-node-lease",
		"kube-public",
	}
	// another safeguard to minimize conflict potential, keep infra stuff in namespaces we expect infra to be in
	// (basically forbidden app namespaces minus forbidden infra ones)
	ALLOWED_INFRA_KUBERNETES_NAMESPACES = []string{
		"aws-observability",
		"cert-manager",
		"istio-system",
		"karpenter",
		"kube-system",
		"logging",
		"monitoring",
		"prometheus-scraper",
		"rafay-infra",
		"rafay-system",
	}

	// all mappings are in here, even when both sides are the same, to avoid whiny logging
	SHORT_LOB_TO_GITLAB_NAMESPACE_MAP = map[string]string{
		"dpt": "data",
		"ifi": "ifi",
		"inf": "infra",
		"ptt": "platform",
		"twd": "twd",
	}

	localhost_METADATA_DOMAIN_PREFIX = fmt.Sprintf("metadata.%s", localhost_DOMAIN_PREFIX)

	TW_ACCOUNT_ID_KEY               = fmt.Sprintf("%s/account-id", localhost_DOMAIN_PREFIX)   // used for EKS clusters only
	TW_ACCOUNT_NAME_KEY             = fmt.Sprintf("%s/account-name", localhost_DOMAIN_PREFIX) // used for EKS clusters only
	TW_BYPASS_SOURCE_REPO_CHECK_KEY = fmt.Sprintf("%s/bypass-source-repo-check", localhost_DOMAIN_PREFIX)
	TW_ARGOCD_CLUSTER_NAME_KEY      = fmt.Sprintf("%s/argocd-cluster-name", localhost_DOMAIN_PREFIX)
	TW_CLUSTER_ARN_KEY              = fmt.Sprintf("%s/cluster-arn", localhost_DOMAIN_PREFIX) // used for EKS clusters only
	TW_CLUSTER_DELETED_MARKER_KEY   = fmt.Sprintf("%s/cluster-deleted-marker", localhost_DOMAIN_PREFIX)
	TW_CLUSTER_NAME_KEY             = fmt.Sprintf("%s/cluster-name", localhost_DOMAIN_PREFIX)
	TW_CLUSTER_REGION_KEY           = fmt.Sprintf("%s/cluster-region", localhost_DOMAIN_PREFIX)
	TW_CLUSTER_TYPE_KEY             = fmt.Sprintf("%s/cluster-type", localhost_DOMAIN_PREFIX)
	TW_ENVIRONMENT_KEY              = fmt.Sprintf("%s/environment", localhost_DOMAIN_PREFIX)
	TW_IGNORE_OWNING_REPOSITORY_KEY = fmt.Sprintf("%s/ignore-owning-repository", localhost_DOMAIN_PREFIX)
	TW_KUBERNETES_GITVERSION_KEY    = fmt.Sprintf("%s/kubernetes-gitversion", localhost_DOMAIN_PREFIX)
	TW_KUBERNETES_VERSION_KEY       = fmt.Sprintf("%s/kubernetes-version", localhost_DOMAIN_PREFIX)
	TW_LOB_SCOPE_KEY                = fmt.Sprintf("%s/lob-scope", localhost_DOMAIN_PREFIX)
	TW_LOCKED_FOR_PROCESSING_KEY    = fmt.Sprintf("%s/locked-for-processing", localhost_DOMAIN_PREFIX)
	TW_MANAGED_RESOURCE_KEY         = fmt.Sprintf("%s/managed-by", localhost_DOMAIN_PREFIX)
	TW_ONPREM_HAPROXY_SERVICE_KEY   = fmt.Sprintf("%s/on-prem-haproxy-lb", localhost_DOMAIN_PREFIX) // used to find Service resources that front HAProxy endpoints to on-prem Kubernetes clusters
	TW_OWNING_REPOSITORY_KEY        = fmt.Sprintf("%s/owning-repository", localhost_DOMAIN_PREFIX)  // inspected by validator to see if the app is being deployed from the same repo

	// any repo URLs in here will be added to all projects
	UNIVERSAL_REPOS = []string{"https://artifactory.localhost.com/artifactory/tweb-helm-dev"}

	// Gitlab namespaces with these suffixes will be accepted
	VALID_ENVIRONMENTS = []string{"dev", "qa", "prod", "sandbox"}
)
