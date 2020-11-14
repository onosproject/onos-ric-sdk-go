// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/ricapi/e2/v1beta1/ricapi_e2.proto

// Package ricapi.e2.v1beta1 defines the interior gRPC interfaces for xApps to interact with E2T.

package v1beta1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1beta1 "github.com/onosproject/onos-e2t/api/ricapi/e2/headers/v1beta1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Indication an indication message
type Indication struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Indication) Reset()         { *m = Indication{} }
func (m *Indication) String() string { return proto.CompactTextString(m) }
func (*Indication) ProtoMessage()    {}
func (*Indication) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27d8777407939bb, []int{0}
}
func (m *Indication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Indication.Unmarshal(m, b)
}
func (m *Indication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Indication.Marshal(b, m, deterministic)
}
func (m *Indication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Indication.Merge(m, src)
}
func (m *Indication) XXX_Size() int {
	return xxx_messageInfo_Indication.Size(m)
}
func (m *Indication) XXX_DiscardUnknown() {
	xxx_messageInfo_Indication.DiscardUnknown(m)
}

var xxx_messageInfo_Indication proto.InternalMessageInfo

// ControlRequest
type ControlRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlRequest) Reset()         { *m = ControlRequest{} }
func (m *ControlRequest) String() string { return proto.CompactTextString(m) }
func (*ControlRequest) ProtoMessage()    {}
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27d8777407939bb, []int{1}
}
func (m *ControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlRequest.Unmarshal(m, b)
}
func (m *ControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlRequest.Marshal(b, m, deterministic)
}
func (m *ControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlRequest.Merge(m, src)
}
func (m *ControlRequest) XXX_Size() int {
	return xxx_messageInfo_ControlRequest.Size(m)
}
func (m *ControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ControlRequest proto.InternalMessageInfo

// ControlResponse
type ControlResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlResponse) Reset()         { *m = ControlResponse{} }
func (m *ControlResponse) String() string { return proto.CompactTextString(m) }
func (*ControlResponse) ProtoMessage()    {}
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27d8777407939bb, []int{2}
}
func (m *ControlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlResponse.Unmarshal(m, b)
}
func (m *ControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlResponse.Marshal(b, m, deterministic)
}
func (m *ControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlResponse.Merge(m, src)
}
func (m *ControlResponse) XXX_Size() int {
	return xxx_messageInfo_ControlResponse.Size(m)
}
func (m *ControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ControlResponse proto.InternalMessageInfo

type StreamRequest struct {
	AppID                AppID          `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3,casttype=AppID" json:"app_id,omitempty"`
	InstanceID           InstanceID     `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3,casttype=InstanceID" json:"instance_id,omitempty"`
	SubscriptionID       SubscriptionID `protobuf:"bytes,3,opt,name=subscription_id,json=subscriptionId,proto3,casttype=SubscriptionID" json:"subscription_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27d8777407939bb, []int{3}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetAppID() AppID {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *StreamRequest) GetInstanceID() InstanceID {
	if m != nil {
		return m.InstanceID
	}
	return ""
}

func (m *StreamRequest) GetSubscriptionID() SubscriptionID {
	if m != nil {
		return m.SubscriptionID
	}
	return ""
}

type StreamResponse struct {
	Header               *v1beta1.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payload              []byte                  `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e27d8777407939bb, []int{4}
}
func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetHeader() *v1beta1.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *StreamResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Indication)(nil), "ricapi.e2.v1beta1.Indication")
	proto.RegisterType((*ControlRequest)(nil), "ricapi.e2.v1beta1.ControlRequest")
	proto.RegisterType((*ControlResponse)(nil), "ricapi.e2.v1beta1.ControlResponse")
	proto.RegisterType((*StreamRequest)(nil), "ricapi.e2.v1beta1.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "ricapi.e2.v1beta1.StreamResponse")
}

func init() {
	proto.RegisterFile("api/ricapi/e2/v1beta1/ricapi_e2.proto", fileDescriptor_e27d8777407939bb)
}

var fileDescriptor_e27d8777407939bb = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xda, 0x30,
	0x18, 0xc6, 0x95, 0x4d, 0x64, 0xda, 0x0b, 0x0b, 0xc3, 0xda, 0x01, 0xa1, 0x49, 0xb0, 0x68, 0xd3,
	0xe0, 0x30, 0x67, 0xa4, 0xaa, 0x7a, 0xea, 0x01, 0xda, 0x4a, 0xcd, 0xa1, 0x87, 0x86, 0x9e, 0x7a,
	0x41, 0x4e, 0x62, 0x81, 0x2b, 0x88, 0x5d, 0xdb, 0x20, 0xf5, 0x1b, 0xf6, 0x53, 0x70, 0xe0, 0x63,
	0xf4, 0x54, 0xc5, 0x4e, 0x80, 0xfe, 0x51, 0x4f, 0x79, 0xdf, 0x27, 0xcf, 0xef, 0x8d, 0xdf, 0x27,
	0x86, 0x3f, 0x44, 0xb0, 0x40, 0xb2, 0xb4, 0x78, 0xd0, 0x30, 0x58, 0x0f, 0x13, 0xaa, 0xc9, 0xb0,
	0x54, 0xa6, 0x34, 0xc4, 0x42, 0x72, 0xcd, 0x51, 0xcb, 0x0a, 0x98, 0x86, 0xb8, 0xb4, 0x74, 0xfe,
	0xee, 0xa9, 0x39, 0x25, 0x19, 0x95, 0x6a, 0x47, 0x97, 0xbd, 0x65, 0x3b, 0x3f, 0x66, 0x7c, 0xc6,
	0x4d, 0x19, 0x14, 0x95, 0x55, 0xfd, 0x06, 0x40, 0x94, 0x67, 0x2c, 0x25, 0x9a, 0xf1, 0xdc, 0xff,
	0x0e, 0xde, 0x19, 0xcf, 0xb5, 0xe4, 0x8b, 0x98, 0xde, 0xaf, 0xa8, 0xd2, 0x7e, 0x0b, 0x9a, 0x3b,
	0x45, 0x09, 0x9e, 0x2b, 0xea, 0x3f, 0x3a, 0xf0, 0x6d, 0xa2, 0x25, 0x25, 0xcb, 0xd2, 0x84, 0x06,
	0xe0, 0x12, 0x21, 0xa6, 0x2c, 0x6b, 0x3b, 0x3d, 0xa7, 0xff, 0x75, 0x8c, 0xb6, 0x9b, 0x6e, 0x6d,
	0x24, 0x44, 0x74, 0xfe, 0x54, 0x15, 0x71, 0x8d, 0x08, 0x11, 0x65, 0xe8, 0x14, 0xea, 0x2c, 0x57,
	0x9a, 0xe4, 0x29, 0x2d, 0xfc, 0x9f, 0x8c, 0xff, 0xe7, 0x76, 0xd3, 0x85, 0xa8, 0x94, 0x0d, 0x74,
	0xd0, 0xc5, 0x50, 0x01, 0x51, 0x86, 0xae, 0xa0, 0xa9, 0x56, 0x89, 0x4a, 0x25, 0x13, 0xc5, 0x81,
	0x8b, 0x11, 0x9f, 0xcd, 0x88, 0xdf, 0xdb, 0x4d, 0xd7, 0x9b, 0x1c, 0xbc, 0x32, 0x63, 0x5e, 0x29,
	0xb1, 0x77, 0x08, 0x47, 0x99, 0xbf, 0x04, 0xaf, 0xda, 0xc4, 0x2e, 0x87, 0x46, 0xe0, 0xda, 0xd8,
	0xcc, 0x2a, 0xf5, 0x70, 0x80, 0xf7, 0x91, 0x57, 0x79, 0x96, 0xf9, 0xe2, 0x0a, 0xba, 0x34, 0x7a,
	0x5c, 0x82, 0xa8, 0x0d, 0x5f, 0x04, 0x79, 0x58, 0x70, 0x62, 0xd7, 0x6b, 0xc4, 0x55, 0x1b, 0x4e,
	0x01, 0x2e, 0xc2, 0x9b, 0x09, 0x95, 0x6b, 0x96, 0x52, 0x74, 0x0d, 0xae, 0xfd, 0x38, 0xea, 0xe1,
	0x37, 0xff, 0x15, 0xbf, 0x48, 0xb8, 0xf3, 0xeb, 0x03, 0x87, 0x3d, 0x44, 0xdf, 0xf9, 0xef, 0x8c,
	0x4f, 0x6e, 0x8f, 0x67, 0x4c, 0xcf, 0x57, 0x09, 0x4e, 0xf9, 0x32, 0xe0, 0x39, 0x57, 0x42, 0xf2,
	0x3b, 0x9a, 0x6a, 0x53, 0xff, 0xa3, 0xa1, 0x0e, 0xde, 0xbd, 0x68, 0x89, 0x6b, 0x6e, 0xc3, 0xd1,
	0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xa2, 0x21, 0x5b, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// E2TServiceClient is the client API for E2TService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type E2TServiceClient interface {
	// Stream opens an indications stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (E2TService_StreamClient, error)
}

type e2TServiceClient struct {
	cc *grpc.ClientConn
}

func NewE2TServiceClient(cc *grpc.ClientConn) E2TServiceClient {
	return &e2TServiceClient{cc}
}

func (c *e2TServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (E2TService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_E2TService_serviceDesc.Streams[0], "/ricapi.e2.v1beta1.E2TService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &e2TServiceStreamClient{stream}
	return x, nil
}

type E2TService_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type e2TServiceStreamClient struct {
	grpc.ClientStream
}

func (x *e2TServiceStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *e2TServiceStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// E2TServiceServer is the server API for E2TService service.
type E2TServiceServer interface {
	// Stream opens an indications stream
	Stream(E2TService_StreamServer) error
}

// UnimplementedE2TServiceServer can be embedded to have forward compatible implementations.
type UnimplementedE2TServiceServer struct {
}

func (*UnimplementedE2TServiceServer) Stream(srv E2TService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterE2TServiceServer(s *grpc.Server, srv E2TServiceServer) {
	s.RegisterService(&_E2TService_serviceDesc, srv)
}

func _E2TService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(E2TServiceServer).Stream(&e2TServiceStreamServer{stream})
}

type E2TService_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type e2TServiceStreamServer struct {
	grpc.ServerStream
}

func (x *e2TServiceStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *e2TServiceStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _E2TService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ricapi.e2.v1beta1.E2TService",
	HandlerType: (*E2TServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _E2TService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/ricapi/e2/v1beta1/ricapi_e2.proto",
}
