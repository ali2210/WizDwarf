// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: encryption_access.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EncryptionAccess struct {
	DefaultKey                  []byte                         `protobuf:"bytes,1,opt,name=default_key,json=defaultKey,proto3" json:"default_key,omitempty"`
	StoreEntries                []*EncryptionAccess_StoreEntry `protobuf:"bytes,2,rep,name=store_entries,json=storeEntries,proto3" json:"store_entries,omitempty"`
	DefaultPathCipher           CipherSuite                    `protobuf:"varint,3,opt,name=default_path_cipher,json=defaultPathCipher,proto3,enum=encryption.CipherSuite" json:"default_path_cipher,omitempty"`
	DefaultEncryptionParameters *EncryptionParameters          `protobuf:"bytes,4,opt,name=default_encryption_parameters,json=defaultEncryptionParameters,proto3" json:"default_encryption_parameters,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                       `json:"-"`
	XXX_unrecognized            []byte                         `json:"-"`
	XXX_sizecache               int32                          `json:"-"`
}

func (m *EncryptionAccess) Reset()         { *m = EncryptionAccess{} }
func (m *EncryptionAccess) String() string { return proto.CompactTextString(m) }
func (*EncryptionAccess) ProtoMessage()    {}
func (*EncryptionAccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_464b1a18bff4a17b, []int{0}
}
func (m *EncryptionAccess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionAccess.Unmarshal(m, b)
}
func (m *EncryptionAccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionAccess.Marshal(b, m, deterministic)
}
func (m *EncryptionAccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionAccess.Merge(m, src)
}
func (m *EncryptionAccess) XXX_Size() int {
	return xxx_messageInfo_EncryptionAccess.Size(m)
}
func (m *EncryptionAccess) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionAccess.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionAccess proto.InternalMessageInfo

func (m *EncryptionAccess) GetDefaultKey() []byte {
	if m != nil {
		return m.DefaultKey
	}
	return nil
}

func (m *EncryptionAccess) GetStoreEntries() []*EncryptionAccess_StoreEntry {
	if m != nil {
		return m.StoreEntries
	}
	return nil
}

func (m *EncryptionAccess) GetDefaultPathCipher() CipherSuite {
	if m != nil {
		return m.DefaultPathCipher
	}
	return CipherSuite_ENC_UNSPECIFIED
}

func (m *EncryptionAccess) GetDefaultEncryptionParameters() *EncryptionParameters {
	if m != nil {
		return m.DefaultEncryptionParameters
	}
	return nil
}

type EncryptionAccess_StoreEntry struct {
	Bucket               []byte                `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	UnencryptedPath      []byte                `protobuf:"bytes,2,opt,name=unencrypted_path,json=unencryptedPath,proto3" json:"unencrypted_path,omitempty"`
	EncryptedPath        []byte                `protobuf:"bytes,3,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	Key                  []byte                `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	PathCipher           CipherSuite           `protobuf:"varint,5,opt,name=path_cipher,json=pathCipher,proto3,enum=encryption.CipherSuite" json:"path_cipher,omitempty"`
	EncryptionParameters *EncryptionParameters `protobuf:"bytes,6,opt,name=encryption_parameters,json=encryptionParameters,proto3" json:"encryption_parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EncryptionAccess_StoreEntry) Reset()         { *m = EncryptionAccess_StoreEntry{} }
func (m *EncryptionAccess_StoreEntry) String() string { return proto.CompactTextString(m) }
func (*EncryptionAccess_StoreEntry) ProtoMessage()    {}
func (*EncryptionAccess_StoreEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_464b1a18bff4a17b, []int{0, 0}
}
func (m *EncryptionAccess_StoreEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptionAccess_StoreEntry.Unmarshal(m, b)
}
func (m *EncryptionAccess_StoreEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptionAccess_StoreEntry.Marshal(b, m, deterministic)
}
func (m *EncryptionAccess_StoreEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptionAccess_StoreEntry.Merge(m, src)
}
func (m *EncryptionAccess_StoreEntry) XXX_Size() int {
	return xxx_messageInfo_EncryptionAccess_StoreEntry.Size(m)
}
func (m *EncryptionAccess_StoreEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptionAccess_StoreEntry.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptionAccess_StoreEntry proto.InternalMessageInfo

func (m *EncryptionAccess_StoreEntry) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *EncryptionAccess_StoreEntry) GetUnencryptedPath() []byte {
	if m != nil {
		return m.UnencryptedPath
	}
	return nil
}

func (m *EncryptionAccess_StoreEntry) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

func (m *EncryptionAccess_StoreEntry) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EncryptionAccess_StoreEntry) GetPathCipher() CipherSuite {
	if m != nil {
		return m.PathCipher
	}
	return CipherSuite_ENC_UNSPECIFIED
}

func (m *EncryptionAccess_StoreEntry) GetEncryptionParameters() *EncryptionParameters {
	if m != nil {
		return m.EncryptionParameters
	}
	return nil
}

func init() {
	proto.RegisterType((*EncryptionAccess)(nil), "encryption_access.EncryptionAccess")
	proto.RegisterType((*EncryptionAccess_StoreEntry)(nil), "encryption_access.EncryptionAccess.StoreEntry")
}

func init() { proto.RegisterFile("encryption_access.proto", fileDescriptor_464b1a18bff4a17b) }

var fileDescriptor_464b1a18bff4a17b = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0x4c, 0x29, 0x1f, 0x87, 0x57, 0xe0, 0x2b, 0x2b, 0x4a, 0x83, 0x31, 0x36, 0x26, 0x26, 0xf5,
	0x52, 0x12, 0xbc, 0x78, 0x55, 0x43, 0x3c, 0x78, 0x21, 0x25, 0x5e, 0xbc, 0x34, 0xa5, 0x3c, 0x43,
	0x45, 0x76, 0x37, 0xbb, 0xdb, 0x43, 0x7f, 0x8a, 0xbf, 0xcb, 0x3f, 0x64, 0xba, 0x2c, 0xb4, 0x02,
	0x31, 0xde, 0xb6, 0x33, 0xb3, 0xb3, 0xf3, 0xe6, 0x15, 0x06, 0x48, 0x53, 0x51, 0x70, 0x95, 0x31,
	0x1a, 0x27, 0x69, 0x8a, 0x52, 0x86, 0x5c, 0x30, 0xc5, 0x48, 0xef, 0x80, 0x18, 0xba, 0x15, 0xb4,
	0x11, 0x5d, 0x7d, 0x35, 0xc1, 0x9d, 0xec, 0xc0, 0x7b, 0x2d, 0x23, 0x97, 0xe0, 0x2c, 0xf0, 0x2d,
	0xc9, 0x3f, 0x54, 0xbc, 0xc2, 0xc2, 0xb3, 0x7c, 0x2b, 0x68, 0x47, 0x60, 0xa0, 0x67, 0x2c, 0xc8,
	0x0c, 0x3a, 0x52, 0x31, 0x81, 0x31, 0x52, 0x25, 0x32, 0x94, 0x5e, 0xc3, 0xb7, 0x03, 0x67, 0x1c,
	0x86, 0x87, 0x59, 0xf6, 0xcd, 0xc3, 0x59, 0x79, 0x71, 0x42, 0x95, 0x28, 0xa2, 0xb6, 0xdc, 0x9e,
	0x33, 0x94, 0xe4, 0x09, 0x4e, 0xb6, 0xaf, 0xf2, 0x44, 0x2d, 0xe3, 0x34, 0xe3, 0x4b, 0x14, 0x9e,
	0xed, 0x5b, 0x41, 0x77, 0x3c, 0xa8, 0x59, 0x87, 0x8f, 0x9a, 0x99, 0xe5, 0x99, 0xc2, 0xa8, 0x67,
	0xee, 0x4c, 0x13, 0xb5, 0xdc, 0xe0, 0x64, 0x01, 0x17, 0x5b, 0xa3, 0x5a, 0x1e, 0x9e, 0x88, 0x64,
	0x8d, 0x0a, 0x85, 0xf4, 0x9a, 0xbe, 0x15, 0x38, 0x63, 0xbf, 0x6e, 0x59, 0xc5, 0x9c, 0xee, 0x74,
	0xd1, 0xb9, 0xb1, 0x39, 0x46, 0x0e, 0x3f, 0x1b, 0x00, 0xd5, 0x2c, 0xe4, 0x0c, 0x5a, 0xf3, 0x3c,
	0x5d, 0xa1, 0x32, 0x75, 0x99, 0x2f, 0x72, 0x03, 0x6e, 0x4e, 0xcd, 0x43, 0xb8, 0xd0, 0x93, 0x79,
	0x0d, 0xad, 0xf8, 0x5f, 0xc3, 0xcb, 0xf4, 0xe4, 0x1a, 0xba, 0x7b, 0x42, 0x5b, 0x0b, 0x3b, 0x3f,
	0x65, 0x2e, 0xd8, 0xe5, 0x56, 0x9a, 0x9a, 0x2b, 0x8f, 0xe4, 0x0e, 0x9c, 0x7a, 0x63, 0xff, 0x7e,
	0x6f, 0x0c, 0x78, 0x55, 0xd5, 0x0b, 0x9c, 0x1e, 0xaf, 0xa8, 0xf5, 0xc7, 0x8a, 0xfa, 0x78, 0x04,
	0x7d, 0xe8, 0xbf, 0x92, 0x72, 0xb5, 0xef, 0x61, 0xc6, 0x46, 0x29, 0x5b, 0xaf, 0x19, 0x1d, 0xf1,
	0xf9, 0xbc, 0xa5, 0x7f, 0xb9, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xa8, 0x25, 0x28,
	0xb2, 0x02, 0x00, 0x00,
}
