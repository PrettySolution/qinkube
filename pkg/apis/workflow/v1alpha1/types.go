package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Kluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KlusterSpec `json:"spec"`
}

type KlusterSpec struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Version string `json:"version"`

	NodePools []NodePool
}

type NodePool struct {
	Size  string `json:"size"`
	Name  string `json:"mane"`
	Count int    `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KlusterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Kluster `json:"items"`
}
