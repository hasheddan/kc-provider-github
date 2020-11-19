/*
Copyright 2020 The Crossplane Authors.

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

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// TeamParameters are the configurable fields of a Team.
type TeamParameters struct {
	Org         string  `json:"org"`
	Description *string `json:"description,omitempty"`
	// +kubebuilder:validation:Enum=secret;closed
	Privacy *string `json:"privacy,omitempty"`
}

// TeamObservation are the observable fields of a Team.
type TeamObservation struct {
	NodeID string `json:"nodeId,omitempty"`
}

// A TeamSpec defines the desired state of a Team.
type TeamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  TeamParameters `json:"forProvider"`
}

// A TeamStatus represents the observed state of a Team.
type TeamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     TeamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Team is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TeamSpec   `json:"spec"`
	Status TeamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamList contains a list of Team
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// MembershipParameters are the configurable fields of a Membership.
type MembershipParameters struct {
	Org  string `json:"org"`
	User string `json:"user"`

	Team         string                     `json:"team,omitempty"`
	TeamRef      *runtimev1alpha1.Reference `json:"teamRef,omitempty"`
	TeamSelector *runtimev1alpha1.Selector  `json:"teamSelector,omitempty"`
}

// MembershipObservation are the observable fields of a Membership.
type MembershipObservation struct {
	State string `json:"state,omitempty"`
}

// A MembershipSpec defines the desired state of a Membership.
type MembershipSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MembershipParameters `json:"forProvider"`
}

// A MembershipStatus represents the observed state of a Membership.
type MembershipStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Membership is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type Membership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MembershipSpec   `json:"spec"`
	Status MembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MembershipList contains a list of Membership
type MembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Membership `json:"items"`
}
