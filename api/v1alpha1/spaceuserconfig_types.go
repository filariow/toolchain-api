package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SpaceVisibility string

const (
	SpaceVisibilityPrivate   SpaceVisibility = "private"
	SpaceVisibilityCommunity SpaceVisibility = "community"
)

// SpaceSpec defines the desired state of Space
// +k8s:openapi-gen=true
type SpaceUserConfigSpec struct {
	// +kubebuilder:validation:Enum:=private;community
	// +kubebuilder:default:=private
	Visibility SpaceVisibility `json:"visibility"`
}

// SpaceStatus defines the observed state of Space
// +k8s:openapi-gen=true
type SpaceUserConfigStatus struct {
	// Conditions is an array of current Space conditions
	// Supported condition types: ConditionReady
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Space is the Schema for the spaces API
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Visibility",type="string",JSONPath=`.spec.visibility`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Space"
type SpaceUserConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpaceUserConfigSpec   `json:"spec,omitempty"`
	Status SpaceUserConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpaceList contains a list of Space
type SpaceUserConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Space `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpaceUserConfig{}, &SpaceUserConfigList{})
}
