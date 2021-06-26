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
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	StatefulSet = "statefulset"
	Deployment  = "deployment"
	CloneSet    = "cloneset"

	ComponentInstalled  ComponentPhase = "installed"
	ComponentFail       ComponentPhase = "failed"
	ComponentProcessing ComponentPhase = "processing"
)

type ComponentPhase string

type Trait struct {
	Name       string               `json:"name"`
	Properties runtime.RawExtension `json:"properties,omitempty"`
}

// ComponentSpec defines the desired state of Component
type ComponentSpec struct {
	ComponentType string               `json:"type"` // needed ?
	Properties    runtime.RawExtension `json:"properties,omitempty"`
	TemplateInfo  runtime.RawExtension `json:"template,omitempty"`
	Traits        []Trait
}

type ComponentCondition struct {
}

// ComponentStatus defines the observed state of Component
type ComponentStatus struct {
	Phase      ComponentPhase       `json:"phase,omitempty"`
	Conditions []ComponentCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	Message    string               `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Reason     string               `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Component is the Schema for the components API
type Component struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentSpec   `json:"spec,omitempty"`
	Status ComponentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ComponentList contains a list of Component
type ComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Component `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Component{}, &ComponentList{})
}
