// Code generated by protoc-gen-go.
// source: runtime/internal/stream_chunk.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	runtime/internal/stream_chunk.proto

It has these top-level messages:
	StreamError
*/
package internal

import proto "github.com/openshift/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StreamError is a response type which is returned when
// streaming rpc returns an error.
type StreamError struct {
	GrpcCode   int32  `protobuf:"varint,1,opt,name=grpc_code,json=grpcCode" json:"grpc_code,omitempty"`
	HttpCode   int32  `protobuf:"varint,2,opt,name=http_code,json=httpCode" json:"http_code,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	HttpStatus string `protobuf:"bytes,4,opt,name=http_status,json=httpStatus" json:"http_status,omitempty"`
}

func (m *StreamError) Reset()                    { *m = StreamError{} }
func (m *StreamError) String() string            { return proto.CompactTextString(m) }
func (*StreamError) ProtoMessage()               {}
func (*StreamError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*StreamError)(nil), "grpc.gateway.runtime.StreamError")
}

func init() { proto.RegisterFile("runtime/internal/stream_chunk.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x2a, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4d, 0xd5, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x2e, 0x29,
	0x4a, 0x4d, 0xcc, 0x8d, 0x4f, 0xce, 0x28, 0xcd, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x49, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x83, 0xea,
	0x50, 0x6a, 0x62, 0xe4, 0xe2, 0x0e, 0x06, 0x2b, 0x76, 0x2d, 0x2a, 0xca, 0x2f, 0x12, 0x92, 0xe6,
	0xe2, 0x04, 0xa9, 0x8b, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0xe2,
	0x00, 0x09, 0x38, 0x03, 0xf9, 0x20, 0xc9, 0x8c, 0x92, 0x92, 0x02, 0x88, 0x24, 0x13, 0x44, 0x12,
	0x24, 0x00, 0x96, 0x94, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x60, 0x06,
	0x4a, 0x71, 0x06, 0xc1, 0xb8, 0x42, 0xf2, 0x5c, 0xdc, 0x60, 0x6d, 0xc5, 0x25, 0x89, 0x25, 0xa5,
	0xc5, 0x12, 0x2c, 0x60, 0x59, 0x2e, 0x90, 0x50, 0x30, 0x58, 0xc4, 0x89, 0x2b, 0x8a, 0x03, 0xe6,
	0xf2, 0x24, 0x36, 0xb0, 0x6b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x07, 0x92, 0xb6,
	0xd4, 0x00, 0x00, 0x00,
}
