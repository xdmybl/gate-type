// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/common/v1/headers.proto

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

// This plugin provides configuration options to append and remove headers from
// requests and responses
// HeaderManipulation can be specified on routes, virtual hosts, or weighted destinations
type HeaderManipulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies a list of HTTP headers that should be added to each request
	// handled by this route or virtual host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	RequestHeadersToAdd []*HeaderValueOption `protobuf:"bytes,1,rep,name=requestHeadersToAdd,proto3" json:"requestHeadersToAdd,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// handled by this route or virtual host.
	RequestHeadersToRemove []string `protobuf:"bytes,2,rep,name=requestHeadersToRemove,proto3" json:"requestHeadersToRemove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response
	// handled by this route or host. For more information, including
	// details on header value syntax, see the
	// [Envoy documentation](https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#config-http-conn-man-headers-custom-request-headers) .
	ResponseHeadersToAdd []*HeaderValueOption `protobuf:"bytes,3,rep,name=responseHeadersToAdd,proto3" json:"responseHeadersToAdd,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// handled by this route or virtual host.
	ResponseHeadersToRemove []string `protobuf:"bytes,4,rep,name=responseHeadersToRemove,proto3" json:"responseHeadersToRemove,omitempty"`
}

func (x *HeaderManipulation) Reset() {
	*x = HeaderManipulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderManipulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderManipulation) ProtoMessage() {}

func (x *HeaderManipulation) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderManipulation.ProtoReflect.Descriptor instead.
func (*HeaderManipulation) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderManipulation) GetRequestHeadersToAdd() []*HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *HeaderManipulation) GetRequestHeadersToRemove() []string {
	if x != nil {
		return x.RequestHeadersToRemove
	}
	return nil
}

func (x *HeaderManipulation) GetResponseHeadersToAdd() []*HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *HeaderManipulation) GetResponseHeadersToRemove() []string {
	if x != nil {
		return x.ResponseHeadersToRemove
	}
	return nil
}

// Header name/value pair plus option to control append behavior.
type HeaderValueOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name/value pair that this option applies to.
	Header *HeaderValue `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Should the value be appended? If true (default), the value is appended to
	// existing values.
	Append bool `protobuf:"varint,2,opt,name=append,proto3" json:"append,omitempty"`
}

func (x *HeaderValueOption) Reset() {
	*x = HeaderValueOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValueOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValueOption) ProtoMessage() {}

func (x *HeaderValueOption) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValueOption.ProtoReflect.Descriptor instead.
func (*HeaderValueOption) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderValueOption) GetHeader() *HeaderValue {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *HeaderValueOption) GetAppend() bool {
	if x != nil {
		return x.Append
	}
	return false
}

// Header name/value pair.
type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Header value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HeaderValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_github_com_xdmybl_gate_type_proto_common_v1_headers_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x02, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x08, 0x52, 0x13, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x12, 0x56, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x92, 0x01, 0x18, 0x10, 0x08, 0x18, 0x01, 0x22, 0x12, 0x72, 0x10,
	0x32, 0x0e, 0x5e, 0x5b, 0x21, 0x2d, 0x7e, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x31, 0x32, 0x7d, 0x24,
	0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x5a, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x08, 0x52, 0x14,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54,
	0x6f, 0x41, 0x64, 0x64, 0x12, 0x58, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b, 0x92, 0x01, 0x18, 0x10, 0x08, 0x18,
	0x01, 0x22, 0x12, 0x72, 0x10, 0x32, 0x0e, 0x5e, 0x5b, 0x21, 0x2d, 0x7e, 0x5d, 0x7b, 0x31, 0x2c,
	0x35, 0x31, 0x32, 0x7d, 0x24, 0x52, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x5b,
	0x0a, 0x11, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x22, 0x58, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x32, 0x0e,
	0x5e, 0x5b, 0x21, 0x2d, 0x7e, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x31, 0x32, 0x7d, 0x24, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x01, 0x28, 0x80, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d, 0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_goTypes = []interface{}{
	(*HeaderManipulation)(nil), // 0: common.v1.HeaderManipulation
	(*HeaderValueOption)(nil),  // 1: common.v1.HeaderValueOption
	(*HeaderValue)(nil),        // 2: common.v1.HeaderValue
}
var file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_depIdxs = []int32{
	1, // 0: common.v1.HeaderManipulation.requestHeadersToAdd:type_name -> common.v1.HeaderValueOption
	1, // 1: common.v1.HeaderManipulation.responseHeadersToAdd:type_name -> common.v1.HeaderValueOption
	2, // 2: common.v1.HeaderValueOption.header:type_name -> common.v1.HeaderValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_init() }
func file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_common_v1_headers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderManipulation); i {
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
		file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValueOption); i {
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
		file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_depIdxs,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_common_v1_headers_proto = out.File
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_headers_proto_depIdxs = nil
}
