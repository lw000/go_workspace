// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	ReqRegisterService
	AckRegisterService
	ReqHeartBeat
	AckHeartBeat
	ReqMsg
	AckMsg
	ReqGroupMsg
	AckGroupMsg
*/
package example

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

type ReqRegisterService struct {
	SrvId   int32 `protobuf:"varint,1,opt,name=SrvId" json:"SrvId,omitempty"`
	SrvType int32 `protobuf:"varint,2,opt,name=SrvType" json:"SrvType,omitempty"`
}

func (m *ReqRegisterService) Reset()                    { *m = ReqRegisterService{} }
func (m *ReqRegisterService) String() string            { return proto.CompactTextString(m) }
func (*ReqRegisterService) ProtoMessage()               {}
func (*ReqRegisterService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqRegisterService) GetSrvId() int32 {
	if m != nil {
		return m.SrvId
	}
	return 0
}

func (m *ReqRegisterService) GetSrvType() int32 {
	if m != nil {
		return m.SrvType
	}
	return 0
}

type AckRegisterService struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *AckRegisterService) Reset()                    { *m = AckRegisterService{} }
func (m *AckRegisterService) String() string            { return proto.CompactTextString(m) }
func (*AckRegisterService) ProtoMessage()               {}
func (*AckRegisterService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AckRegisterService) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AckRegisterService) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReqHeartBeat struct {
	Tm int64 `protobuf:"varint,1,opt,name=tm" json:"tm,omitempty"`
}

func (m *ReqHeartBeat) Reset()                    { *m = ReqHeartBeat{} }
func (m *ReqHeartBeat) String() string            { return proto.CompactTextString(m) }
func (*ReqHeartBeat) ProtoMessage()               {}
func (*ReqHeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqHeartBeat) GetTm() int64 {
	if m != nil {
		return m.Tm
	}
	return 0
}

type AckHeartBeat struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *AckHeartBeat) Reset()                    { *m = AckHeartBeat{} }
func (m *AckHeartBeat) String() string            { return proto.CompactTextString(m) }
func (*AckHeartBeat) ProtoMessage()               {}
func (*AckHeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AckHeartBeat) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ReqMsg struct {
	FromId int64  `protobuf:"varint,1,opt,name=fromId" json:"fromId,omitempty"`
	ToId   int64  `protobuf:"varint,2,opt,name=toId" json:"toId,omitempty"`
	Msg    string `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *ReqMsg) Reset()                    { *m = ReqMsg{} }
func (m *ReqMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqMsg) ProtoMessage()               {}
func (*ReqMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqMsg) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *ReqMsg) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *ReqMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AckMsg struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *AckMsg) Reset()                    { *m = AckMsg{} }
func (m *AckMsg) String() string            { return proto.CompactTextString(m) }
func (*AckMsg) ProtoMessage()               {}
func (*AckMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AckMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ReqGroupMsg struct {
	FromId  int64  `protobuf:"varint,1,opt,name=fromId" json:"fromId,omitempty"`
	ToId    int64  `protobuf:"varint,2,opt,name=toId" json:"toId,omitempty"`
	GroupId int64  `protobuf:"varint,3,opt,name=groupId" json:"groupId,omitempty"`
	Msg     string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *ReqGroupMsg) Reset()                    { *m = ReqGroupMsg{} }
func (m *ReqGroupMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqGroupMsg) ProtoMessage()               {}
func (*ReqGroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqGroupMsg) GetFromId() int64 {
	if m != nil {
		return m.FromId
	}
	return 0
}

func (m *ReqGroupMsg) GetToId() int64 {
	if m != nil {
		return m.ToId
	}
	return 0
}

func (m *ReqGroupMsg) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *ReqGroupMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type AckGroupMsg struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *AckGroupMsg) Reset()                    { *m = AckGroupMsg{} }
func (m *AckGroupMsg) String() string            { return proto.CompactTextString(m) }
func (*AckGroupMsg) ProtoMessage()               {}
func (*AckGroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AckGroupMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqRegisterService)(nil), "example.ReqRegisterService")
	proto.RegisterType((*AckRegisterService)(nil), "example.AckRegisterService")
	proto.RegisterType((*ReqHeartBeat)(nil), "example.ReqHeartBeat")
	proto.RegisterType((*AckHeartBeat)(nil), "example.AckHeartBeat")
	proto.RegisterType((*ReqMsg)(nil), "example.ReqMsg")
	proto.RegisterType((*AckMsg)(nil), "example.AckMsg")
	proto.RegisterType((*ReqGroupMsg)(nil), "example.ReqGroupMsg")
	proto.RegisterType((*AckGroupMsg)(nil), "example.AckGroupMsg")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xfc, 0x40,
	0x10, 0xc5, 0xb9, 0xec, 0x5d, 0xc2, 0x7f, 0xee, 0xf8, 0x23, 0x8b, 0x48, 0x0a, 0x11, 0xdd, 0xca,
	0xca, 0xc6, 0xd6, 0x26, 0x22, 0x6a, 0x0a, 0x9b, 0x89, 0x5f, 0x20, 0x6e, 0xc6, 0x70, 0xc4, 0xb0,
	0xc9, 0x66, 0x0c, 0xfa, 0xed, 0x65, 0xc7, 0xc4, 0x03, 0x49, 0x63, 0xf7, 0xde, 0x2c, 0xef, 0xbd,
	0x1f, 0x2c, 0x00, 0xd3, 0xc0, 0x57, 0x9d, 0x77, 0xec, 0x74, 0x42, 0x1f, 0x65, 0xdb, 0xbd, 0x91,
	0xb9, 0x03, 0x8d, 0xd4, 0x23, 0xd5, 0xfb, 0x81, 0xc9, 0x17, 0xe4, 0xc7, 0xbd, 0x25, 0x7d, 0x0c,
	0x9b, 0xc2, 0x8f, 0x79, 0x95, 0xae, 0xce, 0x57, 0x97, 0x1b, 0xfc, 0x36, 0x3a, 0x85, 0xa4, 0xf0,
	0xe3, 0xf3, 0x67, 0x47, 0x69, 0x24, 0xf7, 0xd9, 0x9a, 0x1b, 0xd0, 0x99, 0x6d, 0x7e, 0xb7, 0x68,
	0x58, 0x5b, 0x57, 0xd1, 0x54, 0x22, 0x3a, 0xdc, 0xaa, 0x92, 0x4b, 0x29, 0xf8, 0x87, 0xa2, 0xcd,
	0x19, 0xec, 0x90, 0xfa, 0x47, 0x2a, 0x3d, 0xdf, 0x52, 0xc9, 0xfa, 0x3f, 0x44, 0xdc, 0x4a, 0x4a,
	0x61, 0xc4, 0xad, 0x31, 0xb0, 0xcb, 0x6c, 0x73, 0x78, 0x5f, 0xe8, 0x35, 0xf7, 0x10, 0x23, 0xf5,
	0x4f, 0x43, 0xad, 0x4f, 0x20, 0x7e, 0xf5, 0xae, 0x9d, 0xe0, 0x15, 0x4e, 0x2e, 0xa4, 0xd8, 0xe5,
	0x95, 0x2c, 0x2b, 0x14, 0xad, 0x8f, 0x40, 0xb5, 0x43, 0x9d, 0x2a, 0x81, 0x09, 0xd2, 0x9c, 0x42,
	0x9c, 0xd9, 0x26, 0xf4, 0x2c, 0xad, 0x10, 0x6c, 0x91, 0xfa, 0x07, 0xef, 0xde, 0xbb, 0xbf, 0x4e,
	0xa5, 0x90, 0xd4, 0x21, 0x97, 0x57, 0x32, 0xa7, 0x70, 0xb6, 0x33, 0xc4, 0xfa, 0x00, 0x71, 0x01,
	0xdb, 0xcc, 0x36, 0x3f, 0x33, 0x0b, 0x24, 0x2f, 0xb1, 0xfc, 0xe3, 0xf5, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x12, 0xa8, 0x1c, 0xd4, 0xd5, 0x01, 0x00, 0x00,
}