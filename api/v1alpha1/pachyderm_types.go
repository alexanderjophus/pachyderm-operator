/*
Copyright 2023.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Pachyderm is the Schema for the pachyderms API
type Pachyderm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PachydermSpec   `json:"spec,omitempty"`
	Status PachydermStatus `json:"status,omitempty"`
}

// PachydermSpec defines the desired state of Pachyderm
type PachydermSpec struct {
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	Description string `json:"description"`
	// +kubebuilder:validation:Required
	Project string `json:"project"`
	// +kubebuilder:validation:Required
	Input *Input `json:"input"`
	// +kubebuilder:validation:Required
	Transform *Transform `json:"transform"`
}

// PachydermStatus defines the observed state of Pachyderm
type PachydermStatus struct {
}

type Input struct {
	// +kubebuilder:validation:Required
	Pfs *Pfs `json:"pfs"`
}

type Pfs struct {
	// +kubebuilder:validation:Required
	Repo string `json:"repo"`
	// +kubebuilder:validation:Required
	Glob string `json:"glob"`
}

type Transform struct {
	// +kubebuilder:validation:Required
	Image string `json:"image"`
	// +kubebuilder:validation:Required
	Cmd []string `json:"cmd"`
	// +kubebuilder:validation:Optional
	StdIn []string `json:"stdin"`
}

//+kubebuilder:object:root=true

// PachydermList contains a list of Pachyderm
type PachydermList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pachyderm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Pachyderm{}, &PachydermList{})
}
