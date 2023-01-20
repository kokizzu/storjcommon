// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contact.proto

package pb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CheckInRequest struct {
	Address  string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Version  *NodeVersion  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Capacity *NodeCapacity `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Operator *NodeOperator `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	// CheckIn may happen over TLS, which means there is not necessarily an
	// opportunity to use NoiseSessionAttestation instead of
	// NoiseKeyAttestation.
	NoiseKeyAttestation  *NoiseKeyAttestation `protobuf:"bytes,5,opt,name=noise_key_attestation,json=noiseKeyAttestation,proto3" json:"noise_key_attestation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CheckInRequest) Reset()         { *m = CheckInRequest{} }
func (m *CheckInRequest) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest) ProtoMessage()    {}
func (*CheckInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}
func (m *CheckInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest.Unmarshal(m, b)
}
func (m *CheckInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest.Marshal(b, m, deterministic)
}
func (m *CheckInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest.Merge(m, src)
}
func (m *CheckInRequest) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest.Size(m)
}
func (m *CheckInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest proto.InternalMessageInfo

func (m *CheckInRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CheckInRequest) GetVersion() *NodeVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *CheckInRequest) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *CheckInRequest) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *CheckInRequest) GetNoiseKeyAttestation() *NoiseKeyAttestation {
	if m != nil {
		return m.NoiseKeyAttestation
	}
	return nil
}

type CheckInResponse struct {
	PingNodeSuccess      bool     `protobuf:"varint,1,opt,name=ping_node_success,json=pingNodeSuccess,proto3" json:"ping_node_success,omitempty"`
	PingErrorMessage     string   `protobuf:"bytes,2,opt,name=ping_error_message,json=pingErrorMessage,proto3" json:"ping_error_message,omitempty"`
	PingNodeSuccessQuic  bool     `protobuf:"varint,3,opt,name=ping_node_success_quic,json=pingNodeSuccessQuic,proto3" json:"ping_node_success_quic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}
func (*CheckInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}
func (m *CheckInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInResponse.Unmarshal(m, b)
}
func (m *CheckInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInResponse.Marshal(b, m, deterministic)
}
func (m *CheckInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInResponse.Merge(m, src)
}
func (m *CheckInResponse) XXX_Size() int {
	return xxx_messageInfo_CheckInResponse.Size(m)
}
func (m *CheckInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInResponse proto.InternalMessageInfo

func (m *CheckInResponse) GetPingNodeSuccess() bool {
	if m != nil {
		return m.PingNodeSuccess
	}
	return false
}

func (m *CheckInResponse) GetPingErrorMessage() string {
	if m != nil {
		return m.PingErrorMessage
	}
	return ""
}

func (m *CheckInResponse) GetPingNodeSuccessQuic() bool {
	if m != nil {
		return m.PingNodeSuccessQuic
	}
	return false
}

type GetTimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}
func (m *GetTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeRequest.Unmarshal(m, b)
}
func (m *GetTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeRequest.Merge(m, src)
}
func (m *GetTimeRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeRequest.Size(m)
}
func (m *GetTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeRequest proto.InternalMessageInfo

type GetTimeResponse struct {
	Timestamp            time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTimeResponse) Reset()         { *m = GetTimeResponse{} }
func (m *GetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeResponse) ProtoMessage()    {}
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}
func (m *GetTimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeResponse.Unmarshal(m, b)
}
func (m *GetTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeResponse.Marshal(b, m, deterministic)
}
func (m *GetTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeResponse.Merge(m, src)
}
func (m *GetTimeResponse) XXX_Size() int {
	return xxx_messageInfo_GetTimeResponse.Size(m)
}
func (m *GetTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeResponse proto.InternalMessageInfo

func (m *GetTimeResponse) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type ContactPingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingRequest) Reset()         { *m = ContactPingRequest{} }
func (m *ContactPingRequest) String() string { return proto.CompactTextString(m) }
func (*ContactPingRequest) ProtoMessage()    {}
func (*ContactPingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}
func (m *ContactPingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingRequest.Unmarshal(m, b)
}
func (m *ContactPingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingRequest.Marshal(b, m, deterministic)
}
func (m *ContactPingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingRequest.Merge(m, src)
}
func (m *ContactPingRequest) XXX_Size() int {
	return xxx_messageInfo_ContactPingRequest.Size(m)
}
func (m *ContactPingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingRequest proto.InternalMessageInfo

type ContactPingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingResponse) Reset()         { *m = ContactPingResponse{} }
func (m *ContactPingResponse) String() string { return proto.CompactTextString(m) }
func (*ContactPingResponse) ProtoMessage()    {}
func (*ContactPingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}
func (m *ContactPingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingResponse.Unmarshal(m, b)
}
func (m *ContactPingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingResponse.Marshal(b, m, deterministic)
}
func (m *ContactPingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingResponse.Merge(m, src)
}
func (m *ContactPingResponse) XXX_Size() int {
	return xxx_messageInfo_ContactPingResponse.Size(m)
}
func (m *ContactPingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingResponse proto.InternalMessageInfo

type PingMeRequest struct {
	Address   string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Transport NodeTransport `protobuf:"varint,2,opt,name=transport,proto3,enum=node.NodeTransport" json:"transport,omitempty"`
	// PingMe may happen over TLS, which means there is not necessarily an
	// opportunity to use NoiseSessionAttestation instead of
	// NoiseKeyAttestation.
	NoiseKeyAttestation  *NoiseKeyAttestation `protobuf:"bytes,3,opt,name=noise_key_attestation,json=noiseKeyAttestation,proto3" json:"noise_key_attestation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PingMeRequest) Reset()         { *m = PingMeRequest{} }
func (m *PingMeRequest) String() string { return proto.CompactTextString(m) }
func (*PingMeRequest) ProtoMessage()    {}
func (*PingMeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{6}
}
func (m *PingMeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMeRequest.Unmarshal(m, b)
}
func (m *PingMeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMeRequest.Marshal(b, m, deterministic)
}
func (m *PingMeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMeRequest.Merge(m, src)
}
func (m *PingMeRequest) XXX_Size() int {
	return xxx_messageInfo_PingMeRequest.Size(m)
}
func (m *PingMeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingMeRequest proto.InternalMessageInfo

func (m *PingMeRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PingMeRequest) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP_TLS_RPC
}

func (m *PingMeRequest) GetNoiseKeyAttestation() *NoiseKeyAttestation {
	if m != nil {
		return m.NoiseKeyAttestation
	}
	return nil
}

type PingMeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMeResponse) Reset()         { *m = PingMeResponse{} }
func (m *PingMeResponse) String() string { return proto.CompactTextString(m) }
func (*PingMeResponse) ProtoMessage()    {}
func (*PingMeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{7}
}
func (m *PingMeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMeResponse.Unmarshal(m, b)
}
func (m *PingMeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMeResponse.Marshal(b, m, deterministic)
}
func (m *PingMeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMeResponse.Merge(m, src)
}
func (m *PingMeResponse) XXX_Size() int {
	return xxx_messageInfo_PingMeResponse.Size(m)
}
func (m *PingMeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingMeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckInRequest)(nil), "contact.CheckInRequest")
	proto.RegisterType((*CheckInResponse)(nil), "contact.CheckInResponse")
	proto.RegisterType((*GetTimeRequest)(nil), "contact.GetTimeRequest")
	proto.RegisterType((*GetTimeResponse)(nil), "contact.GetTimeResponse")
	proto.RegisterType((*ContactPingRequest)(nil), "contact.ContactPingRequest")
	proto.RegisterType((*ContactPingResponse)(nil), "contact.ContactPingResponse")
	proto.RegisterType((*PingMeRequest)(nil), "contact.PingMeRequest")
	proto.RegisterType((*PingMeResponse)(nil), "contact.PingMeResponse")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xdc, 0xf6, 0xab, 0x93, 0x5b, 0x35, 0x49, 0x27, 0xfd, 0xb1, 0x0c, 0x52, 0x2a, 0xaf,
	0x2a, 0x40, 0x8e, 0x48, 0x57, 0x48, 0xdd, 0x90, 0xa8, 0x42, 0x08, 0x35, 0x84, 0x21, 0xb0, 0x60,
	0x63, 0x39, 0xce, 0x60, 0x4c, 0xf1, 0x8c, 0x3b, 0x33, 0x41, 0xca, 0x13, 0xb0, 0xe5, 0x11, 0x78,
	0x00, 0x5e, 0x81, 0x3d, 0x4f, 0x01, 0xaf, 0x82, 0xe6, 0xc7, 0x76, 0x93, 0x80, 0x90, 0xd8, 0x44,
	0x99, 0x73, 0xce, 0xdc, 0xb9, 0xf7, 0xdc, 0x63, 0xd8, 0x4f, 0x18, 0x95, 0x71, 0x22, 0xc3, 0x82,
	0x33, 0xc9, 0x90, 0x6b, 0x8f, 0x3e, 0xa4, 0x2c, 0x65, 0x06, 0xf4, 0x7b, 0x29, 0x63, 0xe9, 0x07,
	0xd2, 0xd7, 0xa7, 0xd9, 0xe2, 0x6d, 0x5f, 0x66, 0x39, 0x11, 0x32, 0xce, 0x0b, 0x2b, 0x00, 0xca,
	0xe6, 0xc4, 0xfe, 0xdf, 0xa3, 0x2c, 0x13, 0xf6, 0x10, 0x7c, 0xda, 0x82, 0xd6, 0xe8, 0x1d, 0x49,
	0xae, 0x9f, 0x52, 0x4c, 0x6e, 0x16, 0x44, 0x48, 0xe4, 0x81, 0x1b, 0xcf, 0xe7, 0x9c, 0x08, 0xe1,
	0x39, 0xa7, 0xce, 0x59, 0x13, 0x97, 0x47, 0x74, 0x1f, 0xdc, 0x8f, 0x84, 0x8b, 0x8c, 0x51, 0x6f,
	0xeb, 0xd4, 0x39, 0xdb, 0x1b, 0x1c, 0x84, 0xba, 0xee, 0x98, 0xcd, 0xc9, 0x6b, 0x43, 0xe0, 0x52,
	0x81, 0x42, 0x68, 0x24, 0x71, 0x11, 0x27, 0x99, 0x5c, 0x7a, 0xdb, 0x5a, 0x8d, 0x6a, 0xf5, 0xc8,
	0x32, 0xb8, 0xd2, 0x28, 0x3d, 0x2b, 0x08, 0x8f, 0x25, 0xe3, 0xde, 0xce, 0xba, 0xfe, 0xb9, 0x65,
	0x70, 0xa5, 0x41, 0x63, 0x38, 0xd2, 0x83, 0x44, 0xd7, 0x64, 0x19, 0xc5, 0x52, 0xaa, 0x79, 0xa5,
	0x6a, 0xed, 0x7f, 0x7d, 0xd9, 0x0f, 0xcd, 0x98, 0x63, 0xf5, 0xfb, 0x8c, 0x2c, 0x1f, 0xd7, 0x0a,
	0xdc, 0xa5, 0x9b, 0x60, 0xf0, 0xc5, 0x81, 0x76, 0xe5, 0x84, 0x28, 0x18, 0x15, 0x04, 0xdd, 0x83,
	0x83, 0x22, 0xa3, 0x69, 0xa4, 0xfa, 0x88, 0xc4, 0x22, 0x49, 0x4a, 0x53, 0x1a, 0xb8, 0xad, 0x08,
	0xd5, 0xda, 0x4b, 0x03, 0xa3, 0x07, 0x80, 0xb4, 0x96, 0x70, 0xce, 0x78, 0x94, 0x13, 0x21, 0xe2,
	0x94, 0x68, 0x9f, 0x9a, 0xb8, 0xa3, 0x98, 0x4b, 0x45, 0x5c, 0x19, 0x1c, 0x9d, 0xc3, 0xf1, 0x46,
	0xe5, 0xe8, 0x66, 0x91, 0x25, 0xda, 0xab, 0x06, 0xee, 0xae, 0x95, 0x7f, 0xb1, 0xc8, 0x92, 0xa0,
	0x03, 0xad, 0x27, 0x44, 0x4e, 0xb3, 0x9c, 0xd8, 0x5d, 0x05, 0xaf, 0xa0, 0x5d, 0x21, 0xb6, 0xe7,
	0x21, 0x34, 0xab, 0xed, 0xeb, 0x5e, 0x95, 0x17, 0x26, 0x1f, 0x61, 0x99, 0x8f, 0x70, 0x5a, 0x2a,
	0x86, 0x8d, 0xef, 0x3f, 0x7a, 0xff, 0x7d, 0xfe, 0xd9, 0x73, 0x70, 0x7d, 0x2d, 0x38, 0x04, 0x34,
	0x32, 0x31, 0x9b, 0x64, 0x34, 0x2d, 0x1f, 0x3b, 0x82, 0xee, 0x0a, 0x6a, 0x1e, 0x0c, 0xbe, 0x3a,
	0xb0, 0xaf, 0x80, 0x2b, 0xf2, 0xf7, 0x04, 0x3d, 0x84, 0xa6, 0xe4, 0x31, 0x15, 0x05, 0xe3, 0x52,
	0x7b, 0xd3, 0x1a, 0x74, 0xeb, 0x2d, 0x4f, 0x4b, 0x0a, 0xd7, 0xaa, 0x3f, 0xef, 0x79, 0xfb, 0xdf,
	0xf6, 0xdc, 0x81, 0x56, 0xd9, 0xad, 0x19, 0x60, 0x30, 0x01, 0xd7, 0xce, 0x85, 0x2e, 0xa1, 0x31,
	0xb1, 0xc6, 0xa3, 0x3b, 0x61, 0xf9, 0xe5, 0x6d, 0x7a, 0xe1, 0xdf, 0xfd, 0x3d, 0x69, 0x2b, 0x7e,
	0x73, 0x60, 0x47, 0xd7, 0xb8, 0x00, 0xd7, 0x66, 0x0a, 0x9d, 0xd4, 0x37, 0x56, 0xbe, 0x37, 0xdf,
	0xdb, 0x24, 0xec, 0x2a, 0x1f, 0xc1, 0xae, 0x69, 0x15, 0x1d, 0x57, 0x9a, 0x15, 0xa7, 0xfd, 0x93,
	0x0d, 0xdc, 0x5e, 0xbd, 0x00, 0xd7, 0x06, 0xe3, 0xd6, 0xc3, 0xab, 0xe1, 0xb9, 0xf5, 0xf0, 0x5a,
	0x86, 0x86, 0x87, 0x6f, 0x90, 0x90, 0x8c, 0xbf, 0x0f, 0x33, 0xd6, 0x4f, 0x58, 0x9e, 0x33, 0xda,
	0x2f, 0x66, 0xb3, 0x5d, 0x1d, 0x9f, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x00, 0x9c,
	0x20, 0x92, 0x04, 0x00, 0x00,
}
