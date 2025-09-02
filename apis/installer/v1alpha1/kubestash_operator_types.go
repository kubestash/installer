/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubestashOperator = "KubestashOperator"
	ResourceKubestashOperator     = "kubestashoperator"
	ResourceKubestashOperators    = "kubestashoperators"
)

// KubestashOperator defines the schama for KubestashOperator Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type KubestashOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubestashOperatorSpec `json:"spec,omitempty"`
}

// KubestashOperatorSpec is the schema for Operator Operator values file
type KubestashOperatorSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string       `json:"fullnameOverride"`
	RegistryFQDN     string       `json:"registryFQDN"`
	ReplicaCount     int32        `json:"replicaCount"`
	Operator         ContianerRef `json:"operator"`
	Cleaner          CleanerRef   `json:"cleaner"`
	// +optional
	TaskQueue       TaskQueue `json:"taskQueue,omitempty"`
	ImagePullPolicy string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	// +optional
	Apiserver  WebHookSpec  `json:"apiserver"`
	Monitoring Monitoring   `json:"monitoring"`
	Security   SecuritySpec `json:"security"`
	// +optional
	Distro DistroSpec `json:"distro"`
	// +optional
	NetVolAccessor NetVolAccessor `json:"netVolAccessor"`
	// +optional
	License string `json:"license"`
	// List of sources to populate environment variables in the container.
	// The keys defined within a source must be a C_IDENTIFIER. All invalid keys
	// will be reported as an event when the container is starting. When a key exists in multiple
	// sources, the value associated with the last source will take precedence.
	// Values defined by an Env with a duplicate key will take precedence.
	// Cannot be updated.
	// +optional
	// +listType=atomic
	EnvFrom []core.EnvFromSource `json:"envFrom"`
	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Env []core.EnvVar `json:"env"`
}

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type CleanerRef struct {
	ImageRef `json:",inline"`
	Skip     bool `json:"skip"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
}

type ContianerRef struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type WebHookSpec struct {
	GroupPriorityMinimum    int32 `json:"groupPriorityMinimum"`
	VersionPriority         int32 `json:"versionPriority"`
	EnableMutatingWebhook   bool  `json:"enableMutatingWebhook"`
	EnableValidatingWebhook bool  `json:"enableValidatingWebhook"`
	//+optional
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

// +kubebuilder:validation:Enum=prometheus.io;prometheus.io/operator;prometheus.io/builtin
type (
	MonitoringAgent string
	Monitoring      struct {
		Agent MonitoringAgent `json:"agent"`
		//+optional
		Backup bool `json:"backup"`
		//+optional
		Operator       bool                 `json:"operator"`
		ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
	}
)

type ServiceMonitorLabels struct {
	//+optional
	Labels map[string]string `json:"labels"`
}

type SecuritySpec struct {
	Apparmor ApparmorSpec `json:"apparmor"`
	Seccomp  SeccompSpec  `json:"seccomp"`
}

type ApparmorSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type SeccompSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type DistroSpec struct {
	Openshift bool `json:"openshift"`
}

type NetVolAccessor struct {
	// +optional
	CPU string `json:"cpu"`
	// +optional
	Memory string `json:"memory"`
	// +optional
	RunAsUser int `json:"runAsUser"`
	// +optional
	Privileged bool `json:"privileged"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubestashOperatorList is a list of KubestashOperators
type KubestashOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubestashOperator CRD objects
	Items []KubestashOperator `json:"items,omitempty"`
}
