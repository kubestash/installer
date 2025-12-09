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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubestashCatalog = "KubestashCatalog"
	ResourceKubestashCatalog     = "kubestashcatalog"
	ResourceKubestashCatalogs    = "kubestashcatalogs"
)

// KubestashCatalog defines the schema for Stash Catalog chart.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubestashcatalogs,singular=kubestashcatalog,categories={stash,appscode}
type KubestashCatalog struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubestashCatalogSpec `json:"spec,omitempty"`
}

// KubestashCatalogSpec is the schema for Stash Postgres values file
type KubestashCatalogSpec struct {
	//+optional
	Proxies        RegistryProxies         `json:"proxies"`
	WaitTimeout    int64                   `json:"waitTimeout"`
	KubeDump       KubeDumpSpec            `json:"kubedump"`
	Pvc            StashPvcSpec            `json:"pvc"`
	Volumesnapshot StashVolumesnapshotSpec `json:"volumesnapshot"`
	Workload       StashWorkloadSpec       `json:"workload"`
	Manifest       StashManifestSpec       `json:"manifest"`
	// +optional
	Distro DistroSpec `json:"distro"`
}

type RegistryProxies struct {
	// ghcr.io
	//+optional
	GHCR string `json:"ghcr"`
}

type KubeDumpSpec struct {
	Enabled bool           `json:"enabled"`
	Backup  KubeDumpBackup `json:"backup"`
}

type KubeDumpBackup struct {
	Sanitize          bool   `json:"sanitize"`
	LabelSelector     string `json:"labelSelector"`
	IncludeDependants bool   `json:"includeDependants"`
}

type StashPvcSpec struct {
	Enabled bool `json:"enabled"`
}

type StashVolumesnapshotSpec struct {
	Enabled bool `json:"enabled"`
}

type StashWorkloadSpec struct {
	Enabled bool `json:"enabled"`
}

type StashManifestSpec struct {
	Enabled bool `json:"enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubestashCatalogList is a list of KubestashCatalogs
type KubestashCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StashPostgres CRD objects
	Items []KubestashCatalog `json:"items,omitempty"`
}
