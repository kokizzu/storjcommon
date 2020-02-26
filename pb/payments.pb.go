// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payments.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"

	drpc "storj.io/drpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PrepareInvoiceRecordsRequest struct {
	Period               time.Time `protobuf:"bytes,1,opt,name=period,proto3,stdtime" json:"period"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PrepareInvoiceRecordsRequest) Reset()         { *m = PrepareInvoiceRecordsRequest{} }
func (m *PrepareInvoiceRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareInvoiceRecordsRequest) ProtoMessage()    {}
func (*PrepareInvoiceRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{0}
}
func (m *PrepareInvoiceRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareInvoiceRecordsRequest.Unmarshal(m, b)
}
func (m *PrepareInvoiceRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareInvoiceRecordsRequest.Marshal(b, m, deterministic)
}
func (m *PrepareInvoiceRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareInvoiceRecordsRequest.Merge(m, src)
}
func (m *PrepareInvoiceRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareInvoiceRecordsRequest.Size(m)
}
func (m *PrepareInvoiceRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareInvoiceRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareInvoiceRecordsRequest proto.InternalMessageInfo

func (m *PrepareInvoiceRecordsRequest) GetPeriod() time.Time {
	if m != nil {
		return m.Period
	}
	return time.Time{}
}

type PrepareInvoiceRecordsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareInvoiceRecordsResponse) Reset()         { *m = PrepareInvoiceRecordsResponse{} }
func (m *PrepareInvoiceRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareInvoiceRecordsResponse) ProtoMessage()    {}
func (*PrepareInvoiceRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{1}
}
func (m *PrepareInvoiceRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareInvoiceRecordsResponse.Unmarshal(m, b)
}
func (m *PrepareInvoiceRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareInvoiceRecordsResponse.Marshal(b, m, deterministic)
}
func (m *PrepareInvoiceRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareInvoiceRecordsResponse.Merge(m, src)
}
func (m *PrepareInvoiceRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareInvoiceRecordsResponse.Size(m)
}
func (m *PrepareInvoiceRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareInvoiceRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareInvoiceRecordsResponse proto.InternalMessageInfo

type ApplyInvoiceRecordsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceRecordsRequest) Reset()         { *m = ApplyInvoiceRecordsRequest{} }
func (m *ApplyInvoiceRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceRecordsRequest) ProtoMessage()    {}
func (*ApplyInvoiceRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{2}
}
func (m *ApplyInvoiceRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceRecordsRequest.Unmarshal(m, b)
}
func (m *ApplyInvoiceRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceRecordsRequest.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceRecordsRequest.Merge(m, src)
}
func (m *ApplyInvoiceRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceRecordsRequest.Size(m)
}
func (m *ApplyInvoiceRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceRecordsRequest proto.InternalMessageInfo

type ApplyInvoiceRecordsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceRecordsResponse) Reset()         { *m = ApplyInvoiceRecordsResponse{} }
func (m *ApplyInvoiceRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceRecordsResponse) ProtoMessage()    {}
func (*ApplyInvoiceRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{3}
}
func (m *ApplyInvoiceRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceRecordsResponse.Unmarshal(m, b)
}
func (m *ApplyInvoiceRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceRecordsResponse.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceRecordsResponse.Merge(m, src)
}
func (m *ApplyInvoiceRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceRecordsResponse.Size(m)
}
func (m *ApplyInvoiceRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceRecordsResponse proto.InternalMessageInfo

type ApplyInvoiceCouponsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceCouponsRequest) Reset()         { *m = ApplyInvoiceCouponsRequest{} }
func (m *ApplyInvoiceCouponsRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceCouponsRequest) ProtoMessage()    {}
func (*ApplyInvoiceCouponsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{4}
}
func (m *ApplyInvoiceCouponsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceCouponsRequest.Unmarshal(m, b)
}
func (m *ApplyInvoiceCouponsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceCouponsRequest.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceCouponsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceCouponsRequest.Merge(m, src)
}
func (m *ApplyInvoiceCouponsRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceCouponsRequest.Size(m)
}
func (m *ApplyInvoiceCouponsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceCouponsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceCouponsRequest proto.InternalMessageInfo

type ApplyInvoiceCouponsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceCouponsResponse) Reset()         { *m = ApplyInvoiceCouponsResponse{} }
func (m *ApplyInvoiceCouponsResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceCouponsResponse) ProtoMessage()    {}
func (*ApplyInvoiceCouponsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{5}
}
func (m *ApplyInvoiceCouponsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceCouponsResponse.Unmarshal(m, b)
}
func (m *ApplyInvoiceCouponsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceCouponsResponse.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceCouponsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceCouponsResponse.Merge(m, src)
}
func (m *ApplyInvoiceCouponsResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceCouponsResponse.Size(m)
}
func (m *ApplyInvoiceCouponsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceCouponsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceCouponsResponse proto.InternalMessageInfo

type ApplyInvoiceCreditsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceCreditsRequest) Reset()         { *m = ApplyInvoiceCreditsRequest{} }
func (m *ApplyInvoiceCreditsRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceCreditsRequest) ProtoMessage()    {}
func (*ApplyInvoiceCreditsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{6}
}
func (m *ApplyInvoiceCreditsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceCreditsRequest.Unmarshal(m, b)
}
func (m *ApplyInvoiceCreditsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceCreditsRequest.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceCreditsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceCreditsRequest.Merge(m, src)
}
func (m *ApplyInvoiceCreditsRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceCreditsRequest.Size(m)
}
func (m *ApplyInvoiceCreditsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceCreditsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceCreditsRequest proto.InternalMessageInfo

type ApplyInvoiceCreditsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceCreditsResponse) Reset()         { *m = ApplyInvoiceCreditsResponse{} }
func (m *ApplyInvoiceCreditsResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceCreditsResponse) ProtoMessage()    {}
func (*ApplyInvoiceCreditsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{7}
}
func (m *ApplyInvoiceCreditsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyInvoiceCreditsResponse.Unmarshal(m, b)
}
func (m *ApplyInvoiceCreditsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyInvoiceCreditsResponse.Marshal(b, m, deterministic)
}
func (m *ApplyInvoiceCreditsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyInvoiceCreditsResponse.Merge(m, src)
}
func (m *ApplyInvoiceCreditsResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyInvoiceCreditsResponse.Size(m)
}
func (m *ApplyInvoiceCreditsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyInvoiceCreditsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyInvoiceCreditsResponse proto.InternalMessageInfo

type CreateInvoicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInvoicesRequest) Reset()         { *m = CreateInvoicesRequest{} }
func (m *CreateInvoicesRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInvoicesRequest) ProtoMessage()    {}
func (*CreateInvoicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{8}
}
func (m *CreateInvoicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInvoicesRequest.Unmarshal(m, b)
}
func (m *CreateInvoicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInvoicesRequest.Marshal(b, m, deterministic)
}
func (m *CreateInvoicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInvoicesRequest.Merge(m, src)
}
func (m *CreateInvoicesRequest) XXX_Size() int {
	return xxx_messageInfo_CreateInvoicesRequest.Size(m)
}
func (m *CreateInvoicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInvoicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInvoicesRequest proto.InternalMessageInfo

type CreateInvoicesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateInvoicesResponse) Reset()         { *m = CreateInvoicesResponse{} }
func (m *CreateInvoicesResponse) String() string { return proto.CompactTextString(m) }
func (*CreateInvoicesResponse) ProtoMessage()    {}
func (*CreateInvoicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{9}
}
func (m *CreateInvoicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateInvoicesResponse.Unmarshal(m, b)
}
func (m *CreateInvoicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateInvoicesResponse.Marshal(b, m, deterministic)
}
func (m *CreateInvoicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateInvoicesResponse.Merge(m, src)
}
func (m *CreateInvoicesResponse) XXX_Size() int {
	return xxx_messageInfo_CreateInvoicesResponse.Size(m)
}
func (m *CreateInvoicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateInvoicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateInvoicesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PrepareInvoiceRecordsRequest)(nil), "nodestats.PrepareInvoiceRecordsRequest")
	proto.RegisterType((*PrepareInvoiceRecordsResponse)(nil), "nodestats.PrepareInvoiceRecordsResponse")
	proto.RegisterType((*ApplyInvoiceRecordsRequest)(nil), "nodestats.ApplyInvoiceRecordsRequest")
	proto.RegisterType((*ApplyInvoiceRecordsResponse)(nil), "nodestats.ApplyInvoiceRecordsResponse")
	proto.RegisterType((*ApplyInvoiceCouponsRequest)(nil), "nodestats.ApplyInvoiceCouponsRequest")
	proto.RegisterType((*ApplyInvoiceCouponsResponse)(nil), "nodestats.ApplyInvoiceCouponsResponse")
	proto.RegisterType((*ApplyInvoiceCreditsRequest)(nil), "nodestats.ApplyInvoiceCreditsRequest")
	proto.RegisterType((*ApplyInvoiceCreditsResponse)(nil), "nodestats.ApplyInvoiceCreditsResponse")
	proto.RegisterType((*CreateInvoicesRequest)(nil), "nodestats.CreateInvoicesRequest")
	proto.RegisterType((*CreateInvoicesResponse)(nil), "nodestats.CreateInvoicesResponse")
}

func init() { proto.RegisterFile("payments.proto", fileDescriptor_a9566e6e864d2854) }

var fileDescriptor_a9566e6e864d2854 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0x7f, 0xfc, 0x34, 0x04, 0xc7, 0x84, 0xc3, 0x1a, 0x94, 0xac, 0x90, 0x62, 0x13, 0x95,
	0x53, 0x49, 0xf0, 0xea, 0x45, 0x38, 0x79, 0x23, 0x44, 0x2f, 0xc6, 0x4b, 0xa1, 0x63, 0x53, 0x03,
	0x9d, 0x75, 0x77, 0x31, 0xe1, 0x2d, 0x7c, 0x2c, 0x9f, 0x42, 0x1f, 0x45, 0x03, 0xdb, 0x36, 0x14,
	0x77, 0xe1, 0xd8, 0x99, 0xef, 0x9f, 0x74, 0x3e, 0x59, 0xa8, 0x8b, 0x70, 0x39, 0xc7, 0x54, 0xab,
	0x40, 0x48, 0xd2, 0xc4, 0x8e, 0x52, 0x8a, 0x50, 0xe9, 0x50, 0x2b, 0x0e, 0x31, 0xc5, 0x64, 0xc6,
	0xdc, 0x8b, 0x89, 0xe2, 0x19, 0xf6, 0xd6, 0x5f, 0x93, 0xc5, 0x4b, 0x4f, 0x27, 0xf3, 0x95, 0x6c,
	0x2e, 0x8c, 0xc0, 0x7f, 0x86, 0xd6, 0x48, 0xa2, 0x08, 0x25, 0xde, 0xa7, 0xef, 0x94, 0x4c, 0x71,
	0x8c, 0x53, 0x92, 0x91, 0x1a, 0xe3, 0xdb, 0x02, 0x95, 0x66, 0xb7, 0x50, 0x15, 0x28, 0x13, 0x8a,
	0x9a, 0x95, 0x4e, 0xa5, 0x7b, 0xdc, 0xe7, 0x81, 0x49, 0x0c, 0xf2, 0xc4, 0xe0, 0x21, 0x4f, 0x1c,
	0xd4, 0x3e, 0xbf, 0xbc, 0x7f, 0x1f, 0xdf, 0x5e, 0x65, 0x9c, 0x79, 0x7c, 0x0f, 0xda, 0x8e, 0x74,
	0x25, 0x28, 0x55, 0xe8, 0xb7, 0x80, 0xdf, 0x09, 0x31, 0x5b, 0x5a, 0xcb, 0xfd, 0x36, 0x9c, 0x5b,
	0xb7, 0x76, 0xf3, 0x90, 0x16, 0xab, 0xb9, 0xc3, 0x5c, 0x6c, 0x1d, 0x66, 0x89, 0x51, 0xa2, 0x9d,
	0xe6, 0x7c, 0x9b, 0x99, 0xcf, 0xa0, 0x31, 0x94, 0x18, 0xea, 0xfc, 0xb7, 0x0a, 0x5f, 0x13, 0x4e,
	0xb7, 0x17, 0xc6, 0xd2, 0xff, 0x39, 0x80, 0xda, 0x28, 0x63, 0xc6, 0x5e, 0xa1, 0x61, 0xbd, 0x0b,
	0xbb, 0x0e, 0x0a, 0x8e, 0xc1, 0x2e, 0x2e, 0xbc, 0xbb, 0x5f, 0x68, 0x8a, 0x59, 0x04, 0x27, 0x96,
	0x23, 0xb2, 0xcb, 0x8d, 0x00, 0x37, 0x02, 0x7e, 0xb5, 0x4f, 0x66, 0x6f, 0xc9, 0xae, 0xed, 0x6c,
	0x29, 0xb3, 0x72, 0xb6, 0x6c, 0x41, 0xfb, 0xd3, 0x62, 0xb0, 0xb8, 0x5b, 0x4a, 0x50, 0xdd, 0x2d,
	0x65, 0xba, 0xec, 0x11, 0xea, 0x65, 0x88, 0xac, 0xb3, 0xe1, 0xb4, 0x82, 0xe7, 0x17, 0x3b, 0x14,
	0x26, 0x76, 0x70, 0xf8, 0xf4, 0x5f, 0x4c, 0x26, 0xd5, 0xf5, 0xc3, 0xb9, 0xf9, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0xdc, 0xb8, 0xac, 0xc1, 0x03, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCPaymentsClient interface {
	DRPCConn() drpc.Conn

	PrepareInvoiceRecords(ctx context.Context, in *PrepareInvoiceRecordsRequest) (*PrepareInvoiceRecordsResponse, error)
	ApplyInvoiceRecords(ctx context.Context, in *ApplyInvoiceRecordsRequest) (*ApplyInvoiceRecordsResponse, error)
	ApplyInvoiceCoupons(ctx context.Context, in *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error)
	ApplyInvoiceCredits(ctx context.Context, in *ApplyInvoiceCreditsRequest) (*ApplyInvoiceCreditsResponse, error)
	CreateInvoices(ctx context.Context, in *CreateInvoicesRequest) (*CreateInvoicesResponse, error)
}

type drpcPaymentsClient struct {
	cc drpc.Conn
}

func NewDRPCPaymentsClient(cc drpc.Conn) DRPCPaymentsClient {
	return &drpcPaymentsClient{cc}
}

func (c *drpcPaymentsClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcPaymentsClient) PrepareInvoiceRecords(ctx context.Context, in *PrepareInvoiceRecordsRequest) (*PrepareInvoiceRecordsResponse, error) {
	out := new(PrepareInvoiceRecordsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/PrepareInvoiceRecords", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPaymentsClient) ApplyInvoiceRecords(ctx context.Context, in *ApplyInvoiceRecordsRequest) (*ApplyInvoiceRecordsResponse, error) {
	out := new(ApplyInvoiceRecordsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/ApplyInvoiceRecords", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPaymentsClient) ApplyInvoiceCoupons(ctx context.Context, in *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error) {
	out := new(ApplyInvoiceCouponsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/ApplyInvoiceCoupons", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPaymentsClient) ApplyInvoiceCredits(ctx context.Context, in *ApplyInvoiceCreditsRequest) (*ApplyInvoiceCreditsResponse, error) {
	out := new(ApplyInvoiceCreditsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/ApplyInvoiceCredits", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcPaymentsClient) CreateInvoices(ctx context.Context, in *CreateInvoicesRequest) (*CreateInvoicesResponse, error) {
	out := new(CreateInvoicesResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/CreateInvoices", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCPaymentsServer interface {
	PrepareInvoiceRecords(context.Context, *PrepareInvoiceRecordsRequest) (*PrepareInvoiceRecordsResponse, error)
	ApplyInvoiceRecords(context.Context, *ApplyInvoiceRecordsRequest) (*ApplyInvoiceRecordsResponse, error)
	ApplyInvoiceCoupons(context.Context, *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error)
	ApplyInvoiceCredits(context.Context, *ApplyInvoiceCreditsRequest) (*ApplyInvoiceCreditsResponse, error)
	CreateInvoices(context.Context, *CreateInvoicesRequest) (*CreateInvoicesResponse, error)
}

type DRPCPaymentsDescription struct{}

func (DRPCPaymentsDescription) NumMethods() int { return 5 }

func (DRPCPaymentsDescription) Method(n int) (string, drpc.Handler, interface{}, bool) {
	switch n {
	case 0:
		return "/nodestats.Payments/PrepareInvoiceRecords",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					PrepareInvoiceRecords(
						ctx,
						in1.(*PrepareInvoiceRecordsRequest),
					)
			}, DRPCPaymentsServer.PrepareInvoiceRecords, true
	case 1:
		return "/nodestats.Payments/ApplyInvoiceRecords",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					ApplyInvoiceRecords(
						ctx,
						in1.(*ApplyInvoiceRecordsRequest),
					)
			}, DRPCPaymentsServer.ApplyInvoiceRecords, true
	case 2:
		return "/nodestats.Payments/ApplyInvoiceCoupons",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					ApplyInvoiceCoupons(
						ctx,
						in1.(*ApplyInvoiceCouponsRequest),
					)
			}, DRPCPaymentsServer.ApplyInvoiceCoupons, true
	case 3:
		return "/nodestats.Payments/ApplyInvoiceCredits",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					ApplyInvoiceCredits(
						ctx,
						in1.(*ApplyInvoiceCreditsRequest),
					)
			}, DRPCPaymentsServer.ApplyInvoiceCredits, true
	case 4:
		return "/nodestats.Payments/CreateInvoices",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					CreateInvoices(
						ctx,
						in1.(*CreateInvoicesRequest),
					)
			}, DRPCPaymentsServer.CreateInvoices, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterPayments(srv drpc.Server, impl DRPCPaymentsServer) {
	srv.Register(impl, DRPCPaymentsDescription{})
}

type DRPCPayments_PrepareInvoiceRecordsStream interface {
	drpc.Stream
	SendAndClose(*PrepareInvoiceRecordsResponse) error
}

type drpcPaymentsPrepareInvoiceRecordsStream struct {
	drpc.Stream
}

func (x *drpcPaymentsPrepareInvoiceRecordsStream) SendAndClose(m *PrepareInvoiceRecordsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPayments_ApplyInvoiceRecordsStream interface {
	drpc.Stream
	SendAndClose(*ApplyInvoiceRecordsResponse) error
}

type drpcPaymentsApplyInvoiceRecordsStream struct {
	drpc.Stream
}

func (x *drpcPaymentsApplyInvoiceRecordsStream) SendAndClose(m *ApplyInvoiceRecordsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPayments_ApplyInvoiceCouponsStream interface {
	drpc.Stream
	SendAndClose(*ApplyInvoiceCouponsResponse) error
}

type drpcPaymentsApplyInvoiceCouponsStream struct {
	drpc.Stream
}

func (x *drpcPaymentsApplyInvoiceCouponsStream) SendAndClose(m *ApplyInvoiceCouponsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPayments_ApplyInvoiceCreditsStream interface {
	drpc.Stream
	SendAndClose(*ApplyInvoiceCreditsResponse) error
}

type drpcPaymentsApplyInvoiceCreditsStream struct {
	drpc.Stream
}

func (x *drpcPaymentsApplyInvoiceCreditsStream) SendAndClose(m *ApplyInvoiceCreditsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCPayments_CreateInvoicesStream interface {
	drpc.Stream
	SendAndClose(*CreateInvoicesResponse) error
}

type drpcPaymentsCreateInvoicesStream struct {
	drpc.Stream
}

func (x *drpcPaymentsCreateInvoicesStream) SendAndClose(m *CreateInvoicesResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
