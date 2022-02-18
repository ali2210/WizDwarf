package proteins

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	fd "github.com/ali2210/wizdwarf/other/jsonpb/jsonledit"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
)


func ProteinsFileGenerateTest(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "Proteins File Generate Test suite")
}

var _ = Describe("Proteins File Generated Automated", func() {
	Context(" Proteins Information store in JSON format", func() {
		It("should file have valid JSON", func() {
			
			molecule := make([]*binary.Micromolecule, 1)
			molecule = append(molecule,&binary.Micromolecule{
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
				})
			
			CreateNewJSONFile(&fd.FileDescriptor{
				Names: "proteins",
				Types: ".json", 
				Molecule : molecule,
				Occurance : 0,		
			})
		})
	})
})