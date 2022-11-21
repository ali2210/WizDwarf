// Code generated by protoc-gen-go. DO NOT EDIT.
// source: other/cloudmedia/media/media.proto

/*
Package media is a generated protocol buffer package.

It is generated from these files:
	other/cloudmedia/media/media.proto

It has these top-level messages:
	IMAGE_METADATA
	MediaStream
*/
package media

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

type TOKEN_CATEGORY int32

const (
	TOKEN_CATEGORY_NON_FUNGIABLE_TOKEN TOKEN_CATEGORY = 0
	TOKEN_CATEGORY_FUNGIABLE_TOKEN     TOKEN_CATEGORY = 1
	TOKEN_CATEGORY_SOUL_BOUND_TOKEN    TOKEN_CATEGORY = 2
)

var TOKEN_CATEGORY_name = map[int32]string{
	0: "NON_FUNGIABLE_TOKEN",
	1: "FUNGIABLE_TOKEN",
	2: "SOUL_BOUND_TOKEN",
}
var TOKEN_CATEGORY_value = map[string]int32{
	"NON_FUNGIABLE_TOKEN": 0,
	"FUNGIABLE_TOKEN":     1,
	"SOUL_BOUND_TOKEN":    2,
}

func (x TOKEN_CATEGORY) String() string {
	return proto.EnumName(TOKEN_CATEGORY_name, int32(x))
}
func (TOKEN_CATEGORY) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Descriptor_Category int32

const (
	Descriptor_Category_Text  Descriptor_Category = 0
	Descriptor_Category_Video Descriptor_Category = 1
	Descriptor_Category_Audio Descriptor_Category = 2
	Descriptor_Category_Image Descriptor_Category = 3
)

var Descriptor_Category_name = map[int32]string{
	0: "Text",
	1: "Video",
	2: "Audio",
	3: "Image",
}
var Descriptor_Category_value = map[string]int32{
	"Text":  0,
	"Video": 1,
	"Audio": 2,
	"Image": 3,
}

func (x Descriptor_Category) String() string {
	return proto.EnumName(Descriptor_Category_name, int32(x))
}
func (Descriptor_Category) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IMAGE_METADATA struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type      string            `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Created   string            `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	Tokens    TOKEN_CATEGORY    `protobuf:"varint,4,opt,name=tokens,enum=media.TOKEN_CATEGORY" json:"tokens,omitempty"`
	MyProfile bool              `protobuf:"varint,5,opt,name=MyProfile,json=myProfile" json:"MyProfile,omitempty"`
	Cdr       map[string]string `protobuf:"bytes,6,rep,name=cdr" json:"cdr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Timeline  string            `protobuf:"bytes,7,opt,name=timeline" json:"timeline,omitempty"`
	Tags      string            `protobuf:"bytes,8,opt,name=tags" json:"tags,omitempty"`
	Signature []string          `protobuf:"bytes,9,rep,name=Signature,json=signature" json:"Signature,omitempty"`
}

func (m *IMAGE_METADATA) Reset()                    { *m = IMAGE_METADATA{} }
func (m *IMAGE_METADATA) String() string            { return proto.CompactTextString(m) }
func (*IMAGE_METADATA) ProtoMessage()               {}
func (*IMAGE_METADATA) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IMAGE_METADATA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IMAGE_METADATA) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *IMAGE_METADATA) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *IMAGE_METADATA) GetTokens() TOKEN_CATEGORY {
	if m != nil {
		return m.Tokens
	}
	return TOKEN_CATEGORY_NON_FUNGIABLE_TOKEN
}

func (m *IMAGE_METADATA) GetMyProfile() bool {
	if m != nil {
		return m.MyProfile
	}
	return false
}

func (m *IMAGE_METADATA) GetCdr() map[string]string {
	if m != nil {
		return m.Cdr
	}
	return nil
}

func (m *IMAGE_METADATA) GetTimeline() string {
	if m != nil {
		return m.Timeline
	}
	return ""
}

func (m *IMAGE_METADATA) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *IMAGE_METADATA) GetSignature() []string {
	if m != nil {
		return m.Signature
	}
	return nil
}

type MediaStream struct {
	Name         string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IdentityCode string              `protobuf:"bytes,2,opt,name=identity_code,json=identityCode" json:"identity_code,omitempty"`
	Datecreated  string              `protobuf:"bytes,3,opt,name=datecreated" json:"datecreated,omitempty"`
	Path         string              `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	Category     Descriptor_Category `protobuf:"varint,5,opt,name=category,enum=media.Descriptor_Category" json:"category,omitempty"`
	Signature    []string            `protobuf:"bytes,6,rep,name=Signature,json=signature" json:"Signature,omitempty"`
}

func (m *MediaStream) Reset()                    { *m = MediaStream{} }
func (m *MediaStream) String() string            { return proto.CompactTextString(m) }
func (*MediaStream) ProtoMessage()               {}
func (*MediaStream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MediaStream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MediaStream) GetIdentityCode() string {
	if m != nil {
		return m.IdentityCode
	}
	return ""
}

func (m *MediaStream) GetDatecreated() string {
	if m != nil {
		return m.Datecreated
	}
	return ""
}

func (m *MediaStream) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MediaStream) GetCategory() Descriptor_Category {
	if m != nil {
		return m.Category
	}
	return Descriptor_Category_Text
}

func (m *MediaStream) GetSignature() []string {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*IMAGE_METADATA)(nil), "media.IMAGE_METADATA")
	proto.RegisterType((*MediaStream)(nil), "media.MediaStream")
	proto.RegisterEnum("media.TOKEN_CATEGORY", TOKEN_CATEGORY_name, TOKEN_CATEGORY_value)
	proto.RegisterEnum("media.Descriptor_Category", Descriptor_Category_name, Descriptor_Category_value)
}

func init() { proto.RegisterFile("other/cloudmedia/media/media.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x2f, 0xcd, 0xb5, 0x97, 0x4c, 0xb5, 0x86, 0xed, 0x89, 0xa1, 0x88, 0x84, 0xfa, 0x52,
	0x0e, 0xec, 0x49, 0x85, 0x43, 0x7c, 0x32, 0xd7, 0xc6, 0x52, 0xbc, 0xb6, 0x92, 0xa6, 0x82, 0x4f,
	0x61, 0xcd, 0x8e, 0xb9, 0xe5, 0x9a, 0x6c, 0xd9, 0x6e, 0xc5, 0x7c, 0x4d, 0x5f, 0xfc, 0x3a, 0x92,
	0x6d, 0x72, 0x5c, 0x8f, 0x7b, 0x09, 0xff, 0xf9, 0x65, 0x98, 0x9d, 0xfd, 0x25, 0xd0, 0x17, 0xea,
	0x16, 0xe5, 0x65, 0xb2, 0x11, 0x7b, 0x96, 0x21, 0xe3, 0xf4, 0xf2, 0xc1, 0x73, 0xb8, 0x95, 0x42,
	0x09, 0xd2, 0xd4, 0x45, 0xff, 0x5f, 0x03, 0x3a, 0xb3, 0xb9, 0x3f, 0x0d, 0xe2, 0x79, 0x10, 0xf9,
	0x13, 0x3f, 0xf2, 0x09, 0x81, 0xd3, 0x9c, 0x66, 0xe8, 0x1a, 0x9e, 0x31, 0xb0, 0x43, 0x9d, 0x4b,
	0xa6, 0x8a, 0x2d, 0xba, 0x8d, 0x03, 0x2b, 0x33, 0x71, 0xe1, 0x2c, 0x91, 0x48, 0x15, 0x32, 0xd7,
	0xd4, 0xb8, 0x2e, 0xc9, 0x3b, 0x68, 0x29, 0x71, 0x87, 0xf9, 0xce, 0x3d, 0xf5, 0x8c, 0x41, 0x67,
	0xf4, 0x72, 0x78, 0x38, 0x39, 0x5a, 0x7e, 0x0d, 0x16, 0xf1, 0xd8, 0x8f, 0x82, 0xe9, 0x32, 0xfc,
	0x11, 0x56, 0x4d, 0xe4, 0x35, 0xd8, 0xf3, 0xe2, 0x9b, 0x14, 0xbf, 0xf8, 0x06, 0xdd, 0xa6, 0x67,
	0x0c, 0xac, 0xd0, 0xce, 0x6a, 0x40, 0xde, 0x83, 0x99, 0x30, 0xe9, 0xb6, 0x3c, 0x73, 0xd0, 0x1e,
	0xbd, 0xa9, 0x26, 0x1d, 0xaf, 0x3c, 0x1c, 0x33, 0x19, 0xe4, 0x4a, 0x16, 0x61, 0xd9, 0x4a, 0x7a,
	0x60, 0x29, 0x9e, 0xe1, 0x86, 0xe7, 0xe8, 0x9e, 0xe9, 0xcd, 0xee, 0x6b, 0x7d, 0x11, 0x9a, 0xee,
	0x5c, 0xab, 0xba, 0x08, 0x4d, 0xf5, 0xf9, 0x2b, 0x9e, 0xe6, 0x54, 0xed, 0x25, 0xba, 0xb6, 0x67,
	0x0e, 0xec, 0xd0, 0xde, 0xd5, 0xa0, 0x77, 0x05, 0x56, 0x3d, 0x9e, 0x38, 0x60, 0xde, 0x61, 0x51,
	0x99, 0x29, 0x23, 0x39, 0x87, 0xe6, 0x6f, 0xba, 0xd9, 0xd7, 0x66, 0x0e, 0xc5, 0xa7, 0xc6, 0x47,
	0xa3, 0xff, 0xd7, 0x80, 0xf6, 0xbc, 0x5c, 0x76, 0xa5, 0x24, 0xd2, 0xec, 0x49, 0xad, 0x6f, 0xe1,
	0x39, 0x67, 0x98, 0x2b, 0xae, 0x8a, 0x38, 0x11, 0xac, 0x9e, 0xf2, 0xac, 0x86, 0x63, 0xc1, 0x90,
	0x78, 0xd0, 0x66, 0x54, 0xe1, 0xb1, 0xeb, 0x87, 0xa8, 0x1c, 0xbd, 0xa5, 0xea, 0x56, 0xdb, 0xb6,
	0x43, 0x9d, 0xc9, 0x15, 0x58, 0x09, 0x55, 0x98, 0x0a, 0x59, 0x68, 0xa7, 0x9d, 0x51, 0xaf, 0x72,
	0x37, 0xc1, 0x5d, 0x22, 0xf9, 0x56, 0x09, 0x19, 0x8f, 0xab, 0x8e, 0xf0, 0xbe, 0xf7, 0x58, 0x46,
	0xeb, 0x91, 0x8c, 0x8b, 0x08, 0x3a, 0xc7, 0x1f, 0x91, 0xbc, 0x82, 0xee, 0x62, 0xb9, 0x88, 0xbf,
	0xac, 0x17, 0xd3, 0x99, 0x7f, 0x7d, 0x13, 0xc4, 0xfa, 0xbd, 0x73, 0x42, 0xba, 0xf0, 0xe2, 0x31,
	0x34, 0xc8, 0x39, 0x38, 0xab, 0xe5, 0xfa, 0x26, 0xbe, 0x5e, 0xae, 0x17, 0x93, 0x8a, 0x36, 0x2e,
	0x3e, 0x43, 0xf7, 0x89, 0xa5, 0x88, 0x05, 0xa7, 0x11, 0xfe, 0x51, 0xce, 0x09, 0xb1, 0xa1, 0xf9,
	0x9d, 0x33, 0x14, 0x8e, 0x51, 0x46, 0x7f, 0xcf, 0xb8, 0x70, 0x1a, 0x65, 0x9c, 0x65, 0x34, 0x45,
	0xc7, 0xfc, 0xd9, 0xd2, 0x3f, 0xf5, 0x87, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xc5, 0x19,
	0x27, 0xfa, 0x02, 0x00, 0x00,
}
