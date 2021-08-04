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

// ComponentDefineSpec defines the desired state of ComponentDefine
type ComponentDefineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Repo     Repository `json:"repo"`
	Provider string     `json:"provider"`
}

type Repository struct {
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	URI    string `json:"uri"`
	Auth   `json:",omitempty,inline"`
}
type Auth struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	BasicAuth string `json:"basic_auth"`
}

// ComponentDefineStatus defines the observed state of ComponentDefine
type ComponentDefineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ComponentDefine is the Schema for the componentdefines API
type ComponentDefine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentDefineSpec   `json:"spec,omitempty"`
	Status ComponentDefineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ComponentDefineList contains a list of ComponentDefine
type ComponentDefineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComponentDefine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComponentDefine{}, &ComponentDefineList{})
}
