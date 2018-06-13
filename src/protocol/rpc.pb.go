// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/rpc.proto

package protocol

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
type CalParam struct {
	A                    int32    `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalParam) Reset()         { *m = CalParam{} }
func (m *CalParam) String() string { return proto.CompactTextString(m) }
func (*CalParam) ProtoMessage()    {}
func (*CalParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_15c43f292e82ff36, []int{0}
}
func (m *CalParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalParam.Unmarshal(m, b)
}
func (m *CalParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalParam.Marshal(b, m, deterministic)
}
func (dst *CalParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalParam.Merge(dst, src)
}
func (m *CalParam) XXX_Size() int {
	return xxx_messageInfo_CalParam.Size(m)
}
func (m *CalParam) XXX_DiscardUnknown() {
	xxx_messageInfo_CalParam.DiscardUnknown(m)
}

var xxx_messageInfo_CalParam proto.InternalMessageInfo

func (m *CalParam) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *CalParam) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

// The response message containing the greetings
type CalResult struct {
	Result               int32    `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalResult) Reset()         { *m = CalResult{} }
func (m *CalResult) String() string { return proto.CompactTextString(m) }
func (*CalResult) ProtoMessage()    {}
func (*CalResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpc_15c43f292e82ff36, []int{1}
}
func (m *CalResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalResult.Unmarshal(m, b)
}
func (m *CalResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalResult.Marshal(b, m, deterministic)
}
func (dst *CalResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalResult.Merge(dst, src)
}
func (m *CalResult) XXX_Size() int {
	return xxx_messageInfo_CalResult.Size(m)
}
func (m *CalResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CalResult.DiscardUnknown(m)
}

var xxx_messageInfo_CalResult proto.InternalMessageInfo

func (m *CalResult) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalParam)(nil), "protocol.CalParam")
	proto.RegisterType((*CalResult)(nil), "protocol.CalResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	Add(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error)
	Sub(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error)
	Mul(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error)
	Div(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error) {
	out := new(CalResult)
	err := grpc.Invoke(ctx, "/protocol.Calculator/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Sub(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error) {
	out := new(CalResult)
	err := grpc.Invoke(ctx, "/protocol.Calculator/Sub", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Mul(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error) {
	out := new(CalResult)
	err := grpc.Invoke(ctx, "/protocol.Calculator/Mul", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Div(ctx context.Context, in *CalParam, opts ...grpc.CallOption) (*CalResult, error) {
	out := new(CalResult)
	err := grpc.Invoke(ctx, "/protocol.Calculator/Div", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	Add(context.Context, *CalParam) (*CalResult, error)
	Sub(context.Context, *CalParam) (*CalResult, error)
	Mul(context.Context, *CalParam) (*CalResult, error)
	Div(context.Context, *CalParam) (*CalResult, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*CalParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Calculator/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Sub(ctx, req.(*CalParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Calculator/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Mul(ctx, req.(*CalParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Calculator/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Div(ctx, req.(*CalParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Calculator_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _Calculator_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _Calculator_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/rpc.proto",
}

func init() { proto.RegisterFile("protocol/rpc.proto", fileDescriptor_rpc_15c43f292e82ff36) }

var fileDescriptor_rpc_15c43f292e82ff36 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0x73, 0x84, 0x38, 0x60, 0x62, 0x4a,
	0x6a, 0x5c, 0x1c, 0xce, 0x89, 0x39, 0x01, 0x89, 0x45, 0x89, 0xb9, 0x42, 0x3c, 0x5c, 0x8c, 0x89,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x8c, 0x89, 0x20, 0x5e, 0x92, 0x04, 0x13, 0x84, 0x97,
	0xa4, 0xa4, 0xcc, 0xc5, 0xe9, 0x9c, 0x98, 0x13, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x22, 0x24, 0xc6,
	0xc5, 0x56, 0x04, 0x66, 0x41, 0x55, 0x43, 0x79, 0x46, 0x57, 0x18, 0xb9, 0xb8, 0x9c, 0x13, 0x73,
	0x92, 0x4b, 0x73, 0x12, 0x4b, 0xf2, 0x8b, 0x84, 0x0c, 0xb8, 0x98, 0x1d, 0x53, 0x52, 0x84, 0x84,
	0xf4, 0x60, 0xb6, 0xe9, 0xc1, 0xac, 0x92, 0x12, 0x46, 0x11, 0x83, 0x18, 0xab, 0xc4, 0x00, 0xd2,
	0x11, 0x5c, 0x9a, 0x44, 0xa2, 0x0e, 0xdf, 0xd2, 0x1c, 0x12, 0x75, 0xb8, 0x64, 0x96, 0x91, 0xa0,
	0x23, 0x89, 0x0d, 0x2c, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x00, 0x85, 0xcc, 0x4a,
	0x01, 0x00, 0x00,
}
