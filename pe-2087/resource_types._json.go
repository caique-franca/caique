package main

import (
	"encoding/json"
	"log"
	"os"

	argocdtypes "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var whitelistedClusterResourcesDefaultjson []metav1.GroupKind
var blacklistedClusterResourcesDefaultjson []metav1.GroupKind
var whitelistedNamespacedResourcesDefaultjson []metav1.GroupKind
var blacklistedNamespacedResourcesDefaultjson []metav1.GroupKind
var syncWindowsDefaultjson []*argocdtypes.SyncWindow

func ResourcesTypesJson() {
	// Path to the JSON file
	configFile := "resources_config.json"

	// Load the JSON file
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration from JSON: %v", err)
	}

	var config struct {
		WhitelistedClusterResources    []metav1.GroupKind        `json:"whitelistedClusterResources"`
		BlacklistedClusterResources    []metav1.GroupKind        `json:"blacklistedClusterResources"`
		WhitelistedNamespacedResources []metav1.GroupKind        `json:"whitelistedNamespacedResources"`
		BlacklistedNamespacedResources []metav1.GroupKind        `json:"blacklistedNamespacedResources"`
		SyncWindows                    []*argocdtypes.SyncWindow `json:"syncWindows"`
	}

	// Deserialize the JSON
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Assign the loaded values to the global variables
	whitelistedClusterResourcesDefaultjson = config.WhitelistedClusterResources
	blacklistedClusterResourcesDefaultjson = config.BlacklistedClusterResources
	whitelistedNamespacedResourcesDefaultjson = config.WhitelistedNamespacedResources
	blacklistedNamespacedResourcesDefaultjson = config.BlacklistedNamespacedResources
	syncWindowsDefaultjson = config.SyncWindows
}
