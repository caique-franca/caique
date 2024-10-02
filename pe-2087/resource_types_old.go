package main

import (
	argocdtypes "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var whitelistedClusterResourcesDefaultold = []metav1.GroupKind{
	{
		Group: "karpenter.k8s.aws",
		Kind:  "EC2NodeClass",
	},
	{
		Group: "karpenter.sh",
		Kind:  "NodePool",
	},
	{
		Group: "storage.k8s.io",
		Kind:  "StorageClass",
	},
}

var blacklistedClusterResourcesDefaultold = []metav1.GroupKind{}

var blacklistedNamespacedResourcesDefaultold = []metav1.GroupKind{
	{
		Group: "",
		Kind:  "LimitRange",
	},
	{
		Group: "networking.k8s.io",
		Kind:  "NetworkPolicy",
	},
	{
		Group: "",
		Kind:  "ResourceQuota",
	},
}

var whitelistedNamespacedResourcesDefaultold = []metav1.GroupKind{
	{
		Group: "",
		Kind:  "ConfigMap",
	},
	{
		Group: "",
		Kind:  "PersistentVolumeClaim",
	},
	// NB: users should not deploy Pods directly, but this needs to be here so that it shows up in the UI
	{
		Group: "",
		Kind:  "Pod",
	},
	{
		Group: "",
		Kind:  "Secret",
	},
	{
		Group: "",
		Kind:  "Service",
	},
	{
		Group: "",
		Kind:  "Endpoint",
	},
	{
		Group: "",
		Kind:  "ServiceAccount",
	},
	{
		Group: "apps",
		Kind:  "Deployment",
	},
	// NB: users should not deploy ReplicaSets directly, but this needs to be here so that it shows up in the UI
	{
		Group: "apps",
		Kind:  "ReplicaSet",
	},
	{
		Group: "apps",
		Kind:  "StatefulSet",
	},
	{
		Group: "autoscaling",
		Kind:  "HorizontalPodAutoscaler",
	},
	{
		Group: "batch",
		Kind:  "CronJob",
	},
	{
		Group: "batch",
		Kind:  "Job",
	},
	{
		Group: "bitnami.com",
		Kind:  "SealedSecret",
	},
	{
		Group: "cert-manager.io",
		Kind:  "Certificate",
	},
	{
		Group: "discovery.k8s.io",
		Kind:  "EndpointSlice",
	},
	{
		Group: "networking.k8s.io",
		Kind:  "Ingress",
	},
	{
		Group: "opentelemetry.io",
		Kind:  "Instrumentation",
	},
	{
		Group: "opentelemetry.io",
		Kind:  "OpenTelemetryCollector",
	},
	{
		Group: "policy",
		Kind:  "PodDisruptionBudget",
	},
	{
		Group: "rbac.authorization.k8s.io",
		Kind:  "Role",
	},
	{
		Group: "rbac.authorization.k8s.io",
		Kind:  "RoleBinding",
	},
	{
		Group: "secrets-store.csi.x-k8s.io",
		Kind:  "SecretProviderClass",
	},
	{
		Group: "autoscaling.k8s.io",
		Kind:  "VerticalPodAutoscaler",
	},
	{
		Group: "external-secrets.io",
		Kind:  "SecretStore",
	},
	{
		Group: "external-secrets.io",
		Kind:  "ExternalSecret",
	},
	{
		Group: "external-secrets.io",
		Kind:  "PushSecret",
	},
	{
		Group: "traefik.io",
		Kind:  "Middleware",
	},
}

var syncWindowsDefaultold = []*argocdtypes.SyncWindow{}
