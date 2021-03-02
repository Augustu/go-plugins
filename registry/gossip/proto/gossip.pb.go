// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Augustu/go-plugins/registry/gossip/proto/gossip.proto

package gossip

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

// Update is the message broadcast
type Update struct {
	// time to live for entry
	Expires uint64 `protobuf:"varint,1,opt,name=expires,proto3" json:"expires,omitempty"`
	// type of update
	Type int32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	// what action is taken
	Action int32 `protobuf:"varint,3,opt,name=action,proto3" json:"action,omitempty"`
	// any other associated metadata about the data
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the payload data;
	Data                 []byte   `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Update) Reset()         { *m = Update{} }
func (m *Update) String() string { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()    {}
func (*Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_e81db501087fb3b4, []int{0}
}

func (m *Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Update.Unmarshal(m, b)
}
func (m *Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Update.Marshal(b, m, deterministic)
}
func (m *Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Update.Merge(m, src)
}
func (m *Update) XXX_Size() int {
	return xxx_messageInfo_Update.Size(m)
}
func (m *Update) XXX_DiscardUnknown() {
	xxx_messageInfo_Update.DiscardUnknown(m)
}

var xxx_messageInfo_Update proto.InternalMessageInfo

func (m *Update) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *Update) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Update) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *Update) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Update) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Update)(nil), "gossip.Update")
	proto.RegisterMapType((map[string]string)(nil), "gossip.Update.MetadataEntry")
}

func init() {
	proto.RegisterFile("Augustu/go-plugins/registry/gossip/proto/gossip.proto", fileDescriptor_e81db501087fb3b4)
}

var fileDescriptor_e81db501087fb3b4 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x49, 0xb7, 0x4d, 0xed, 0xa8, 0x20, 0x83, 0x48, 0x10, 0x0f, 0x8b, 0xa7, 0xbd, 0xb8,
	0x01, 0x7b, 0x29, 0x7a, 0xf6, 0xe8, 0x25, 0xe0, 0x0f, 0x88, 0x6d, 0x08, 0xc1, 0x76, 0x13, 0x92,
	0xa9, 0x98, 0x9f, 0xea, 0xbf, 0x91, 0x26, 0x51, 0xf0, 0xf6, 0xbe, 0x99, 0x37, 0xbc, 0x37, 0xb0,
	0x3e, 0xb8, 0x6d, 0xf4, 0xd2, 0xfa, 0x87, 0xb0, 0x3f, 0x5a, 0x37, 0x25, 0x19, 0x8d, 0x75, 0x89,
	0x62, 0x96, 0xd6, 0xa7, 0xe4, 0x82, 0x0c, 0xd1, 0x93, 0x6f, 0x30, 0x16, 0x40, 0x5e, 0xe9, 0xfe,
	0x9b, 0x01, 0x7f, 0x0b, 0x3b, 0x4d, 0x06, 0x05, 0x2c, 0xcd, 0x57, 0x70, 0xd1, 0x24, 0xc1, 0x7a,
	0x36, 0xcc, 0xd5, 0x2f, 0x22, 0xc2, 0x9c, 0x72, 0x30, 0x62, 0xd6, 0xb3, 0x61, 0xa1, 0x8a, 0xc6,
	0x1b, 0xe0, 0x7a, 0x4b, 0xce, 0x4f, 0xa2, 0x2b, 0xd3, 0x46, 0xb8, 0x81, 0xb3, 0x83, 0x21, 0xbd,
	0xd3, 0xa4, 0x05, 0xef, 0xbb, 0xe1, 0xfc, 0xf1, 0x6e, 0x6c, 0xc9, 0x35, 0x67, 0x7c, 0x6d, 0xeb,
	0x97, 0x89, 0x62, 0x56, 0x7f, 0xee, 0x53, 0x4a, 0xb9, 0x5a, 0xf6, 0x6c, 0xb8, 0x50, 0x45, 0xdf,
	0x3e, 0xc3, 0xe5, 0x3f, 0x3b, 0x5e, 0x41, 0xf7, 0x61, 0x72, 0x29, 0xb8, 0x52, 0x27, 0x89, 0xd7,
	0xb0, 0xf8, 0xd4, 0xfb, 0x63, 0x6d, 0xb7, 0x52, 0x15, 0x9e, 0x66, 0x1b, 0xf6, 0xce, 0xcb, 0xab,
	0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xfb, 0xd3, 0xd6, 0x21, 0x01, 0x00, 0x00,
}
