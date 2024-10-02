package main

import (
	"fmt"

	"github.com/go-logr/stdr"
)

func PrintConfig() {
	logger := stdr.New(nil)
	config, err := LoadResourcesConfigv2(logger, "resources_config.yaml")
	if err != nil {
		logger.Error(err, "Failed to load configuration")
		return
	}

	whitelistedClusterResourcesFinal := config.WhitelistedClusterResources
	blacklistedClusterResourcesFinal := config.BlacklistedClusterResources
	whitelistedNamespacedResourcesFinal := config.WhitelistedNamespacedResources
	blacklistedNamespacedResourcesFinal := config.BlacklistedNamespacedResources
	syncWindowsDefaultfinal := config.SyncWindows

	fmt.Println("Whitelisted Cluster Resources:")
	for _, resource := range whitelistedClusterResourcesFinal {
		fmt.Printf("  Group: %s, Kind: %s\n", resource.Group, resource.Kind)
	}

	fmt.Println("\nBlacklisted Cluster Resources:")
	for _, resource := range blacklistedClusterResourcesFinal {
		fmt.Printf("  Group: %s, Kind: %s\n", resource.Group, resource.Kind)
	}

	fmt.Println("\nWhitelisted Namespaced Resources:")
	for _, resource := range whitelistedNamespacedResourcesFinal {
		fmt.Printf("  Group: %s, Kind: %s\n", resource.Group, resource.Kind)
	}

	fmt.Println("\nBlacklisted Namespaced Resources:")
	for _, resource := range blacklistedNamespacedResourcesFinal {
		fmt.Printf("  Group: %s, Kind: %s\n", resource.Group, resource.Kind)
	}

	fmt.Println("\nSync Windows:")
	for _, syncWindow := range syncWindowsDefaultfinal {
		fmt.Printf("  Window: %v\n", syncWindow)
	}
}
