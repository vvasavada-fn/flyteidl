// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/audit.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Request_Mode int32

const (
	Request_READ_ONLY  Request_Mode = 0
	Request_READ_WRITE Request_Mode = 1
)

var Request_Mode_name = map[int32]string{
	0: "READ_ONLY",
	1: "READ_WRITE",
}

var Request_Mode_value = map[string]int32{
	"READ_ONLY":  0,
	"READ_WRITE": 1,
}

func (x Request_Mode) String() string {
	return proto.EnumName(Request_Mode_name, int32(x))
}

func (Request_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{2, 0}
}

// Audit log that is emitted on each user request.
type AuditLog struct {
	Principal            *Principal `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	ClientIp             string     `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Request              *Request   `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Response             *Response  `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}
func (*AuditLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{0}
}

func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return xxx_messageInfo_AuditLog.Size(m)
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetPrincipal() *Principal {
	if m != nil {
		return m.Principal
	}
	return nil
}

func (m *AuditLog) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *AuditLog) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AuditLog) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

// Details about the authenticated user who issued a service request.
type Principal struct {
	// Identifies authenticated end-user
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// The client that initiated the auth flow.
	ClientId             string               `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TokenIssuedAt        *timestamp.Timestamp `protobuf:"bytes,3,opt,name=token_issued_at,json=tokenIssuedAt,proto3" json:"token_issued_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{1}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Principal) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Principal) GetTokenIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.TokenIssuedAt
	}
	return nil
}

// Details about a specific request issued by a user.
type Request struct {
	// Service method endpoint e.g. GetWorkflowExecution
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// In the case of HTTP service calls.
	HttpVerb string `protobuf:"bytes,2,opt,name=http_verb,json=httpVerb,proto3" json:"http_verb,omitempty"`
	// Includes parameters submitted in the request.
	Parameters           map[string]string    `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mode                 Request_Mode         `protobuf:"varint,4,opt,name=mode,proto3,enum=flyteidl.admin.Request_Mode" json:"mode,omitempty"`
	ReceivedAt           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetHttpVerb() string {
	if m != nil {
		return m.HttpVerb
	}
	return ""
}

func (m *Request) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Request) GetMode() Request_Mode {
	if m != nil {
		return m.Mode
	}
	return Request_READ_ONLY
}

func (m *Request) GetReceivedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

// Summary of service response details.
type Response struct {
	// e.g. gRPC status code
	ResponseCode         string               `protobuf:"bytes,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	SentAt               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f7dfb38018d590, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseCode() string {
	if m != nil {
		return m.ResponseCode
	}
	return ""
}

func (m *Response) GetSentAt() *timestamp.Timestamp {
	if m != nil {
		return m.SentAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.Request_Mode", Request_Mode_name, Request_Mode_value)
	proto.RegisterType((*AuditLog)(nil), "flyteidl.admin.AuditLog")
	proto.RegisterType((*Principal)(nil), "flyteidl.admin.Principal")
	proto.RegisterType((*Request)(nil), "flyteidl.admin.Request")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.Request.ParametersEntry")
	proto.RegisterType((*Response)(nil), "flyteidl.admin.Response")
}

func init() { proto.RegisterFile("flyteidl/admin/audit.proto", fileDescriptor_11f7dfb38018d590) }

var fileDescriptor_11f7dfb38018d590 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xed, 0x8a, 0xd3, 0x40,
	0x14, 0x35, 0xdb, 0x6e, 0xdb, 0xdc, 0xda, 0x6e, 0x19, 0x44, 0x63, 0x15, 0x2c, 0x15, 0xb1, 0x7f,
	0x4c, 0xdc, 0x56, 0x50, 0x14, 0x7f, 0x74, 0xb5, 0x48, 0x61, 0xd5, 0x65, 0x58, 0x14, 0xfd, 0x13,
	0xf2, 0x71, 0x9b, 0x8e, 0x9b, 0x64, 0xe2, 0x64, 0x52, 0xe8, 0x0b, 0xf8, 0x6e, 0xbe, 0x83, 0x0f,
	0x23, 0x99, 0x64, 0xea, 0x6e, 0x65, 0xf5, 0x5f, 0xee, 0xbd, 0xe7, 0x9c, 0x9c, 0x73, 0x2f, 0x03,
	0xc3, 0x55, 0xbc, 0x95, 0xc8, 0xc2, 0xd8, 0xf1, 0xc2, 0x84, 0xa5, 0x8e, 0x57, 0x84, 0x4c, 0xda,
	0x99, 0xe0, 0x92, 0x93, 0xbe, 0x9e, 0xd9, 0x6a, 0x36, 0x7c, 0x10, 0x71, 0x1e, 0xc5, 0xe8, 0xa8,
	0xa9, 0x5f, 0xac, 0x1c, 0xc9, 0x12, 0xcc, 0xa5, 0x97, 0x64, 0x15, 0x61, 0xfc, 0xd3, 0x80, 0xce,
	0xbc, 0x14, 0x38, 0xe5, 0x11, 0x79, 0x0e, 0x66, 0x26, 0x58, 0x1a, 0xb0, 0xcc, 0x8b, 0x2d, 0x63,
	0x64, 0x4c, 0xba, 0xd3, 0xbb, 0xf6, 0x55, 0x45, 0xfb, 0x4c, 0x03, 0xe8, 0x1f, 0x2c, 0xb9, 0x07,
	0x66, 0x10, 0x33, 0x4c, 0xa5, 0xcb, 0x32, 0xeb, 0x60, 0x64, 0x4c, 0x4c, 0xda, 0xa9, 0x1a, 0xcb,
	0x8c, 0x1c, 0x43, 0x5b, 0xe0, 0xf7, 0x02, 0x73, 0x69, 0x35, 0x94, 0xe6, 0x9d, 0x7d, 0x4d, 0x5a,
	0x8d, 0xa9, 0xc6, 0x91, 0x67, 0xd0, 0x11, 0x98, 0x67, 0x3c, 0xcd, 0xd1, 0x6a, 0x2a, 0x8e, 0xf5,
	0x37, 0xa7, 0x9a, 0xd3, 0x1d, 0x72, 0xfc, 0xc3, 0x00, 0x73, 0x67, 0x8f, 0x58, 0xd0, 0xce, 0x0b,
	0xff, 0x1b, 0x06, 0x52, 0x45, 0x31, 0xa9, 0x2e, 0x2f, 0xbb, 0x0d, 0xf7, 0xdc, 0x86, 0xe4, 0x04,
	0x8e, 0x24, 0xbf, 0xc0, 0xd4, 0x65, 0x79, 0x5e, 0x60, 0xe8, 0x7a, 0xda, 0xf5, 0xd0, 0xae, 0x76,
	0x69, 0xeb, 0x5d, 0xda, 0xe7, 0x7a, 0x97, 0xb4, 0xa7, 0x28, 0x4b, 0xc5, 0x98, 0xcb, 0xf1, 0xaf,
	0x03, 0x68, 0xd7, 0x99, 0xc8, 0x6d, 0x68, 0x25, 0x28, 0xd7, 0x3c, 0xac, 0x5d, 0xd4, 0x55, 0x69,
	0x62, 0x2d, 0x65, 0xe6, 0x6e, 0x50, 0xf8, 0xda, 0x44, 0xd9, 0xf8, 0x84, 0xc2, 0x27, 0xef, 0x00,
	0x32, 0x4f, 0x78, 0x09, 0x4a, 0x14, 0xb9, 0xd5, 0x18, 0x35, 0x26, 0xdd, 0xe9, 0xe3, 0x6b, 0xb6,
	0x66, 0x9f, 0xed, 0x90, 0x8b, 0x54, 0x8a, 0x2d, 0xbd, 0x44, 0x25, 0x4f, 0xa1, 0x99, 0xf0, 0xb0,
	0x5a, 0x62, 0x7f, 0x7a, 0xff, 0x3a, 0x89, 0xf7, 0x3c, 0x44, 0xaa, 0x90, 0xe4, 0x15, 0x74, 0x05,
	0x06, 0xc8, 0x36, 0x55, 0xf6, 0xc3, 0xff, 0x66, 0x07, 0x0d, 0x9f, 0xcb, 0xe1, 0x6b, 0x38, 0xda,
	0x73, 0x43, 0x06, 0xd0, 0xb8, 0xc0, 0x6d, 0x1d, 0xbe, 0xfc, 0x24, 0xb7, 0xe0, 0x70, 0xe3, 0xc5,
	0x05, 0xd6, 0xa9, 0xab, 0xe2, 0xe5, 0xc1, 0x0b, 0x63, 0xfc, 0x08, 0x9a, 0xa5, 0x13, 0xd2, 0x03,
	0x93, 0x2e, 0xe6, 0x6f, 0xdd, 0x8f, 0x1f, 0x4e, 0xbf, 0x0c, 0x6e, 0x90, 0x3e, 0x80, 0x2a, 0x3f,
	0xd3, 0xe5, 0xf9, 0x62, 0x60, 0x8c, 0x43, 0xe8, 0xe8, 0xeb, 0x93, 0x87, 0xd0, 0xd3, 0xf7, 0x77,
	0x83, 0x32, 0x69, 0xf5, 0xa3, 0x9b, 0xba, 0xf9, 0xa6, 0xd4, 0x9b, 0x41, 0x3b, 0x2f, 0xcf, 0xed,
	0x49, 0xf5, 0xcf, 0x7f, 0xe7, 0x69, 0x95, 0xd0, 0xb9, 0x3c, 0x99, 0x7d, 0x3d, 0x8e, 0x98, 0x5c,
	0x17, 0xbe, 0x1d, 0xf0, 0xc4, 0x89, 0xb7, 0x2b, 0xe9, 0xec, 0x1e, 0x5e, 0x84, 0xa9, 0x93, 0xf9,
	0x4f, 0x22, 0xee, 0x5c, 0x7d, 0x8b, 0x7e, 0x4b, 0x09, 0xce, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x05, 0xcf, 0xf5, 0x5a, 0xa4, 0x03, 0x00, 0x00,
}