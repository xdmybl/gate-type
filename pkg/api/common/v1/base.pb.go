// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/common/v1/base.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CommonInfo_Source int32

const (
	CommonInfo_CustomDefinition CommonInfo_Source = 0
	CommonInfo_AdminDefinition  CommonInfo_Source = 1
)

// Enum value maps for CommonInfo_Source.
var (
	CommonInfo_Source_name = map[int32]string{
		0: "CustomDefinition",
		1: "AdminDefinition",
	}
	CommonInfo_Source_value = map[string]int32{
		"CustomDefinition": 0,
		"AdminDefinition":  1,
	}
)

func (x CommonInfo_Source) Enum() *CommonInfo_Source {
	p := new(CommonInfo_Source)
	*p = x
	return p
}

func (x CommonInfo_Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonInfo_Source) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_enumTypes[0].Descriptor()
}

func (CommonInfo_Source) Type() protoreflect.EnumType {
	return &file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_enumTypes[0]
}

func (x CommonInfo_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonInfo_Source.Descriptor instead.
func (CommonInfo_Source) EnumDescriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescGZIP(), []int{1, 0}
}

type ResourceRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defaults to `default`
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Group     string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Version   string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Kind      string `protobuf:"bytes,5,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *ResourceRef) Reset() {
	*x = ResourceRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRef) ProtoMessage() {}

func (x *ResourceRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRef.ProtoReflect.Descriptor instead.
func (*ResourceRef) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ResourceRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceRef) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *ResourceRef) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ResourceRef) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type CommonInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description     string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Source          CommonInfo_Source `protobuf:"varint,3,opt,name=source,proto3,enum=common.v1.CommonInfo_Source" json:"source,omitempty"`
	Disabled        bool              `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
	UpdatedDatetime *v1.Timestamp     `protobuf:"bytes,5,opt,name=updatedDatetime,proto3" json:"updatedDatetime,omitempty"`
}

func (x *CommonInfo) Reset() {
	*x = CommonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonInfo) ProtoMessage() {}

func (x *CommonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonInfo.ProtoReflect.Descriptor instead.
func (*CommonInfo) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescGZIP(), []int{1}
}

func (x *CommonInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommonInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommonInfo) GetSource() CommonInfo_Source {
	if x != nil {
		return x.Source
	}
	return CommonInfo_CustomDefinition
}

func (x *CommonInfo) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *CommonInfo) GetUpdatedDatetime() *v1.Timestamp {
	if x != nil {
		return x.UpdatedDatetime
	}
	return nil
}

var File_github_com_xdmybl_gate_type_proto_common_v1_base_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x68, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0xfa, 0x42, 0x47, 0x72,
	0x45, 0x18, 0x3f, 0x32, 0x41, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b,
	0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x29, 0x3f, 0x28, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b,
	0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x4a, 0xfa, 0x42, 0x47, 0x72, 0x45, 0x18, 0x3f, 0x32, 0x41, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x28, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xa5, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80,
	0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x0f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d, 0x79,
	0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_goTypes = []interface{}{
	(CommonInfo_Source)(0), // 0: common.v1.CommonInfo.Source
	(*ResourceRef)(nil),    // 1: common.v1.ResourceRef
	(*CommonInfo)(nil),     // 2: common.v1.CommonInfo
	(*v1.Timestamp)(nil),   // 3: core.v1.Timestamp
}
var file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_depIdxs = []int32{
	0, // 0: common.v1.CommonInfo.source:type_name -> common.v1.CommonInfo.Source
	3, // 1: common.v1.CommonInfo.updatedDatetime:type_name -> core.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_init() }
func file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_common_v1_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRef); i {
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
		file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonInfo); i {
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
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_depIdxs,
		EnumInfos:         file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_enumTypes,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_common_v1_base_proto = out.File
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_common_v1_base_proto_depIdxs = nil
}
