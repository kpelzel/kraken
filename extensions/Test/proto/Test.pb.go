// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Test.proto

package proto

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

type Test_Test int32

const (
	Test_UNKNOWN Test_Test = 0
	Test_LOW     Test_Test = 1
	Test_MED     Test_Test = 2
	Test_HIGH    Test_Test = 3
)

var Test_Test_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOW",
	2: "MED",
	3: "HIGH",
}
var Test_Test_value = map[string]int32{
	"UNKNOWN": 0,
	"LOW":     1,
	"MED":     2,
	"HIGH":    3,
}

func (x Test_Test) String() string {
	return proto.EnumName(Test_Test_name, int32(x))
}
func (Test_Test) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Test_8c9f0586176b7a63, []int{0, 0}
}

type Test struct {
	State                Test_Test `protobuf:"varint,1,opt,name=state,proto3,enum=proto.Test_Test" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_Test_8c9f0586176b7a63, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetState() Test_Test {
	if m != nil {
		return m.State
	}
	return Test_UNKNOWN
}

func init() {
	proto.RegisterType((*Test)(nil), "proto.Test")
	proto.RegisterEnum("proto.Test_Test", Test_Test_name, Test_Test_value)
}

func init() { proto.RegisterFile("Test.proto", fileDescriptor_Test_8c9f0586176b7a63) }

var fileDescriptor_Test_8c9f0586176b7a63 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xf1, 0x5c, 0x2c, 0x20, 0x41,
	0x21, 0x35, 0x2e, 0xd6, 0xe2, 0x92, 0xc4, 0x92, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23,
	0x01, 0x88, 0x2a, 0x3d, 0xb0, 0x06, 0x10, 0x11, 0x04, 0x91, 0x56, 0xd2, 0x87, 0xaa, 0xe7, 0xe6,
	0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0x62, 0xe7, 0x62, 0xf6, 0xf1,
	0x0f, 0x17, 0x60, 0x04, 0x31, 0x7c, 0x5d, 0x5d, 0x04, 0x98, 0x84, 0x38, 0xb8, 0x58, 0x3c, 0x3c,
	0xdd, 0x3d, 0x04, 0x98, 0x93, 0xd8, 0xc0, 0x06, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0xaf, 0x41, 0xae, 0x7c, 0x00, 0x00, 0x00,
}
