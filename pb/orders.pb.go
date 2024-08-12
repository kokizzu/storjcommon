// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orders.proto

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

// PieceAction is an enumeration of all possible executed actions on storage node.
type PieceAction int32

const (
	PieceAction_INVALID           PieceAction = 0
	PieceAction_PUT               PieceAction = 1
	PieceAction_GET               PieceAction = 2
	PieceAction_GET_AUDIT         PieceAction = 3
	PieceAction_GET_REPAIR        PieceAction = 4
	PieceAction_PUT_REPAIR        PieceAction = 5
	PieceAction_DELETE            PieceAction = 6
	PieceAction_PUT_GRACEFUL_EXIT PieceAction = 7
)

var PieceAction_name = map[int32]string{
	0: "INVALID",
	1: "PUT",
	2: "GET",
	3: "GET_AUDIT",
	4: "GET_REPAIR",
	5: "PUT_REPAIR",
	6: "DELETE",
	7: "PUT_GRACEFUL_EXIT",
}

var PieceAction_value = map[string]int32{
	"INVALID":           0,
	"PUT":               1,
	"GET":               2,
	"GET_AUDIT":         3,
	"GET_REPAIR":        4,
	"PUT_REPAIR":        5,
	"DELETE":            6,
	"PUT_GRACEFUL_EXIT": 7,
}

func (x PieceAction) String() string {
	return proto.EnumName(PieceAction_name, int32(x))
}

// PieceHashAlgorithm defines how the hashes of the pieces are calculated.
type PieceHashAlgorithm int32

const (
	PieceHashAlgorithm_SHA256 PieceHashAlgorithm = 0
	PieceHashAlgorithm_BLAKE3 PieceHashAlgorithm = 1
)

var PieceHashAlgorithm_name = map[int32]string{
	0: "SHA256",
	1: "BLAKE3",
}

var PieceHashAlgorithm_value = map[string]int32{
	"SHA256": 0,
	"BLAKE3": 1,
}

func (x PieceHashAlgorithm) String() string {
	return proto.EnumName(PieceHashAlgorithm_name, int32(x))
}

type SettlementWithWindowResponse_Status int32

const (
	SettlementWithWindowResponse_ACCEPTED SettlementWithWindowResponse_Status = 0
	SettlementWithWindowResponse_REJECTED SettlementWithWindowResponse_Status = 1
)

var SettlementWithWindowResponse_Status_name = map[int32]string{
	0: "ACCEPTED",
	1: "REJECTED",
}

var SettlementWithWindowResponse_Status_value = map[string]int32{
	"ACCEPTED": 0,
	"REJECTED": 1,
}

func (x SettlementWithWindowResponse_Status) String() string {
	return proto.EnumName(SettlementWithWindowResponse_Status_name, int32(x))
}

// OrderLimit is provided by satellite to execute specific action on storage node within some limits.
type OrderLimit struct {
	// unique serial to avoid replay attacks
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// satellite who issued this order limit allowing orderer to do the specified action
	SatelliteId NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// uplink who requested or whom behalf the order limit to do an action
	DeprecatedUplinkId *NodeID `protobuf:"bytes,3,opt,name=deprecated_uplink_id,json=deprecatedUplinkId,proto3,customtype=NodeID" json:"deprecated_uplink_id,omitempty"`
	// public key that will be used to sign orders and piece hash
	UplinkPublicKey PiecePublicKey `protobuf:"bytes,13,opt,name=uplink_public_key,json=uplinkPublicKey,proto3,customtype=PiecePublicKey" json:"uplink_public_key"`
	// storage node who can re claimthe order limit specified by serial
	StorageNodeId NodeID `protobuf:"bytes,4,opt,name=storage_node_id,json=storageNodeId,proto3,customtype=NodeID" json:"storage_node_id"`
	// piece which is allowed to be touched
	PieceId PieceID `protobuf:"bytes,5,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// limit in bytes how much can be changed
	Limit                  int64       `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Action                 PieceAction `protobuf:"varint,7,opt,name=action,proto3,enum=orders.PieceAction" json:"action,omitempty"`
	PieceExpiration        time.Time   `protobuf:"bytes,8,opt,name=piece_expiration,json=pieceExpiration,proto3,stdtime" json:"piece_expiration"`
	OrderExpiration        time.Time   `protobuf:"bytes,9,opt,name=order_expiration,json=orderExpiration,proto3,stdtime" json:"order_expiration"`
	OrderCreation          time.Time   `protobuf:"bytes,12,opt,name=order_creation,json=orderCreation,proto3,stdtime" json:"order_creation"`
	EncryptedMetadataKeyId []byte      `protobuf:"bytes,14,opt,name=encrypted_metadata_key_id,json=encryptedMetadataKeyId,proto3" json:"encrypted_metadata_key_id,omitempty"`
	EncryptedMetadata      []byte      `protobuf:"bytes,15,opt,name=encrypted_metadata,json=encryptedMetadata,proto3" json:"encrypted_metadata,omitempty"`
	SatelliteSignature     []byte      `protobuf:"bytes,10,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	// this allows a storage node to find a satellite and handshake with it to get its key.
	DeprecatedSatelliteAddress *NodeAddress `protobuf:"bytes,11,opt,name=deprecated_satellite_address,json=deprecatedSatelliteAddress,proto3" json:"deprecated_satellite_address,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}     `json:"-"`
	XXX_unrecognized           []byte       `json:"-"`
	XXX_sizecache              int32        `json:"-"`
}

func (m *OrderLimit) Reset()         { *m = OrderLimit{} }
func (m *OrderLimit) String() string { return proto.CompactTextString(m) }
func (*OrderLimit) ProtoMessage()    {}

func (m *OrderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimit.Unmarshal(m, b)
}
func (m *OrderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimit.Marshal(b, m, deterministic)
}
func (m *OrderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimit.Merge(m, src)
}
func (m *OrderLimit) XXX_Size() int {
	return xxx_messageInfo_OrderLimit.Size(m)
}
func (m *OrderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimit proto.InternalMessageInfo

func (m *OrderLimit) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderLimit) GetAction() PieceAction {
	if m != nil {
		return m.Action
	}
	return PieceAction_INVALID
}

func (m *OrderLimit) GetPieceExpiration() time.Time {
	if m != nil {
		return m.PieceExpiration
	}
	return time.Time{}
}

func (m *OrderLimit) GetOrderExpiration() time.Time {
	if m != nil {
		return m.OrderExpiration
	}
	return time.Time{}
}

func (m *OrderLimit) GetOrderCreation() time.Time {
	if m != nil {
		return m.OrderCreation
	}
	return time.Time{}
}

func (m *OrderLimit) GetEncryptedMetadataKeyId() []byte {
	if m != nil {
		return m.EncryptedMetadataKeyId
	}
	return nil
}

func (m *OrderLimit) GetEncryptedMetadata() []byte {
	if m != nil {
		return m.EncryptedMetadata
	}
	return nil
}

func (m *OrderLimit) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

func (m *OrderLimit) GetDeprecatedSatelliteAddress() *NodeAddress {
	if m != nil {
		return m.DeprecatedSatelliteAddress
	}
	return nil
}

// OrderLimitSigning provides OrderLimit signing serialization.
//
// It is never used for sending across the network, it is
// used in signing to ensure that nullable=false fields get handled properly.
// Its purpose is to solidify the format of how we serialize for
// signing, to handle some backwards compatibility considerations.
type OrderLimitSigning struct {
	// unique serial to avoid replay attacks
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// satellite who issued this order limit allowing orderer to do the specified action
	SatelliteId NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// uplink who requested or whom behalf the order limit to do an action
	DeprecatedUplinkId *NodeID `protobuf:"bytes,3,opt,name=deprecated_uplink_id,json=deprecatedUplinkId,proto3,customtype=NodeID" json:"deprecated_uplink_id,omitempty"`
	// public key that will be used to sign orders and piece hash
	UplinkPublicKey *PiecePublicKey `protobuf:"bytes,13,opt,name=uplink_public_key,json=uplinkPublicKey,proto3,customtype=PiecePublicKey" json:"uplink_public_key,omitempty"`
	// storage node who can re claimthe order limit specified by serial
	StorageNodeId NodeID `protobuf:"bytes,4,opt,name=storage_node_id,json=storageNodeId,proto3,customtype=NodeID" json:"storage_node_id"`
	// piece which is allowed to be touched
	PieceId PieceID `protobuf:"bytes,5,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// limit in bytes how much can be changed
	Limit                  int64       `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Action                 PieceAction `protobuf:"varint,7,opt,name=action,proto3,enum=orders.PieceAction" json:"action,omitempty"`
	PieceExpiration        *time.Time  `protobuf:"bytes,8,opt,name=piece_expiration,json=pieceExpiration,proto3,stdtime" json:"piece_expiration,omitempty"`
	OrderExpiration        *time.Time  `protobuf:"bytes,9,opt,name=order_expiration,json=orderExpiration,proto3,stdtime" json:"order_expiration,omitempty"`
	OrderCreation          *time.Time  `protobuf:"bytes,12,opt,name=order_creation,json=orderCreation,proto3,stdtime" json:"order_creation,omitempty"`
	EncryptedMetadataKeyId []byte      `protobuf:"bytes,14,opt,name=encrypted_metadata_key_id,json=encryptedMetadataKeyId,proto3" json:"encrypted_metadata_key_id,omitempty"`
	EncryptedMetadata      []byte      `protobuf:"bytes,15,opt,name=encrypted_metadata,json=encryptedMetadata,proto3" json:"encrypted_metadata,omitempty"`
	SatelliteSignature     []byte      `protobuf:"bytes,10,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	// this allows a storage node to find a satellite and handshake with it to get its key.
	DeprecatedSatelliteAddress *NodeAddress `protobuf:"bytes,11,opt,name=deprecated_satellite_address,json=deprecatedSatelliteAddress,proto3" json:"deprecated_satellite_address,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}     `json:"-"`
	XXX_unrecognized           []byte       `json:"-"`
	XXX_sizecache              int32        `json:"-"`
}

func (m *OrderLimitSigning) Reset()         { *m = OrderLimitSigning{} }
func (m *OrderLimitSigning) String() string { return proto.CompactTextString(m) }
func (*OrderLimitSigning) ProtoMessage()    {}

func (m *OrderLimitSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderLimitSigning.Unmarshal(m, b)
}
func (m *OrderLimitSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderLimitSigning.Marshal(b, m, deterministic)
}
func (m *OrderLimitSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderLimitSigning.Merge(m, src)
}
func (m *OrderLimitSigning) XXX_Size() int {
	return xxx_messageInfo_OrderLimitSigning.Size(m)
}
func (m *OrderLimitSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderLimitSigning.DiscardUnknown(m)
}

var xxx_messageInfo_OrderLimitSigning proto.InternalMessageInfo

func (m *OrderLimitSigning) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderLimitSigning) GetAction() PieceAction {
	if m != nil {
		return m.Action
	}
	return PieceAction_INVALID
}

func (m *OrderLimitSigning) GetPieceExpiration() *time.Time {
	if m != nil {
		return m.PieceExpiration
	}
	return nil
}

func (m *OrderLimitSigning) GetOrderExpiration() *time.Time {
	if m != nil {
		return m.OrderExpiration
	}
	return nil
}

func (m *OrderLimitSigning) GetOrderCreation() *time.Time {
	if m != nil {
		return m.OrderCreation
	}
	return nil
}

func (m *OrderLimitSigning) GetEncryptedMetadataKeyId() []byte {
	if m != nil {
		return m.EncryptedMetadataKeyId
	}
	return nil
}

func (m *OrderLimitSigning) GetEncryptedMetadata() []byte {
	if m != nil {
		return m.EncryptedMetadata
	}
	return nil
}

func (m *OrderLimitSigning) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

func (m *OrderLimitSigning) GetDeprecatedSatelliteAddress() *NodeAddress {
	if m != nil {
		return m.DeprecatedSatelliteAddress
	}
	return nil
}

// Order is a one step of fullfilling Amount number of bytes from an OrderLimit with SerialNumber.
type Order struct {
	// serial of the order limit that was signed
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// amount to be signed for
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// signature
	UplinkSignature      []byte   `protobuf:"bytes,3,opt,name=uplink_signature,json=uplinkSignature,proto3" json:"uplink_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Order) GetUplinkSignature() []byte {
	if m != nil {
		return m.UplinkSignature
	}
	return nil
}

// OrderSigning provides Order signing format.
//
// It is never used for sending across the network, it is
// used in signing to ensure that nullable=false fields get handled properly.
// Its purpose is to solidify the format of how we serialize for
// signing, to handle some backwards compatibility considerations.
type OrderSigning struct {
	// serial of the order limit that was signed
	SerialNumber SerialNumber `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3,customtype=SerialNumber" json:"serial_number"`
	// amount to be signed for
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// signature
	UplinkSignature      []byte   `protobuf:"bytes,3,opt,name=uplink_signature,json=uplinkSignature,proto3" json:"uplink_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderSigning) Reset()         { *m = OrderSigning{} }
func (m *OrderSigning) String() string { return proto.CompactTextString(m) }
func (*OrderSigning) ProtoMessage()    {}

func (m *OrderSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderSigning.Unmarshal(m, b)
}
func (m *OrderSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderSigning.Marshal(b, m, deterministic)
}
func (m *OrderSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderSigning.Merge(m, src)
}
func (m *OrderSigning) XXX_Size() int {
	return xxx_messageInfo_OrderSigning.Size(m)
}
func (m *OrderSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderSigning.DiscardUnknown(m)
}

var xxx_messageInfo_OrderSigning proto.InternalMessageInfo

func (m *OrderSigning) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OrderSigning) GetUplinkSignature() []byte {
	if m != nil {
		return m.UplinkSignature
	}
	return nil
}

type PieceHash struct {
	// piece id
	PieceId PieceID `protobuf:"bytes,1,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// hash of the piece that was/is uploaded
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// size of uploaded piece
	PieceSize int64 `protobuf:"varint,4,opt,name=piece_size,json=pieceSize,proto3" json:"piece_size,omitempty"`
	// timestamp when upload occurred
	Timestamp time.Time `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// signature either satellite or storage node
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// hash algorithm used
	HashAlgorithm        PieceHashAlgorithm `protobuf:"varint,6,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=orders.PieceHashAlgorithm" json:"hash_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PieceHash) Reset()         { *m = PieceHash{} }
func (m *PieceHash) String() string { return proto.CompactTextString(m) }
func (*PieceHash) ProtoMessage()    {}

func (m *PieceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceHash.Unmarshal(m, b)
}
func (m *PieceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceHash.Marshal(b, m, deterministic)
}
func (m *PieceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceHash.Merge(m, src)
}
func (m *PieceHash) XXX_Size() int {
	return xxx_messageInfo_PieceHash.Size(m)
}
func (m *PieceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceHash.DiscardUnknown(m)
}

var xxx_messageInfo_PieceHash proto.InternalMessageInfo

func (m *PieceHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PieceHash) GetPieceSize() int64 {
	if m != nil {
		return m.PieceSize
	}
	return 0
}

func (m *PieceHash) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *PieceHash) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PieceHash) GetHashAlgorithm() PieceHashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return PieceHashAlgorithm_SHA256
}

// PieceHashSigning provides piece hash signing format.
//
// It is never used for sending across the network, it is
// used in signing to ensure that nullable=false fields get handled properly.
// Its purpose is to solidify the format of how we serialize for
// signing, to handle some backwards compatibility considerations.
type PieceHashSigning struct {
	// piece id
	PieceId PieceID `protobuf:"bytes,1,opt,name=piece_id,json=pieceId,proto3,customtype=PieceID" json:"piece_id"`
	// hash of the piece that was/is uploaded
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// size of uploaded piece
	PieceSize int64 `protobuf:"varint,4,opt,name=piece_size,json=pieceSize,proto3" json:"piece_size,omitempty"`
	// timestamp when upload occurred
	Timestamp *time.Time `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	// signature either satellite or storage node
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// hash algorithm used
	HashAlgorithm        PieceHashAlgorithm `protobuf:"varint,6,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=orders.PieceHashAlgorithm" json:"hash_algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PieceHashSigning) Reset()         { *m = PieceHashSigning{} }
func (m *PieceHashSigning) String() string { return proto.CompactTextString(m) }
func (*PieceHashSigning) ProtoMessage()    {}

func (m *PieceHashSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceHashSigning.Unmarshal(m, b)
}
func (m *PieceHashSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceHashSigning.Marshal(b, m, deterministic)
}
func (m *PieceHashSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceHashSigning.Merge(m, src)
}
func (m *PieceHashSigning) XXX_Size() int {
	return xxx_messageInfo_PieceHashSigning.Size(m)
}
func (m *PieceHashSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceHashSigning.DiscardUnknown(m)
}

var xxx_messageInfo_PieceHashSigning proto.InternalMessageInfo

func (m *PieceHashSigning) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PieceHashSigning) GetPieceSize() int64 {
	if m != nil {
		return m.PieceSize
	}
	return 0
}

func (m *PieceHashSigning) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PieceHashSigning) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PieceHashSigning) GetHashAlgorithm() PieceHashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return PieceHashAlgorithm_SHA256
}

// Expected order of messages from storagenode:
//   go repeated
//      SettlementRequest -> (async)
//   go repeated
//      <- SettlementResponse
type SettlementRequest struct {
	Limit                *OrderLimit `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Order                *Order      `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SettlementRequest) Reset()         { *m = SettlementRequest{} }
func (m *SettlementRequest) String() string { return proto.CompactTextString(m) }
func (*SettlementRequest) ProtoMessage()    {}

func (m *SettlementRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlementRequest.Unmarshal(m, b)
}
func (m *SettlementRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlementRequest.Marshal(b, m, deterministic)
}
func (m *SettlementRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementRequest.Merge(m, src)
}
func (m *SettlementRequest) XXX_Size() int {
	return xxx_messageInfo_SettlementRequest.Size(m)
}
func (m *SettlementRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementRequest proto.InternalMessageInfo

func (m *SettlementRequest) GetLimit() *OrderLimit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *SettlementRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

// SettlementWithWindowResponse returns the status and the amounts
// settled summed by piece action.
// Accepted status means that the orders were successfully processed
// (or they were previously processed)
// Rejected status means that the orders were previously processed and
// the summed settled amounts don't match
type SettlementWithWindowResponse struct {
	Status               SettlementWithWindowResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=orders.SettlementWithWindowResponse_Status" json:"status,omitempty"`
	ActionSettled        map[int32]int64                     `protobuf:"bytes,2,rep,name=action_settled,json=actionSettled,proto3" json:"action_settled,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *SettlementWithWindowResponse) Reset()         { *m = SettlementWithWindowResponse{} }
func (m *SettlementWithWindowResponse) String() string { return proto.CompactTextString(m) }
func (*SettlementWithWindowResponse) ProtoMessage()    {}

func (m *SettlementWithWindowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettlementWithWindowResponse.Unmarshal(m, b)
}
func (m *SettlementWithWindowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettlementWithWindowResponse.Marshal(b, m, deterministic)
}
func (m *SettlementWithWindowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementWithWindowResponse.Merge(m, src)
}
func (m *SettlementWithWindowResponse) XXX_Size() int {
	return xxx_messageInfo_SettlementWithWindowResponse.Size(m)
}
func (m *SettlementWithWindowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementWithWindowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementWithWindowResponse proto.InternalMessageInfo

func (m *SettlementWithWindowResponse) GetStatus() SettlementWithWindowResponse_Status {
	if m != nil {
		return m.Status
	}
	return SettlementWithWindowResponse_ACCEPTED
}

func (m *SettlementWithWindowResponse) GetActionSettled() map[int32]int64 {
	if m != nil {
		return m.ActionSettled
	}
	return nil
}
