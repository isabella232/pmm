// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/ia/receivers.proto

package iav1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HTTPConfig struct {
	BasicAuth       *HTTPConfig_BasicAuth `protobuf:"bytes,1,opt,name=basic_auth,json=basicAuth,proto3" json:"basic_auth,omitempty"`
	BearerToken     string                `protobuf:"bytes,2,opt,name=bearer_token,json=bearerToken,proto3" json:"bearer_token,omitempty"`
	BearerTokenFile string                `protobuf:"bytes,3,opt,name=bearer_token_file,json=bearerTokenFile,proto3" json:"bearer_token_file,omitempty"`
	// TODO TLSConfig tls_config = 4;
	ProxyUrl             string   `protobuf:"bytes,5,opt,name=proxy_url,json=proxyUrl,proto3" json:"proxy_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPConfig) Reset()         { *m = HTTPConfig{} }
func (m *HTTPConfig) String() string { return proto.CompactTextString(m) }
func (*HTTPConfig) ProtoMessage()    {}
func (*HTTPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{0}
}

func (m *HTTPConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPConfig.Unmarshal(m, b)
}
func (m *HTTPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPConfig.Marshal(b, m, deterministic)
}
func (m *HTTPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPConfig.Merge(m, src)
}
func (m *HTTPConfig) XXX_Size() int {
	return xxx_messageInfo_HTTPConfig.Size(m)
}
func (m *HTTPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPConfig proto.InternalMessageInfo

func (m *HTTPConfig) GetBasicAuth() *HTTPConfig_BasicAuth {
	if m != nil {
		return m.BasicAuth
	}
	return nil
}

func (m *HTTPConfig) GetBearerToken() string {
	if m != nil {
		return m.BearerToken
	}
	return ""
}

func (m *HTTPConfig) GetBearerTokenFile() string {
	if m != nil {
		return m.BearerTokenFile
	}
	return ""
}

func (m *HTTPConfig) GetProxyUrl() string {
	if m != nil {
		return m.ProxyUrl
	}
	return ""
}

type HTTPConfig_BasicAuth struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordFile         string   `protobuf:"bytes,3,opt,name=password_file,json=passwordFile,proto3" json:"password_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPConfig_BasicAuth) Reset()         { *m = HTTPConfig_BasicAuth{} }
func (m *HTTPConfig_BasicAuth) String() string { return proto.CompactTextString(m) }
func (*HTTPConfig_BasicAuth) ProtoMessage()    {}
func (*HTTPConfig_BasicAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{0, 0}
}

func (m *HTTPConfig_BasicAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPConfig_BasicAuth.Unmarshal(m, b)
}
func (m *HTTPConfig_BasicAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPConfig_BasicAuth.Marshal(b, m, deterministic)
}
func (m *HTTPConfig_BasicAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPConfig_BasicAuth.Merge(m, src)
}
func (m *HTTPConfig_BasicAuth) XXX_Size() int {
	return xxx_messageInfo_HTTPConfig_BasicAuth.Size(m)
}
func (m *HTTPConfig_BasicAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPConfig_BasicAuth.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPConfig_BasicAuth proto.InternalMessageInfo

func (m *HTTPConfig_BasicAuth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *HTTPConfig_BasicAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *HTTPConfig_BasicAuth) GetPasswordFile() string {
	if m != nil {
		return m.PasswordFile
	}
	return ""
}

type EmailConfig struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailConfig) Reset()         { *m = EmailConfig{} }
func (m *EmailConfig) String() string { return proto.CompactTextString(m) }
func (*EmailConfig) ProtoMessage()    {}
func (*EmailConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{1}
}

func (m *EmailConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailConfig.Unmarshal(m, b)
}
func (m *EmailConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailConfig.Marshal(b, m, deterministic)
}
func (m *EmailConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailConfig.Merge(m, src)
}
func (m *EmailConfig) XXX_Size() int {
	return xxx_messageInfo_EmailConfig.Size(m)
}
func (m *EmailConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EmailConfig proto.InternalMessageInfo

func (m *EmailConfig) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type SlackConfig struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlackConfig) Reset()         { *m = SlackConfig{} }
func (m *SlackConfig) String() string { return proto.CompactTextString(m) }
func (*SlackConfig) ProtoMessage()    {}
func (*SlackConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{2}
}

func (m *SlackConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlackConfig.Unmarshal(m, b)
}
func (m *SlackConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlackConfig.Marshal(b, m, deterministic)
}
func (m *SlackConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlackConfig.Merge(m, src)
}
func (m *SlackConfig) XXX_Size() int {
	return xxx_messageInfo_SlackConfig.Size(m)
}
func (m *SlackConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SlackConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SlackConfig proto.InternalMessageInfo

func (m *SlackConfig) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type WebhookConfig struct {
	SendResolved         bool        `protobuf:"varint,1,opt,name=send_resolved,json=sendResolved,proto3" json:"send_resolved,omitempty"`
	Url                  string      `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	HttpConfig           *HTTPConfig `protobuf:"bytes,3,opt,name=http_config,json=httpConfig,proto3" json:"http_config,omitempty"`
	MaxAlerts            int32       `protobuf:"varint,4,opt,name=max_alerts,json=maxAlerts,proto3" json:"max_alerts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WebhookConfig) Reset()         { *m = WebhookConfig{} }
func (m *WebhookConfig) String() string { return proto.CompactTextString(m) }
func (*WebhookConfig) ProtoMessage()    {}
func (*WebhookConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{3}
}

func (m *WebhookConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebhookConfig.Unmarshal(m, b)
}
func (m *WebhookConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebhookConfig.Marshal(b, m, deterministic)
}
func (m *WebhookConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebhookConfig.Merge(m, src)
}
func (m *WebhookConfig) XXX_Size() int {
	return xxx_messageInfo_WebhookConfig.Size(m)
}
func (m *WebhookConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WebhookConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WebhookConfig proto.InternalMessageInfo

func (m *WebhookConfig) GetSendResolved() bool {
	if m != nil {
		return m.SendResolved
	}
	return false
}

func (m *WebhookConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebhookConfig) GetHttpConfig() *HTTPConfig {
	if m != nil {
		return m.HttpConfig
	}
	return nil
}

func (m *WebhookConfig) GetMaxAlerts() int32 {
	if m != nil {
		return m.MaxAlerts
	}
	return 0
}

type Receiver struct {
	Editable    bool           `protobuf:"varint,1,opt,name=editable,proto3" json:"editable,omitempty"`
	Name        string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EmailConfig []*EmailConfig `protobuf:"bytes,3,rep,name=email_config,json=emailConfig,proto3" json:"email_config,omitempty"`
	// pagerduty_configs = 3;
	// pushover_configs = 4;
	SlackConfigs []*SlackConfig `protobuf:"bytes,5,rep,name=slack_configs,json=slackConfigs,proto3" json:"slack_configs,omitempty"`
	// opsgenie_configs = 6;
	WebhookConfigs       []*WebhookConfig `protobuf:"bytes,7,rep,name=webhook_configs,json=webhookConfigs,proto3" json:"webhook_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Receiver) Reset()         { *m = Receiver{} }
func (m *Receiver) String() string { return proto.CompactTextString(m) }
func (*Receiver) ProtoMessage()    {}
func (*Receiver) Descriptor() ([]byte, []int) {
	return fileDescriptor_aed4d28027f7f0c9, []int{4}
}

func (m *Receiver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receiver.Unmarshal(m, b)
}
func (m *Receiver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receiver.Marshal(b, m, deterministic)
}
func (m *Receiver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receiver.Merge(m, src)
}
func (m *Receiver) XXX_Size() int {
	return xxx_messageInfo_Receiver.Size(m)
}
func (m *Receiver) XXX_DiscardUnknown() {
	xxx_messageInfo_Receiver.DiscardUnknown(m)
}

var xxx_messageInfo_Receiver proto.InternalMessageInfo

func (m *Receiver) GetEditable() bool {
	if m != nil {
		return m.Editable
	}
	return false
}

func (m *Receiver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Receiver) GetEmailConfig() []*EmailConfig {
	if m != nil {
		return m.EmailConfig
	}
	return nil
}

func (m *Receiver) GetSlackConfigs() []*SlackConfig {
	if m != nil {
		return m.SlackConfigs
	}
	return nil
}

func (m *Receiver) GetWebhookConfigs() []*WebhookConfig {
	if m != nil {
		return m.WebhookConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*HTTPConfig)(nil), "ia.v1beta1.HTTPConfig")
	proto.RegisterType((*HTTPConfig_BasicAuth)(nil), "ia.v1beta1.HTTPConfig.BasicAuth")
	proto.RegisterType((*EmailConfig)(nil), "ia.v1beta1.EmailConfig")
	proto.RegisterType((*SlackConfig)(nil), "ia.v1beta1.SlackConfig")
	proto.RegisterType((*WebhookConfig)(nil), "ia.v1beta1.WebhookConfig")
	proto.RegisterType((*Receiver)(nil), "ia.v1beta1.Receiver")
}

func init() {
	proto.RegisterFile("managementpb/ia/receivers.proto", fileDescriptor_aed4d28027f7f0c9)
}

var fileDescriptor_aed4d28027f7f0c9 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0xda, 0x95, 0x35, 0x9f, 0xdb, 0x0d, 0x7c, 0x80, 0x30, 0x34, 0x51, 0xba, 0x03, 0x15,
	0x87, 0x4e, 0x2b, 0x07, 0x24, 0x84, 0x84, 0x56, 0x04, 0xe2, 0x88, 0x4c, 0x11, 0x12, 0x97, 0xc8,
	0x69, 0xbf, 0x2d, 0xd6, 0x9c, 0x38, 0xb2, 0xdd, 0x1f, 0xfc, 0x33, 0x5c, 0xf8, 0x27, 0x39, 0x22,
	0xbb, 0x4e, 0x1a, 0xc4, 0x6e, 0xdf, 0x7b, 0x7e, 0x2f, 0x7e, 0xdf, 0x4b, 0x02, 0xcf, 0x0b, 0x5e,
	0xf2, 0x5b, 0x2c, 0xb0, 0xb4, 0x55, 0x76, 0x29, 0xf8, 0xa5, 0xc6, 0x25, 0x8a, 0x0d, 0x6a, 0x33,
	0xad, 0xb4, 0xb2, 0x8a, 0x82, 0xe0, 0xd3, 0xcd, 0x55, 0x86, 0x96, 0x5f, 0x8d, 0x7f, 0x77, 0x00,
	0x3e, 0x2f, 0x16, 0x5f, 0x3e, 0xa8, 0xf2, 0x46, 0xdc, 0xd2, 0xf7, 0x00, 0x19, 0x37, 0x62, 0x99,
	0xf2, 0xb5, 0xcd, 0x93, 0x68, 0x14, 0x4d, 0xc8, 0x6c, 0x34, 0x3d, 0xe8, 0xa7, 0x07, 0xed, 0x74,
	0xee, 0x84, 0xd7, 0x6b, 0x9b, 0xb3, 0x38, 0xab, 0x47, 0xfa, 0x02, 0x06, 0x19, 0x72, 0x8d, 0x3a,
	0xb5, 0xea, 0x0e, 0xcb, 0xa4, 0x33, 0x8a, 0x26, 0x31, 0x23, 0x7b, 0x6e, 0xe1, 0x28, 0xfa, 0x0a,
	0x1e, 0xb5, 0x25, 0xe9, 0x8d, 0x90, 0x98, 0x74, 0xbd, 0xee, 0xb4, 0xa5, 0xfb, 0x24, 0x24, 0xd2,
	0x67, 0x10, 0x57, 0x5a, 0xed, 0x7e, 0xa6, 0x6b, 0x2d, 0x93, 0x9e, 0xd7, 0xf4, 0x3d, 0xf1, 0x4d,
	0xcb, 0xb3, 0x1c, 0xe2, 0x26, 0x03, 0x3d, 0x83, 0xfe, 0xda, 0xa0, 0x2e, 0x79, 0x81, 0x3e, 0x77,
	0xcc, 0x1a, 0xec, 0xce, 0x2a, 0x6e, 0xcc, 0x56, 0xe9, 0x55, 0x08, 0xd4, 0x60, 0x7a, 0x01, 0xc3,
	0x7a, 0x6e, 0x27, 0x19, 0xd4, 0xa4, 0x8b, 0x31, 0x3e, 0x07, 0xf2, 0xb1, 0xe0, 0x42, 0x86, 0x96,
	0x4e, 0xa0, 0x63, 0x55, 0xb8, 0xa5, 0x63, 0xd5, 0xf8, 0x25, 0x90, 0xaf, 0x92, 0x2f, 0xef, 0xc2,
	0x71, 0x02, 0xc7, 0xcb, 0x9c, 0x97, 0x25, 0xca, 0xa0, 0xa9, 0xe1, 0xf8, 0x57, 0x04, 0xc3, 0xef,
	0x98, 0xe5, 0x4a, 0xd5, 0xda, 0x0b, 0x18, 0x1a, 0x2c, 0x57, 0xa9, 0x46, 0xa3, 0xe4, 0x06, 0x57,
	0xde, 0xd1, 0x67, 0x03, 0x47, 0xb2, 0xc0, 0xd1, 0x87, 0xd0, 0x75, 0xfb, 0xef, 0xa3, 0xbb, 0x91,
	0xbe, 0x01, 0x92, 0x5b, 0x5b, 0xa5, 0x4b, 0xff, 0x14, 0x9f, 0x99, 0xcc, 0x1e, 0xdf, 0xff, 0xa2,
	0x18, 0x38, 0x69, 0xb8, 0xef, 0x1c, 0xa0, 0xe0, 0xbb, 0x94, 0x4b, 0xd4, 0xd6, 0x24, 0x47, 0xa3,
	0x68, 0xd2, 0x63, 0x71, 0xc1, 0x77, 0xd7, 0x9e, 0x18, 0xff, 0x89, 0xa0, 0xcf, 0xc2, 0xe7, 0xe2,
	0x6a, 0xc3, 0x95, 0xb0, 0x3c, 0x93, 0x18, 0x62, 0x35, 0x98, 0x52, 0x38, 0xf2, 0x55, 0xef, 0x33,
	0xf9, 0x99, 0xbe, 0x85, 0x01, 0xba, 0x96, 0x0e, 0xa9, 0xba, 0x13, 0x32, 0x7b, 0xd2, 0x4e, 0xd5,
	0x6a, 0x91, 0x11, 0x6c, 0x55, 0xfa, 0x0e, 0x86, 0xc6, 0x55, 0x18, 0xbc, 0x26, 0xe9, 0xfd, 0x6f,
	0x6e, 0x75, 0xcc, 0x06, 0xe6, 0x00, 0x0c, 0x9d, 0xc3, 0xe9, 0x76, 0x5f, 0x6b, 0xe3, 0x3f, 0xf6,
	0xfe, 0xa7, 0x6d, 0xff, 0x3f, 0xcd, 0xb3, 0x93, 0x6d, 0x1b, 0x9a, 0x19, 0x81, 0xb8, 0xde, 0xdc,
	0xcc, 0xc9, 0x8f, 0x58, 0xf0, 0xe0, 0xcb, 0x1e, 0xf8, 0xdf, 0xe6, 0xf5, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x40, 0xe6, 0xad, 0x59, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReceiversClient is the client API for Receivers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReceiversClient interface {
}

type receiversClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiversClient(cc grpc.ClientConnInterface) ReceiversClient {
	return &receiversClient{cc}
}

// ReceiversServer is the server API for Receivers service.
type ReceiversServer interface {
}

// UnimplementedReceiversServer can be embedded to have forward compatible implementations.
type UnimplementedReceiversServer struct {
}

func RegisterReceiversServer(s *grpc.Server, srv ReceiversServer) {
	s.RegisterService(&_Receivers_serviceDesc, srv)
}

var _Receivers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ia.v1beta1.Receivers",
	HandlerType: (*ReceiversServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "managementpb/ia/receivers.proto",
}
