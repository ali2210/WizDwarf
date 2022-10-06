package jsonpb

import (
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	user "github.com/ali2210/wizdwarf/other/users/register"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtojsonMarshaler(v *binary.Micromolecule) ([]byte, error) {
	return protojson.Marshal(v)
}

func ProtojsonUnmarshaler(b []byte, m *binary.Micromolecule) error {
	return protojson.Unmarshal(b, m)
}

type ProtocEncode interface {
	NewReflect(any *user.New_User) protoreflect.Message
	UpdateReflect(any *user.Updated_User) protoreflect.Message
}

type ProtocEncoder struct{}

func New_ProtocEncoder() ProtocEncode { return &ProtocEncoder{} }

func (enc *ProtocEncoder) NewReflect(any *user.New_User) protoreflect.Message {

	return proto.MessageReflect(any)
}

func (enc *ProtocEncoder) UpdateReflect(any *user.Updated_User) protoreflect.Message {
	return proto.MessageReflect(any)
}
