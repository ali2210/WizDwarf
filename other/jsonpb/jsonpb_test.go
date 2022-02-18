package jsonpb

import(
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
)


func ProtoBufSchemeTest(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "Protocol Encode Schema test suite")
}

var _ = Describe(" Protocol Buffer Schema Testing", func() {
	Context("Marshalling Json object into Protocol Buffer", func() {
		It("should marshal valid", func() {
			Expect(ProtojsonMarshaler(&binary.Micromolecule{
				Symbol : "",
				Mass : 0.00,
				Composition : &binary.Element{
					C : 0,
					H : 0,
					S : 0,
					O : 0,
					N : 0,
				},
				Molecule : &binary.Traits{
					A : "0.00",
					B : "0.00",
					Magnetic_Field: "0.00",
				},
				Abundance : 0,
			})).ShouldNot(BeEmpty())
		})
	})
	Context (" Unmarshal Json object into Protocol Buffer", func() {
		It("should unmarshal", func() {
			Expect(ProtojsonUnmarshaler([]byte{}, &binary.Micromolecule{
				Symbol : "",
				Mass : 0.00,
				Composition : &binary.Element{
					C : 0,
					H : 0,
					S : 0,
					O : 0,
					N : 0,
				},
				Molecule : &binary.Traits{
					A : "0.00",
					B : "0.00",
					Magnetic_Field: "0.00",
				},
				Abundance : 0,
			})).Should(BeNil())		
		})
	})
})