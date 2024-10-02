package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	argocdtypes "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/go-logr/stdr"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	whitelistedClusterResourcesDefault    []metav1.GroupKind
	blacklistedClusterResourcesDefault    []metav1.GroupKind
	whitelistedNamespacedResourcesDefault []metav1.GroupKind
	blacklistedNamespacedResourcesDefault []metav1.GroupKind
	syncWindowsDefault                    []*argocdtypes.SyncWindow
)

type ResourcesConfig struct {
	WhitelistedClusterResources    []metav1.GroupKind        `yaml:"whitelistedClusterResources"`
	BlacklistedClusterResources    []metav1.GroupKind        `yaml:"blacklistedClusterResources"`
	WhitelistedNamespacedResources []metav1.GroupKind        `yaml:"whitelistedNamespacedResources"`
	BlacklistedNamespacedResources []metav1.GroupKind        `yaml:"blacklistedNamespacedResources"`
	SyncWindows                    []*argocdtypes.SyncWindow `yaml:"syncWindows"`
}

func main() {
	config, err := loadResourcesConfig("resources_config.yaml")
	if err != nil {
		log.Fatalf("Failed to load resources configuration: %v", err)
	}

	logger := stdr.New(nil)
	configv2, _ := LoadResourcesConfigv2(logger, "resources_config.yaml")

	whitelistedClusterResourcesDefault = config.WhitelistedClusterResources
	blacklistedClusterResourcesDefault = config.BlacklistedClusterResources
	whitelistedNamespacedResourcesDefault = config.WhitelistedNamespacedResources
	blacklistedNamespacedResourcesDefault = config.BlacklistedNamespacedResources
	syncWindowsDefault = config.SyncWindows

	ResourcesTypesJson()
	PrintConfig()

	if CompareVariables(whitelistedClusterResourcesDefaultold, configv2.WhitelistedClusterResources) {
		fmt.Println("The variables are equal.")
	} else {
		fmt.Println("The variables are different.")
	}

	if CompareVariables(blacklistedClusterResourcesDefaultold, configv2.BlacklistedClusterResources) {
		fmt.Println("The variables are equal.")
	} else {
		fmt.Println("The variables are different.")
	}

	if CompareVariables(whitelistedNamespacedResourcesDefaultold, configv2.WhitelistedNamespacedResources) {
		fmt.Println("The variables are equal.")
	} else {
		fmt.Println("The variables are different.")
	}

	if CompareVariables(blacklistedNamespacedResourcesDefaultold, configv2.BlacklistedNamespacedResources) {
		fmt.Println("The variables are equal.")
	} else {
		fmt.Println("The variables are different.")
	}

	if CompareSyncWindows(syncWindowsDefaultold, configv2.SyncWindows) {
		fmt.Println("The variables are equal.")
	} else {
		fmt.Println("The variables are different.")
	}

}

func CompareVariables(var1, var2 []metav1.GroupKind) bool {
	return reflect.DeepEqual(var1, var2)
}

func CompareSyncWindows(var1, var2 []*argocdtypes.SyncWindow) bool {
	return reflect.DeepEqual(var1, var2)
}

func loadResourcesConfig(path string) (*ResourcesConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	var config ResourcesConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal YAML: %w", err)
	}

	return &config, nil
}
