// Code generated by protoc-gen-go.
// source: practice_app.proto
// DO NOT EDIT!

/*
Package routes is a generated protocol buffer package.

It is generated from these files:
	practice_app.proto

It has these top-level messages:
	NumberRequest
	NumberResponse
*/
package routes

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the number
type NumberRequest struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *NumberRequest) Reset()                    { *m = NumberRequest{} }
func (m *NumberRequest) String() string            { return proto.CompactTextString(m) }
func (*NumberRequest) ProtoMessage()               {}
func (*NumberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NumberRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// The response message containing the new number
type NumberResponse struct {
	Number int32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *NumberResponse) Reset()                    { *m = NumberResponse{} }
func (m *NumberResponse) String() string            { return proto.CompactTextString(m) }
func (*NumberResponse) ProtoMessage()               {}
func (*NumberResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NumberResponse) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*NumberRequest)(nil), "routes.NumberRequest")
	proto.RegisterType((*NumberResponse)(nil), "routes.NumberResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	AddOne(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error)
	Square(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) AddOne(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error) {
	out := new(NumberResponse)
	err := grpc.Invoke(ctx, "/routes.Calculator/AddOne", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Square(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error) {
	out := new(NumberResponse)
	err := grpc.Invoke(ctx, "/routes.Calculator/Square", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	AddOne(context.Context, *NumberRequest) (*NumberResponse, error)
	Square(context.Context, *NumberRequest) (*NumberResponse, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_AddOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).AddOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routes.Calculator/AddOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).AddOne(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Square_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Square(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routes.Calculator/Square",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Square(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routes.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOne",
			Handler:    _Calculator_AddOne_Handler,
		},
		{
			MethodName: "Square",
			Handler:    _Calculator_Square_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "practice_app.proto",
}

func init() { proto.RegisterFile("practice_app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0x4a, 0x4c,
	0x2e, 0xc9, 0x4c, 0x4e, 0x8d, 0x4f, 0x2c, 0x28, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2b, 0xca, 0x2f, 0x2d, 0x49, 0x2d, 0x56, 0x52, 0xe7, 0xe2, 0xf5, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d,
	0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xcb, 0x03, 0x0b, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x4a, 0x1a, 0x5c, 0x7c, 0x30, 0x85, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0xb8, 0x54, 0x1a, 0x35, 0x31, 0x72, 0x71, 0x39, 0x27, 0xe6, 0x24, 0x97, 0xe6,
	0x24, 0x96, 0xe4, 0x17, 0x09, 0x59, 0x72, 0xb1, 0x39, 0xa6, 0xa4, 0xf8, 0xe7, 0xa5, 0x0a, 0x89,
	0xea, 0x41, 0x2c, 0xd5, 0x43, 0xb1, 0x51, 0x4a, 0x0c, 0x5d, 0x18, 0x62, 0xbe, 0x12, 0x03, 0x48,
	0x6b, 0x70, 0x61, 0x69, 0x62, 0x11, 0xe9, 0x5a, 0x93, 0xd8, 0xc0, 0xde, 0x34, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x47, 0x90, 0x53, 0xb9, 0xfc, 0x00, 0x00, 0x00,
}