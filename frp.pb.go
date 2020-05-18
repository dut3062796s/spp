// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frp.proto

package main

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FrpFrame_TYPE int32

const (
	FrpFrame_LOGIN FrpFrame_TYPE = 0
	FrpFrame_DATA  FrpFrame_TYPE = 1
	FrpFrame_HB    FrpFrame_TYPE = 2
	FrpFrame_OPEN  FrpFrame_TYPE = 3
	FrpFrame_CLOSR FrpFrame_TYPE = 4
)

var FrpFrame_TYPE_name = map[int32]string{
	0: "LOGIN",
	1: "DATA",
	2: "HB",
	3: "OPEN",
	4: "CLOSR",
}

var FrpFrame_TYPE_value = map[string]int32{
	"LOGIN": 0,
	"DATA":  1,
	"HB":    2,
	"OPEN":  3,
	"CLOSR": 4,
}

func (x FrpFrame_TYPE) String() string {
	return proto.EnumName(FrpFrame_TYPE_name, int32(x))
}

func (FrpFrame_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{6, 0}
}

type FrpLoginFrame struct {
	Remote               string   `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpLoginFrame) Reset()         { *m = FrpLoginFrame{} }
func (m *FrpLoginFrame) String() string { return proto.CompactTextString(m) }
func (*FrpLoginFrame) ProtoMessage()    {}
func (*FrpLoginFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{0}
}

func (m *FrpLoginFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpLoginFrame.Unmarshal(m, b)
}
func (m *FrpLoginFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpLoginFrame.Marshal(b, m, deterministic)
}
func (m *FrpLoginFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpLoginFrame.Merge(m, src)
}
func (m *FrpLoginFrame) XXX_Size() int {
	return xxx_messageInfo_FrpLoginFrame.Size(m)
}
func (m *FrpLoginFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpLoginFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpLoginFrame proto.InternalMessageInfo

func (m *FrpLoginFrame) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

type FrpLoginRspFrame struct {
	Ret                  bool     `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpLoginRspFrame) Reset()         { *m = FrpLoginRspFrame{} }
func (m *FrpLoginRspFrame) String() string { return proto.CompactTextString(m) }
func (*FrpLoginRspFrame) ProtoMessage()    {}
func (*FrpLoginRspFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{1}
}

func (m *FrpLoginRspFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpLoginRspFrame.Unmarshal(m, b)
}
func (m *FrpLoginRspFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpLoginRspFrame.Marshal(b, m, deterministic)
}
func (m *FrpLoginRspFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpLoginRspFrame.Merge(m, src)
}
func (m *FrpLoginRspFrame) XXX_Size() int {
	return xxx_messageInfo_FrpLoginRspFrame.Size(m)
}
func (m *FrpLoginRspFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpLoginRspFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpLoginRspFrame proto.InternalMessageInfo

func (m *FrpLoginRspFrame) GetRet() bool {
	if m != nil {
		return m.Ret
	}
	return false
}

func (m *FrpLoginRspFrame) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type FrpHBFrame struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpHBFrame) Reset()         { *m = FrpHBFrame{} }
func (m *FrpHBFrame) String() string { return proto.CompactTextString(m) }
func (*FrpHBFrame) ProtoMessage()    {}
func (*FrpHBFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{2}
}

func (m *FrpHBFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpHBFrame.Unmarshal(m, b)
}
func (m *FrpHBFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpHBFrame.Marshal(b, m, deterministic)
}
func (m *FrpHBFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpHBFrame.Merge(m, src)
}
func (m *FrpHBFrame) XXX_Size() int {
	return xxx_messageInfo_FrpHBFrame.Size(m)
}
func (m *FrpHBFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpHBFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpHBFrame proto.InternalMessageInfo

func (m *FrpHBFrame) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type FrpOpenFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpOpenFrame) Reset()         { *m = FrpOpenFrame{} }
func (m *FrpOpenFrame) String() string { return proto.CompactTextString(m) }
func (*FrpOpenFrame) ProtoMessage()    {}
func (*FrpOpenFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{3}
}

func (m *FrpOpenFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpOpenFrame.Unmarshal(m, b)
}
func (m *FrpOpenFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpOpenFrame.Marshal(b, m, deterministic)
}
func (m *FrpOpenFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpOpenFrame.Merge(m, src)
}
func (m *FrpOpenFrame) XXX_Size() int {
	return xxx_messageInfo_FrpOpenFrame.Size(m)
}
func (m *FrpOpenFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpOpenFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpOpenFrame proto.InternalMessageInfo

func (m *FrpOpenFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FrpCloseFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpCloseFrame) Reset()         { *m = FrpCloseFrame{} }
func (m *FrpCloseFrame) String() string { return proto.CompactTextString(m) }
func (*FrpCloseFrame) ProtoMessage()    {}
func (*FrpCloseFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{4}
}

func (m *FrpCloseFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpCloseFrame.Unmarshal(m, b)
}
func (m *FrpCloseFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpCloseFrame.Marshal(b, m, deterministic)
}
func (m *FrpCloseFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpCloseFrame.Merge(m, src)
}
func (m *FrpCloseFrame) XXX_Size() int {
	return xxx_messageInfo_FrpCloseFrame.Size(m)
}
func (m *FrpCloseFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpCloseFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpCloseFrame proto.InternalMessageInfo

func (m *FrpCloseFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FrpDataFrame struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Compress             bool     `protobuf:"varint,2,opt,name=compress,proto3" json:"compress,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrpDataFrame) Reset()         { *m = FrpDataFrame{} }
func (m *FrpDataFrame) String() string { return proto.CompactTextString(m) }
func (*FrpDataFrame) ProtoMessage()    {}
func (*FrpDataFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{5}
}

func (m *FrpDataFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpDataFrame.Unmarshal(m, b)
}
func (m *FrpDataFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpDataFrame.Marshal(b, m, deterministic)
}
func (m *FrpDataFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpDataFrame.Merge(m, src)
}
func (m *FrpDataFrame) XXX_Size() int {
	return xxx_messageInfo_FrpDataFrame.Size(m)
}
func (m *FrpDataFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpDataFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpDataFrame proto.InternalMessageInfo

func (m *FrpDataFrame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FrpDataFrame) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

func (m *FrpDataFrame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FrpFrame struct {
	Type                 int32             `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	LoginFrame           *FrpLoginFrame    `protobuf:"bytes,2,opt,name=loginFrame,proto3" json:"loginFrame,omitempty"`
	LoginRspFrame        *FrpLoginRspFrame `protobuf:"bytes,3,opt,name=loginRspFrame,proto3" json:"loginRspFrame,omitempty"`
	DataFrame            *FrpDataFrame     `protobuf:"bytes,4,opt,name=dataFrame,proto3" json:"dataFrame,omitempty"`
	HbFrame              *FrpHBFrame       `protobuf:"bytes,5,opt,name=hbFrame,proto3" json:"hbFrame,omitempty"`
	OpenFrame            *FrpOpenFrame     `protobuf:"bytes,6,opt,name=openFrame,proto3" json:"openFrame,omitempty"`
	CloseFrame           *FrpCloseFrame    `protobuf:"bytes,7,opt,name=closeFrame,proto3" json:"closeFrame,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FrpFrame) Reset()         { *m = FrpFrame{} }
func (m *FrpFrame) String() string { return proto.CompactTextString(m) }
func (*FrpFrame) ProtoMessage()    {}
func (*FrpFrame) Descriptor() ([]byte, []int) {
	return fileDescriptor_d316925312aba732, []int{6}
}

func (m *FrpFrame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrpFrame.Unmarshal(m, b)
}
func (m *FrpFrame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrpFrame.Marshal(b, m, deterministic)
}
func (m *FrpFrame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrpFrame.Merge(m, src)
}
func (m *FrpFrame) XXX_Size() int {
	return xxx_messageInfo_FrpFrame.Size(m)
}
func (m *FrpFrame) XXX_DiscardUnknown() {
	xxx_messageInfo_FrpFrame.DiscardUnknown(m)
}

var xxx_messageInfo_FrpFrame proto.InternalMessageInfo

func (m *FrpFrame) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FrpFrame) GetLoginFrame() *FrpLoginFrame {
	if m != nil {
		return m.LoginFrame
	}
	return nil
}

func (m *FrpFrame) GetLoginRspFrame() *FrpLoginRspFrame {
	if m != nil {
		return m.LoginRspFrame
	}
	return nil
}

func (m *FrpFrame) GetDataFrame() *FrpDataFrame {
	if m != nil {
		return m.DataFrame
	}
	return nil
}

func (m *FrpFrame) GetHbFrame() *FrpHBFrame {
	if m != nil {
		return m.HbFrame
	}
	return nil
}

func (m *FrpFrame) GetOpenFrame() *FrpOpenFrame {
	if m != nil {
		return m.OpenFrame
	}
	return nil
}

func (m *FrpFrame) GetCloseFrame() *FrpCloseFrame {
	if m != nil {
		return m.CloseFrame
	}
	return nil
}

func init() {
	proto.RegisterEnum("FrpFrame_TYPE", FrpFrame_TYPE_name, FrpFrame_TYPE_value)
	proto.RegisterType((*FrpLoginFrame)(nil), "FrpLoginFrame")
	proto.RegisterType((*FrpLoginRspFrame)(nil), "FrpLoginRspFrame")
	proto.RegisterType((*FrpHBFrame)(nil), "FrpHBFrame")
	proto.RegisterType((*FrpOpenFrame)(nil), "FrpOpenFrame")
	proto.RegisterType((*FrpCloseFrame)(nil), "FrpCloseFrame")
	proto.RegisterType((*FrpDataFrame)(nil), "FrpDataFrame")
	proto.RegisterType((*FrpFrame)(nil), "FrpFrame")
}

func init() { proto.RegisterFile("frp.proto", fileDescriptor_d316925312aba732) }

var fileDescriptor_d316925312aba732 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcb, 0x6b, 0xab, 0x40,
	0x14, 0xc6, 0xaf, 0x8f, 0x18, 0x3d, 0x79, 0xe0, 0x9d, 0xc5, 0x45, 0xee, 0xa2, 0x0d, 0x42, 0x69,
	0xa0, 0xe0, 0x22, 0x85, 0xb6, 0xdb, 0xbc, 0x6c, 0x0a, 0x21, 0x86, 0xd3, 0x6c, 0xda, 0x9d, 0x89,
	0x36, 0x15, 0x62, 0x1c, 0x46, 0x37, 0xfd, 0xb7, 0xfb, 0x17, 0x94, 0x19, 0xa3, 0x63, 0x4a, 0x76,
	0xc7, 0x39, 0xbf, 0xf3, 0xcd, 0x77, 0xbe, 0x11, 0xac, 0x0f, 0x46, 0x3d, 0xca, 0xb2, 0x22, 0x73,
	0x6f, 0xa1, 0xe7, 0x33, 0xba, 0xcc, 0xf6, 0xc9, 0xd1, 0x67, 0x61, 0x1a, 0x93, 0x7f, 0x60, 0xb0,
	0x38, 0xcd, 0x8a, 0xd8, 0x51, 0x06, 0xca, 0xd0, 0xc2, 0xd3, 0x97, 0xfb, 0x00, 0x76, 0x05, 0x62,
	0x4e, 0x4b, 0xd6, 0x06, 0x8d, 0xc5, 0x85, 0x00, 0x4d, 0xe4, 0x25, 0x3f, 0x49, 0xf3, 0xbd, 0xa3,
	0x8a, 0x51, 0x5e, 0xba, 0x03, 0x00, 0x9f, 0xd1, 0xc5, 0xa4, 0x9c, 0x20, 0xa0, 0x17, 0x49, 0x5a,
	0x6a, 0x6b, 0x28, 0x6a, 0xf7, 0x0a, 0xba, 0x3e, 0xa3, 0x01, 0x8d, 0x4f, 0x0e, 0xfa, 0xa0, 0x26,
	0xd1, 0xe9, 0x76, 0x35, 0x89, 0xdc, 0x6b, 0x61, 0x71, 0x7a, 0xc8, 0xf2, 0xf8, 0x32, 0xb0, 0x12,
	0x02, 0xb3, 0xb0, 0x08, 0x2f, 0xf6, 0xc9, 0x7f, 0x30, 0x77, 0x59, 0x4a, 0x59, 0x9c, 0xe7, 0xc2,
	0x99, 0x89, 0xf5, 0x37, 0x37, 0x14, 0x85, 0x45, 0xe8, 0x68, 0x03, 0x65, 0xd8, 0x45, 0x51, 0xbb,
	0xdf, 0x2a, 0x98, 0x3e, 0xa3, 0xd2, 0xf1, 0x17, 0x2d, 0x1d, 0xb7, 0x50, 0xd4, 0xc4, 0x03, 0x38,
	0xd4, 0x89, 0x09, 0xc9, 0xce, 0xa8, 0xef, 0x9d, 0xe5, 0x88, 0x0d, 0x82, 0x3c, 0x42, 0xef, 0xd0,
	0x0c, 0x4e, 0xdc, 0xd6, 0x19, 0xfd, 0xf5, 0x7e, 0x27, 0x8a, 0xe7, 0x1c, 0xb9, 0x03, 0x2b, 0xaa,
	0xd6, 0x72, 0x74, 0x31, 0xd4, 0xf3, 0x9a, 0xbb, 0xa2, 0xec, 0x93, 0x1b, 0x68, 0x7f, 0x6e, 0x4b,
	0xb4, 0x25, 0xd0, 0x8e, 0x27, 0x93, 0xc7, 0xaa, 0xc7, 0x35, 0xb3, 0x2a, 0x6b, 0xc7, 0x90, 0x9a,
	0xf5, 0x03, 0xa0, 0xec, 0xf3, 0x4d, 0x77, 0x75, 0xf0, 0x4e, 0x5b, 0x6e, 0x2a, 0x9f, 0x03, 0x1b,
	0x84, 0xfb, 0x04, 0xfa, 0xe6, 0x6d, 0x3d, 0x27, 0x16, 0xb4, 0x96, 0xc1, 0xf3, 0xcb, 0xca, 0xfe,
	0x43, 0x4c, 0xd0, 0x67, 0xe3, 0xcd, 0xd8, 0x56, 0x88, 0x01, 0xea, 0x62, 0x62, 0xab, 0xfc, 0x24,
	0x58, 0xcf, 0x57, 0xb6, 0xc6, 0xb1, 0xe9, 0x32, 0x78, 0x45, 0x5b, 0x9f, 0x18, 0xef, 0x7a, 0x1a,
	0x26, 0xc7, 0xad, 0x21, 0xfe, 0xcb, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x50, 0xcf,
	0x1e, 0xa4, 0x02, 0x00, 0x00,
}
