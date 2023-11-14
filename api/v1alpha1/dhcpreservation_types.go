/*
Copyright 2023 402-wtf.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DHCPReservationSpec defines the desired state of DHCPReservation
type DHCPReservationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// IPAddress is the reserved IP address
	IPAddress string `json:"ipAddress"`
}

// DHCPReservationStatus defines the observed state of DHCPReservation
type DHCPReservationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DHCPReservation is the Schema for the dhcpreservations API
type DHCPReservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DHCPReservationSpec   `json:"spec,omitempty"`
	Status DHCPReservationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DHCPReservationList contains a list of DHCPReservation
type DHCPReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DHCPReservation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DHCPReservation{}, &DHCPReservationList{})
}
