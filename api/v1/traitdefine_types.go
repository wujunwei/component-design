/*
Copyright 2021.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TraitDefineSpec defines the desired state of TraitDefine
type TraitDefineSpec struct {

	// ApplyFor is the suitable type of component for this trait
	ApplyFor []string `json:"apply_for"`
	//Templates is the template  for helm package the values will be rendered by component ‘s values combined trait's values
	Templates []TraitTemplate `json:"templates,omitempty"`

	IsPatch bool `json:"is_patch"`
}

type TraitTemplate struct {
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}

// TraitDefineStatus defines the observed state of TraitDefine
type TraitDefineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TraitDefine is the Schema for the traitdefines API
type TraitDefine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TraitDefineSpec   `json:"spec,omitempty"`
	Status TraitDefineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TraitDefineList contains a list of TraitDefine
type TraitDefineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TraitDefine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TraitDefine{}, &TraitDefineList{})
}
