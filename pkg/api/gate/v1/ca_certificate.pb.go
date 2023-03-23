// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/gate/v1/ca_certificate.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/xdmybl/gate-type/pkg/api/common/v1"
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

type CaCertificateSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonInfo              *v1.CommonInfo    `protobuf:"bytes,1,opt,name=commonInfo,proto3" json:"commonInfo,omitempty"`
	CertificateInfo         *v1.Certificate   `protobuf:"bytes,2,opt,name=certificateInfo,proto3" json:"certificateInfo,omitempty"`
	Ca                      string            `protobuf:"bytes,3,opt,name=ca,proto3" json:"ca,omitempty"`
	Crl                     string            `protobuf:"bytes,5,opt,name=crl,proto3" json:"crl,omitempty"`
	AllowExpiredCertificate bool              `protobuf:"varint,6,opt,name=allowExpiredCertificate,proto3" json:"allowExpiredCertificate,omitempty"`
	TlsParameters           *v1.TlsParameters `protobuf:"bytes,7,opt,name=tlsParameters,proto3" json:"tlsParameters,omitempty"`
	// Set Application Level Protocol Negotiation
	// If empty, defaults to ["h2", "http/1.1"].
	AlpnProtocols []string `protobuf:"bytes,8,rep,name=alpnProtocols,proto3" json:"alpnProtocols,omitempty"`
}

func (x *CaCertificateSpec) Reset() {
	*x = CaCertificateSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaCertificateSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaCertificateSpec) ProtoMessage() {}

func (x *CaCertificateSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaCertificateSpec.ProtoReflect.Descriptor instead.
func (*CaCertificateSpec) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CaCertificateSpec) GetCommonInfo() *v1.CommonInfo {
	if x != nil {
		return x.CommonInfo
	}
	return nil
}

func (x *CaCertificateSpec) GetCertificateInfo() *v1.Certificate {
	if x != nil {
		return x.CertificateInfo
	}
	return nil
}

func (x *CaCertificateSpec) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

func (x *CaCertificateSpec) GetCrl() string {
	if x != nil {
		return x.Crl
	}
	return ""
}

func (x *CaCertificateSpec) GetAllowExpiredCertificate() bool {
	if x != nil {
		return x.AllowExpiredCertificate
	}
	return false
}

func (x *CaCertificateSpec) GetTlsParameters() *v1.TlsParameters {
	if x != nil {
		return x.TlsParameters
	}
	return nil
}

func (x *CaCertificateSpec) GetAlpnProtocols() []string {
	if x != nil {
		return x.AlpnProtocols
	}
	return nil
}

var File_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x11, 0x43, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x35, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x17, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x74, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x0d, 0x61, 0x6c, 0x70, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x25, 0xfa, 0x42, 0x22,
	0x92, 0x01, 0x1f, 0x18, 0x01, 0x22, 0x19, 0x72, 0x17, 0x52, 0x0b, 0x68, 0x32, 0x2c, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x31, 0x2e, 0x31, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x31, 0x2e, 0x31,
	0x28, 0x01, 0x52, 0x0d, 0x61, 0x6c, 0x70, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x64, 0x6d, 0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_goTypes = []interface{}{
	(*CaCertificateSpec)(nil), // 0: gate.CaCertificateSpec
	(*v1.CommonInfo)(nil),     // 1: common.v1.CommonInfo
	(*v1.Certificate)(nil),    // 2: common.v1.Certificate
	(*v1.TlsParameters)(nil),  // 3: common.v1.TlsParameters
}
var file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_depIdxs = []int32{
	1, // 0: gate.CaCertificateSpec.commonInfo:type_name -> common.v1.CommonInfo
	2, // 1: gate.CaCertificateSpec.certificateInfo:type_name -> common.v1.Certificate
	3, // 2: gate.CaCertificateSpec.tlsParameters:type_name -> common.v1.TlsParameters
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_init() }
func file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaCertificateSpec); i {
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
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_depIdxs,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto = out.File
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_gate_v1_ca_certificate_proto_depIdxs = nil
}
