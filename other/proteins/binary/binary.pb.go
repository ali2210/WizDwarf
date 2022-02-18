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
	Mass        float64  `protobuf:"fixed64,2,opt,name=Mass,proto3" json:"Mass,omitempty"`
	Composition *Element `protobuf:"bytes,3,opt,name=Composition,proto3" json:"Composition,omitempty"`
	Molecule    *Traits  `protobuf:"bytes,4,opt,name=Molecule,proto3" json:"Molecule,omitempty"`
	Abundance   int64    `protobuf:"varint,5,opt,name=Abundance,proto3" json:"Abundance,omitempty"`
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

func (x *Micromolecule) GetMass() float64 {
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

func (x *Micromolecule) GetAbundance() int64 {
	if x != nil {
		return x.Abundance
	}
	return 0
}

var File_other_proteins_binary_binary_proto protoreflect.FileDescriptor

var file_other_proteins_binary_binary_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x73,
	0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x07,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x43, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x43, 0x12, 0x0c, 0x0a, 0x01, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x53, 0x12, 0x0c, 0x0a, 0x01, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x48, 0x12, 0x0c, 0x0a, 0x01, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x4e, 0x12,
	0x0c, 0x0a, 0x01, 0x4f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x4f, 0x22, 0x4b, 0x0a,
	0x06, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x41, 0x12, 0x0c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x42, 0x12, 0x25, 0x0a, 0x0e, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x5f,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4d, 0x61, 0x67,
	0x6e, 0x65, 0x74, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x0d, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x6d, 0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x04, 0x4d, 0x61, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x08, 0x4d,
	0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x73, 0x52, 0x08, 0x4d,
	0x6f, 0x6c, 0x65, 0x63, 0x75, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x62, 0x75, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x62, 0x75, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69, 0x32, 0x32, 0x31, 0x30, 0x2f, 0x77, 0x69, 0x7a, 0x64,
	0x77, 0x61, 0x72, 0x66, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x69, 0x6e, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_other_proteins_binary_binary_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_other_proteins_binary_binary_proto_goTypes = []interface{}{
	(*Element)(nil),       // 0: binary.Element
	(*Traits)(nil),        // 1: binary.Traits
	(*Micromolecule)(nil), // 2: binary.Micromolecule
}
var file_other_proteins_binary_binary_proto_depIdxs = []int32{
	0, // 0: binary.Micromolecule.Composition:type_name -> binary.Element
	1, // 1: binary.Micromolecule.Molecule:type_name -> binary.Traits
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_other_proteins_binary_binary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
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
