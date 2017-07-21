// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	FetchAttestationDataRequest
	FetchAttestationDataResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type FetchAttestationDataRequest struct {
}

func (m *FetchAttestationDataRequest) Reset()                    { *m = FetchAttestationDataRequest{} }
func (m *FetchAttestationDataRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchAttestationDataRequest) ProtoMessage()               {}
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FetchAttestationDataResponse struct {
	AttestationData []byte `protobuf:"bytes,1,opt,name=attestationData,proto3" json:"attestationData,omitempty"`
}

func (m *FetchAttestationDataResponse) Reset()                    { *m = FetchAttestationDataResponse{} }
func (m *FetchAttestationDataResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchAttestationDataResponse) ProtoMessage()               {}
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchAttestationDataResponse) GetAttestationData() []byte {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func init() {
	proto1.RegisterType((*FetchAttestationDataRequest)(nil), "proto.FetchAttestationDataRequest")
	proto1.RegisterType((*FetchAttestationDataResponse)(nil), "proto.FetchAttestationDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error) {
	out := new(FetchAttestationDataResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/FetchAttestationData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	FetchAttestationData(context.Context, *FetchAttestationDataRequest) (*FetchAttestationDataResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAttestationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/FetchAttestationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, req.(*FetchAttestationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAttestationData",
			Handler:    _NodeAttestor_FetchAttestationData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto1.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcb, 0x4f, 0x49,
	0x8d, 0x4f, 0x2c, 0x29, 0x49, 0x2d, 0x2e, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x53, 0x4a, 0xb2, 0x5c, 0xd2, 0x6e, 0xa9, 0x25, 0xc9, 0x19, 0x8e, 0x60, 0xd9, 0xc4,
	0x92, 0xcc, 0xfc, 0x3c, 0x97, 0xc4, 0x92, 0xc4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25,
	0x0f, 0x2e, 0x19, 0xec, 0xd2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x1a, 0x5c, 0xfc, 0x89,
	0xa8, 0x52, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0xe8, 0xc2, 0x46, 0xf9, 0x5c, 0x3c, 0x7e,
	0xf9, 0x29, 0xa9, 0x8e, 0x50, 0x57, 0x08, 0xc5, 0x73, 0x89, 0x60, 0x33, 0x59, 0x48, 0x09, 0xe2,
	0x3e, 0x3d, 0x3c, 0xae, 0x92, 0x52, 0xc6, 0xab, 0x06, 0xe2, 0xb4, 0x24, 0x36, 0xb0, 0x1a, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x79, 0x7d, 0x39, 0xfe, 0x00, 0x00, 0x00,
}
