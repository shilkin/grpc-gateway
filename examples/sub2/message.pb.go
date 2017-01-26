// Code generated by protoc-gen-go.
// source: examples/sub2/message.proto
// DO NOT EDIT!

/*
Package sub2 is a generated protocol buffer package.

It is generated from these files:
	examples/sub2/message.proto

It has these top-level messages:
	IdMessage
*/
package sub2

import proto "github.com/golang/protobuf/proto"
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

type IdMessage struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *IdMessage) Reset()                    { *m = IdMessage{} }
func (m *IdMessage) String() string            { return proto.CompactTextString(m) }
func (*IdMessage) ProtoMessage()               {}
func (*IdMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdMessage) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func init() {
	proto.RegisterType((*IdMessage)(nil), "sub2.IdMessage")
}

func init() { proto.RegisterFile("examples/sub2/message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0x4d, 0x32, 0xd2, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x29, 0xc9, 0x73, 0x71, 0x7a,
	0xa6, 0xf8, 0x42, 0x24, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x4b, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x27, 0xfd, 0x28, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0xfd, 0xe2, 0x8c, 0xcc, 0x9c, 0xec, 0xcc, 0x3c, 0xfd, 0xf4, 0xa2, 0x82, 0x64,
	0xdd, 0xf4, 0xc4, 0x92, 0xd4, 0xf2, 0xc4, 0x4a, 0x7d, 0x14, 0x5b, 0x92, 0xd8, 0xc0, 0xc6, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x39, 0x5e, 0xb3, 0xa1, 0x7d, 0x00, 0x00, 0x00,
}
