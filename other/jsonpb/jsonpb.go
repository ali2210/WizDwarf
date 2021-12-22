package jsonpb

import (
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProtojsonMarshaler(v *binary.Micromolecule) ([]byte, error) {
	return protojson.Marshal(v)
}

func ProtojsonUnmarshaler(b []byte, m *binary.Micromolecule) error {
	return protojson.Unmarshal(b, m)
}
