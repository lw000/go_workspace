// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	AddRequest
	AddReply
	SubRequest
	SubReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AddRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AddReply struct {
	C int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SubRequest struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *SubRequest) Reset()                    { *m = SubRequest{} }
func (m *SubRequest) String() string            { return proto.CompactTextString(m) }
func (*SubRequest) ProtoMessage()               {}
func (*SubRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SubReply struct {
	C int32 `protobuf:"varint,1,opt,name=c" json:"c,omitempty"`
}

func (m *SubReply) Reset()                    { *m = SubReply{} }
func (m *SubReply) String() string            { return proto.CompactTextString(m) }
func (*SubReply) ProtoMessage()               {}
func (*SubReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*AddRequest)(nil), "helloworld.AddRequest")
	proto.RegisterType((*AddReply)(nil), "helloworld.AddReply")
	proto.RegisterType((*SubRequest)(nil), "helloworld.SubRequest")
	proto.RegisterType((*SubReply)(nil), "helloworld.SubReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Sub(ctx context.Context, in *SubRequest, opts ...grpc.CallOption) (*SubReply, error) {
	out := new(SubReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Sub", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
	Sub(context.Context, *SubRequest) (*SubReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Sub(ctx, req.(*SubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Greeter_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Greeter_Sub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0x16, 0x68, 0x39, 0x65, 0xa8, 0x4e, 0x08, 0x59, 0x9d, 0x90, 0x07, 0xd4, 0x29,
	0x03, 0x88, 0x15, 0xa9, 0x2c, 0x74, 0x60, 0x4a, 0x06, 0x66, 0x3b, 0x3e, 0x15, 0x24, 0x07, 0x1b,
	0x27, 0x11, 0xe4, 0xef, 0xf0, 0x27, 0xf8, 0x7b, 0xc8, 0x06, 0x93, 0x46, 0x65, 0x60, 0xf3, 0x3b,
	0x7f, 0xef, 0xc9, 0xcf, 0x07, 0xcb, 0x27, 0x32, 0xc6, 0xbe, 0x59, 0x6f, 0x74, 0xee, 0xbc, 0x6d,
	0x2d, 0xc2, 0x30, 0x11, 0x02, 0xb2, 0x6d, 0x50, 0x05, 0xbd, 0x76, 0xd4, 0xb4, 0x88, 0x70, 0xf4,
	0x22, 0x6b, 0xe2, 0xec, 0x82, 0xad, 0x4f, 0x8b, 0x78, 0x16, 0x97, 0x00, 0x3f, 0x8c, 0x33, 0x3d,
	0x72, 0x98, 0xd7, 0xd4, 0x34, 0x72, 0x97, 0xa0, 0x24, 0xc5, 0x1a, 0x60, 0xa3, 0x75, 0x4a, 0xca,
	0x80, 0xc9, 0x48, 0x1c, 0x17, 0x4c, 0x06, 0xa5, 0xf8, 0xf4, 0x5b, 0x29, 0xc1, 0x61, 0x11, 0xc9,
	0x90, 0x97, 0x01, 0xab, 0x12, 0x57, 0x85, 0x8c, 0xb2, 0x53, 0xff, 0xcc, 0x88, 0xe4, 0x41, 0xc6,
	0xd5, 0x27, 0x83, 0xf9, 0xbd, 0x27, 0x6a, 0xc9, 0xe3, 0x2d, 0x2c, 0x4a, 0xd9, 0xc7, 0xe7, 0x23,
	0xcf, 0xf7, 0xbe, 0x62, 0xbf, 0xf5, 0xea, 0xfc, 0x8f, 0x1b, 0x67, 0x7a, 0x31, 0xc1, 0x1b, 0x98,
	0x6d, 0xb4, 0xc6, 0x11, 0x30, 0x94, 0x5c, 0x9d, 0x1d, 0xcc, 0x7f, 0x6d, 0x65, 0xa7, 0xc6, 0xb6,
	0xa1, 0xd7, 0xd8, 0x96, 0x5a, 0x88, 0xc9, 0x1d, 0x87, 0xe5, 0xb3, 0xcd, 0x77, 0xde, 0x55, 0x39,
	0xbd, 0xcb, 0xda, 0x19, 0x6a, 0x3e, 0xa6, 0xb3, 0xed, 0xc3, 0xa3, 0x3a, 0x89, 0xab, 0xbb, 0xfe,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x33, 0xa2, 0xeb, 0xce, 0x01, 0x00, 0x00,
}
