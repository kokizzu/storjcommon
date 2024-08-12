// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gracefulexit.proto

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

type TransferFailed_Error int32

const (
	TransferFailed_NOT_FOUND                TransferFailed_Error = 0
	TransferFailed_STORAGE_NODE_UNAVAILABLE TransferFailed_Error = 1
	TransferFailed_HASH_VERIFICATION        TransferFailed_Error = 2
	TransferFailed_UNKNOWN                  TransferFailed_Error = 10
)

var TransferFailed_Error_name = map[int32]string{
	0:  "NOT_FOUND",
	1:  "STORAGE_NODE_UNAVAILABLE",
	2:  "HASH_VERIFICATION",
	10: "UNKNOWN",
}

var TransferFailed_Error_value = map[string]int32{
	"NOT_FOUND":                0,
	"STORAGE_NODE_UNAVAILABLE": 1,
	"HASH_VERIFICATION":        2,
	"UNKNOWN":                  10,
}

func (x TransferFailed_Error) String() string {
	return proto.EnumName(TransferFailed_Error_name, int32(x))
}

type ExitFailed_Reason int32

const (
	ExitFailed_VERIFICATION_FAILED                 ExitFailed_Reason = 0
	ExitFailed_INACTIVE_TIMEFRAME_EXCEEDED         ExitFailed_Reason = 1
	ExitFailed_OVERALL_FAILURE_PERCENTAGE_EXCEEDED ExitFailed_Reason = 2
)

var ExitFailed_Reason_name = map[int32]string{
	0: "VERIFICATION_FAILED",
	1: "INACTIVE_TIMEFRAME_EXCEEDED",
	2: "OVERALL_FAILURE_PERCENTAGE_EXCEEDED",
}

var ExitFailed_Reason_value = map[string]int32{
	"VERIFICATION_FAILED":                 0,
	"INACTIVE_TIMEFRAME_EXCEEDED":         1,
	"OVERALL_FAILURE_PERCENTAGE_EXCEEDED": 2,
}

func (x ExitFailed_Reason) String() string {
	return proto.EnumName(ExitFailed_Reason_name, int32(x))
}

type GracefulExitFeasibilityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GracefulExitFeasibilityRequest) Reset()         { *m = GracefulExitFeasibilityRequest{} }
func (m *GracefulExitFeasibilityRequest) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityRequest) ProtoMessage()    {}

func (m *GracefulExitFeasibilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityRequest.Merge(m, src)
}
func (m *GracefulExitFeasibilityRequest) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Size(m)
}
func (m *GracefulExitFeasibilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityRequest proto.InternalMessageInfo

type GracefulExitFeasibilityResponse struct {
	JoinedAt             time.Time `protobuf:"bytes,1,opt,name=joined_at,json=joinedAt,proto3,stdtime" json:"joined_at"`
	MonthsRequired       int32     `protobuf:"varint,2,opt,name=months_required,json=monthsRequired,proto3" json:"months_required,omitempty"`
	IsAllowed            bool      `protobuf:"varint,3,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GracefulExitFeasibilityResponse) Reset()         { *m = GracefulExitFeasibilityResponse{} }
func (m *GracefulExitFeasibilityResponse) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityResponse) ProtoMessage()    {}

func (m *GracefulExitFeasibilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityResponse.Merge(m, src)
}
func (m *GracefulExitFeasibilityResponse) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Size(m)
}
func (m *GracefulExitFeasibilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityResponse proto.InternalMessageInfo

func (m *GracefulExitFeasibilityResponse) GetJoinedAt() time.Time {
	if m != nil {
		return m.JoinedAt
	}
	return time.Time{}
}

func (m *GracefulExitFeasibilityResponse) GetMonthsRequired() int32 {
	if m != nil {
		return m.MonthsRequired
	}
	return 0
}

func (m *GracefulExitFeasibilityResponse) GetIsAllowed() bool {
	if m != nil {
		return m.IsAllowed
	}
	return false
}

type TransferSucceeded struct {
	OriginalOrderLimit   *OrderLimit `protobuf:"bytes,1,opt,name=original_order_limit,json=originalOrderLimit,proto3" json:"original_order_limit,omitempty"`
	OriginalPieceHash    *PieceHash  `protobuf:"bytes,2,opt,name=original_piece_hash,json=originalPieceHash,proto3" json:"original_piece_hash,omitempty"`
	ReplacementPieceHash *PieceHash  `protobuf:"bytes,3,opt,name=replacement_piece_hash,json=replacementPieceHash,proto3" json:"replacement_piece_hash,omitempty"`
	OriginalPieceId      PieceID     `protobuf:"bytes,4,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransferSucceeded) Reset()         { *m = TransferSucceeded{} }
func (m *TransferSucceeded) String() string { return proto.CompactTextString(m) }
func (*TransferSucceeded) ProtoMessage()    {}

func (m *TransferSucceeded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferSucceeded.Unmarshal(m, b)
}
func (m *TransferSucceeded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferSucceeded.Marshal(b, m, deterministic)
}
func (m *TransferSucceeded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferSucceeded.Merge(m, src)
}
func (m *TransferSucceeded) XXX_Size() int {
	return xxx_messageInfo_TransferSucceeded.Size(m)
}
func (m *TransferSucceeded) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferSucceeded.DiscardUnknown(m)
}

var xxx_messageInfo_TransferSucceeded proto.InternalMessageInfo

func (m *TransferSucceeded) GetOriginalOrderLimit() *OrderLimit {
	if m != nil {
		return m.OriginalOrderLimit
	}
	return nil
}

func (m *TransferSucceeded) GetOriginalPieceHash() *PieceHash {
	if m != nil {
		return m.OriginalPieceHash
	}
	return nil
}

func (m *TransferSucceeded) GetReplacementPieceHash() *PieceHash {
	if m != nil {
		return m.ReplacementPieceHash
	}
	return nil
}

type TransferFailed struct {
	OriginalPieceId      PieceID              `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	Error                TransferFailed_Error `protobuf:"varint,2,opt,name=error,proto3,enum=gracefulexit.TransferFailed_Error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransferFailed) Reset()         { *m = TransferFailed{} }
func (m *TransferFailed) String() string { return proto.CompactTextString(m) }
func (*TransferFailed) ProtoMessage()    {}

func (m *TransferFailed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFailed.Unmarshal(m, b)
}
func (m *TransferFailed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFailed.Marshal(b, m, deterministic)
}
func (m *TransferFailed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFailed.Merge(m, src)
}
func (m *TransferFailed) XXX_Size() int {
	return xxx_messageInfo_TransferFailed.Size(m)
}
func (m *TransferFailed) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFailed.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFailed proto.InternalMessageInfo

func (m *TransferFailed) GetError() TransferFailed_Error {
	if m != nil {
		return m.Error
	}
	return TransferFailed_NOT_FOUND
}

type StorageNodeMessage struct {
	// Types that are valid to be assigned to Message:
	//	*StorageNodeMessage_Succeeded
	//	*StorageNodeMessage_Failed
	Message              isStorageNodeMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StorageNodeMessage) Reset()         { *m = StorageNodeMessage{} }
func (m *StorageNodeMessage) String() string { return proto.CompactTextString(m) }
func (*StorageNodeMessage) ProtoMessage()    {}

func (m *StorageNodeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageNodeMessage.Unmarshal(m, b)
}
func (m *StorageNodeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageNodeMessage.Marshal(b, m, deterministic)
}
func (m *StorageNodeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageNodeMessage.Merge(m, src)
}
func (m *StorageNodeMessage) XXX_Size() int {
	return xxx_messageInfo_StorageNodeMessage.Size(m)
}
func (m *StorageNodeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageNodeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StorageNodeMessage proto.InternalMessageInfo

type isStorageNodeMessage_Message interface {
	isStorageNodeMessage_Message()
}

type StorageNodeMessage_Succeeded struct {
	Succeeded *TransferSucceeded `protobuf:"bytes,1,opt,name=succeeded,proto3,oneof" json:"succeeded,omitempty"`
}
type StorageNodeMessage_Failed struct {
	Failed *TransferFailed `protobuf:"bytes,2,opt,name=failed,proto3,oneof" json:"failed,omitempty"`
}

func (*StorageNodeMessage_Succeeded) isStorageNodeMessage_Message() {}
func (*StorageNodeMessage_Failed) isStorageNodeMessage_Message()    {}

func (m *StorageNodeMessage) GetMessage() isStorageNodeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *StorageNodeMessage) GetSucceeded() *TransferSucceeded {
	if x, ok := m.GetMessage().(*StorageNodeMessage_Succeeded); ok {
		return x.Succeeded
	}
	return nil
}

func (m *StorageNodeMessage) GetFailed() *TransferFailed {
	if x, ok := m.GetMessage().(*StorageNodeMessage_Failed); ok {
		return x.Failed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StorageNodeMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StorageNodeMessage_Succeeded)(nil),
		(*StorageNodeMessage_Failed)(nil),
	}
}

type NotReady struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotReady) Reset()         { *m = NotReady{} }
func (m *NotReady) String() string { return proto.CompactTextString(m) }
func (*NotReady) ProtoMessage()    {}

func (m *NotReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotReady.Unmarshal(m, b)
}
func (m *NotReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotReady.Marshal(b, m, deterministic)
}
func (m *NotReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotReady.Merge(m, src)
}
func (m *NotReady) XXX_Size() int {
	return xxx_messageInfo_NotReady.Size(m)
}
func (m *NotReady) XXX_DiscardUnknown() {
	xxx_messageInfo_NotReady.DiscardUnknown(m)
}

var xxx_messageInfo_NotReady proto.InternalMessageInfo

type TransferPiece struct {
	OriginalPieceId PieceID         `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	PrivateKey      PiecePrivateKey `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3,customtype=PiecePrivateKey" json:"private_key"`
	// addressed_order_limit contains the new piece id.
	AddressedOrderLimit  *AddressedOrderLimit `protobuf:"bytes,3,opt,name=addressed_order_limit,json=addressedOrderLimit,proto3" json:"addressed_order_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransferPiece) Reset()         { *m = TransferPiece{} }
func (m *TransferPiece) String() string { return proto.CompactTextString(m) }
func (*TransferPiece) ProtoMessage()    {}

func (m *TransferPiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferPiece.Unmarshal(m, b)
}
func (m *TransferPiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferPiece.Marshal(b, m, deterministic)
}
func (m *TransferPiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferPiece.Merge(m, src)
}
func (m *TransferPiece) XXX_Size() int {
	return xxx_messageInfo_TransferPiece.Size(m)
}
func (m *TransferPiece) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferPiece.DiscardUnknown(m)
}

var xxx_messageInfo_TransferPiece proto.InternalMessageInfo

func (m *TransferPiece) GetAddressedOrderLimit() *AddressedOrderLimit {
	if m != nil {
		return m.AddressedOrderLimit
	}
	return nil
}

type DeletePiece struct {
	OriginalPieceId      PieceID  `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePiece) Reset()         { *m = DeletePiece{} }
func (m *DeletePiece) String() string { return proto.CompactTextString(m) }
func (*DeletePiece) ProtoMessage()    {}

func (m *DeletePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePiece.Unmarshal(m, b)
}
func (m *DeletePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePiece.Marshal(b, m, deterministic)
}
func (m *DeletePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePiece.Merge(m, src)
}
func (m *DeletePiece) XXX_Size() int {
	return xxx_messageInfo_DeletePiece.Size(m)
}
func (m *DeletePiece) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePiece.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePiece proto.InternalMessageInfo

type ExitCompleted struct {
	// when everything is completed
	ExitCompleteSignature []byte `protobuf:"bytes,1,opt,name=exit_complete_signature,json=exitCompleteSignature,proto3" json:"exit_complete_signature,omitempty"`
	// satellite who issued this exit completed
	SatelliteId NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// storage node this exit completed was issued to
	NodeId NodeID `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	// timestamp when the exit completed
	Completed            time.Time `protobuf:"bytes,4,opt,name=completed,proto3,stdtime" json:"completed"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExitCompleted) Reset()         { *m = ExitCompleted{} }
func (m *ExitCompleted) String() string { return proto.CompactTextString(m) }
func (*ExitCompleted) ProtoMessage()    {}

func (m *ExitCompleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitCompleted.Unmarshal(m, b)
}
func (m *ExitCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitCompleted.Marshal(b, m, deterministic)
}
func (m *ExitCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitCompleted.Merge(m, src)
}
func (m *ExitCompleted) XXX_Size() int {
	return xxx_messageInfo_ExitCompleted.Size(m)
}
func (m *ExitCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_ExitCompleted proto.InternalMessageInfo

func (m *ExitCompleted) GetExitCompleteSignature() []byte {
	if m != nil {
		return m.ExitCompleteSignature
	}
	return nil
}

func (m *ExitCompleted) GetCompleted() time.Time {
	if m != nil {
		return m.Completed
	}
	return time.Time{}
}

type ExitFailed struct {
	// on failure
	ExitFailureSignature []byte            `protobuf:"bytes,1,opt,name=exit_failure_signature,json=exitFailureSignature,proto3" json:"exit_failure_signature,omitempty"`
	Reason               ExitFailed_Reason `protobuf:"varint,2,opt,name=reason,proto3,enum=gracefulexit.ExitFailed_Reason" json:"reason,omitempty"`
	// satellite who issued this exit failed
	SatelliteId NodeID `protobuf:"bytes,3,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// storage node this exit failed was issued to
	NodeId NodeID `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	// timestamp when the exit failed
	Failed               time.Time `protobuf:"bytes,5,opt,name=failed,proto3,stdtime" json:"failed"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExitFailed) Reset()         { *m = ExitFailed{} }
func (m *ExitFailed) String() string { return proto.CompactTextString(m) }
func (*ExitFailed) ProtoMessage()    {}

func (m *ExitFailed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitFailed.Unmarshal(m, b)
}
func (m *ExitFailed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitFailed.Marshal(b, m, deterministic)
}
func (m *ExitFailed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitFailed.Merge(m, src)
}
func (m *ExitFailed) XXX_Size() int {
	return xxx_messageInfo_ExitFailed.Size(m)
}
func (m *ExitFailed) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitFailed.DiscardUnknown(m)
}

var xxx_messageInfo_ExitFailed proto.InternalMessageInfo

func (m *ExitFailed) GetExitFailureSignature() []byte {
	if m != nil {
		return m.ExitFailureSignature
	}
	return nil
}

func (m *ExitFailed) GetReason() ExitFailed_Reason {
	if m != nil {
		return m.Reason
	}
	return ExitFailed_VERIFICATION_FAILED
}

func (m *ExitFailed) GetFailed() time.Time {
	if m != nil {
		return m.Failed
	}
	return time.Time{}
}

type SatelliteMessage struct {
	// Types that are valid to be assigned to Message:
	//	*SatelliteMessage_NotReady
	//	*SatelliteMessage_TransferPiece
	//	*SatelliteMessage_DeletePiece
	//	*SatelliteMessage_ExitCompleted
	//	*SatelliteMessage_ExitFailed
	Message              isSatelliteMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SatelliteMessage) Reset()         { *m = SatelliteMessage{} }
func (m *SatelliteMessage) String() string { return proto.CompactTextString(m) }
func (*SatelliteMessage) ProtoMessage()    {}

func (m *SatelliteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SatelliteMessage.Unmarshal(m, b)
}
func (m *SatelliteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SatelliteMessage.Marshal(b, m, deterministic)
}
func (m *SatelliteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SatelliteMessage.Merge(m, src)
}
func (m *SatelliteMessage) XXX_Size() int {
	return xxx_messageInfo_SatelliteMessage.Size(m)
}
func (m *SatelliteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SatelliteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SatelliteMessage proto.InternalMessageInfo

type isSatelliteMessage_Message interface {
	isSatelliteMessage_Message()
}

type SatelliteMessage_NotReady struct {
	NotReady *NotReady `protobuf:"bytes,1,opt,name=not_ready,json=notReady,proto3,oneof" json:"not_ready,omitempty"`
}
type SatelliteMessage_TransferPiece struct {
	TransferPiece *TransferPiece `protobuf:"bytes,2,opt,name=transfer_piece,json=transferPiece,proto3,oneof" json:"transfer_piece,omitempty"`
}
type SatelliteMessage_DeletePiece struct {
	DeletePiece *DeletePiece `protobuf:"bytes,3,opt,name=delete_piece,json=deletePiece,proto3,oneof" json:"delete_piece,omitempty"`
}
type SatelliteMessage_ExitCompleted struct {
	ExitCompleted *ExitCompleted `protobuf:"bytes,4,opt,name=exit_completed,json=exitCompleted,proto3,oneof" json:"exit_completed,omitempty"`
}
type SatelliteMessage_ExitFailed struct {
	ExitFailed *ExitFailed `protobuf:"bytes,5,opt,name=exit_failed,json=exitFailed,proto3,oneof" json:"exit_failed,omitempty"`
}

func (*SatelliteMessage_NotReady) isSatelliteMessage_Message()      {}
func (*SatelliteMessage_TransferPiece) isSatelliteMessage_Message() {}
func (*SatelliteMessage_DeletePiece) isSatelliteMessage_Message()   {}
func (*SatelliteMessage_ExitCompleted) isSatelliteMessage_Message() {}
func (*SatelliteMessage_ExitFailed) isSatelliteMessage_Message()    {}

func (m *SatelliteMessage) GetMessage() isSatelliteMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SatelliteMessage) GetNotReady() *NotReady {
	if x, ok := m.GetMessage().(*SatelliteMessage_NotReady); ok {
		return x.NotReady
	}
	return nil
}

func (m *SatelliteMessage) GetTransferPiece() *TransferPiece {
	if x, ok := m.GetMessage().(*SatelliteMessage_TransferPiece); ok {
		return x.TransferPiece
	}
	return nil
}

func (m *SatelliteMessage) GetDeletePiece() *DeletePiece {
	if x, ok := m.GetMessage().(*SatelliteMessage_DeletePiece); ok {
		return x.DeletePiece
	}
	return nil
}

func (m *SatelliteMessage) GetExitCompleted() *ExitCompleted {
	if x, ok := m.GetMessage().(*SatelliteMessage_ExitCompleted); ok {
		return x.ExitCompleted
	}
	return nil
}

func (m *SatelliteMessage) GetExitFailed() *ExitFailed {
	if x, ok := m.GetMessage().(*SatelliteMessage_ExitFailed); ok {
		return x.ExitFailed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SatelliteMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SatelliteMessage_NotReady)(nil),
		(*SatelliteMessage_TransferPiece)(nil),
		(*SatelliteMessage_DeletePiece)(nil),
		(*SatelliteMessage_ExitCompleted)(nil),
		(*SatelliteMessage_ExitFailed)(nil),
	}
}
