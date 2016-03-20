// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Position
	SetPositionRequest
	GetPositionRequest
	Empty
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion1

type Position struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
	H float64 `protobuf:"fixed64,3,opt,name=h" json:"h,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SetPositionRequest struct {
	HardwareAddress string    `protobuf:"bytes,1,opt,name=hardwareAddress" json:"hardwareAddress,omitempty"`
	Position        *Position `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
}

func (m *SetPositionRequest) Reset()                    { *m = SetPositionRequest{} }
func (m *SetPositionRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPositionRequest) ProtoMessage()               {}
func (*SetPositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SetPositionRequest) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

type GetPositionRequest struct {
	HardwareAddress string `protobuf:"bytes,1,opt,name=hardwareAddress" json:"hardwareAddress,omitempty"`
}

func (m *GetPositionRequest) Reset()                    { *m = GetPositionRequest{} }
func (m *GetPositionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPositionRequest) ProtoMessage()               {}
func (*GetPositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*SetPositionRequest)(nil), "pb.SetPositionRequest")
	proto.RegisterType((*GetPositionRequest)(nil), "pb.GetPositionRequest")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for PositionService service

type PositionServiceClient interface {
	SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*Empty, error)
	GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*Position, error)
}

type positionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPositionServiceClient(cc *grpc.ClientConn) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.PositionService/SetPosition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetPosition(ctx context.Context, in *GetPositionRequest, opts ...grpc.CallOption) (*Position, error) {
	out := new(Position)
	err := grpc.Invoke(ctx, "/pb.PositionService/GetPosition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PositionService service

type PositionServiceServer interface {
	SetPosition(context.Context, *SetPositionRequest) (*Empty, error)
	GetPosition(context.Context, *GetPositionRequest) (*Position, error)
}

func RegisterPositionServiceServer(s *grpc.Server, srv PositionServiceServer) {
	s.RegisterService(&_PositionService_serviceDesc, srv)
}

func _PositionService_SetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PositionServiceServer).SetPosition(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PositionService_GetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PositionServiceServer).GetPosition(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPosition",
			Handler:    _PositionService_SetPosition_Handler,
		},
		{
			MethodName: "GetPosition",
			Handler:    _PositionService_GetPosition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe1,
	0xe2, 0x08, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0x13, 0xe2, 0xe1, 0x62, 0xac, 0x90, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0c, 0x62, 0xac, 0x00, 0xf1, 0x2a, 0x25, 0x98, 0x20, 0xbc, 0x4a, 0x10,
	0x2f, 0x43, 0x82, 0x19, 0xc2, 0xcb, 0x50, 0xca, 0xe0, 0x12, 0x0a, 0x4e, 0x2d, 0x81, 0x69, 0x0c,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xd2, 0xe0, 0xe2, 0xcf, 0x48, 0x2c, 0x4a, 0x29, 0x4f,
	0x2c, 0x4a, 0x75, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x06, 0x9b, 0xc6, 0x19, 0x84, 0x2e, 0x0c,
	0x54, 0xc9, 0x51, 0x00, 0xd5, 0x0c, 0xb6, 0x82, 0xdb, 0x88, 0x47, 0xaf, 0x20, 0x49, 0x0f, 0x6e,
	0x20, 0x5c, 0x56, 0xc9, 0x8e, 0x4b, 0xc8, 0x9d, 0x02, 0x9b, 0x94, 0xd8, 0xb9, 0x58, 0x5d, 0x73,
	0x0b, 0x4a, 0x2a, 0x8d, 0x2a, 0xb8, 0xf8, 0x61, 0xa6, 0x04, 0x43, 0x42, 0x41, 0xc8, 0x80, 0x8b,
	0x1b, 0xc9, 0x17, 0x42, 0x62, 0x20, 0x27, 0x60, 0x7a, 0x4b, 0x8a, 0x13, 0x24, 0x0e, 0x36, 0x44,
	0xc8, 0x98, 0x8b, 0xdb, 0x1d, 0x5d, 0x07, 0xa6, 0xf3, 0xa4, 0x50, 0x3c, 0x93, 0xc4, 0x06, 0x0e,
	0x6d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x94, 0x41, 0x5f, 0x7e, 0x01, 0x00, 0x00,
}