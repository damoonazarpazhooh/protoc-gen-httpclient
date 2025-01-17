// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example.proto

package example

import (
	fmt "fmt"
	_ "github.com/bifrostcloud/protoc-gen-httpclient/proto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type StopRequest struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	SomeOther            string   `protobuf:"bytes,3,opt,name=some_other,json=someOther,proto3" json:"some_other,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopRequest) Reset()         { *m = StopRequest{} }
func (m *StopRequest) String() string { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()    {}
func (*StopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}
func (m *StopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopRequest.Unmarshal(m, b)
}
func (m *StopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopRequest.Marshal(b, m, deterministic)
}
func (m *StopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopRequest.Merge(m, src)
}
func (m *StopRequest) XXX_Size() int {
	return xxx_messageInfo_StopRequest.Size(m)
}
func (m *StopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopRequest proto.InternalMessageInfo

func (m *StopRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StopRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StopRequest) GetSomeOther() string {
	if m != nil {
		return m.SomeOther
	}
	return ""
}

type StopResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopResponse) Reset()         { *m = StopResponse{} }
func (m *StopResponse) String() string { return proto.CompactTextString(m) }
func (*StopResponse) ProtoMessage()    {}
func (*StopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}
func (m *StopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopResponse.Unmarshal(m, b)
}
func (m *StopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopResponse.Marshal(b, m, deterministic)
}
func (m *StopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopResponse.Merge(m, src)
}
func (m *StopResponse) XXX_Size() int {
	return xxx_messageInfo_StopResponse.Size(m)
}
func (m *StopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StopResponse proto.InternalMessageInfo

func (m *StopResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *StopResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type VersionRequest struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Boolean              bool     `protobuf:"varint,2,opt,name=boolean,proto3" json:"boolean,omitempty"`
	Integer              uint32   `protobuf:"varint,3,opt,name=integer,proto3" json:"integer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func (m *VersionRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *VersionRequest) GetBoolean() bool {
	if m != nil {
		return m.Boolean
	}
	return false
}

func (m *VersionRequest) GetInteger() uint32 {
	if m != nil {
		return m.Integer
	}
	return 0
}

type VersionResponse struct {
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*StopRequest)(nil), "example.StopRequest")
	proto.RegisterType((*StopResponse)(nil), "example.StopResponse")
	proto.RegisterType((*VersionRequest)(nil), "example.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "example.VersionResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0xa9, 0x69, 0x1b, 0x7b, 0xdc, 0x5a, 0x19, 0x5c, 0x0d, 0x01, 0x45, 0x23, 0xc8, 0xb2,
	0x92, 0x66, 0xad, 0x5e, 0x2c, 0x7b, 0xe3, 0xd7, 0x82, 0x78, 0x25, 0xb4, 0xe0, 0x95, 0x28, 0x93,
	0xf4, 0x34, 0x1d, 0x9a, 0xcc, 0x89, 0x99, 0x93, 0xe2, 0x3e, 0x81, 0x0f, 0xb2, 0x2f, 0xb7, 0xf8,
	0x08, 0xde, 0x28, 0x93, 0x8f, 0x65, 0x57, 0x11, 0x74, 0x2f, 0xff, 0xff, 0x13, 0x7e, 0xbf, 0x33,
	0x13, 0x06, 0xc6, 0xf8, 0x55, 0xe6, 0x45, 0x86, 0xd3, 0xa2, 0x24, 0x26, 0xe1, 0xb6, 0xd1, 0x7f,
	0x91, 0x2a, 0x5e, 0x57, 0xf1, 0x34, 0xa1, 0x3c, 0x8a, 0xd5, 0xaa, 0x24, 0xc3, 0x49, 0x46, 0xd5,
	0x32, 0xaa, 0xbf, 0x4b, 0xc2, 0x14, 0x75, 0xb8, 0x66, 0x2e, 0x92, 0x4c, 0xa1, 0xe6, 0xa6, 0x8d,
	0xf8, 0xa4, 0x40, 0xd3, 0x90, 0x82, 0x4f, 0x70, 0x63, 0xc1, 0x54, 0xcc, 0xf1, 0x4b, 0x85, 0x86,
	0xc5, 0x1d, 0x18, 0x1a, 0x96, 0x5c, 0x19, 0xaf, 0xf7, 0xa0, 0xb7, 0x37, 0x9a, 0xb7, 0x49, 0x78,
	0xe0, 0xe6, 0x68, 0x8c, 0x4c, 0xd1, 0xbb, 0x56, 0x0f, 0xba, 0x28, 0xee, 0x01, 0x18, 0xca, 0xf1,
	0x33, 0xf1, 0x1a, 0x4b, 0xcf, 0xa9, 0x87, 0x23, 0xdb, 0xbc, 0xb7, 0x45, 0xf0, 0x12, 0x76, 0x1a,
	0xbe, 0x29, 0x48, 0x1b, 0xfc, 0x7f, 0x41, 0xf0, 0x11, 0x6e, 0x7e, 0xc0, 0xd2, 0x28, 0xd2, 0xff,
	0xb0, 0x64, 0x4c, 0x94, 0xa1, 0xd4, 0x35, 0xe3, 0xfa, 0xbc, 0x8b, 0x76, 0xa2, 0x34, 0x63, 0xda,
	0x6e, 0x38, 0x9e, 0x77, 0x31, 0x78, 0x02, 0x93, 0x73, 0x7a, 0xbb, 0xa2, 0x07, 0xee, 0xb6, 0xa9,
	0xda, 0xe3, 0x74, 0x71, 0xf6, 0xd3, 0x81, 0xf1, 0xb1, 0xc4, 0x9c, 0xf4, 0x02, 0xcb, 0xad, 0x4a,
	0x50, 0x28, 0xe8, 0xdb, 0xe3, 0x89, 0xdb, 0xd3, 0xee, 0x07, 0x5d, 0xb8, 0x4d, 0x7f, 0xf7, 0xb7,
	0xb6, 0x11, 0x04, 0xcf, 0xbf, 0x9f, 0x9e, 0x39, 0x11, 0x4c, 0x12, 0x55, 0x26, 0x95, 0xe2, 0x30,
	0x2e, 0x51, 0x6e, 0xb0, 0x14, 0x3b, 0xd1, 0xb2, 0x86, 0x47, 0x86, 0xa9, 0xf0, 0x47, 0x0b, 0x25,
	0xc3, 0x57, 0x29, 0x6a, 0x0e, 0x9c, 0x14, 0x59, 0x6c, 0xc0, 0x6d, 0x37, 0x15, 0x77, 0xcf, 0xb9,
	0x97, 0x6f, 0xc6, 0xf7, 0xfe, 0x1c, 0xb4, 0xce, 0xd0, 0x3a, 0xf7, 0x60, 0x10, 0x4b, 0xa3, 0x12,
	0x31, 0xe9, 0x4c, 0xed, 0xc9, 0x2e, 0xc9, 0x8a, 0x8a, 0x45, 0x0a, 0x70, 0x8c, 0x19, 0x32, 0xbe,
	0xd3, 0x2b, 0xba, 0x8a, 0xef, 0xb1, 0xf5, 0x3d, 0xec, 0x7c, 0x83, 0x68, 0x4d, 0x86, 0x2f, 0x5a,
	0x86, 0xcb, 0x1a, 0x2f, 0xbe, 0xf5, 0x60, 0xf4, 0x86, 0xf4, 0x4a, 0xa5, 0x55, 0x89, 0x57, 0x11,
	0xbd, 0xb5, 0xa2, 0xd7, 0x7f, 0x17, 0xf5, 0x0b, 0x32, 0xbc, 0x3f, 0x81, 0xfe, 0x06, 0x4f, 0x42,
	0xe1, 0x6e, 0x65, 0x56, 0x61, 0xf8, 0x74, 0xff, 0x16, 0x0c, 0x6c, 0x31, 0xeb, 0x9a, 0x99, 0xff,
	0xe8, 0xc7, 0xe9, 0x99, 0x73, 0x1f, 0x76, 0xed, 0x73, 0x39, 0x8a, 0xa2, 0x8c, 0x12, 0x99, 0x59,
	0xd8, 0xd1, 0xe1, 0xc1, 0xe1, 0x81, 0x68, 0xf8, 0xf1, 0xb0, 0x7e, 0x35, 0xcf, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xb1, 0xfe, 0x63, 0x74, 0x90, 0x03, 0x00, 0x00,
}
