// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: gracefulexit.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_gracefulexit_proto struct{}

func (drpcEncoding_File_gracefulexit_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_gracefulexit_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_gracefulexit_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_gracefulexit_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCSatelliteGracefulExitClient interface {
	DRPCConn() drpc.Conn

	Process(ctx context.Context) (DRPCSatelliteGracefulExit_ProcessClient, error)
	GracefulExitFeasibility(ctx context.Context, in *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error)
}

type drpcSatelliteGracefulExitClient struct {
	cc drpc.Conn
}

func NewDRPCSatelliteGracefulExitClient(cc drpc.Conn) DRPCSatelliteGracefulExitClient {
	return &drpcSatelliteGracefulExitClient{cc}
}

func (c *drpcSatelliteGracefulExitClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcSatelliteGracefulExitClient) Process(ctx context.Context) (DRPCSatelliteGracefulExit_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, "/gracefulexit.SatelliteGracefulExit/Process", drpcEncoding_File_gracefulexit_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcSatelliteGracefulExit_ProcessClient{stream}
	return x, nil
}

type DRPCSatelliteGracefulExit_ProcessClient interface {
	drpc.Stream
	Send(*StorageNodeMessage) error
	Recv() (*SatelliteMessage, error)
}

type drpcSatelliteGracefulExit_ProcessClient struct {
	drpc.Stream
}

func (x *drpcSatelliteGracefulExit_ProcessClient) GetStream() drpc.Stream {
	return x.Stream
}

func (x *drpcSatelliteGracefulExit_ProcessClient) Send(m *StorageNodeMessage) error {
	return x.MsgSend(m, drpcEncoding_File_gracefulexit_proto{})
}

func (x *drpcSatelliteGracefulExit_ProcessClient) Recv() (*SatelliteMessage, error) {
	m := new(SatelliteMessage)
	if err := x.MsgRecv(m, drpcEncoding_File_gracefulexit_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcSatelliteGracefulExit_ProcessClient) RecvMsg(m *SatelliteMessage) error {
	return x.MsgRecv(m, drpcEncoding_File_gracefulexit_proto{})
}

func (c *drpcSatelliteGracefulExitClient) GracefulExitFeasibility(ctx context.Context, in *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error) {
	out := new(GracefulExitFeasibilityResponse)
	err := c.cc.Invoke(ctx, "/gracefulexit.SatelliteGracefulExit/GracefulExitFeasibility", drpcEncoding_File_gracefulexit_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCSatelliteGracefulExitServer interface {
	Process(DRPCSatelliteGracefulExit_ProcessStream) error
	GracefulExitFeasibility(context.Context, *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error)
}

type DRPCSatelliteGracefulExitUnimplementedServer struct{}

func (s *DRPCSatelliteGracefulExitUnimplementedServer) Process(DRPCSatelliteGracefulExit_ProcessStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCSatelliteGracefulExitUnimplementedServer) GracefulExitFeasibility(context.Context, *GracefulExitFeasibilityRequest) (*GracefulExitFeasibilityResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCSatelliteGracefulExitDescription struct{}

func (DRPCSatelliteGracefulExitDescription) NumMethods() int { return 2 }

func (DRPCSatelliteGracefulExitDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/gracefulexit.SatelliteGracefulExit/Process", drpcEncoding_File_gracefulexit_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCSatelliteGracefulExitServer).
					Process(
						&drpcSatelliteGracefulExit_ProcessStream{in1.(drpc.Stream)},
					)
			}, DRPCSatelliteGracefulExitServer.Process, true
	case 1:
		return "/gracefulexit.SatelliteGracefulExit/GracefulExitFeasibility", drpcEncoding_File_gracefulexit_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCSatelliteGracefulExitServer).
					GracefulExitFeasibility(
						ctx,
						in1.(*GracefulExitFeasibilityRequest),
					)
			}, DRPCSatelliteGracefulExitServer.GracefulExitFeasibility, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterSatelliteGracefulExit(mux drpc.Mux, impl DRPCSatelliteGracefulExitServer) error {
	return mux.Register(impl, DRPCSatelliteGracefulExitDescription{})
}

type DRPCSatelliteGracefulExit_ProcessStream interface {
	drpc.Stream
	Send(*SatelliteMessage) error
	Recv() (*StorageNodeMessage, error)
}

type drpcSatelliteGracefulExit_ProcessStream struct {
	drpc.Stream
}

func (x *drpcSatelliteGracefulExit_ProcessStream) Send(m *SatelliteMessage) error {
	return x.MsgSend(m, drpcEncoding_File_gracefulexit_proto{})
}

func (x *drpcSatelliteGracefulExit_ProcessStream) Recv() (*StorageNodeMessage, error) {
	m := new(StorageNodeMessage)
	if err := x.MsgRecv(m, drpcEncoding_File_gracefulexit_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcSatelliteGracefulExit_ProcessStream) RecvMsg(m *StorageNodeMessage) error {
	return x.MsgRecv(m, drpcEncoding_File_gracefulexit_proto{})
}

type DRPCSatelliteGracefulExit_GracefulExitFeasibilityStream interface {
	drpc.Stream
	SendAndClose(*GracefulExitFeasibilityResponse) error
}

type drpcSatelliteGracefulExit_GracefulExitFeasibilityStream struct {
	drpc.Stream
}

func (x *drpcSatelliteGracefulExit_GracefulExitFeasibilityStream) SendAndClose(m *GracefulExitFeasibilityResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_gracefulexit_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
