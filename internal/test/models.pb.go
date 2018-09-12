// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import geo "github.com/brocaar/loraserver/api/geo"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResolveTDOATestSuite struct {
	Tests                []*TDOATest `protobuf:"bytes,1,rep,name=tests,proto3" json:"tests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResolveTDOATestSuite) Reset()         { *m = ResolveTDOATestSuite{} }
func (m *ResolveTDOATestSuite) String() string { return proto.CompactTextString(m) }
func (*ResolveTDOATestSuite) ProtoMessage()    {}
func (*ResolveTDOATestSuite) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}
func (m *ResolveTDOATestSuite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveTDOATestSuite.Unmarshal(m, b)
}
func (m *ResolveTDOATestSuite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveTDOATestSuite.Marshal(b, m, deterministic)
}
func (dst *ResolveTDOATestSuite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveTDOATestSuite.Merge(dst, src)
}
func (m *ResolveTDOATestSuite) XXX_Size() int {
	return xxx_messageInfo_ResolveTDOATestSuite.Size(m)
}
func (m *ResolveTDOATestSuite) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveTDOATestSuite.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveTDOATestSuite proto.InternalMessageInfo

func (m *ResolveTDOATestSuite) GetTests() []*TDOATest {
	if m != nil {
		return m.Tests
	}
	return nil
}

type TDOATest struct {
	Request              *geo.ResolveTDOARequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	ExpectedResult       *geo.ResolveResult      `protobuf:"bytes,2,opt,name=expected_result,json=expectedResult,proto3" json:"expected_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TDOATest) Reset()         { *m = TDOATest{} }
func (m *TDOATest) String() string { return proto.CompactTextString(m) }
func (*TDOATest) ProtoMessage()    {}
func (*TDOATest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}
func (m *TDOATest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TDOATest.Unmarshal(m, b)
}
func (m *TDOATest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TDOATest.Marshal(b, m, deterministic)
}
func (dst *TDOATest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TDOATest.Merge(dst, src)
}
func (m *TDOATest) XXX_Size() int {
	return xxx_messageInfo_TDOATest.Size(m)
}
func (m *TDOATest) XXX_DiscardUnknown() {
	xxx_messageInfo_TDOATest.DiscardUnknown(m)
}

var xxx_messageInfo_TDOATest proto.InternalMessageInfo

func (m *TDOATest) GetRequest() *geo.ResolveTDOARequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *TDOATest) GetExpectedResult() *geo.ResolveResult {
	if m != nil {
		return m.ExpectedResult
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveTDOATestSuite)(nil), "test.ResolveTDOATestSuite")
	proto.RegisterType((*TDOATest)(nil), "test.TDOATest")
}

func init() { proto.RegisterFile("models.proto", fileDescriptor_0b5431a010549573) }

var fileDescriptor_0b5431a010549573 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0xff, 0xc9, 0xca, 0x0a, 0x41, 0xb0, 0x78, 0x5a, 0x16, 0x0f, 0x7b, 0x4a, 0x70,
	0x3d, 0xea, 0x45, 0xf0, 0x2e, 0xc4, 0xbd, 0x4b, 0xda, 0x1d, 0x6a, 0x21, 0x65, 0xea, 0xcc, 0xa4,
	0x88, 0x9f, 0x7e, 0x49, 0x43, 0xa0, 0x87, 0x39, 0xbc, 0x79, 0xef, 0xf7, 0x86, 0x51, 0xb7, 0x03,
	0x1e, 0x21, 0xb0, 0x19, 0x09, 0x05, 0xf5, 0x85, 0x00, 0xcb, 0xa3, 0xed, 0x7a, 0xf9, 0x89, 0x8d,
	0x69, 0x71, 0xb0, 0x0d, 0x61, 0xeb, 0x3d, 0xd9, 0x80, 0xe4, 0x19, 0x68, 0x02, 0xb2, 0x7e, 0xec,
	0x6d, 0x07, 0x98, 0x26, 0x63, 0xdb, 0x37, 0x75, 0xef, 0x80, 0x31, 0x4c, 0x70, 0xf8, 0xf8, 0x7c,
	0x3f, 0x00, 0xcb, 0x57, 0xec, 0x05, 0xf4, 0x93, 0xba, 0x4c, 0x85, 0x5c, 0x57, 0x9b, 0xf3, 0xdd,
	0x6a, 0xbf, 0x36, 0x49, 0x99, 0x92, 0x71, 0xd9, 0xdc, 0xfe, 0xab, 0x9b, 0xb2, 0xd2, 0xcf, 0xea,
	0x9a, 0xe0, 0x37, 0x02, 0x4b, 0x5d, 0x6d, 0xaa, 0xdd, 0x6a, 0xff, 0x60, 0xd2, 0x99, 0x45, 0xbb,
	0xcb, 0xb6, 0x2b, 0x39, 0xfd, 0xaa, 0xee, 0xe0, 0x6f, 0x84, 0x56, 0xe0, 0xf8, 0x4d, 0xc0, 0x31,
	0x48, 0x7d, 0x36, 0xa3, 0x7a, 0x89, 0xba, 0xd9, 0x71, 0xeb, 0x12, 0xcd, 0xba, 0xb9, 0x9a, 0x1f,
	0x78, 0x39, 0x05, 0x00, 0x00, 0xff, 0xff, 0x34, 0x35, 0x56, 0xe1, 0x07, 0x01, 0x00, 0x00,
}
