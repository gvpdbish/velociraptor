// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "www.velocidex.com/golang/velociraptor/actions/proto"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Artifacts struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_e251903582d0bcfb, []int{0}
}
func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (dst *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(dst, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ArtifactParameters struct {
	Env                  []*proto1.VQLEnv `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Files                []*proto1.VQLEnv `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactParameters) Reset()         { *m = ArtifactParameters{} }
func (m *ArtifactParameters) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameters) ProtoMessage()    {}
func (*ArtifactParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_e251903582d0bcfb, []int{1}
}
func (m *ArtifactParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameters.Unmarshal(m, b)
}
func (m *ArtifactParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameters.Marshal(b, m, deterministic)
}
func (dst *ArtifactParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameters.Merge(dst, src)
}
func (m *ArtifactParameters) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameters.Size(m)
}
func (m *ArtifactParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameters proto.InternalMessageInfo

func (m *ArtifactParameters) GetEnv() []*proto1.VQLEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ArtifactParameters) GetFiles() []*proto1.VQLEnv {
	if m != nil {
		return m.Files
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Artifacts            *Artifacts          `protobuf:"bytes,1,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Parameters           *ArtifactParameters `protobuf:"bytes,5,opt,name=parameters,proto3" json:"parameters,omitempty"`
	OpsPerSecond         float32             `protobuf:"fixed32,6,opt,name=ops_per_second,json=opsPerSecond,proto3" json:"ops_per_second,omitempty"`
	Timeout              uint64              `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_e251903582d0bcfb, []int{2}
}
func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (dst *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(dst, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetArtifacts() *Artifacts {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetParameters() *ArtifactParameters {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ArtifactCollectorArgs) GetOpsPerSecond() float32 {
	if m != nil {
		return m.OpsPerSecond
	}
	return 0
}

func (m *ArtifactCollectorArgs) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// This is stored in the ArtifactCollector state.
type ArtifactCompressionDict struct {
	Substs               []*proto1.VQLEnv `protobuf:"bytes,1,rep,name=substs,proto3" json:"substs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ArtifactCompressionDict) Reset()         { *m = ArtifactCompressionDict{} }
func (m *ArtifactCompressionDict) String() string { return proto.CompactTextString(m) }
func (*ArtifactCompressionDict) ProtoMessage()    {}
func (*ArtifactCompressionDict) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_e251903582d0bcfb, []int{3}
}
func (m *ArtifactCompressionDict) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCompressionDict.Unmarshal(m, b)
}
func (m *ArtifactCompressionDict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCompressionDict.Marshal(b, m, deterministic)
}
func (dst *ArtifactCompressionDict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCompressionDict.Merge(dst, src)
}
func (m *ArtifactCompressionDict) XXX_Size() int {
	return xxx_messageInfo_ArtifactCompressionDict.Size(m)
}
func (m *ArtifactCompressionDict) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCompressionDict.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCompressionDict proto.InternalMessageInfo

func (m *ArtifactCompressionDict) GetSubsts() []*proto1.VQLEnv {
	if m != nil {
		return m.Substs
	}
	return nil
}

func init() {
	proto.RegisterType((*Artifacts)(nil), "proto.Artifacts")
	proto.RegisterType((*ArtifactParameters)(nil), "proto.ArtifactParameters")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
	proto.RegisterType((*ArtifactCompressionDict)(nil), "proto.ArtifactCompressionDict")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_e251903582d0bcfb) }

var fileDescriptor_artifacts_e251903582d0bcfb = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0x15, 0xba, 0xb6, 0x9a, 0x07, 0x03, 0x59, 0x42, 0x84, 0x89, 0x8b, 0xa3, 0x20, 0xa4,
	0xc0, 0x45, 0x5a, 0x3a, 0xb1, 0x0b, 0xa4, 0x21, 0xd6, 0xad, 0x70, 0xc1, 0xc4, 0x46, 0x98, 0x10,
	0x77, 0x93, 0x9b, 0x9e, 0xb4, 0x06, 0xc7, 0xce, 0x8e, 0xdd, 0x6c, 0x3c, 0x16, 0xef, 0xc0, 0x93,
	0xc0, 0x0d, 0x0f, 0xc1, 0x05, 0xaa, 0xd3, 0x7f, 0x4c, 0x20, 0x91, 0x9b, 0x44, 0xc7, 0xf6, 0xf7,
	0xfd, 0x8e, 0x73, 0xd8, 0x6d, 0x41, 0x4e, 0xe6, 0x22, 0x73, 0x36, 0x29, 0xc9, 0x38, 0xc3, 0x9b,
	0xfe, 0xb5, 0xf3, 0xfc, 0xf2, 0xf2, 0x32, 0xa9, 0x50, 0x99, 0x4c, 0x8e, 0xf0, 0x2a, 0xc9, 0x4c,
	0xd1, 0x19, 0x1b, 0x25, 0xf4, 0xb8, 0x53, 0x17, 0x49, 0x94, 0xce, 0x50, 0xc7, 0x6f, 0xee, 0x58,
	0x2c, 0x84, 0x76, 0x32, 0xab, 0x11, 0x3b, 0xfb, 0xff, 0x77, 0x56, 0x64, 0x4e, 0x1a, 0x6d, 0xe7,
	0x8c, 0xea, 0x42, 0xd5, 0xc7, 0xa3, 0x43, 0xb6, 0x79, 0xb0, 0x08, 0xc5, 0xf7, 0x58, 0x53, 0x8b,
	0x02, 0x6d, 0x18, 0x40, 0x23, 0xde, 0xec, 0xc3, 0xf7, 0x5f, 0x3f, 0xbe, 0x05, 0x3b, 0x3c, 0x3c,
	0x9b, 0x20, 0x2c, 0xa3, 0x83, 0x33, 0xa0, 0xc4, 0x54, 0x67, 0x93, 0x24, 0xad, 0xb7, 0x47, 0x3f,
	0x03, 0xc6, 0x17, 0x94, 0x53, 0x41, 0xa2, 0x40, 0x87, 0x64, 0xf9, 0x88, 0x35, 0x50, 0x57, 0x61,
	0x03, 0x1a, 0xf1, 0x56, 0xef, 0x56, 0x2d, 0x4c, 0x3e, 0xbc, 0x3b, 0x1e, 0xe8, 0xaa, 0x7f, 0xe8,
	0xd9, 0xfb, 0x7c, 0x77, 0xa0, 0x2b, 0x49, 0x46, 0x17, 0xa8, 0x1d, 0x54, 0x82, 0xa4, 0x18, 0x2a,
	0xf4, 0x8e, 0x21, 0x42, 0x49, 0xa6, 0x92, 0x23, 0x1c, 0x41, 0x6e, 0x08, 0xdc, 0x04, 0xe1, 0x62,
	0x8a, 0xf4, 0x25, 0x89, 0x5a, 0x5e, 0x62, 0xd3, 0x19, 0x9e, 0x2b, 0xd6, 0xcc, 0xa5, 0x42, 0x1b,
	0x6e, 0xfc, 0xcd, 0xf3, 0xda, 0x7b, 0x0e, 0x78, 0x67, 0x95, 0x6b, 0x0e, 0xcf, 0xa5, 0x52, 0x38,
	0x02, 0xa9, 0x21, 0x27, 0x53, 0x00, 0x5e, 0x95, 0x86, 0xdc, 0xcc, 0x35, 0x83, 0x25, 0xd1, 0xf6,
	0x60, 0x51, 0x78, 0x35, 0x2b, 0xa4, 0xb5, 0x24, 0xfa, 0xda, 0x60, 0x77, 0x17, 0xad, 0x1e, 0x1a,
	0xa5, 0x30, 0x73, 0x86, 0x0e, 0x68, 0x6c, 0xf9, 0x47, 0xb6, 0xb9, 0xbc, 0xa3, 0x30, 0x80, 0x20,
	0xde, 0xea, 0xdd, 0x99, 0x67, 0x59, 0xde, 0x70, 0x3f, 0xf6, 0x71, 0xa2, 0x7f, 0x5f, 0x69, 0xd4,
	0x3a, 0xf6, 0x1f, 0xe9, 0x0a, 0xc6, 0x3f, 0x33, 0x56, 0x2e, 0xd3, 0x87, 0x4d, 0x8f, 0xbe, 0x7f,
	0x0d, 0xbd, 0x6a, 0xaf, 0xdf, 0xf5, 0x8e, 0x27, 0xfc, 0xc1, 0x9f, 0x2d, 0xbb, 0x75, 0x63, 0x12,
	0xb1, 0xd5, 0x6a, 0xba, 0x86, 0xe7, 0x9f, 0xd8, 0xb6, 0x29, 0xed, 0x79, 0x89, 0x74, 0x6e, 0x31,
	0x33, 0x7a, 0x14, 0xb6, 0x20, 0x88, 0x6f, 0xf4, 0x8f, 0x3c, 0xf5, 0x05, 0x7f, 0x78, 0x52, 0x22,
	0x09, 0x3f, 0x4c, 0x50, 0x22, 0x41, 0xbd, 0x09, 0xe2, 0xb3, 0x09, 0x19, 0xe7, 0x94, 0xd4, 0xe3,
	0xc7, 0x49, 0xb4, 0x7d, 0x52, 0x5a, 0x38, 0x45, 0x82, 0xf7, 0x7e, 0xb5, 0xd7, 0x7e, 0xda, 0xf5,
	0x4f, 0x7a, 0xd3, 0x94, 0xf6, 0x14, 0xa9, 0x2e, 0x73, 0x64, 0x6d, 0x27, 0x0b, 0x34, 0x53, 0x17,
	0xb6, 0x21, 0x88, 0x37, 0xfa, 0x6f, 0xbc, 0x64, 0xc0, 0x9f, 0xbd, 0x9d, 0x16, 0x43, 0x24, 0x30,
	0xf9, 0x9c, 0xef, 0x3b, 0xa0, 0xa9, 0x86, 0x21, 0xe6, 0x86, 0x10, 0x32, 0xa1, 0x33, 0x54, 0x33,
	0xdb, 0xfa, 0x5c, 0xb4, 0xcf, 0x6a, 0x5a, 0xaf, 0xb1, 0xd7, 0xed, 0xa6, 0x0b, 0x76, 0xf4, 0x92,
	0xdd, 0x5b, 0xfd, 0xb2, 0xa2, 0x24, 0xb4, 0x56, 0x1a, 0x7d, 0x24, 0x33, 0xc7, 0x1f, 0xb1, 0x96,
	0x9d, 0x0e, 0xad, 0xab, 0x47, 0xfe, 0xfa, 0xf4, 0xa4, 0xf3, 0xc5, 0x61, 0xcb, 0x57, 0x77, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x8e, 0x31, 0x9b, 0xc1, 0x03, 0x00, 0x00,
}
