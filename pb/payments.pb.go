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
	grpc "google.golang.org/grpc"

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

type PrepareInvoiceCouponsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareInvoiceCouponsRequest) Reset()         { *m = PrepareInvoiceCouponsRequest{} }
func (m *PrepareInvoiceCouponsRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareInvoiceCouponsRequest) ProtoMessage()    {}
func (*PrepareInvoiceCouponsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{4}
}
func (m *PrepareInvoiceCouponsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareInvoiceCouponsRequest.Unmarshal(m, b)
}
func (m *PrepareInvoiceCouponsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareInvoiceCouponsRequest.Marshal(b, m, deterministic)
}
func (m *PrepareInvoiceCouponsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareInvoiceCouponsRequest.Merge(m, src)
}
func (m *PrepareInvoiceCouponsRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareInvoiceCouponsRequest.Size(m)
}
func (m *PrepareInvoiceCouponsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareInvoiceCouponsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareInvoiceCouponsRequest proto.InternalMessageInfo

type PrepareInvoiceCouponsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareInvoiceCouponsResponse) Reset()         { *m = PrepareInvoiceCouponsResponse{} }
func (m *PrepareInvoiceCouponsResponse) String() string { return proto.CompactTextString(m) }
func (*PrepareInvoiceCouponsResponse) ProtoMessage()    {}
func (*PrepareInvoiceCouponsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{5}
}
func (m *PrepareInvoiceCouponsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareInvoiceCouponsResponse.Unmarshal(m, b)
}
func (m *PrepareInvoiceCouponsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareInvoiceCouponsResponse.Marshal(b, m, deterministic)
}
func (m *PrepareInvoiceCouponsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareInvoiceCouponsResponse.Merge(m, src)
}
func (m *PrepareInvoiceCouponsResponse) XXX_Size() int {
	return xxx_messageInfo_PrepareInvoiceCouponsResponse.Size(m)
}
func (m *PrepareInvoiceCouponsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareInvoiceCouponsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareInvoiceCouponsResponse proto.InternalMessageInfo

type ApplyInvoiceCouponsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyInvoiceCouponsRequest) Reset()         { *m = ApplyInvoiceCouponsRequest{} }
func (m *ApplyInvoiceCouponsRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyInvoiceCouponsRequest) ProtoMessage()    {}
func (*ApplyInvoiceCouponsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9566e6e864d2854, []int{6}
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
	return fileDescriptor_a9566e6e864d2854, []int{7}
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
	proto.RegisterType((*PrepareInvoiceCouponsRequest)(nil), "nodestats.PrepareInvoiceCouponsRequest")
	proto.RegisterType((*PrepareInvoiceCouponsResponse)(nil), "nodestats.PrepareInvoiceCouponsResponse")
	proto.RegisterType((*ApplyInvoiceCouponsRequest)(nil), "nodestats.ApplyInvoiceCouponsRequest")
	proto.RegisterType((*ApplyInvoiceCouponsResponse)(nil), "nodestats.ApplyInvoiceCouponsResponse")
	proto.RegisterType((*CreateInvoicesRequest)(nil), "nodestats.CreateInvoicesRequest")
	proto.RegisterType((*CreateInvoicesResponse)(nil), "nodestats.CreateInvoicesResponse")
}

func init() { proto.RegisterFile("payments.proto", fileDescriptor_a9566e6e864d2854) }

var fileDescriptor_a9566e6e864d2854 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0x15, 0x43, 0x08, 0x5e, 0x13, 0x16, 0x63, 0x50, 0x32, 0x82, 0xc5, 0x26, 0x22, 0xab, 0x92,
	0xe0, 0xd6, 0x8d, 0xb0, 0x72, 0x47, 0x88, 0x6e, 0x8c, 0x9b, 0x42, 0xaf, 0x0d, 0x06, 0x7a, 0xc7,
	0xce, 0x60, 0xc2, 0x07, 0xb8, 0xf7, 0xb3, 0xfc, 0x0a, 0xfd, 0x15, 0x03, 0x33, 0x25, 0x30, 0x0c,
	0x65, 0xd9, 0xde, 0xf3, 0xe8, 0x39, 0x27, 0x85, 0x8a, 0x08, 0x17, 0x33, 0x4c, 0x94, 0x0c, 0x44,
	0x4a, 0x8a, 0xd8, 0x49, 0x42, 0x11, 0x4a, 0x15, 0x2a, 0xc9, 0x21, 0xa6, 0x98, 0xf4, 0x6b, 0xee,
	0xc5, 0x44, 0xf1, 0x14, 0x3b, 0xab, 0xa7, 0xd1, 0xfc, 0xad, 0xa3, 0x26, 0xb3, 0x25, 0x6c, 0x26,
	0x34, 0xc0, 0x7f, 0x85, 0xfa, 0x20, 0x45, 0x11, 0xa6, 0xf8, 0x98, 0x7c, 0xd2, 0x64, 0x8c, 0x43,
	0x1c, 0x53, 0x1a, 0xc9, 0x21, 0x7e, 0xcc, 0x51, 0x2a, 0x76, 0x0f, 0x25, 0x81, 0xe9, 0x84, 0xa2,
	0x5a, 0xa1, 0x59, 0x68, 0x9f, 0x76, 0x79, 0xa0, 0x15, 0x83, 0x4c, 0x31, 0x78, 0xca, 0x14, 0x7b,
	0xe5, 0x9f, 0x5f, 0xef, 0xe8, 0xfb, 0xcf, 0x2b, 0x0c, 0x0d, 0xc7, 0xf7, 0xa0, 0xb1, 0x47, 0x5d,
	0x0a, 0x4a, 0x24, 0xfa, 0x75, 0xe0, 0x0f, 0x42, 0x4c, 0x17, 0x4e, 0x73, 0xbf, 0x01, 0x97, 0xce,
	0xab, 0x21, 0x5f, 0xd9, 0xdf, 0xde, 0xa7, 0xf9, 0xf2, 0x92, 0xd1, 0x77, 0xdc, 0xd7, 0x77, 0xb7,
	0xbb, 0x45, 0xb7, 0xdc, 0x6d, 0xf2, 0x05, 0x54, 0xfb, 0x29, 0x86, 0x2a, 0x13, 0x5f, 0xf3, 0x6a,
	0x70, 0x6e, 0x1f, 0x34, 0xa5, 0xfb, 0x55, 0x84, 0xf2, 0xc0, 0xec, 0xc6, 0xde, 0xa1, 0xea, 0xec,
	0x86, 0xdd, 0x06, 0xeb, 0x2d, 0x83, 0xbc, 0x6d, 0x78, 0xfb, 0x30, 0x50, 0x1b, 0xb3, 0x08, 0xce,
	0x1c, 0x45, 0xb2, 0x9b, 0x0d, 0x81, 0xfd, 0x33, 0xf0, 0xd6, 0x21, 0x98, 0x71, 0xd9, 0x49, 0x64,
	0x2a, 0xcb, 0x49, 0xb4, 0x5d, 0x79, 0x4e, 0x22, 0xab, 0x7d, 0x3b, 0x51, 0xe6, 0xb4, 0x2f, 0x91,
	0xe5, 0xd3, 0x3a, 0x04, 0x33, 0x2e, 0xcf, 0x50, 0xd9, 0x9e, 0x92, 0x35, 0x37, 0x98, 0xce, 0xf9,
	0xf9, 0x75, 0x0e, 0x42, 0xcb, 0xf6, 0x8a, 0x2f, 0xc7, 0x62, 0x34, 0x2a, 0xad, 0x7e, 0xa1, 0xbb,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xde, 0xaa, 0x1d, 0xcb, 0x03, 0x00, 0x00,
}

type DRPCPaymentsClient interface {
	DRPCConn() drpc.Conn

	PrepareInvoiceRecords(ctx context.Context, in *PrepareInvoiceRecordsRequest) (*PrepareInvoiceRecordsResponse, error)
	ApplyInvoiceRecords(ctx context.Context, in *ApplyInvoiceRecordsRequest) (*ApplyInvoiceRecordsResponse, error)
	PrepareInvoiceCoupons(ctx context.Context, in *PrepareInvoiceCouponsRequest) (*PrepareInvoiceCouponsResponse, error)
	ApplyInvoiceCoupons(ctx context.Context, in *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error)
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

func (c *drpcPaymentsClient) PrepareInvoiceCoupons(ctx context.Context, in *PrepareInvoiceCouponsRequest) (*PrepareInvoiceCouponsResponse, error) {
	out := new(PrepareInvoiceCouponsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/PrepareInvoiceCoupons", in, out)
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
	PrepareInvoiceCoupons(context.Context, *PrepareInvoiceCouponsRequest) (*PrepareInvoiceCouponsResponse, error)
	ApplyInvoiceCoupons(context.Context, *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error)
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
		return "/nodestats.Payments/PrepareInvoiceCoupons",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					PrepareInvoiceCoupons(
						ctx,
						in1.(*PrepareInvoiceCouponsRequest),
					)
			}, DRPCPaymentsServer.PrepareInvoiceCoupons, true
	case 3:
		return "/nodestats.Payments/ApplyInvoiceCoupons",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCPaymentsServer).
					ApplyInvoiceCoupons(
						ctx,
						in1.(*ApplyInvoiceCouponsRequest),
					)
			}, DRPCPaymentsServer.ApplyInvoiceCoupons, true
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

type DRPCPayments_PrepareInvoiceCouponsStream interface {
	drpc.Stream
	SendAndClose(*PrepareInvoiceCouponsResponse) error
}

type drpcPaymentsPrepareInvoiceCouponsStream struct {
	drpc.Stream
}

func (x *drpcPaymentsPrepareInvoiceCouponsStream) SendAndClose(m *PrepareInvoiceCouponsResponse) error {
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentsClient is the client API for Payments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsClient interface {
	PrepareInvoiceRecords(ctx context.Context, in *PrepareInvoiceRecordsRequest, opts ...grpc.CallOption) (*PrepareInvoiceRecordsResponse, error)
	ApplyInvoiceRecords(ctx context.Context, in *ApplyInvoiceRecordsRequest, opts ...grpc.CallOption) (*ApplyInvoiceRecordsResponse, error)
	PrepareInvoiceCoupons(ctx context.Context, in *PrepareInvoiceCouponsRequest, opts ...grpc.CallOption) (*PrepareInvoiceCouponsResponse, error)
	ApplyInvoiceCoupons(ctx context.Context, in *ApplyInvoiceCouponsRequest, opts ...grpc.CallOption) (*ApplyInvoiceCouponsResponse, error)
	CreateInvoices(ctx context.Context, in *CreateInvoicesRequest, opts ...grpc.CallOption) (*CreateInvoicesResponse, error)
}

type paymentsClient struct {
	cc *grpc.ClientConn
}

func NewPaymentsClient(cc *grpc.ClientConn) PaymentsClient {
	return &paymentsClient{cc}
}

func (c *paymentsClient) PrepareInvoiceRecords(ctx context.Context, in *PrepareInvoiceRecordsRequest, opts ...grpc.CallOption) (*PrepareInvoiceRecordsResponse, error) {
	out := new(PrepareInvoiceRecordsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/PrepareInvoiceRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) ApplyInvoiceRecords(ctx context.Context, in *ApplyInvoiceRecordsRequest, opts ...grpc.CallOption) (*ApplyInvoiceRecordsResponse, error) {
	out := new(ApplyInvoiceRecordsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/ApplyInvoiceRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) PrepareInvoiceCoupons(ctx context.Context, in *PrepareInvoiceCouponsRequest, opts ...grpc.CallOption) (*PrepareInvoiceCouponsResponse, error) {
	out := new(PrepareInvoiceCouponsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/PrepareInvoiceCoupons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) ApplyInvoiceCoupons(ctx context.Context, in *ApplyInvoiceCouponsRequest, opts ...grpc.CallOption) (*ApplyInvoiceCouponsResponse, error) {
	out := new(ApplyInvoiceCouponsResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/ApplyInvoiceCoupons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentsClient) CreateInvoices(ctx context.Context, in *CreateInvoicesRequest, opts ...grpc.CallOption) (*CreateInvoicesResponse, error) {
	out := new(CreateInvoicesResponse)
	err := c.cc.Invoke(ctx, "/nodestats.Payments/CreateInvoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServer is the server API for Payments service.
type PaymentsServer interface {
	PrepareInvoiceRecords(context.Context, *PrepareInvoiceRecordsRequest) (*PrepareInvoiceRecordsResponse, error)
	ApplyInvoiceRecords(context.Context, *ApplyInvoiceRecordsRequest) (*ApplyInvoiceRecordsResponse, error)
	PrepareInvoiceCoupons(context.Context, *PrepareInvoiceCouponsRequest) (*PrepareInvoiceCouponsResponse, error)
	ApplyInvoiceCoupons(context.Context, *ApplyInvoiceCouponsRequest) (*ApplyInvoiceCouponsResponse, error)
	CreateInvoices(context.Context, *CreateInvoicesRequest) (*CreateInvoicesResponse, error)
}

func RegisterPaymentsServer(s *grpc.Server, srv PaymentsServer) {
	s.RegisterService(&_Payments_serviceDesc, srv)
}

func _Payments_PrepareInvoiceRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareInvoiceRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).PrepareInvoiceRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.Payments/PrepareInvoiceRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).PrepareInvoiceRecords(ctx, req.(*PrepareInvoiceRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_ApplyInvoiceRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyInvoiceRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).ApplyInvoiceRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.Payments/ApplyInvoiceRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).ApplyInvoiceRecords(ctx, req.(*ApplyInvoiceRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_PrepareInvoiceCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareInvoiceCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).PrepareInvoiceCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.Payments/PrepareInvoiceCoupons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).PrepareInvoiceCoupons(ctx, req.(*PrepareInvoiceCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_ApplyInvoiceCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyInvoiceCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).ApplyInvoiceCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.Payments/ApplyInvoiceCoupons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).ApplyInvoiceCoupons(ctx, req.(*ApplyInvoiceCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payments_CreateInvoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInvoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).CreateInvoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.Payments/CreateInvoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).CreateInvoices(ctx, req.(*CreateInvoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Payments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodestats.Payments",
	HandlerType: (*PaymentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareInvoiceRecords",
			Handler:    _Payments_PrepareInvoiceRecords_Handler,
		},
		{
			MethodName: "ApplyInvoiceRecords",
			Handler:    _Payments_ApplyInvoiceRecords_Handler,
		},
		{
			MethodName: "PrepareInvoiceCoupons",
			Handler:    _Payments_PrepareInvoiceCoupons_Handler,
		},
		{
			MethodName: "ApplyInvoiceCoupons",
			Handler:    _Payments_ApplyInvoiceCoupons_Handler,
		},
		{
			MethodName: "CreateInvoices",
			Handler:    _Payments_CreateInvoices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payments.proto",
}
