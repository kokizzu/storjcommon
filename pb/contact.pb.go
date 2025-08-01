// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contact.proto

package pb

import (
	time "time"

	proto "github.com/gogo/protobuf/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LostPieceReason int32

const (
	LostPieceReason_ZERO_HASH     LostPieceReason = 0
	LostPieceReason_HASH_MISMATCH LostPieceReason = 1
)

var LostPieceReason_name = map[int32]string{
	0: "ZERO_HASH",
	1: "HASH_MISMATCH",
}

var LostPieceReason_value = map[string]int32{
	"ZERO_HASH":     0,
	"HASH_MISMATCH": 1,
}

func (x LostPieceReason) String() string {
	return proto.EnumName(LostPieceReason_name, int32(x))
}

type CheckInRequest struct {
	Address  string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Version  *NodeVersion  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Capacity *NodeCapacity `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Operator *NodeOperator `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	// CheckIn may happen over TLS, which means there is not necessarily an
	// opportunity to use NoiseSessionAttestation instead of
	// NoiseKeyAttestation.
	NoiseKeyAttestation *NoiseKeyAttestation `protobuf:"bytes,5,opt,name=noise_key_attestation,json=noiseKeyAttestation,proto3" json:"noise_key_attestation,omitempty"`
	DebounceLimit       int32                `protobuf:"varint,6,opt,name=debounce_limit,json=debounceLimit,proto3" json:"debounce_limit,omitempty"`
	// features is a set of bit flags of the NodeAddress.Features enum.
	Features             uint64             `protobuf:"varint,7,opt,name=features,proto3" json:"features,omitempty"`
	SignedTags           *SignedNodeTagSets `protobuf:"bytes,8,opt,name=signed_tags,json=signedTags,proto3" json:"signed_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CheckInRequest) Reset()         { *m = CheckInRequest{} }
func (m *CheckInRequest) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest) ProtoMessage()    {}

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

func (m *CheckInRequest) GetDebounceLimit() int32 {
	if m != nil {
		return m.DebounceLimit
	}
	return 0
}

func (m *CheckInRequest) GetFeatures() uint64 {
	if m != nil {
		return m.Features
	}
	return 0
}

func (m *CheckInRequest) GetSignedTags() *SignedNodeTagSets {
	if m != nil {
		return m.SignedTags
	}
	return nil
}

type HashstoreSettings struct {
	ActiveMigrate        bool     `protobuf:"varint,1,opt,name=active_migrate,json=activeMigrate,proto3" json:"active_migrate,omitempty"`
	PassiveMigrate       bool     `protobuf:"varint,2,opt,name=passive_migrate,json=passiveMigrate,proto3" json:"passive_migrate,omitempty"`
	WriteToNew           bool     `protobuf:"varint,3,opt,name=write_to_new,json=writeToNew,proto3" json:"write_to_new,omitempty"`
	ReadNewFirst         bool     `protobuf:"varint,4,opt,name=read_new_first,json=readNewFirst,proto3" json:"read_new_first,omitempty"`
	TtlToNew             bool     `protobuf:"varint,5,opt,name=ttl_to_new,json=ttlToNew,proto3" json:"ttl_to_new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashstoreSettings) Reset()         { *m = HashstoreSettings{} }
func (m *HashstoreSettings) String() string { return proto.CompactTextString(m) }
func (*HashstoreSettings) ProtoMessage()    {}

func (m *HashstoreSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashstoreSettings.Unmarshal(m, b)
}
func (m *HashstoreSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashstoreSettings.Marshal(b, m, deterministic)
}
func (m *HashstoreSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashstoreSettings.Merge(m, src)
}
func (m *HashstoreSettings) XXX_Size() int {
	return xxx_messageInfo_HashstoreSettings.Size(m)
}
func (m *HashstoreSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_HashstoreSettings.DiscardUnknown(m)
}

var xxx_messageInfo_HashstoreSettings proto.InternalMessageInfo

func (m *HashstoreSettings) GetActiveMigrate() bool {
	if m != nil {
		return m.ActiveMigrate
	}
	return false
}

func (m *HashstoreSettings) GetPassiveMigrate() bool {
	if m != nil {
		return m.PassiveMigrate
	}
	return false
}

func (m *HashstoreSettings) GetWriteToNew() bool {
	if m != nil {
		return m.WriteToNew
	}
	return false
}

func (m *HashstoreSettings) GetReadNewFirst() bool {
	if m != nil {
		return m.ReadNewFirst
	}
	return false
}

func (m *HashstoreSettings) GetTtlToNew() bool {
	if m != nil {
		return m.TtlToNew
	}
	return false
}

type CheckInResponse struct {
	PingNodeSuccess      bool               `protobuf:"varint,1,opt,name=ping_node_success,json=pingNodeSuccess,proto3" json:"ping_node_success,omitempty"`
	PingErrorMessage     string             `protobuf:"bytes,2,opt,name=ping_error_message,json=pingErrorMessage,proto3" json:"ping_error_message,omitempty"`
	PingNodeSuccessQuic  bool               `protobuf:"varint,3,opt,name=ping_node_success_quic,json=pingNodeSuccessQuic,proto3" json:"ping_node_success_quic,omitempty"`
	NodeTagSuccess       bool               `protobuf:"varint,4,opt,name=node_tag_success,json=nodeTagSuccess,proto3" json:"node_tag_success,omitempty"`
	NodeTagErrorMessage  string             `protobuf:"bytes,5,opt,name=node_tag_error_message,json=nodeTagErrorMessage,proto3" json:"node_tag_error_message,omitempty"`
	HashstoreSettings    *HashstoreSettings `protobuf:"bytes,6,opt,name=hashstore_settings,json=hashstoreSettings,proto3" json:"hashstore_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}

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

func (m *CheckInResponse) GetNodeTagSuccess() bool {
	if m != nil {
		return m.NodeTagSuccess
	}
	return false
}

func (m *CheckInResponse) GetNodeTagErrorMessage() string {
	if m != nil {
		return m.NodeTagErrorMessage
	}
	return ""
}

func (m *CheckInResponse) GetHashstoreSettings() *HashstoreSettings {
	if m != nil {
		return m.HashstoreSettings
	}
	return nil
}

type GetTimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}

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

type AmnestyReportRequest struct {
	LostPieces           []*LostPiece `protobuf:"bytes,1,rep,name=lost_pieces,json=lostPieces,proto3" json:"lost_pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AmnestyReportRequest) Reset()         { *m = AmnestyReportRequest{} }
func (m *AmnestyReportRequest) String() string { return proto.CompactTextString(m) }
func (*AmnestyReportRequest) ProtoMessage()    {}

func (m *AmnestyReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmnestyReportRequest.Unmarshal(m, b)
}
func (m *AmnestyReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmnestyReportRequest.Marshal(b, m, deterministic)
}
func (m *AmnestyReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmnestyReportRequest.Merge(m, src)
}
func (m *AmnestyReportRequest) XXX_Size() int {
	return xxx_messageInfo_AmnestyReportRequest.Size(m)
}
func (m *AmnestyReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AmnestyReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AmnestyReportRequest proto.InternalMessageInfo

func (m *AmnestyReportRequest) GetLostPieces() []*LostPiece {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

type LostPiece struct {
	PieceId              PieceID         `protobuf:"bytes,1,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	Reason               LostPieceReason `protobuf:"varint,2,opt,name=reason,proto3,enum=contact.LostPieceReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LostPiece) Reset()         { *m = LostPiece{} }
func (m *LostPiece) String() string { return proto.CompactTextString(m) }
func (*LostPiece) ProtoMessage()    {}

func (m *LostPiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LostPiece.Unmarshal(m, b)
}
func (m *LostPiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LostPiece.Marshal(b, m, deterministic)
}
func (m *LostPiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LostPiece.Merge(m, src)
}
func (m *LostPiece) XXX_Size() int {
	return xxx_messageInfo_LostPiece.Size(m)
}
func (m *LostPiece) XXX_DiscardUnknown() {
	xxx_messageInfo_LostPiece.DiscardUnknown(m)
}

var xxx_messageInfo_LostPiece proto.InternalMessageInfo

func (m *LostPiece) GetReason() LostPieceReason {
	if m != nil {
		return m.Reason
	}
	return LostPieceReason_ZERO_HASH
}

type AmnestyReportResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmnestyReportResponse) Reset()         { *m = AmnestyReportResponse{} }
func (m *AmnestyReportResponse) String() string { return proto.CompactTextString(m) }
func (*AmnestyReportResponse) ProtoMessage()    {}

func (m *AmnestyReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmnestyReportResponse.Unmarshal(m, b)
}
func (m *AmnestyReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmnestyReportResponse.Marshal(b, m, deterministic)
}
func (m *AmnestyReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmnestyReportResponse.Merge(m, src)
}
func (m *AmnestyReportResponse) XXX_Size() int {
	return xxx_messageInfo_AmnestyReportResponse.Size(m)
}
func (m *AmnestyReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AmnestyReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AmnestyReportResponse proto.InternalMessageInfo
