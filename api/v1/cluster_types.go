/*
Copyright 2023 Nguyen Thanh Nguyen.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Blueprint Spec of Cluster Resource referred
type BlueprintInfoSpec struct {
	// Name of kind Blueprint
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`
	// Type Blueprint
	// +kubebuilder:validation:Required
	Type string `json:"type,omitempty"`
	// Revision of Blueprint
	// +kubebuilder:validation:Required
	Revision string `json:"revision,omitempty"`
	// Published Version of Blueprint
	// +kubebuilder:validation:Required
	Version string `json:"version,omitempty"`
}

// Content of Blueprint Packages
type BlueprintInfo struct {
	// Name of Blueprint

	Name string `json:"name,omitempty"`
	// Spec

	Spec BlueprintInfoSpec `json:"spec,omitempty"`
	// Override field of blueprint
	// +kubebuilder:validation:Optional
	Override map[string]string `json:"override,omitempty"`
}

type BlueprintInfoList struct {
	Items []Blueprint `json:"items,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	Infrastructure []BlueprintInfo `json:"infrastructure,omitempty"`
	Software       []BlueprintInfo `json:"software,omitempty"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Status of cluster
	Status string `json:"status,omitempty"`
	// Sync status of cluster
	Sync string `json:"sync,omitempty"`
	// SHA of cluster package
	SHA string `json:"sha,omitempty"`
	// Repo contains cluster package
	Repo string `json:"repo,omitempty"`
	// Version  of cluster package
	Version string `json:"version,omitempty"`
	// Revision of cluster package
	Revision string `json:"revision,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:pruning:PreserveUnknownFields
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="Time duration since creation of Cluster"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.status",description="Cluster status"
// +kubebuilder:printcolumn:name="SHA",type="string",JSONPath=".status.sha",description="SHA"
// +kubebuilder:printcolumn:name="Repo",type="string",JSONPath=".status.repo",description="Repo"
// +kubebuilder:printcolumn:name="Sync",type="string",JSONPath=".status.sync",description="Sync"
// Cluster is the Schema for the clusters API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ClusterList contains a list of Cluster
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
