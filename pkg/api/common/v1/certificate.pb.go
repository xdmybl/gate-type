// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/common/v1/certificate.proto

package v1

import (
	v1 "github.com/xdmybl/gate-type/pkg/api/core/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 证书文件bytes
	Raw                    []byte           `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Data                   *CertificateData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	SignatureAlgorithm     string           `protobuf:"bytes,3,opt,name=signatureAlgorithm,proto3" json:"signatureAlgorithm,omitempty"`
	SignatureHashAlgorithm string           `protobuf:"bytes,4,opt,name=signatureHashAlgorithm,proto3" json:"signatureHashAlgorithm,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *Certificate) GetData() *CertificateData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Certificate) GetSignatureAlgorithm() string {
	if x != nil {
		return x.SignatureAlgorithm
	}
	return ""
}

func (x *Certificate) GetSignatureHashAlgorithm() string {
	if x != nil {
		return x.SignatureHashAlgorithm
	}
	return ""
}

type CertificateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	SerialNumber           string `protobuf:"bytes,2,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	SignatureAlgorithm     string `protobuf:"bytes,3,opt,name=signatureAlgorithm,proto3" json:"signatureAlgorithm,omitempty"`
	SignatureHashAlgorithm string `protobuf:"bytes,4,opt,name=signatureHashAlgorithm,proto3" json:"signatureHashAlgorithm,omitempty"`
	// 发行者 发行证书的证书认证中心 (CA) 的标识信息
	Issuer   []*CertSubject `protobuf:"bytes,5,rep,name=issuer,proto3" json:"issuer,omitempty"`
	Validity *Validity      `protobuf:"bytes,6,opt,name=validity,proto3" json:"validity,omitempty"`
	// subject
	Subject              []*CertSubject        `protobuf:"bytes,7,rep,name=subject,proto3" json:"subject,omitempty"`
	SubjectPublicKeyInfo *SubjectPublicKeyInfo `protobuf:"bytes,8,opt,name=subjectPublicKeyInfo,proto3" json:"subjectPublicKeyInfo,omitempty"`
	Extensions           *Extensions           `protobuf:"bytes,9,opt,name=extensions,proto3" json:"extensions,omitempty"`
}

func (x *CertificateData) Reset() {
	*x = CertificateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateData) ProtoMessage() {}

func (x *CertificateData) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateData.ProtoReflect.Descriptor instead.
func (*CertificateData) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CertificateData) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *CertificateData) GetSignatureAlgorithm() string {
	if x != nil {
		return x.SignatureAlgorithm
	}
	return ""
}

func (x *CertificateData) GetSignatureHashAlgorithm() string {
	if x != nil {
		return x.SignatureHashAlgorithm
	}
	return ""
}

func (x *CertificateData) GetIssuer() []*CertSubject {
	if x != nil {
		return x.Issuer
	}
	return nil
}

func (x *CertificateData) GetValidity() *Validity {
	if x != nil {
		return x.Validity
	}
	return nil
}

func (x *CertificateData) GetSubject() []*CertSubject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *CertificateData) GetSubjectPublicKeyInfo() *SubjectPublicKeyInfo {
	if x != nil {
		return x.SubjectPublicKeyInfo
	}
	return nil
}

func (x *CertificateData) GetExtensions() *Extensions {
	if x != nil {
		return x.Extensions
	}
	return nil
}

type CertSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CN
	CommonName string `protobuf:"bytes,1,opt,name=commonName,proto3" json:"commonName,omitempty"`
	// OU
	OrganizationUint string `protobuf:"bytes,2,opt,name=organizationUint,proto3" json:"organizationUint,omitempty"`
	// O
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
	// C
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	// ST
	StateOrProvince string `protobuf:"bytes,5,opt,name=stateOrProvince,proto3" json:"stateOrProvince,omitempty"`
	// L
	Locality string `protobuf:"bytes,6,opt,name=locality,proto3" json:"locality,omitempty"`
}

func (x *CertSubject) Reset() {
	*x = CertSubject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertSubject) ProtoMessage() {}

func (x *CertSubject) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertSubject.ProtoReflect.Descriptor instead.
func (*CertSubject) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *CertSubject) GetCommonName() string {
	if x != nil {
		return x.CommonName
	}
	return ""
}

func (x *CertSubject) GetOrganizationUint() string {
	if x != nil {
		return x.OrganizationUint
	}
	return ""
}

func (x *CertSubject) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *CertSubject) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CertSubject) GetStateOrProvince() string {
	if x != nil {
		return x.StateOrProvince
	}
	return ""
}

func (x *CertSubject) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

type Validity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotBefore *v1.Timestamp `protobuf:"bytes,1,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter  *v1.Timestamp `protobuf:"bytes,2,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
}

func (x *Validity) Reset() {
	*x = Validity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validity) ProtoMessage() {}

func (x *Validity) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validity.ProtoReflect.Descriptor instead.
func (*Validity) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{3}
}

func (x *Validity) GetNotBefore() *v1.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *Validity) GetNotAfter() *v1.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

type SubjectPublicKeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm           string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Length              string `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	Thumbprint          string `protobuf:"bytes,3,opt,name=thumbprint,proto3" json:"thumbprint,omitempty"`
	ThumbprintAlgorithm string `protobuf:"bytes,4,opt,name=thumbprintAlgorithm,proto3" json:"thumbprintAlgorithm,omitempty"`
}

func (x *SubjectPublicKeyInfo) Reset() {
	*x = SubjectPublicKeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectPublicKeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectPublicKeyInfo) ProtoMessage() {}

func (x *SubjectPublicKeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectPublicKeyInfo.ProtoReflect.Descriptor instead.
func (*SubjectPublicKeyInfo) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{4}
}

func (x *SubjectPublicKeyInfo) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *SubjectPublicKeyInfo) GetLength() string {
	if x != nil {
		return x.Length
	}
	return ""
}

func (x *SubjectPublicKeyInfo) GetThumbprint() string {
	if x != nil {
		return x.Thumbprint
	}
	return ""
}

func (x *SubjectPublicKeyInfo) GetThumbprintAlgorithm() string {
	if x != nil {
		return x.ThumbprintAlgorithm
	}
	return ""
}

type Extensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectAlternativeNames []string `protobuf:"bytes,4,rep,name=subjectAlternativeNames,proto3" json:"subjectAlternativeNames,omitempty"`
}

func (x *Extensions) Reset() {
	*x = Extensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensions) ProtoMessage() {}

func (x *Extensions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensions.ProtoReflect.Descriptor instead.
func (*Extensions) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP(), []int{5}
}

func (x *Extensions) GetSubjectAlternativeNames() []string {
	if x != nil {
		return x.SubjectAlternativeNames
	}
	return nil
}

var File_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0xd6,
	0x03, 0x0a, 0x0f, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x36, 0x0a, 0x16, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x2e, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x53, 0x0a, 0x14, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x35, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74,
	0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x69, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x6c, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6e, 0x6f, 0x74,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0x46, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x6c,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_goTypes = []interface{}{
	(*Certificate)(nil),          // 0: common.v1.Certificate
	(*CertificateData)(nil),      // 1: common.v1.CertificateData
	(*CertSubject)(nil),          // 2: common.v1.CertSubject
	(*Validity)(nil),             // 3: common.v1.Validity
	(*SubjectPublicKeyInfo)(nil), // 4: common.v1.SubjectPublicKeyInfo
	(*Extensions)(nil),           // 5: common.v1.Extensions
	(*v1.Timestamp)(nil),         // 6: core.v1.Timestamp
}
var file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_depIdxs = []int32{
	1, // 0: common.v1.Certificate.data:type_name -> common.v1.CertificateData
	2, // 1: common.v1.CertificateData.issuer:type_name -> common.v1.CertSubject
	3, // 2: common.v1.CertificateData.validity:type_name -> common.v1.Validity
	2, // 3: common.v1.CertificateData.subject:type_name -> common.v1.CertSubject
	4, // 4: common.v1.CertificateData.subjectPublicKeyInfo:type_name -> common.v1.SubjectPublicKeyInfo
	5, // 5: common.v1.CertificateData.extensions:type_name -> common.v1.Extensions
	6, // 6: common.v1.Validity.notBefore:type_name -> core.v1.Timestamp
	6, // 7: common.v1.Validity.notAfter:type_name -> core.v1.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_init() }
func file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertSubject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectPublicKeyInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_depIdxs,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto = out.File
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_certificate_proto_depIdxs = nil
}
