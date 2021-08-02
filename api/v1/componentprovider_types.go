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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type WorkLoadType string

const (
	Deployment  WorkLoadType = "deployment"
	StatefulSet WorkLoadType = "statefulSet"
	CloneSet    WorkLoadType = "cloneSet"
)

type Workload struct {
	//Type 'deployment' by default
	Type           WorkLoadType `json:"type,omitempty"`
	Replicas       *int32       `json:"replicas,omitempty" `
	UpdateStrategy string       `json:"updatestrategy"`
	// Template describes the pods that will be created.
	Template v1.PodTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template"`
}
type Service struct {
	Ports        []v1.ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port"`
	Type         v1.ServiceType   `json:"type"`
	ExternalName string           `json:"externalIPs,omitempty"`
}

// ComponentProviderSpec defines the desired state of ComponentProvider
type ComponentProviderSpec struct {
	Workload Workload          `json:"workload"`
	Service  Service           `json:"service"`
	Selector map[string]string `json:"selector,omitempty"`
}

// ComponentProviderStatus defines the observed state of ComponentProvider
type ComponentProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ComponentProvider is the Schema for the componentproviders API
type ComponentProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComponentProviderSpec   `json:"spec,omitempty"`
	Status ComponentProviderStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ComponentProviderList contains a list of ComponentProvider
type ComponentProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComponentProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComponentProvider{}, &ComponentProviderList{})
}
