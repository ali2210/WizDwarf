// This codebase desgin according to mozilla open source license.
//Redistribution , contribution and improve codebase under license
//convensions. @contact Ali Hassan AliMatrixCode@protonmail.com

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: other/genetic/binary/codebank.proto

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

type Errors int32

const (
	Errors_OK    Errors = 0
	Errors_Error Errors = 1
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "OK",
		1: "Error",
	}
	Errors_value = map[string]int32{
		"OK":    0,
		"Error": 1,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_other_genetic_binary_codebank_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_other_genetic_binary_codebank_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_other_genetic_binary_codebank_proto_rawDescGZIP(), []int{0}
}

type Lifecode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genes string `protobuf:"bytes,1,opt,name=Genes,proto3" json:"Genes,omitempty"`
	Pkk   string `protobuf:"bytes,2,opt,name=Pkk,proto3" json:"Pkk,omitempty"`
}

func (x *Lifecode) Reset() {
	*x = Lifecode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_genetic_binary_codebank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lifecode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lifecode) ProtoMessage() {}

func (x *Lifecode) ProtoReflect() protoreflect.Message {
	mi := &file_other_genetic_binary_codebank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lifecode.ProtoReflect.Descriptor instead.
func (*Lifecode) Descriptor() ([]byte, []int) {
	return file_other_genetic_binary_codebank_proto_rawDescGZIP(), []int{0}
}

func (x *Lifecode) GetGenes() string {
	if x != nil {
		return x.Genes
	}
	return ""
}

func (x *Lifecode) GetPkk() string {
	if x != nil {
		return x.Pkk
	}
	return ""
}

type Lifecode_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    bool   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error     string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	ErrorCode Errors `protobuf:"varint,3,opt,name=ErrorCode,proto3,enum=binary.Errors" json:"ErrorCode,omitempty"`
}

func (x *Lifecode_Status) Reset() {
	*x = Lifecode_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_genetic_binary_codebank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lifecode_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lifecode_Status) ProtoMessage() {}

func (x *Lifecode_Status) ProtoReflect() protoreflect.Message {
	mi := &file_other_genetic_binary_codebank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lifecode_Status.ProtoReflect.Descriptor instead.
func (*Lifecode_Status) Descriptor() ([]byte, []int) {
	return file_other_genetic_binary_codebank_proto_rawDescGZIP(), []int{1}
}

func (x *Lifecode_Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Lifecode_Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Lifecode_Status) GetErrorCode() Errors {
	if x != nil {
		return x.ErrorCode
	}
	return Errors_OK
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_genetic_binary_codebank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_other_genetic_binary_codebank_proto_msgTypes[2]
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
	return file_other_genetic_binary_codebank_proto_rawDescGZIP(), []int{2}
}

var File_other_genetic_binary_codebank_proto protoreflect.FileDescriptor

var file_other_genetic_binary_codebank_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x2f,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x32, 0x0a,
	0x08, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x47, 0x65, 0x6e,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x50, 0x6b, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x6b,
	0x6b, 0x22, 0x6d, 0x0a, 0x0f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x1b, 0x0a, 0x06, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x32, 0x6b, 0x0a, 0x05, 0x47, 0x42, 0x61, 0x6e,
	0x6b, 0x12, 0x34, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x2e, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x17,
	0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0f, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x66,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69, 0x32, 0x32, 0x31, 0x30, 0x2f, 0x77, 0x69, 0x7a, 0x64,
	0x77, 0x61, 0x72, 0x66, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x74,
	0x69, 0x63, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_other_genetic_binary_codebank_proto_rawDescOnce sync.Once
	file_other_genetic_binary_codebank_proto_rawDescData = file_other_genetic_binary_codebank_proto_rawDesc
)

func file_other_genetic_binary_codebank_proto_rawDescGZIP() []byte {
	file_other_genetic_binary_codebank_proto_rawDescOnce.Do(func() {
		file_other_genetic_binary_codebank_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_genetic_binary_codebank_proto_rawDescData)
	})
	return file_other_genetic_binary_codebank_proto_rawDescData
}

var file_other_genetic_binary_codebank_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_other_genetic_binary_codebank_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_other_genetic_binary_codebank_proto_goTypes = []interface{}{
	(Errors)(0),             // 0: binary.Errors
	(*Lifecode)(nil),        // 1: binary.Lifecode
	(*Lifecode_Status)(nil), // 2: binary.Lifecode_Status
	(*Request)(nil),         // 3: binary.Request
}
var file_other_genetic_binary_codebank_proto_depIdxs = []int32{
	0, // 0: binary.Lifecode_Status.ErrorCode:type_name -> binary.Errors
	1, // 1: binary.GBank.AddCode:input_type -> binary.Lifecode
	3, // 2: binary.GBank.GetCode:input_type -> binary.Request
	2, // 3: binary.GBank.AddCode:output_type -> binary.Lifecode_Status
	1, // 4: binary.GBank.GetCode:output_type -> binary.Lifecode
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_other_genetic_binary_codebank_proto_init() }
func file_other_genetic_binary_codebank_proto_init() {
	if File_other_genetic_binary_codebank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_genetic_binary_codebank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lifecode); i {
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
		file_other_genetic_binary_codebank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lifecode_Status); i {
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
		file_other_genetic_binary_codebank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_other_genetic_binary_codebank_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_other_genetic_binary_codebank_proto_goTypes,
		DependencyIndexes: file_other_genetic_binary_codebank_proto_depIdxs,
		EnumInfos:         file_other_genetic_binary_codebank_proto_enumTypes,
		MessageInfos:      file_other_genetic_binary_codebank_proto_msgTypes,
	}.Build()
	File_other_genetic_binary_codebank_proto = out.File
	file_other_genetic_binary_codebank_proto_rawDesc = nil
	file_other_genetic_binary_codebank_proto_goTypes = nil
	file_other_genetic_binary_codebank_proto_depIdxs = nil
}
