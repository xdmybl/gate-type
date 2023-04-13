// Code generated by engine gate build no edit

// Definitions for the Kubernetes types
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for CaCertificate
var CaCertificateGVK = schema.GroupVersionKind{
	Group:   "gate.xdmybl.io",
	Version: "v1",
	Kind:    "CaCertificate",
}

// CaCertificate is the Schema for the caCertificate API
type CaCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CaCertificateSpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CaCertificate) GVK() schema.GroupVersionKind {
	return CaCertificateGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CaCertificateList contains a list of CaCertificate
type CaCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaCertificate `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for Certificate
var CertificateGVK = schema.GroupVersionKind{
	Group:   "gate.xdmybl.io",
	Version: "v1",
	Kind:    "Certificate",
}

// Certificate is the Schema for the certificate API
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CertificateSpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Certificate) GVK() schema.GroupVersionKind {
	return CertificateGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateList contains a list of Certificate
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for Upstream
var UpstreamGVK = schema.GroupVersionKind{
	Group:   "gate.xdmybl.io",
	Version: "v1",
	Kind:    "Upstream",
}

// Upstream is the Schema for the upstream API
type Upstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec UpstreamSpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Upstream) GVK() schema.GroupVersionKind {
	return UpstreamGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpstreamList contains a list of Upstream
type UpstreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Upstream `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for Gateway
var GatewayGVK = schema.GroupVersionKind{
	Group:   "gate.xdmybl.io",
	Version: "v1",
	Kind:    "Gateway",
}

// Gateway is the Schema for the gateway API
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GatewaySpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Gateway) GVK() schema.GroupVersionKind {
	return GatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayList contains a list of Gateway
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +genclient:noStatus

// GroupVersionKind for Filter
var FilterGVK = schema.GroupVersionKind{
	Group:   "gate.xdmybl.io",
	Version: "v1",
	Kind:    "Filter",
}

// Filter is the Schema for the filter API
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FilterSpec `json:"spec,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Filter) GVK() schema.GroupVersionKind {
	return FilterGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FilterList contains a list of Filter
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Filter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CaCertificate{}, &CaCertificateList{})
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
	SchemeBuilder.Register(&Upstream{}, &UpstreamList{})
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
	SchemeBuilder.Register(&Filter{}, &FilterList{})
}