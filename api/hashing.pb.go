// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hashing.proto

package api

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

type HashRequest struct {
	Plain                []byte   `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashRequest) Reset()         { *m = HashRequest{} }
func (m *HashRequest) String() string { return proto.CompactTextString(m) }
func (*HashRequest) ProtoMessage()    {}
func (*HashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hashing_3ac2c063ec583c21, []int{0}
}
func (m *HashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashRequest.Unmarshal(m, b)
}
func (m *HashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashRequest.Marshal(b, m, deterministic)
}
func (dst *HashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashRequest.Merge(dst, src)
}
func (m *HashRequest) XXX_Size() int {
	return xxx_messageInfo_HashRequest.Size(m)
}
func (m *HashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HashRequest proto.InternalMessageInfo

func (m *HashRequest) GetPlain() []byte {
	if m != nil {
		return m.Plain
	}
	return nil
}

type HashResponse struct {
	Hash                 uint64   `protobuf:"varint,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashResponse) Reset()         { *m = HashResponse{} }
func (m *HashResponse) String() string { return proto.CompactTextString(m) }
func (*HashResponse) ProtoMessage()    {}
func (*HashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_hashing_3ac2c063ec583c21, []int{1}
}
func (m *HashResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashResponse.Unmarshal(m, b)
}
func (m *HashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashResponse.Marshal(b, m, deterministic)
}
func (dst *HashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashResponse.Merge(dst, src)
}
func (m *HashResponse) XXX_Size() int {
	return xxx_messageInfo_HashResponse.Size(m)
}
func (m *HashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HashResponse proto.InternalMessageInfo

func (m *HashResponse) GetHash() uint64 {
	if m != nil {
		return m.Hash
	}
	return 0
}

func init() {
	proto.RegisterType((*HashRequest)(nil), "api.HashRequest")
	proto.RegisterType((*HashResponse)(nil), "api.HashResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HashingClient is the client API for Hashing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HashingClient interface {
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
}

type hashingClient struct {
	cc *grpc.ClientConn
}

func NewHashingClient(cc *grpc.ClientConn) HashingClient {
	return &hashingClient{cc}
}

func (c *hashingClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	out := new(HashResponse)
	err := c.cc.Invoke(ctx, "/api.Hashing/Hash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashingServer is the server API for Hashing service.
type HashingServer interface {
	Hash(context.Context, *HashRequest) (*HashResponse, error)
}

func RegisterHashingServer(s *grpc.Server, srv HashingServer) {
	s.RegisterService(&_Hashing_serviceDesc, srv)
}

func _Hashing_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashingServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Hashing/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashingServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hashing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Hashing",
	HandlerType: (*HashingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _Hashing_Hash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hashing.proto",
}

func init() { proto.RegisterFile("hashing.proto", fileDescriptor_hashing_3ac2c063ec583c21) }

var fileDescriptor_hashing_3ac2c063ec583c21 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x48, 0x2c, 0xce,
	0xc8, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52,
	0xe6, 0xe2, 0xf6, 0x48, 0x2c, 0xce, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1,
	0x62, 0x2d, 0xc8, 0x49, 0xcc, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x70, 0x94,
	0x94, 0xb8, 0x78, 0x20, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x40,
	0x46, 0x81, 0x15, 0xb1, 0x04, 0x81, 0xd9, 0x46, 0x66, 0x5c, 0xec, 0x1e, 0x10, 0xe3, 0x85, 0xb4,
	0xb9, 0x58, 0x40, 0x4c, 0x21, 0x01, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x24, 0xe3, 0xa5, 0x04, 0x91,
	0x44, 0x20, 0x66, 0x25, 0xb1, 0x81, 0x1d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x32,
	0x9b, 0xd4, 0x9d, 0x00, 0x00, 0x00,
}
