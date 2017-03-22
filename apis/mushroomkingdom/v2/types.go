package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient=true
// +genapi=true
// +resource=peachescastles
// +k8s:openapi-gen=true
type PeachesCastle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeachesCastleSpec   `json:"spec,omitempty"`
	Status PeachesCastleStatus `json:"status,omitempty"`
}

type PeachesCastleSpec struct {
	Mushrooms int `json:"mushrooms,omitempty"`
}

type PeachesCastleStatus struct {
}

// +genapi=subresource
type ScaleCastle struct {
	metav1.TypeMeta
	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta
}
