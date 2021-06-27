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

type Workload struct {
	//WorkloadType 'deployment' by default
	WorkloadType string `json:"type,omitempty"`
	Replicas     *int32 `json:"replicas,omitempty" `
	// Template describes the pods that will be created.
	Template v1.PodTemplateSpec `json:"template" protobuf:"bytes,3,opt,name=template"`
}
type Service struct {
	Ports       []v1.ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
	Type        v1.ServiceType   `json:"type"`
	ExternalIPs []string         `json:"externalIPs,omitempty" protobuf:"bytes,3,rep,name=externalIPs"`
}

// ComponentDefineSpec defines the desired state of ComponentDefine
type ComponentDefineSpec struct {
	Workload Workload          `json:"workload"`
	Selector map[string]string `json:"selector,omitempty" protobuf:"bytes,2,rep,name=selector"`
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
