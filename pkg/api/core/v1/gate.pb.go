// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: github.com/xdmybl/gate-type/proto/core/v1/gate.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any1 "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes[0].Descriptor()
}

func (CommonInfo_Source) Type() protoreflect.EnumType {
	return &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes[0]
}

func (x CommonInfo_Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonInfo_Source.Descriptor instead.
func (CommonInfo_Source) EnumDescriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{5, 0}
}

type StatefulSession_Type int32

const (
	StatefulSession_UNKNOWN StatefulSession_Type = 0
	StatefulSession_COOKIE  StatefulSession_Type = 1
	StatefulSession_SESSION StatefulSession_Type = 2
)

// Enum value maps for StatefulSession_Type.
var (
	StatefulSession_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "COOKIE",
		2: "SESSION",
	}
	StatefulSession_Type_value = map[string]int32{
		"UNKNOWN": 0,
		"COOKIE":  1,
		"SESSION": 2,
	}
)

func (x StatefulSession_Type) Enum() *StatefulSession_Type {
	p := new(StatefulSession_Type)
	*p = x
	return p
}

func (x StatefulSession_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatefulSession_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes[1].Descriptor()
}

func (StatefulSession_Type) Type() protoreflect.EnumType {
	return &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes[1]
}

func (x StatefulSession_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatefulSession_Type.Descriptor instead.
func (StatefulSession_Type) EnumDescriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{6, 0}
}

type Int64Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start of the range (inclusive)
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Int64Range) Reset() {
	*x = Int64Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Range) ProtoMessage() {}

func (x *Int64Range) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Range.ProtoReflect.Descriptor instead.
func (*Int64Range) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{0}
}

func (x *Int64Range) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Int64Range) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

// Specifies the int32 start and end of the range using half-open interval semantics [start,
// end).
type Int32Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start of the range (inclusive)
	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Int32Range) Reset() {
	*x = Int32Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Range) ProtoMessage() {}

func (x *Int32Range) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Range.ProtoReflect.Descriptor instead.
func (*Int32Range) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{1}
}

func (x *Int32Range) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Int32Range) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

// Specifies the double start and end of the range using half-open interval semantics [start,
// end).
type DoubleRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start of the range (inclusive)
	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	// end of the range (exclusive)
	End float64 `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DoubleRange) Reset() {
	*x = DoubleRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleRange) ProtoMessage() {}

func (x *DoubleRange) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleRange.ProtoReflect.Descriptor instead.
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{2}
}

func (x *DoubleRange) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DoubleRange) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{3}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
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
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRef) ProtoMessage() {}

func (x *ResourceRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[4]
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
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{4}
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
	Source          CommonInfo_Source `protobuf:"varint,3,opt,name=source,proto3,enum=core.v1.CommonInfo_Source" json:"source,omitempty"`
	UpdatedDatetime *Timestamp        `protobuf:"bytes,4,opt,name=updatedDatetime,proto3" json:"updatedDatetime,omitempty"`
}

func (x *CommonInfo) Reset() {
	*x = CommonInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonInfo) ProtoMessage() {}

func (x *CommonInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[5]
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
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{5}
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

func (x *CommonInfo) GetUpdatedDatetime() *Timestamp {
	if x != nil {
		return x.UpdatedDatetime
	}
	return nil
}

type StatefulSession struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type StatefulSession_Type `protobuf:"varint,1,opt,name=type,proto3,enum=core.v1.StatefulSession_Type" json:"type,omitempty"`
	Name string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// only for COOKIE
	Ttl *Timestamp `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// only for COOKIE
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *StatefulSession) Reset() {
	*x = StatefulSession{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatefulSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSession) ProtoMessage() {}

func (x *StatefulSession) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSession.ProtoReflect.Descriptor instead.
func (*StatefulSession) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{6}
}

func (x *StatefulSession) GetType() StatefulSession_Type {
	if x != nil {
		return x.Type
	}
	return StatefulSession_UNKNOWN
}

func (x *StatefulSession) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatefulSession) GetTtl() *Timestamp {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *StatefulSession) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address             string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	LoadBalancingWeight uint32 `protobuf:"varint,3,opt,name=loadBalancingWeight,proto3" json:"loadBalancingWeight,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{7}
}

func (x *Endpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Endpoint) GetLoadBalancingWeight() uint32 {
	if x != nil {
		return x.LoadBalancingWeight
	}
	return 0
}

// https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/base.proto#config-core-v3-metadata
type GateMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilterMetadata      map[string]*_struct.Struct `protobuf:"bytes,1,rep,name=filter_metadata,json=filterMetadata,proto3" json:"filter_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TypedFilterMetadata map[string]*any1.Any       `protobuf:"bytes,2,rep,name=typed_filter_metadata,json=typedFilterMetadata,proto3" json:"typed_filter_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GateMetadata) Reset() {
	*x = GateMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateMetadata) ProtoMessage() {}

func (x *GateMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateMetadata.ProtoReflect.Descriptor instead.
func (*GateMetadata) Descriptor() ([]byte, []int) {
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP(), []int{8}
}

func (x *GateMetadata) GetFilterMetadata() map[string]*_struct.Struct {
	if x != nil {
		return x.FilterMetadata
	}
	return nil
}

func (x *GateMetadata) GetTypedFilterMetadata() map[string]*any1.Any {
	if x != nil {
		return x.TypedFilterMetadata
	}
	return nil
}

var File_github_com_xdmybl_gate_type_proto_core_v1_gate_proto protoreflect.FileDescriptor

var file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x64, 0x6d,
	0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x34, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x35, 0x0a,
	0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f,
	0x73, 0x22, 0x9b, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x66, 0x12, 0x68, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0xfa, 0x42, 0x47, 0x72, 0x45, 0x18, 0x3f, 0x32, 0x41, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x28, 0x5c, 0x2e,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x29, 0x2a, 0x24,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0xfa, 0x42, 0x47, 0x72, 0x45,
	0x18, 0x3f, 0x32, 0x41, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x3f, 0x28, 0x5c, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x5b, 0x2d,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d,
	0x29, 0x3f, 0x29, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22,
	0x87, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0x18, 0x80, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0x80, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x22, 0xe2, 0x01, 0x0a, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x66, 0x75, 0x6c, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x20,
	0x01, 0x28, 0x80, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x20, 0x01, 0x28, 0x80, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x8c,
	0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xa8, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa,
	0x42, 0x08, 0x2a, 0x06, 0x18, 0xff, 0xff, 0x03, 0x28, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x3b, 0x0a, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x2a, 0x04, 0x28, 0x01, 0x40, 0x01, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x80, 0x03,
	0x0a, 0x0c, 0x47, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x52,
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x62, 0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x13, 0x74, 0x79, 0x70, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x5a, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x5c, 0x0a, 0x18, 0x54, 0x79, 0x70, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x64, 0x6d, 0x79, 0x62, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescOnce sync.Once
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescData = file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDesc
)

func file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescGZIP() []byte {
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescOnce.Do(func() {
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescData)
	})
	return file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDescData
}

var file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_goTypes = []interface{}{
	(CommonInfo_Source)(0),    // 0: core.v1.CommonInfo.Source
	(StatefulSession_Type)(0), // 1: core.v1.StatefulSession.Type
	(*Int64Range)(nil),        // 2: core.v1.Int64Range
	(*Int32Range)(nil),        // 3: core.v1.Int32Range
	(*DoubleRange)(nil),       // 4: core.v1.DoubleRange
	(*Timestamp)(nil),         // 5: core.v1.Timestamp
	(*ResourceRef)(nil),       // 6: core.v1.ResourceRef
	(*CommonInfo)(nil),        // 7: core.v1.CommonInfo
	(*StatefulSession)(nil),   // 8: core.v1.StatefulSession
	(*Endpoint)(nil),          // 9: core.v1.Endpoint
	(*GateMetadata)(nil),      // 10: core.v1.GateMetadata
	nil,                       // 11: core.v1.GateMetadata.FilterMetadataEntry
	nil,                       // 12: core.v1.GateMetadata.TypedFilterMetadataEntry
	(*_struct.Struct)(nil),    // 13: google.protobuf.Struct
	(*any1.Any)(nil),          // 14: google.protobuf.Any
}
var file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_depIdxs = []int32{
	0,  // 0: core.v1.CommonInfo.source:type_name -> core.v1.CommonInfo.Source
	5,  // 1: core.v1.CommonInfo.updatedDatetime:type_name -> core.v1.Timestamp
	1,  // 2: core.v1.StatefulSession.type:type_name -> core.v1.StatefulSession.Type
	5,  // 3: core.v1.StatefulSession.ttl:type_name -> core.v1.Timestamp
	11, // 4: core.v1.GateMetadata.filter_metadata:type_name -> core.v1.GateMetadata.FilterMetadataEntry
	12, // 5: core.v1.GateMetadata.typed_filter_metadata:type_name -> core.v1.GateMetadata.TypedFilterMetadataEntry
	13, // 6: core.v1.GateMetadata.FilterMetadataEntry.value:type_name -> google.protobuf.Struct
	14, // 7: core.v1.GateMetadata.TypedFilterMetadataEntry.value:type_name -> google.protobuf.Any
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_init() }
func file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_init() {
	if File_github_com_xdmybl_gate_type_proto_core_v1_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Range); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Range); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleRange); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatefulSession); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateMetadata); i {
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
			RawDescriptor: file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_goTypes,
		DependencyIndexes: file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_depIdxs,
		EnumInfos:         file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_enumTypes,
		MessageInfos:      file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_msgTypes,
	}.Build()
	File_github_com_xdmybl_gate_type_proto_core_v1_gate_proto = out.File
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_rawDesc = nil
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_goTypes = nil
	file_github_com_xdmybl_gate_type_proto_core_v1_gate_proto_depIdxs = nil
}
