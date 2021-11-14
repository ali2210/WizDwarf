// This codebase desgin according to mozilla open source license.
//Redistribution , contribution and improve codebase under license
//convensions. @contact Ali Hassan AliMatrixCode@protonmail.com

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: other/proteins/binary/binary.proto

package binary

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Element have all the resuide element which take part in marcomoleculaes
type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// carbon
	C int64 `protobuf:"varint,1,opt,name=C,proto3" json:"C,omitempty"`
	// sulpfur
	S int64 `protobuf:"varint,2,opt,name=S,proto3" json:"S,omitempty"`
	// hydrogen
	H int64 `protobuf:"varint,3,opt,name=H,proto3" json:"H,omitempty"`
	// nitrogen
	N int64 `protobuf:"varint,4,opt,name=N,proto3" json:"N,omitempty"`
	// oxygen
	O int64 `protobuf:"varint,5,opt,name=O,proto3" json:"O,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{0}
}

func (x *Element) GetC() int64 {
	if x != nil {
		return x.C
	}
	return 0
}

func (x *Element) GetS() int64 {
	if x != nil {
		return x.S
	}
	return 0
}

func (x *Element) GetH() int64 {
	if x != nil {
		return x.H
	}
	return 0
}

func (x *Element) GetN() int64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *Element) GetO() int64 {
	if x != nil {
		return x.O
	}
	return 0
}

type Traits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// acidic_a = "A"
	A string `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	// acidic_b = "B"
	B string `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	// magnetic field
	Magnetic_Field string `protobuf:"bytes,3,opt,name=Magnetic_Field,json=MagneticField,proto3" json:"Magnetic_Field,omitempty"`
}

func (x *Traits) Reset() {
	*x = Traits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traits) ProtoMessage() {}

func (x *Traits) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traits.ProtoReflect.Descriptor instead.
func (*Traits) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{1}
}

func (x *Traits) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Traits) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

func (x *Traits) GetMagnetic_Field() string {
	if x != nil {
		return x.Magnetic_Field
	}
	return ""
}

// micromolecule is a subset of the marcomolecule
type Micromolecule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol      string   `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Mass        float32  `protobuf:"fixed32,2,opt,name=Mass,proto3" json:"Mass,omitempty"`
	Composition *Element `protobuf:"bytes,3,opt,name=Composition,proto3" json:"Composition,omitempty"`
	Molecule    *Traits  `protobuf:"bytes,4,opt,name=Molecule,proto3" json:"Molecule,omitempty"`
}

func (x *Micromolecule) Reset() {
	*x = Micromolecule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Micromolecule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Micromolecule) ProtoMessage() {}

func (x *Micromolecule) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Micromolecule.ProtoReflect.Descriptor instead.
func (*Micromolecule) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{2}
}

func (x *Micromolecule) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Micromolecule) GetMass() float32 {
	if x != nil {
		return x.Mass
	}
	return 0
}

func (x *Micromolecule) GetComposition() *Element {
	if x != nil {
		return x.Composition
	}
	return nil
}

func (x *Micromolecule) GetMolecule() *Traits {
	if x != nil {
		return x.Molecule
	}
	return nil
}

// micromolecules exist in the form of long chain and have common features
type Micromolecule_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peplide []*Micromolecule `protobuf:"bytes,1,rep,name=Peplide,proto3" json:"Peplide,omitempty"`
}

func (x *Micromolecule_List) Reset() {
	*x = Micromolecule_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Micromolecule_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Micromolecule_List) ProtoMessage() {}

func (x *Micromolecule_List) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Micromolecule_List.ProtoReflect.Descriptor instead.
func (*Micromolecule_List) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{3}
}

func (x *Micromolecule_List) GetPeplide() []*Micromolecule {
	if x != nil {
		return x.Peplide
	}
	return nil
}

// during transciption and other process that take will happen or have happened either failed or not. In case process failed hold what the caused behind the failed process.
type MolecularState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State bool   `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *MolecularState) Reset() {
	*x = MolecularState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MolecularState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MolecularState) ProtoMessage() {}

func (x *MolecularState) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MolecularState.ProtoReflect.Descriptor instead.
func (*MolecularState) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{4}
}

func (x *MolecularState) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

func (x *MolecularState) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Empty message
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_proteins_binary_binary_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_other_proteins_binary_binary_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_other_proteins_binary_binary_proto_rawDescGZIP(), []int{5}
}

var File_other_proteins_binary_binary_proto protoreflect.FileDescriptor

var file_other_proteins_binary_binary_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73,
	0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x22, 0x4f,
	0x0a, 0x07, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x43, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x43, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x53, 0x12, 0x0c, 0x0a, 0x01, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x48, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x4e, 0x12, 0x0c, 0x0a, 0x01, 0x4f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x4f, 0x22,
	0x4b, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x42, 0x12, 0x25, 0x0a, 0x0e, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69,
	0x63, 0x5f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4d,
	0x61, 0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x9e, 0x01, 0x0a,
	0x0d, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x4d, 0x61, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x08, 0x4d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x52, 0x08, 0x4d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x22, 0x47, 0x0a,
	0x12, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x5f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x50, 0x65, 0x70, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2e,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x50,
	0x65, 0x70, 0x6c, 0x69, 0x64, 0x65, 0x22, 0x3c, 0x0a, 0x0e, 0x4d, 0x6f, 0x6c, 0x65, 0x63, 0x75,
	0x6c, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32,
	0x8d, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x42, 0x6e, 0x6b, 0x12, 0x40,
	0x0a, 0x06, 0x41, 0x64, 0x64, 0x50, 0x44, 0x42, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x69, 0x6e, 0x73, 0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c,
	0x65, 0x5f, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e,
	0x73, 0x2e, 0x4d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x44, 0x42, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2e, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x5f, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x69, 0x32, 0x32, 0x31, 0x30, 0x2f, 0x77, 0x69, 0x7a, 0x64, 0x77, 0x61, 0x72, 0x66, 0x2f, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73, 0x2f, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_other_proteins_binary_binary_proto_rawDescOnce sync.Once
	file_other_proteins_binary_binary_proto_rawDescData = file_other_proteins_binary_binary_proto_rawDesc
)

func file_other_proteins_binary_binary_proto_rawDescGZIP() []byte {
	file_other_proteins_binary_binary_proto_rawDescOnce.Do(func() {
		file_other_proteins_binary_binary_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_proteins_binary_binary_proto_rawDescData)
	})
	return file_other_proteins_binary_binary_proto_rawDescData
}

var file_other_proteins_binary_binary_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_other_proteins_binary_binary_proto_goTypes = []interface{}{
	(*Element)(nil),            // 0: proteins.Element
	(*Traits)(nil),             // 1: proteins.Traits
	(*Micromolecule)(nil),      // 2: proteins.Micromolecule
	(*Micromolecule_List)(nil), // 3: proteins.Micromolecule_List
	(*MolecularState)(nil),     // 4: proteins.MolecularState
	(*Request)(nil),            // 5: proteins.Request
}
var file_other_proteins_binary_binary_proto_depIdxs = []int32{
	0, // 0: proteins.Micromolecule.Composition:type_name -> proteins.Element
	1, // 1: proteins.Micromolecule.Molecule:type_name -> proteins.Traits
	2, // 2: proteins.Micromolecule_List.Peplide:type_name -> proteins.Micromolecule
	3, // 3: proteins.ProteinBnk.AddPDB:input_type -> proteins.Micromolecule_List
	5, // 4: proteins.ProteinBnk.DisplayPDB:input_type -> proteins.Request
	4, // 5: proteins.ProteinBnk.AddPDB:output_type -> proteins.MolecularState
	3, // 6: proteins.ProteinBnk.DisplayPDB:output_type -> proteins.Micromolecule_List
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_other_proteins_binary_binary_proto_init() }
func file_other_proteins_binary_binary_proto_init() {
	if File_other_proteins_binary_binary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_proteins_binary_binary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_other_proteins_binary_binary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Traits); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_other_proteins_binary_binary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Micromolecule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_other_proteins_binary_binary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Micromolecule_List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_other_proteins_binary_binary_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MolecularState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_other_proteins_binary_binary_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_other_proteins_binary_binary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_other_proteins_binary_binary_proto_goTypes,
		DependencyIndexes: file_other_proteins_binary_binary_proto_depIdxs,
		MessageInfos:      file_other_proteins_binary_binary_proto_msgTypes,
	}.Build()
	File_other_proteins_binary_binary_proto = out.File
	file_other_proteins_binary_binary_proto_rawDesc = nil
	file_other_proteins_binary_binary_proto_goTypes = nil
	file_other_proteins_binary_binary_proto_depIdxs = nil
}
