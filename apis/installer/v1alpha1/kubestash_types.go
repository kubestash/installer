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
	ResourceKindKubestash = "Kubestash"
	ResourceKubestash     = "kubestash"
	ResourceKubestashs    = "kubestashs"
)

// Kubestash defines the schama for Kubestash one-click installer.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubestashs,singular=kubestash,categories={kubestash,appscode}
type Kubestash struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubestashSpec `json:"spec,omitempty"`
}

// KubestashSpec is the schema for kubestash chart values file
type KubestashSpec struct {
	Global GlobalValues `json:"global"`

	//+optional
	Operator KubestashOperatorValues `json:"kubestash-operator"`

	//+optional
	Catalog KubestashCatalogValues `json:"kubestash-catalog"`

	//+optional
	Metrics KubestashMetricsValues `json:"kubestash-metrics"`
}

type KubestashOperatorValues struct {
	Enabled                *bool `json:"enabled"`
	*KubestashOperatorSpec `json:",inline"`
}

type KubestashCatalogValues struct {
	Enabled               *bool `json:"enabled"`
	*KubestashCatalogSpec `json:",inline,omitempty"`
}

type KubestashMetricsValues struct {
	Enabled *bool `json:"enabled"`
}

type GlobalValues struct {
	License      string `json:"license"`
	Registry     string `json:"registry"`
	RegistryFQDN string `json:"registryFQDN"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	SkipCleaner      bool                        `json:"skipCleaner"`
	// +optional
	NetworkPolicy NetworkPolicy `json:"networkPolicy"`
}

type NetworkPolicy struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubestashList is a list of Kubestashs
type KubestashList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Kubestash CRD objects
	Items []Kubestash `json:"items,omitempty"`
}
