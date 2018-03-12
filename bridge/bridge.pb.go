// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bridge.proto

/*
Package bridge is a generated protocol buffer package.

It is generated from these files:
	bridge.proto

It has these top-level messages:
	SubscribeToStreamRequest
	PublishRequest
	PublishResponse
	Msg
*/
package bridge

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubscribeToStreamRequest_SubscribeToStreamOption int32

const (
	SubscribeToStreamRequest_DEFAULT                  SubscribeToStreamRequest_SubscribeToStreamOption = 0
	SubscribeToStreamRequest_START_WITH_LAST_RECEIVED SubscribeToStreamRequest_SubscribeToStreamOption = 1
	SubscribeToStreamRequest_DELIVER_ALL_AVAILABLE    SubscribeToStreamRequest_SubscribeToStreamOption = 2
	SubscribeToStreamRequest_START_AT_SEQUENCE        SubscribeToStreamRequest_SubscribeToStreamOption = 3
	SubscribeToStreamRequest_START_AT_TIME            SubscribeToStreamRequest_SubscribeToStreamOption = 4
	SubscribeToStreamRequest_START_AT_TIME_DELTA      SubscribeToStreamRequest_SubscribeToStreamOption = 5
)

var SubscribeToStreamRequest_SubscribeToStreamOption_name = map[int32]string{
	0: "DEFAULT",
	1: "START_WITH_LAST_RECEIVED",
	2: "DELIVER_ALL_AVAILABLE",
	3: "START_AT_SEQUENCE",
	4: "START_AT_TIME",
	5: "START_AT_TIME_DELTA",
}
var SubscribeToStreamRequest_SubscribeToStreamOption_value = map[string]int32{
	"DEFAULT":                  0,
	"START_WITH_LAST_RECEIVED": 1,
	"DELIVER_ALL_AVAILABLE":    2,
	"START_AT_SEQUENCE":        3,
	"START_AT_TIME":            4,
	"START_AT_TIME_DELTA":      5,
}

func (x SubscribeToStreamRequest_SubscribeToStreamOption) String() string {
	return proto.EnumName(SubscribeToStreamRequest_SubscribeToStreamOption_name, int32(x))
}
func (SubscribeToStreamRequest_SubscribeToStreamOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// SubscriptionRequest contains the parameters for subscribing to a NATS queue.
type SubscribeToStreamRequest struct {
	// Required subject of the queue
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	// Option queue group
	QueueGroup string `protobuf:"bytes,2,opt,name=queue_group,json=queueGroup" json:"queue_group,omitempty"`
	// Subscription option
	SubscriptionType SubscribeToStreamRequest_SubscribeToStreamOption `protobuf:"varint,3,opt,name=subscription_type,json=subscriptionType,enum=bridge.SubscribeToStreamRequest_SubscribeToStreamOption" json:"subscription_type,omitempty"`
	// Set when type is START_AT_SEQUENCE
	StartAtSequence uint64 `protobuf:"varint,4,opt,name=start_at_sequence,json=startAtSequence" json:"start_at_sequence,omitempty"`
	// Set when type is START_AT_TIME
	StartAtTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=start_at_time,json=startAtTime" json:"start_at_time,omitempty"`
	// Set when type is START_AT_TIME_DELTA, in nanoseconds.
	StartAtTimeDeltaNs int64 `protobuf:"varint,6,opt,name=start_at_time_delta_ns,json=startAtTimeDeltaNs" json:"start_at_time_delta_ns,omitempty"`
}

func (m *SubscribeToStreamRequest) Reset()                    { *m = SubscribeToStreamRequest{} }
func (m *SubscribeToStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeToStreamRequest) ProtoMessage()               {}
func (*SubscribeToStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubscribeToStreamRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *SubscribeToStreamRequest) GetQueueGroup() string {
	if m != nil {
		return m.QueueGroup
	}
	return ""
}

func (m *SubscribeToStreamRequest) GetSubscriptionType() SubscribeToStreamRequest_SubscribeToStreamOption {
	if m != nil {
		return m.SubscriptionType
	}
	return SubscribeToStreamRequest_DEFAULT
}

func (m *SubscribeToStreamRequest) GetStartAtSequence() uint64 {
	if m != nil {
		return m.StartAtSequence
	}
	return 0
}

func (m *SubscribeToStreamRequest) GetStartAtTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartAtTime
	}
	return nil
}

func (m *SubscribeToStreamRequest) GetStartAtTimeDeltaNs() int64 {
	if m != nil {
		return m.StartAtTimeDeltaNs
	}
	return 0
}

// Request for publishing a message.
type PublishRequest struct {
	// Required subject of queue to publish to.
	Subject string `protobuf:"bytes,1,opt,name=subject" json:"subject,omitempty"`
	// The data to publish.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublishRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *PublishRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Response to publishing a message.
type PublishResponse struct {
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// This is a copy of a stan.MsgProto, since Go doesn't allow type casts of pointers for some reason.
type Msg struct {
	Sequence    uint64 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Subject     string `protobuf:"bytes,2,opt,name=subject" json:"subject,omitempty"`
	Reply       string `protobuf:"bytes,3,opt,name=reply" json:"reply,omitempty"`
	Data        []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp   int64  `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	Redelivered bool   `protobuf:"varint,6,opt,name=redelivered" json:"redelivered,omitempty"`
	Crc32       uint32 `protobuf:"varint,7,opt,name=crc32" json:"crc32,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Msg) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Msg) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Msg) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Msg) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Msg) GetRedelivered() bool {
	if m != nil {
		return m.Redelivered
	}
	return false
}

func (m *Msg) GetCrc32() uint32 {
	if m != nil {
		return m.Crc32
	}
	return 0
}

func init() {
	proto.RegisterType((*SubscribeToStreamRequest)(nil), "bridge.SubscribeToStreamRequest")
	proto.RegisterType((*PublishRequest)(nil), "bridge.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "bridge.PublishResponse")
	proto.RegisterType((*Msg)(nil), "bridge.Msg")
	proto.RegisterEnum("bridge.SubscribeToStreamRequest_SubscribeToStreamOption", SubscribeToStreamRequest_SubscribeToStreamOption_name, SubscribeToStreamRequest_SubscribeToStreamOption_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Bridge service

type BridgeClient interface {
	// Publish a message
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// Subscribe to a stream
	SubscribeToStream(ctx context.Context, in *SubscribeToStreamRequest, opts ...grpc.CallOption) (Bridge_SubscribeToStreamClient, error)
}

type bridgeClient struct {
	cc *grpc.ClientConn
}

func NewBridgeClient(cc *grpc.ClientConn) BridgeClient {
	return &bridgeClient{cc}
}

func (c *bridgeClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/bridge.Bridge/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) SubscribeToStream(ctx context.Context, in *SubscribeToStreamRequest, opts ...grpc.CallOption) (Bridge_SubscribeToStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Bridge_serviceDesc.Streams[0], c.cc, "/bridge.Bridge/SubscribeToStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bridgeSubscribeToStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bridge_SubscribeToStreamClient interface {
	Recv() (*Msg, error)
	grpc.ClientStream
}

type bridgeSubscribeToStreamClient struct {
	grpc.ClientStream
}

func (x *bridgeSubscribeToStreamClient) Recv() (*Msg, error) {
	m := new(Msg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Bridge service

type BridgeServer interface {
	// Publish a message
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	// Subscribe to a stream
	SubscribeToStream(*SubscribeToStreamRequest, Bridge_SubscribeToStreamServer) error
}

func RegisterBridgeServer(s *grpc.Server, srv BridgeServer) {
	s.RegisterService(&_Bridge_serviceDesc, srv)
}

func _Bridge_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bridge.Bridge/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_SubscribeToStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BridgeServer).SubscribeToStream(m, &bridgeSubscribeToStreamServer{stream})
}

type Bridge_SubscribeToStreamServer interface {
	Send(*Msg) error
	grpc.ServerStream
}

type bridgeSubscribeToStreamServer struct {
	grpc.ServerStream
}

func (x *bridgeSubscribeToStreamServer) Send(m *Msg) error {
	return x.ServerStream.SendMsg(m)
}

var _Bridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bridge.Bridge",
	HandlerType: (*BridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Bridge_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToStream",
			Handler:       _Bridge_SubscribeToStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bridge.proto",
}

func init() { proto.RegisterFile("bridge.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x6e, 0x9b, 0x40,
	0x18, 0xec, 0xc6, 0xc4, 0x4e, 0x3e, 0xf2, 0x03, 0xdb, 0x26, 0xa1, 0x56, 0xa4, 0x20, 0x9f, 0x50,
	0x0f, 0x4e, 0xe5, 0x5c, 0xaa, 0x1e, 0x22, 0x91, 0xb0, 0x6d, 0x2d, 0x91, 0xb4, 0x5d, 0x48, 0x7a,
	0x5c, 0x81, 0xbd, 0xa5, 0x54, 0xb6, 0x21, 0xec, 0x52, 0xc9, 0xc7, 0x5e, 0xfb, 0x18, 0x7d, 0x8d,
	0xbe, 0x5c, 0xc5, 0x62, 0x13, 0x5b, 0x56, 0x94, 0x1b, 0xdf, 0x7c, 0x3f, 0xc3, 0xce, 0x0c, 0xec,
	0xc5, 0x45, 0x3a, 0x4e, 0x78, 0x3f, 0x2f, 0x32, 0x99, 0xe1, 0x76, 0x5d, 0x75, 0xcf, 0x92, 0x2c,
	0x4b, 0x26, 0xfc, 0x5c, 0xa1, 0x71, 0xf9, 0xfd, 0x5c, 0xa6, 0x53, 0x2e, 0x64, 0x34, 0xcd, 0xeb,
	0xc1, 0xde, 0x6f, 0x0d, 0xac, 0xa0, 0x8c, 0xc5, 0xa8, 0x48, 0x63, 0x1e, 0x66, 0x81, 0x2c, 0x78,
	0x34, 0xa5, 0xfc, 0xa1, 0xe4, 0x42, 0x62, 0x0b, 0x3a, 0xa2, 0x8c, 0x7f, 0xf2, 0x91, 0xb4, 0x90,
	0x8d, 0x9c, 0x5d, 0xba, 0x2c, 0xf1, 0x19, 0xe8, 0x0f, 0x25, 0x2f, 0x39, 0x4b, 0x8a, 0xac, 0xcc,
	0xad, 0x2d, 0xd5, 0x05, 0x05, 0x7d, 0xac, 0x10, 0xcc, 0xc1, 0x14, 0xf5, 0xd9, 0x5c, 0xa6, 0xd9,
	0x8c, 0xc9, 0x79, 0xce, 0xad, 0x96, 0x8d, 0x9c, 0x83, 0xc1, 0xbb, 0xfe, 0xe2, 0x57, 0x9f, 0xe2,
	0xdd, 0x6c, 0x7c, 0x56, 0x47, 0xa8, 0xb1, 0x7a, 0x32, 0x9c, 0xe7, 0x1c, 0xbf, 0x01, 0x53, 0xc8,
	0xa8, 0x90, 0x2c, 0x92, 0x4c, 0x54, 0xdb, 0xb3, 0x11, 0xb7, 0x34, 0x1b, 0x39, 0x1a, 0x3d, 0x54,
	0x0d, 0x57, 0x06, 0x0b, 0x18, 0x5f, 0xc2, 0x7e, 0x33, 0x5b, 0xc9, 0x60, 0x6d, 0xdb, 0xc8, 0xd1,
	0x07, 0xdd, 0x7e, 0xad, 0x51, 0x7f, 0xa9, 0x51, 0x3f, 0x5c, 0x6a, 0x44, 0xf5, 0xc5, 0x8d, 0x0a,
	0xc1, 0x03, 0x38, 0x5e, 0xdb, 0x67, 0x63, 0x3e, 0x91, 0x11, 0x9b, 0x09, 0xab, 0x6d, 0x23, 0xa7,
	0x45, 0xf1, 0xca, 0xb0, 0x57, 0xb5, 0x6e, 0x45, 0xef, 0x2f, 0x82, 0x93, 0x27, 0x5e, 0x83, 0x75,
	0xe8, 0x78, 0xe4, 0x83, 0x7b, 0xe7, 0x87, 0xc6, 0x0b, 0x7c, 0x0a, 0x56, 0x10, 0xba, 0x34, 0x64,
	0xdf, 0x86, 0xe1, 0x27, 0xe6, 0xbb, 0x41, 0xc8, 0x28, 0xb9, 0x26, 0xc3, 0x7b, 0xe2, 0x19, 0x08,
	0xbf, 0x86, 0x23, 0x8f, 0xf8, 0xc3, 0x7b, 0x42, 0x99, 0xeb, 0xfb, 0xcc, 0xbd, 0x77, 0x87, 0xbe,
	0x7b, 0xe5, 0x13, 0x63, 0x0b, 0x1f, 0x81, 0x59, 0x2f, 0xba, 0x21, 0x0b, 0xc8, 0xd7, 0x3b, 0x72,
	0x7b, 0x4d, 0x8c, 0x16, 0x36, 0x61, 0xbf, 0x81, 0xc3, 0xe1, 0x0d, 0x31, 0x34, 0x7c, 0x02, 0x2f,
	0xd7, 0x20, 0xe6, 0x11, 0x3f, 0x74, 0x8d, 0xed, 0xde, 0x25, 0x1c, 0x7c, 0x29, 0xe3, 0x49, 0x2a,
	0x7e, 0x3c, 0x6f, 0x3c, 0x06, 0x6d, 0x1c, 0xc9, 0x48, 0x39, 0xbe, 0x47, 0xd5, 0x77, 0xcf, 0x84,
	0xc3, 0x66, 0x5f, 0xe4, 0xd9, 0x4c, 0xf0, 0xde, 0x3f, 0x04, 0xad, 0x1b, 0x91, 0xe0, 0x2e, 0xec,
	0x34, 0xb6, 0x20, 0x65, 0x4b, 0x53, 0xaf, 0x92, 0x6c, 0xad, 0x93, 0xbc, 0x82, 0xed, 0x82, 0xe7,
	0x93, 0xb9, 0x0a, 0xcc, 0x2e, 0xad, 0x8b, 0x86, 0x5a, 0x7b, 0xa4, 0xc6, 0xa7, 0xb0, 0xdb, 0x24,
	0x5a, 0xf9, 0xd9, 0xa2, 0x8f, 0x00, 0xb6, 0x41, 0x2f, 0xf8, 0x98, 0x4f, 0xd2, 0x5f, 0xbc, 0xe0,
	0x63, 0x65, 0xd3, 0x0e, 0x5d, 0x85, 0x2a, 0xa6, 0x51, 0x31, 0xba, 0x18, 0x58, 0x1d, 0x1b, 0x39,
	0xfb, 0xb4, 0x2e, 0x06, 0x7f, 0x10, 0xb4, 0xaf, 0x54, 0x46, 0xf1, 0x7b, 0xe8, 0x2c, 0xde, 0x86,
	0x8f, 0x97, 0xb9, 0x5d, 0x17, 0xab, 0x7b, 0xb2, 0x81, 0xd7, 0x22, 0x60, 0x0f, 0xcc, 0x0d, 0xef,
	0xb1, 0xfd, 0x5c, 0xfa, 0xbb, 0xfa, 0x72, 0xe2, 0x46, 0x24, 0x6f, 0x51, 0xdc, 0x56, 0xb9, 0xbc,
	0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xd0, 0x97, 0x0b, 0xe1, 0x03, 0x00, 0x00,
}
