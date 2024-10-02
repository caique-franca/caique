package main

import (
	"os"

	argocdtypes "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourcesConfigv2 struct {
	WhitelistedClusterResources    []metav1.GroupKind        `yaml:"whitelistedClusterResources"`
	BlacklistedClusterResources    []metav1.GroupKind        `yaml:"blacklistedClusterResources"`
	WhitelistedNamespacedResources []metav1.GroupKind        `yaml:"whitelistedNamespacedResources"`
	BlacklistedNamespacedResources []metav1.GroupKind        `yaml:"blacklistedNamespacedResources"`
	SyncWindows                    []*argocdtypes.SyncWindow `yaml:"syncWindows"`
}

// LoadResourcesConfig loads the resource configuration from a YAML file and returns a ResourcesConfig struct
func LoadResourcesConfigv2(logger logr.Logger, path string) (*ResourcesConfigv2, error) {
	logger.Info("Loading allowlists/denylists resources configuration", "path", path)

	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err, "Failed to read configuration file", "path", path)
		return nil, err
	}

	var configv2 ResourcesConfigv2
	err = yaml.Unmarshal(data, &configv2)
	if err != nil {
		logger.Error(err, "Failed to unmarshal YAML", "path", path)
		return nil, err
	}

	logger.Info("Configuration allowlists/denylists resources loaded successfully")
	return &configv2, nil
}
