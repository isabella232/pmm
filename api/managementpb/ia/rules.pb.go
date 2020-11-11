// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/ia/rules.proto

package iav1beta1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ParamUnit Integrated Alerting rule parameter unit.
type ParamUnit int32

const (
	ParamUnit_PARAM_UNIT_INVALID ParamUnit = 0
	ParamUnit_PERCENTAGE         ParamUnit = 1
)

// Enum value maps for ParamUnit.
var (
	ParamUnit_name = map[int32]string{
		0: "PARAM_UNIT_INVALID",
		1: "PERCENTAGE",
	}
	ParamUnit_value = map[string]int32{
		"PARAM_UNIT_INVALID": 0,
		"PERCENTAGE":         1,
	}
)

func (x ParamUnit) Enum() *ParamUnit {
	p := new(ParamUnit)
	*p = x
	return p
}

func (x ParamUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_ia_rules_proto_enumTypes[0].Descriptor()
}

func (ParamUnit) Type() protoreflect.EnumType {
	return &file_managementpb_ia_rules_proto_enumTypes[0]
}

func (x ParamUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamUnit.Descriptor instead.
func (ParamUnit) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{0}
}

// ParamType Integrated Alerting rule parameter type.
type ParamType int32

const (
	ParamType_PARAM_TYPE_INVALID ParamType = 0
	ParamType_FLOAT              ParamType = 1
)

// Enum value maps for ParamType.
var (
	ParamType_name = map[int32]string{
		0: "PARAM_TYPE_INVALID",
		1: "FLOAT",
	}
	ParamType_value = map[string]int32{
		"PARAM_TYPE_INVALID": 0,
		"FLOAT":              1,
	}
)

func (x ParamType) Enum() *ParamType {
	p := new(ParamType)
	*p = x
	return p
}

func (x ParamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParamType) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_ia_rules_proto_enumTypes[1].Descriptor()
}

func (ParamType) Type() protoreflect.EnumType {
	return &file_managementpb_ia_rules_proto_enumTypes[1]
}

func (x ParamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParamType.Descriptor instead.
func (ParamType) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{1}
}

// Rule represents Integrated Alerting rule.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rule name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Rules status: enabled or disabled.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Rule description.
	Help string `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`
	// Rule parameters.
	Params []*Rule_Param `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
	// Rule default duration.
	DefaultFor *duration.Duration `protobuf:"bytes,5,opt,name=default_for,json=defaultFor,proto3" json:"default_for,omitempty"`
	// Rule set duration.
	For *duration.Duration `protobuf:"bytes,6,opt,name=for,proto3" json:"for,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Rule) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Rule) GetParams() []*Rule_Param {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Rule) GetDefaultFor() *duration.Duration {
	if x != nil {
		return x.DefaultFor
	}
	return nil
}

func (x *Rule) GetFor() *duration.Duration {
	if x != nil {
		return x.For
	}
	return nil
}

type ListAlertingRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, Integrated Alerting rule files will be re-read from disk.
	Reload bool `protobuf:"varint,1,opt,name=reload,proto3" json:"reload,omitempty"`
}

func (x *ListAlertingRulesRequest) Reset() {
	*x = ListAlertingRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertingRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertingRulesRequest) ProtoMessage() {}

func (x *ListAlertingRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertingRulesRequest.ProtoReflect.Descriptor instead.
func (*ListAlertingRulesRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{1}
}

func (x *ListAlertingRulesRequest) GetReload() bool {
	if x != nil {
		return x.Reload
	}
	return false
}

type ListAlertingRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Integrated Alerting rules.
	Rules []*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ListAlertingRulesResponse) Reset() {
	*x = ListAlertingRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlertingRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlertingRulesResponse) ProtoMessage() {}

func (x *ListAlertingRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlertingRulesResponse.ProtoReflect.Descriptor instead.
func (*ListAlertingRulesResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{2}
}

func (x *ListAlertingRulesResponse) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type ChangeAlertingRulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rule name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// New rule status.
	Enabled1 BooleanFlag `protobuf:"varint,2,opt,name=enabled1,proto3,enum=ia.v1beta1.BooleanFlag" json:"enabled1,omitempty"`
	// New rule status.
	Enabled2 *wrappers.BoolValue `protobuf:"bytes,20,opt,name=enabled2,proto3" json:"enabled2,omitempty"`
	// Parameters to change. Missing parameters will not be changed.
	Params []*ChangeAlertingRulesRequest_Param `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	// Rule set duration. Zero value will not change it.
	// FIXME ?
	For *duration.Duration `protobuf:"bytes,4,opt,name=for,proto3" json:"for,omitempty"`
	// Remove set duration (reset to default).
	RemoveFor bool `protobuf:"varint,5,opt,name=remove_for,json=removeFor,proto3" json:"remove_for,omitempty"` // FIXME API for labels? separete API for severity?
}

func (x *ChangeAlertingRulesRequest) Reset() {
	*x = ChangeAlertingRulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAlertingRulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAlertingRulesRequest) ProtoMessage() {}

func (x *ChangeAlertingRulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAlertingRulesRequest.ProtoReflect.Descriptor instead.
func (*ChangeAlertingRulesRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeAlertingRulesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeAlertingRulesRequest) GetEnabled1() BooleanFlag {
	if x != nil {
		return x.Enabled1
	}
	return BooleanFlag_DO_NOT_CHANGE
}

func (x *ChangeAlertingRulesRequest) GetEnabled2() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled2
	}
	return nil
}

func (x *ChangeAlertingRulesRequest) GetParams() []*ChangeAlertingRulesRequest_Param {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *ChangeAlertingRulesRequest) GetFor() *duration.Duration {
	if x != nil {
		return x.For
	}
	return nil
}

func (x *ChangeAlertingRulesRequest) GetRemoveFor() bool {
	if x != nil {
		return x.RemoveFor
	}
	return false
}

type ChangeAlertingRulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Changed rule information.
	Rules *Rule `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
}

func (x *ChangeAlertingRulesResponse) Reset() {
	*x = ChangeAlertingRulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAlertingRulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAlertingRulesResponse) ProtoMessage() {}

func (x *ChangeAlertingRulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAlertingRulesResponse.ProtoReflect.Descriptor instead.
func (*ChangeAlertingRulesResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeAlertingRulesResponse) GetRules() *Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// Param repsesents a single Integrated Alerting rule parameter.
type Rule_Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Parameter name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Parameter description.
	Help string `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	// Parameter unit.
	Unit ParamUnit `protobuf:"varint,3,opt,name=unit,proto3,enum=ia.v1beta1.ParamUnit" json:"unit,omitempty"`
	// Parameter type.
	Type ParamType `protobuf:"varint,4,opt,name=type,proto3,enum=ia.v1beta1.ParamType" json:"type,omitempty"`
	// Parameter minimal value (float).
	MinValue float32 `protobuf:"fixed32,5,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	// Parameter maximum value (float).
	MaxValue float32 `protobuf:"fixed32,6,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	// Parameter set value (float).
	DefaultValue float32 `protobuf:"fixed32,7,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Parameter set value (float).
	Value float32 `protobuf:"fixed32,8,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Rule_Param) Reset() {
	*x = Rule_Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule_Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule_Param) ProtoMessage() {}

func (x *Rule_Param) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule_Param.ProtoReflect.Descriptor instead.
func (*Rule_Param) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Rule_Param) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule_Param) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Rule_Param) GetUnit() ParamUnit {
	if x != nil {
		return x.Unit
	}
	return ParamUnit_PARAM_UNIT_INVALID
}

func (x *Rule_Param) GetType() ParamType {
	if x != nil {
		return x.Type
	}
	return ParamType_PARAM_TYPE_INVALID
}

func (x *Rule_Param) GetMinValue() float32 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *Rule_Param) GetMaxValue() float32 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *Rule_Param) GetDefaultValue() float32 {
	if x != nil {
		return x.DefaultValue
	}
	return 0
}

func (x *Rule_Param) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Param repsesents a single Integrated Alerting rule parameter change.
type ChangeAlertingRulesRequest_Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Parameter name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Parameter set value (float).
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	// Remove set value (reset to default).
	RemoveValue bool `protobuf:"varint,3,opt,name=remove_value,json=removeValue,proto3" json:"remove_value,omitempty"`
}

func (x *ChangeAlertingRulesRequest_Param) Reset() {
	*x = ChangeAlertingRulesRequest_Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_ia_rules_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAlertingRulesRequest_Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAlertingRulesRequest_Param) ProtoMessage() {}

func (x *ChangeAlertingRulesRequest_Param) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_ia_rules_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAlertingRulesRequest_Param.ProtoReflect.Descriptor instead.
func (*ChangeAlertingRulesRequest_Param) Descriptor() ([]byte, []int) {
	return file_managementpb_ia_rules_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ChangeAlertingRulesRequest_Param) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeAlertingRulesRequest_Param) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ChangeAlertingRulesRequest_Param) GetRemoveValue() bool {
	if x != nil {
		return x.RemoveValue
	}
	return false
}

var File_managementpb_ia_rules_proto protoreflect.FileDescriptor

var file_managementpb_ia_rules_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69,
	0x61, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69,
	0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x69, 0x61,
	0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65,
	0x6c, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x2e,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x3a,
	0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x03, 0x66, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x03, 0x66, 0x6f, 0x72, 0x1a, 0xfa, 0x01, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x43, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x95, 0x03,
	0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x69, 0x61, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x31, 0x12, 0x36, 0x0a,
	0x08, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x32, 0x12, 0x44, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x66,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x03, 0x66, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x1a, 0x5c, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x45, 0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2a, 0x33, 0x0a, 0x09,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x41, 0x52,
	0x41, 0x4d, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41, 0x47, 0x45, 0x10,
	0x01, 0x2a, 0x2e, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10,
	0x01, 0x32, 0xa7, 0x02, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x24, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x61, 0x2f, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x91, 0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x26, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x61, 0x2f, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x1f, 0x5a, 0x1d, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f,
	0x69, 0x61, 0x3b, 0x69, 0x61, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_ia_rules_proto_rawDescOnce sync.Once
	file_managementpb_ia_rules_proto_rawDescData = file_managementpb_ia_rules_proto_rawDesc
)

func file_managementpb_ia_rules_proto_rawDescGZIP() []byte {
	file_managementpb_ia_rules_proto_rawDescOnce.Do(func() {
		file_managementpb_ia_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_ia_rules_proto_rawDescData)
	})
	return file_managementpb_ia_rules_proto_rawDescData
}

var file_managementpb_ia_rules_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_managementpb_ia_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_managementpb_ia_rules_proto_goTypes = []interface{}{
	(ParamUnit)(0),                           // 0: ia.v1beta1.ParamUnit
	(ParamType)(0),                           // 1: ia.v1beta1.ParamType
	(*Rule)(nil),                             // 2: ia.v1beta1.Rule
	(*ListAlertingRulesRequest)(nil),         // 3: ia.v1beta1.ListAlertingRulesRequest
	(*ListAlertingRulesResponse)(nil),        // 4: ia.v1beta1.ListAlertingRulesResponse
	(*ChangeAlertingRulesRequest)(nil),       // 5: ia.v1beta1.ChangeAlertingRulesRequest
	(*ChangeAlertingRulesResponse)(nil),      // 6: ia.v1beta1.ChangeAlertingRulesResponse
	(*Rule_Param)(nil),                       // 7: ia.v1beta1.Rule.Param
	(*ChangeAlertingRulesRequest_Param)(nil), // 8: ia.v1beta1.ChangeAlertingRulesRequest.Param
	(*duration.Duration)(nil),                // 9: google.protobuf.Duration
	(BooleanFlag)(0),                         // 10: ia.v1beta1.BooleanFlag
	(*wrappers.BoolValue)(nil),               // 11: google.protobuf.BoolValue
}
var file_managementpb_ia_rules_proto_depIdxs = []int32{
	7,  // 0: ia.v1beta1.Rule.params:type_name -> ia.v1beta1.Rule.Param
	9,  // 1: ia.v1beta1.Rule.default_for:type_name -> google.protobuf.Duration
	9,  // 2: ia.v1beta1.Rule.for:type_name -> google.protobuf.Duration
	2,  // 3: ia.v1beta1.ListAlertingRulesResponse.rules:type_name -> ia.v1beta1.Rule
	10, // 4: ia.v1beta1.ChangeAlertingRulesRequest.enabled1:type_name -> ia.v1beta1.BooleanFlag
	11, // 5: ia.v1beta1.ChangeAlertingRulesRequest.enabled2:type_name -> google.protobuf.BoolValue
	8,  // 6: ia.v1beta1.ChangeAlertingRulesRequest.params:type_name -> ia.v1beta1.ChangeAlertingRulesRequest.Param
	9,  // 7: ia.v1beta1.ChangeAlertingRulesRequest.for:type_name -> google.protobuf.Duration
	2,  // 8: ia.v1beta1.ChangeAlertingRulesResponse.rules:type_name -> ia.v1beta1.Rule
	0,  // 9: ia.v1beta1.Rule.Param.unit:type_name -> ia.v1beta1.ParamUnit
	1,  // 10: ia.v1beta1.Rule.Param.type:type_name -> ia.v1beta1.ParamType
	3,  // 11: ia.v1beta1.Rules.ListAlertingRules:input_type -> ia.v1beta1.ListAlertingRulesRequest
	5,  // 12: ia.v1beta1.Rules.ChangeAlertingRules:input_type -> ia.v1beta1.ChangeAlertingRulesRequest
	4,  // 13: ia.v1beta1.Rules.ListAlertingRules:output_type -> ia.v1beta1.ListAlertingRulesResponse
	6,  // 14: ia.v1beta1.Rules.ChangeAlertingRules:output_type -> ia.v1beta1.ChangeAlertingRulesResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_managementpb_ia_rules_proto_init() }
func file_managementpb_ia_rules_proto_init() {
	if File_managementpb_ia_rules_proto != nil {
		return
	}
	file_managementpb_ia_boolean_flag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_ia_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_managementpb_ia_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertingRulesRequest); i {
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
		file_managementpb_ia_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlertingRulesResponse); i {
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
		file_managementpb_ia_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAlertingRulesRequest); i {
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
		file_managementpb_ia_rules_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAlertingRulesResponse); i {
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
		file_managementpb_ia_rules_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule_Param); i {
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
		file_managementpb_ia_rules_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAlertingRulesRequest_Param); i {
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
			RawDescriptor: file_managementpb_ia_rules_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_ia_rules_proto_goTypes,
		DependencyIndexes: file_managementpb_ia_rules_proto_depIdxs,
		EnumInfos:         file_managementpb_ia_rules_proto_enumTypes,
		MessageInfos:      file_managementpb_ia_rules_proto_msgTypes,
	}.Build()
	File_managementpb_ia_rules_proto = out.File
	file_managementpb_ia_rules_proto_rawDesc = nil
	file_managementpb_ia_rules_proto_goTypes = nil
	file_managementpb_ia_rules_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	// ListAlertingRules returns a list of all Integrated Alerting rules.
	ListAlertingRules(ctx context.Context, in *ListAlertingRulesRequest, opts ...grpc.CallOption) (*ListAlertingRulesResponse, error)
	// ChangeAlertingRules changes Integrated Alerting rule.
	ChangeAlertingRules(ctx context.Context, in *ChangeAlertingRulesRequest, opts ...grpc.CallOption) (*ChangeAlertingRulesResponse, error)
}

type rulesClient struct {
	cc grpc.ClientConnInterface
}

func NewRulesClient(cc grpc.ClientConnInterface) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) ListAlertingRules(ctx context.Context, in *ListAlertingRulesRequest, opts ...grpc.CallOption) (*ListAlertingRulesResponse, error) {
	out := new(ListAlertingRulesResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/ListAlertingRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ChangeAlertingRules(ctx context.Context, in *ChangeAlertingRulesRequest, opts ...grpc.CallOption) (*ChangeAlertingRulesResponse, error) {
	out := new(ChangeAlertingRulesResponse)
	err := c.cc.Invoke(ctx, "/ia.v1beta1.Rules/ChangeAlertingRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	// ListAlertingRules returns a list of all Integrated Alerting rules.
	ListAlertingRules(context.Context, *ListAlertingRulesRequest) (*ListAlertingRulesResponse, error)
	// ChangeAlertingRules changes Integrated Alerting rule.
	ChangeAlertingRules(context.Context, *ChangeAlertingRulesRequest) (*ChangeAlertingRulesResponse, error)
}

// UnimplementedRulesServer can be embedded to have forward compatible implementations.
type UnimplementedRulesServer struct {
}

func (*UnimplementedRulesServer) ListAlertingRules(context.Context, *ListAlertingRulesRequest) (*ListAlertingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertingRules not implemented")
}
func (*UnimplementedRulesServer) ChangeAlertingRules(context.Context, *ChangeAlertingRulesRequest) (*ChangeAlertingRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAlertingRules not implemented")
}

func RegisterRulesServer(s *grpc.Server, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_ListAlertingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListAlertingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/ListAlertingRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListAlertingRules(ctx, req.(*ListAlertingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ChangeAlertingRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAlertingRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ChangeAlertingRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ia.v1beta1.Rules/ChangeAlertingRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ChangeAlertingRules(ctx, req.(*ChangeAlertingRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ia.v1beta1.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlertingRules",
			Handler:    _Rules_ListAlertingRules_Handler,
		},
		{
			MethodName: "ChangeAlertingRules",
			Handler:    _Rules_ChangeAlertingRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/ia/rules.proto",
}
