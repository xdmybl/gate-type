// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/common/v1/tls.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type TlsParameters_ProtocolVersion int32

const (
	// Envoy will choose the optimal TLS version.
	TlsParameters_TLS_AUTO TlsParameters_ProtocolVersion = 0
	// TLS 1.0
	TlsParameters_TLSv1_0 TlsParameters_ProtocolVersion = 1
	// TLS 1.1
	TlsParameters_TLSv1_1 TlsParameters_ProtocolVersion = 2
	// TLS 1.2
	TlsParameters_TLSv1_2 TlsParameters_ProtocolVersion = 3
	// TLS 1.3
	TlsParameters_TLSv1_3 TlsParameters_ProtocolVersion = 4
)

// Enum value maps for TlsParameters_ProtocolVersion.
var (
	TlsParameters_ProtocolVersion_name = map[int32]string{
		0: "TLS_AUTO",
		1: "TLSv1_0",
		2: "TLSv1_1",
		3: "TLSv1_2",
		4: "TLSv1_3",
	}
	TlsParameters_ProtocolVersion_value = map[string]int32{
		"TLS_AUTO": 0,
		"TLSv1_0":  1,
		"TLSv1_1":  2,
		"TLSv1_2":  3,
		"TLSv1_3":  4,
	}
)

func (x TlsParameters_ProtocolVersion) Enum() *TlsParameters_ProtocolVersion {
	p := new(TlsParameters_ProtocolVersion)
	*p = x
	return p
}

func (x TlsParameters_ProtocolVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TlsParameters_ProtocolVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_enumTypes[0].Descriptor()
}

func (TlsParameters_ProtocolVersion) Type() protoreflect.EnumType {
	return &file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_enumTypes[0]
}

func (x TlsParameters_ProtocolVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TlsParameters_ProtocolVersion.Descriptor instead.
func (TlsParameters_ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescGZIP(), []int{0, 0}
}

// General TLS parameters. See the [envoy docs](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/auth/cert.proto#envoy-api-enum-auth-tlsparameters-tlsprotocol)
// for more information on the meaning of these values.
// nochange
type TlsParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinimumProtocolVersion TlsParameters_ProtocolVersion `protobuf:"varint,1,opt,name=minimumProtocolVersion,proto3,enum=common.v1.TlsParameters_ProtocolVersion" json:"minimumProtocolVersion,omitempty"`
	MaximumProtocolVersion TlsParameters_ProtocolVersion `protobuf:"varint,2,opt,name=maximumProtocolVersion,proto3,enum=common.v1.TlsParameters_ProtocolVersion" json:"maximumProtocolVersion,omitempty"`
	CipherSuites           []string                      `protobuf:"bytes,3,rep,name=cipherSuites,proto3" json:"cipherSuites,omitempty"`
	EcdhCurves             []string                      `protobuf:"bytes,4,rep,name=ecdhCurves,proto3" json:"ecdhCurves,omitempty"`
}

func (x *TlsParameters) Reset() {
	*x = TlsParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TlsParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TlsParameters) ProtoMessage() {}

func (x *TlsParameters) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TlsParameters.ProtoReflect.Descriptor instead.
func (*TlsParameters) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescGZIP(), []int{0}
}

func (x *TlsParameters) GetMinimumProtocolVersion() TlsParameters_ProtocolVersion {
	if x != nil {
		return x.MinimumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (x *TlsParameters) GetMaximumProtocolVersion() TlsParameters_ProtocolVersion {
	if x != nil {
		return x.MaximumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (x *TlsParameters) GetCipherSuites() []string {
	if x != nil {
		return x.CipherSuites
	}
	return nil
}

func (x *TlsParameters) GetEcdhCurves() []string {
	if x != nil {
		return x.EcdhCurves
	}
	return nil
}

var File_github_com_xdmybl_gate_type_proto_common_v1_tls_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x06, 0x0a, 0x0d,
	0x54, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x6a, 0x0a,
	0x16, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6c, 0x73, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x16, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x6a, 0x0a, 0x16, 0x6d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x16, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x87, 0x03, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0xe2, 0x02, 0xfa,
	0x42, 0xde, 0x02, 0x92, 0x01, 0xda, 0x02, 0x18, 0x01, 0x22, 0xd3, 0x02, 0x72, 0xd0, 0x02, 0x52,
	0x3d, 0x5b, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x45, 0x43, 0x44, 0x53, 0x41, 0x2d, 0x41, 0x45,
	0x53, 0x31, 0x32, 0x38, 0x2d, 0x47, 0x43, 0x4d, 0x2d, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x7c,
	0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x45, 0x43, 0x44, 0x53, 0x41, 0x2d, 0x43, 0x48, 0x41, 0x43,
	0x48, 0x41, 0x32, 0x30, 0x2d, 0x50, 0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x5d, 0x52, 0x39,
	0x5b, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x52, 0x53, 0x41, 0x2d, 0x41, 0x45, 0x53, 0x31, 0x32,
	0x38, 0x2d, 0x47, 0x43, 0x4d, 0x2d, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x7c, 0x45, 0x43, 0x44,
	0x48, 0x45, 0x2d, 0x52, 0x53, 0x41, 0x2d, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x2d,
	0x50, 0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x5d, 0x52, 0x16, 0x45, 0x43, 0x44, 0x48, 0x45,
	0x2d, 0x45, 0x43, 0x44, 0x53, 0x41, 0x2d, 0x41, 0x45, 0x53, 0x31, 0x32, 0x38, 0x2d, 0x53, 0x48,
	0x41, 0x52, 0x14, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x52, 0x53, 0x41, 0x2d, 0x41, 0x45, 0x53,
	0x31, 0x32, 0x38, 0x2d, 0x53, 0x48, 0x41, 0x52, 0x11, 0x41, 0x45, 0x53, 0x31, 0x32, 0x38, 0x2d,
	0x47, 0x43, 0x4d, 0x2d, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x52, 0x0a, 0x41, 0x45, 0x53, 0x31,
	0x32, 0x38, 0x2d, 0x53, 0x48, 0x41, 0x52, 0x1d, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x45, 0x43,
	0x44, 0x53, 0x41, 0x2d, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x47, 0x43, 0x4d, 0x2d, 0x53,
	0x48, 0x41, 0x33, 0x38, 0x34, 0x52, 0x1b, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x52, 0x53, 0x41,
	0x2d, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x47, 0x43, 0x4d, 0x2d, 0x53, 0x48, 0x41, 0x33,
	0x38, 0x34, 0x52, 0x16, 0x45, 0x43, 0x44, 0x48, 0x45, 0x2d, 0x45, 0x43, 0x44, 0x53, 0x41, 0x2d,
	0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x53, 0x48, 0x41, 0x52, 0x14, 0x45, 0x43, 0x44, 0x48,
	0x45, 0x2d, 0x52, 0x53, 0x41, 0x2d, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x53, 0x48, 0x41,
	0x52, 0x11, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x47, 0x43, 0x4d, 0x2d, 0x53, 0x48, 0x41,
	0x33, 0x38, 0x34, 0x52, 0x0a, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x2d, 0x53, 0x48, 0x41, 0x28,
	0x01, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x12,
	0x3d, 0x0a, 0x0a, 0x65, 0x63, 0x64, 0x68, 0x43, 0x75, 0x72, 0x76, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x1d, 0xfa, 0x42, 0x1a, 0x92, 0x01, 0x17, 0x18, 0x01, 0x22, 0x11, 0x72,
	0x0f, 0x52, 0x06, 0x58, 0x32, 0x35, 0x35, 0x31, 0x39, 0x52, 0x05, 0x50, 0x2d, 0x32, 0x35, 0x36,
	0x28, 0x01, 0x52, 0x0a, 0x65, 0x63, 0x64, 0x68, 0x43, 0x75, 0x72, 0x76, 0x65, 0x73, 0x22, 0x53,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4c, 0x53, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x4c, 0x53, 0x76, 0x31, 0x5f, 0x30, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x4c, 0x53, 0x76, 0x31, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x4c, 0x53,
	0x76, 0x31, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x4c, 0x53, 0x76, 0x31, 0x5f,
	0x33, 0x10, 0x04, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x64, 0x6d, 0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_goTypes = []interface{}{
	(TlsParameters_ProtocolVersion)(0), // 0: common.v1.TlsParameters.ProtocolVersion
	(*TlsParameters)(nil),              // 1: common.v1.TlsParameters
}
var file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_depIdxs = []int32{
	0, // 0: common.v1.TlsParameters.minimumProtocolVersion:type_name -> common.v1.TlsParameters.ProtocolVersion
	0, // 1: common.v1.TlsParameters.maximumProtocolVersion:type_name -> common.v1.TlsParameters.ProtocolVersion
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_init() }
func file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_common_v1_tls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TlsParameters); i {
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
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_depIdxs,
		EnumInfos:         file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_enumTypes,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_common_v1_tls_proto = out.File
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_tls_proto_depIdxs = nil
}