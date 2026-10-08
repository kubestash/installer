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
	"kmodules.xyz/resource-metadata/apis/shared"
)

const (
	ResourceKindKubestashOpscenter = "KubestashOpscenter"
	ResourceKubestashOpscenter     = "kubestashopscenter"
	ResourceKubestashOpscenters    = "kubestashopscenters"
)

// KubestashOpscenter defines the schama for KubeStash Opscenter installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubestashopscenters,singular=kubestashopscenter,categories={kubestash,appscode}
type KubestashOpscenter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubestashOpscenterSpec `json:"spec,omitempty"`
}

// KubestashOpscenterSpec is the schema for kubestash-opscenter chart values file
type KubestashOpscenterSpec struct {
	Global OpscenterGlobalValues `json:"global"`
	//+optional
	Metrics KubestashMetricsValues `json:"kubestash-metrics"`
	//+optional
	UiServer KubestashUiServerValues `json:"kubestash-ui-server"`
	//+optional
	AceUserRoles AceUserRolesValues `json:"ace-user-roles"`
}

type KubestashUiServerValues struct {
	Enabled                *bool `json:"enabled"`
	*KubestashUiServerSpec `json:",inline,omitempty"`
}

type OpscenterGlobalValues struct {
	License string `json:"license"`
	//+optional
	LicenseSecretName string `json:"licenseSecretName"`
	Registry          string `json:"registry"`
	RegistryFQDN      string `json:"registryFQDN"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	Monitoring       EASMonitoring               `json:"monitoring"`
	// +optional
	NetworkPolicy NetworkPolicy `json:"networkPolicy"`
	// +optional
	Distro shared.DistroSpec `json:"distro"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubestashOpscenterList is a list of KubestashOpscenters
type KubestashOpscenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubestashOpscenter CRD objects
	Items []KubestashOpscenter `json:"items,omitempty"`
}
