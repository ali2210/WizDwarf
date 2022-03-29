package molecules

import (
	"testing"

	"github.com/ali2210/wizdwarf/other/proteins/binary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func MoleculeTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Molecules unit tests")
}

var _ = Describe("Molecules test started ", func() {
	Context("Run molecule test", func() {
		molecules := &binary.Micromolecule{Symbol: "G", Mass: 75.067,
			Composition: &binary.Element{C: 2, H: 5, O: 2, N: 1, S: 0},
			Molecule:    &binary.Traits{A: "2.34", B: "9.6", Magnetic_Field: "-40.3·10−6"}}
		It("Should this is a valid molecule", func() {
			Expect(Molecular(molecules)).ShouldNot(BeTrue())
		})
		It("Should molecule have name  ", func() {
			Expect(Symbol(molecules)).ShouldNot(BeEmpty())
		})
		It("Should molecule have mass ", func() {
			Expect(Mass(molecules)).ShouldNot(BeZero())
		})
		It("Should molecule have unique family:", func() {
			Expect(Occurance(molecules)).Should(BeTrue())
		})
		It("Should molecule have any element:", func() {
			Expect(GetMoleculesState(molecules)).Should(BeTrue())
		})
		It("Should molecule have any additional properties:", func() {
			Expect(GetCompositionState(molecules)).Should(BeTrue())
		})
	})
})
