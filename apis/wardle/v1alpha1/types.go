

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// +genclient=true
// +genapi=true
// +resource=flunders
// +k8s:openapi-gen=true
type Flunder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlunderSpec   `json:"spec,omitempty"`
	Status FlunderStatus `json:"status,omitempty"`
}

type FlunderSpec struct {
}

type FlunderStatus struct {
}
