// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/coherence/v1.CoherenceCluster":           schema_pkg_apis_coherence_v1_CoherenceCluster(ref),
		"./pkg/apis/coherence/v1.CoherenceClusterSpec":       schema_pkg_apis_coherence_v1_CoherenceClusterSpec(ref),
		"./pkg/apis/coherence/v1.CoherenceClusterStatus":     schema_pkg_apis_coherence_v1_CoherenceClusterStatus(ref),
		"./pkg/apis/coherence/v1.CoherenceInternal":          schema_pkg_apis_coherence_v1_CoherenceInternal(ref),
		"./pkg/apis/coherence/v1.CoherenceInternalSpec":      schema_pkg_apis_coherence_v1_CoherenceInternalSpec(ref),
		"./pkg/apis/coherence/v1.CoherenceInternalStatus":    schema_pkg_apis_coherence_v1_CoherenceInternalStatus(ref),
		"./pkg/apis/coherence/v1.CoherenceInternalStoreSpec": schema_pkg_apis_coherence_v1_CoherenceInternalStoreSpec(ref),
		"./pkg/apis/coherence/v1.CoherenceRole":              schema_pkg_apis_coherence_v1_CoherenceRole(ref),
		"./pkg/apis/coherence/v1.CoherenceRoleSpec":          schema_pkg_apis_coherence_v1_CoherenceRoleSpec(ref),
		"./pkg/apis/coherence/v1.CoherenceRoleStatus":        schema_pkg_apis_coherence_v1_CoherenceRoleStatus(ref),
		"./pkg/apis/coherence/v1.FluentdApplicationSpec":     schema_pkg_apis_coherence_v1_FluentdApplicationSpec(ref),
		"./pkg/apis/coherence/v1.FluentdImageSpec":           schema_pkg_apis_coherence_v1_FluentdImageSpec(ref),
		"./pkg/apis/coherence/v1.ImageSpec":                  schema_pkg_apis_coherence_v1_ImageSpec(ref),
		"./pkg/apis/coherence/v1.Images":                     schema_pkg_apis_coherence_v1_Images(ref),
		"./pkg/apis/coherence/v1.ReadinessProbeSpec":         schema_pkg_apis_coherence_v1_ReadinessProbeSpec(ref),
		"./pkg/apis/coherence/v1.UserArtifactsImageSpec":     schema_pkg_apis_coherence_v1_UserArtifactsImageSpec(ref),
	}
}

func schema_pkg_apis_coherence_v1_CoherenceCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceCluster is the Schema for the coherenceclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/coherence/v1.CoherenceClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/coherence/v1.CoherenceClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.CoherenceClusterSpec", "./pkg/apis/coherence/v1.CoherenceClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceClusterSpec defines the desired state of CoherenceCluster",
				Properties: map[string]spec.Schema{
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "The secrets to be used when pulling images. Secrets must be manually created in the target namespace. ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name to use for the service account to use when RBAC is enabled The role bindings must already have been created as this chart does not create them it just sets the serviceAccountName value in the Pod spec.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"role": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of this role. This value will be used to set the Coherence role property for all members of this role",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "The desired number of cluster members of this role. This is a pointer to distinguish between explicit zero and not specified. Default value is 3.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Description: "Details of the Docker images used in the role",
							Ref:         ref("./pkg/apis/coherence/v1.Images"),
						},
					},
					"storageEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "A boolean flag indicating whether members of this role are storage enabled. This value will set the corresponding coherence.distributed.localstorage System property. If not specified the default value is true. This flag is also used to configure the ScalingPolicy value if a value is not specified. If the StorageEnabled field is not specified or is true the scaling will be safe, if StorageEnabled is set to false scaling will be parallel.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"scalingPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "ScalingPolicy describes how the replicas of the cluster role will be scaled. The default if not specified is based upon the value of the StorageEnabled field. If StorageEnabled field is not specified or is true the default scaling will be safe, if StorageEnabled is set to false the default scaling will be parallel.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "The readiness probe config to be used for the Pods in this role. ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/",
							Ref:         ref("./pkg/apis/coherence/v1.ReadinessProbeSpec"),
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "The extra labels to add to the all of the Pods in this roles. Labels here will add to or override those defined for the cluster. More info: http://kubernetes.io/docs/user-guide/labels",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cacheConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "CacheConfig is the name of the cache configuration file to use",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pofConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "PofConfig is the name of the POF configuration file to use when using POF serializer",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"overrideConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "OverrideConfig is name of the Coherence operational configuration override file, the default is tangosol-coherence-override.xml",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"maxHeap": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxHeap is the min/max heap value to pass to the JVM. The format should be the same as that used for Java's -Xms and -Xmx JVM options. If not set the JVM defaults are used.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"jvmArgs": {
						SchemaProps: spec.SchemaProps{
							Description: "JvmArgs specifies the options to pass to the Coherence JVM. The default is to use the G1 collector.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"javaOpts": {
						SchemaProps: spec.SchemaProps{
							Description: "JavaOpts is miscellaneous JVM options to pass to the Coherence store container This options will override the system options computed in the start up script.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"roles": {
						SchemaProps: spec.SchemaProps{
							Description: "Roles is the list of different roles in the cluster There must be at least one role in a cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/coherence/v1.CoherenceRoleSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"overrideConfig", "maxHeap", "jvmArgs", "javaOpts"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.CoherenceRoleSpec", "./pkg/apis/coherence/v1.Images", "./pkg/apis/coherence/v1.ReadinessProbeSpec"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceClusterStatus defines the observed state of CoherenceCluster",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceInternal(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceInternal is the Schema for the coherenceinternal API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Spec contains the specification for a Coherence cluster. The format is the same as the values file for the Coherence Helm chart.",
							Ref:         ref("./pkg/apis/coherence/v1.CoherenceInternalSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/coherence/v1.CoherenceInternalStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.CoherenceInternalSpec", "./pkg/apis/coherence/v1.CoherenceInternalStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceInternalSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceInternalSpec defines the desired state of CoherenceInternal",
				Properties: map[string]spec.Schema{
					"fullnameOverride": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"clusterSize": {
						SchemaProps: spec.SchemaProps{
							Description: "The size of the cluster",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"role": {
						SchemaProps: spec.SchemaProps{
							Description: "The role name of a Coherence cluster member",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Description: "The name to use for the service account to use when RBAC is enabled The role bindings must already have been created as this chart does not create them it just sets the serviceAccountName value in the Pod spec.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "The secrets to be used when pulling images. Secrets must be manually created in the target namespace. ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"coherence": {
						SchemaProps: spec.SchemaProps{
							Description: "The Coherence Docker image settings",
							Ref:         ref("./pkg/apis/coherence/v1.ImageSpec"),
						},
					},
					"coherenceUtils": {
						SchemaProps: spec.SchemaProps{
							Description: "The Coherence Utilities Docker image settings",
							Ref:         ref("./pkg/apis/coherence/v1.ImageSpec"),
						},
					},
					"store": {
						SchemaProps: spec.SchemaProps{
							Description: "The store settings",
							Ref:         ref("./pkg/apis/coherence/v1.CoherenceInternalStoreSpec"),
						},
					},
					"logCaptureEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Controls whether or not log capture via EFK stack is enabled.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"fluentd": {
						SchemaProps: spec.SchemaProps{
							Description: "Specify the fluentd image These parameters are ignored if 'LogCaptureEnabled' is false.",
							Ref:         ref("./pkg/apis/coherence/v1.FluentdImageSpec"),
						},
					},
					"userArtifacts": {
						SchemaProps: spec.SchemaProps{
							Description: "The user artifacts image settings",
							Ref:         ref("./pkg/apis/coherence/v1.UserArtifactsImageSpec"),
						},
					},
				},
				Required: []string{"cluster"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.CoherenceInternalStoreSpec", "./pkg/apis/coherence/v1.FluentdImageSpec", "./pkg/apis/coherence/v1.ImageSpec", "./pkg/apis/coherence/v1.UserArtifactsImageSpec"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceInternalStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceInternalStatus defines the observed state of CoherenceInternal",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceInternalStoreSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceInternalStoreSpec defines the desired state of CoherenceInternal stores",
				Properties: map[string]spec.Schema{
					"storageEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "A boolean flag indicating whether members of this role are storage enabled. If not specified the default value is true.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"wka": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of the headless service used for Coherence WKA",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "The extra labels to add to the Coherence Pod. More info: http://kubernetes.io/docs/user-guide/labels",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "The readiness probe config. ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/",
							Ref:         ref("./pkg/apis/coherence/v1.ReadinessProbeSpec"),
						},
					},
					"cacheConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "CacheConfig is the name of the cache configuration file to use",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pofConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "PofConfig is the name of the POF configuration file to use when using POF serializer",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"overrideConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "OverrideConfig is name of the Coherence operational configuration override file, the default is tangosol-coherence-override.xml",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"maxHeap": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxHeap is the min/max heap value to pass to the JVM. The format should be the same as that used for Java's -Xms and -Xmx JVM options. If not set the JVM defaults are used.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"jvmArgs": {
						SchemaProps: spec.SchemaProps{
							Description: "JvmArgs specifies the options to pass to the Coherence JVM. The default is to use the G1 collector.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"javaOpts": {
						SchemaProps: spec.SchemaProps{
							Description: "JavaOpts is miscellaneous JVM options to pass to the Coherence store container This options will override the system options computed in the start up script.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"overrideConfig", "maxHeap", "jvmArgs", "javaOpts"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.ReadinessProbeSpec"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceRole(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceRole is the Schema for the coherenceroles API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/coherence/v1.CoherenceRoleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/coherence/v1.CoherenceRoleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.CoherenceRoleSpec", "./pkg/apis/coherence/v1.CoherenceRoleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceRoleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceRoleSpec defines a role in a Coherence cluster. A role is one or more Pods that perform the same functionality, for example storage members.",
				Properties: map[string]spec.Schema{
					"role": {
						SchemaProps: spec.SchemaProps{
							Description: "The name of this role. This value will be used to set the Coherence role property for all members of this role",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "The desired number of cluster members of this role. This is a pointer to distinguish between explicit zero and not specified. Default value is 3.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Description: "Details of the Docker images used in the role",
							Ref:         ref("./pkg/apis/coherence/v1.Images"),
						},
					},
					"storageEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "A boolean flag indicating whether members of this role are storage enabled. This value will set the corresponding coherence.distributed.localstorage System property. If not specified the default value is true. This flag is also used to configure the ScalingPolicy value if a value is not specified. If the StorageEnabled field is not specified or is true the scaling will be safe, if StorageEnabled is set to false scaling will be parallel.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"scalingPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "ScalingPolicy describes how the replicas of the cluster role will be scaled. The default if not specified is based upon the value of the StorageEnabled field. If StorageEnabled field is not specified or is true the default scaling will be safe, if StorageEnabled is set to false the default scaling will be parallel.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "The readiness probe config to be used for the Pods in this role. ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/",
							Ref:         ref("./pkg/apis/coherence/v1.ReadinessProbeSpec"),
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "The extra labels to add to the all of the Pods in this roles. Labels here will add to or override those defined for the cluster. More info: http://kubernetes.io/docs/user-guide/labels",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cacheConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "CacheConfig is the name of the cache configuration file to use",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pofConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "PofConfig is the name of the POF configuration file to use when using POF serializer",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"overrideConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "OverrideConfig is name of the Coherence operational configuration override file, the default is tangosol-coherence-override.xml",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"maxHeap": {
						SchemaProps: spec.SchemaProps{
							Description: "MaxHeap is the min/max heap value to pass to the JVM. The format should be the same as that used for Java's -Xms and -Xmx JVM options. If not set the JVM defaults are used.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"jvmArgs": {
						SchemaProps: spec.SchemaProps{
							Description: "JvmArgs specifies the options to pass to the Coherence JVM. The default is to use the G1 collector.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"javaOpts": {
						SchemaProps: spec.SchemaProps{
							Description: "JavaOpts is miscellaneous JVM options to pass to the Coherence store container This options will override the system options computed in the start up script.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"overrideConfig", "maxHeap", "jvmArgs", "javaOpts"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.Images", "./pkg/apis/coherence/v1.ReadinessProbeSpec"},
	}
}

func schema_pkg_apis_coherence_v1_CoherenceRoleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceRoleStatus defines the observed state of CoherenceRole",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "The current status.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the desired size of the Coherence cluster.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"currentReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "CurrentReplicas is the current size of the Coherence cluster.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"readyReplicas": {
						SchemaProps: spec.SchemaProps{
							Description: "ReadyReplicas is the number of Pods created by the StatefulSet.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"selector": {
						SchemaProps: spec.SchemaProps{
							Description: "label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"replicas", "currentReplicas", "readyReplicas"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_FluentdApplicationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FluentdImageSpec defines the settings for the fluentd application",
				Properties: map[string]spec.Schema{
					"configFile": {
						SchemaProps: spec.SchemaProps{
							Description: "The fluentd configuration file configuring source for application log.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tag": {
						SchemaProps: spec.SchemaProps{
							Description: "This value should be source.tag from fluentd.application.configFile.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_FluentdImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "FluentdImageSpec defines the settings for the fluentd image",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Image pull policy. One of Always, Never, IfNotPresent. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"application": {
						SchemaProps: spec.SchemaProps{
							Description: "The fluentd application configuration",
							Ref:         ref("./pkg/apis/coherence/v1.FluentdApplicationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.FluentdApplicationSpec"},
	}
}

func schema_pkg_apis_coherence_v1_ImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CoherenceInternalImageSpec defines the settings for a Docker image",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Image pull policy. One of Always, Never, IfNotPresent. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_Images(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Images defines the different Docker images used in the role",
				Properties: map[string]spec.Schema{
					"coherence": {
						SchemaProps: spec.SchemaProps{
							Description: "CoherenceImage is the details of the Coherence image to be used",
							Ref:         ref("./pkg/apis/coherence/v1.ImageSpec"),
						},
					},
					"coherenceUtils": {
						SchemaProps: spec.SchemaProps{
							Description: "CoherenceUtils is the details of the Coherence utilities image to be used",
							Ref:         ref("./pkg/apis/coherence/v1.ImageSpec"),
						},
					},
					"userArtifacts": {
						SchemaProps: spec.SchemaProps{
							Description: "UserArtifacts configures the image containing jar files and configuration files that are added to the Coherence JVM's classpath.",
							Ref:         ref("./pkg/apis/coherence/v1.UserArtifactsImageSpec"),
						},
					},
					"fluentd": {
						SchemaProps: spec.SchemaProps{
							Description: "Fluentd defines the settings for the fluentd image",
							Ref:         ref("./pkg/apis/coherence/v1.FluentdImageSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/coherence/v1.FluentdImageSpec", "./pkg/apis/coherence/v1.ImageSpec", "./pkg/apis/coherence/v1.UserArtifactsImageSpec"},
	}
}

func schema_pkg_apis_coherence_v1_ReadinessProbeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReadinessProbeSpec defines the settings for the Coherence Pod readiness probe",
				Properties: map[string]spec.Schema{
					"initialDelaySeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"timeoutSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of seconds after which the probe times out. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"periodSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "How often (in seconds) to perform the probe.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"successThreshold": {
						SchemaProps: spec.SchemaProps{
							Description: "Minimum consecutive successes for the probe to be considered successful after having failed.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"failureThreshold": {
						SchemaProps: spec.SchemaProps{
							Description: "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_coherence_v1_UserArtifactsImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UserArtifactsImageSpec defines the settings for the user artifacts image",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Image pull policy. One of Always, Never, IfNotPresent. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"libDir": {
						SchemaProps: spec.SchemaProps{
							Description: "The folder in the custom artifacts Docker image containing jar files to be added to the classpath of the Coherence container. If not set the libDir is \"/files/lib\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configDir": {
						SchemaProps: spec.SchemaProps{
							Description: "The folder in the custom artifacts Docker image containing configuration files to be added to the classpath of the Coherence container. If not set the configDir is \"/files/conf\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
