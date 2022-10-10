/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package

package proteins

// libraries
import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Proteins_Suit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maromolecules tests suits starting .....")
}

var _ = Describe("Marcomolecules physical tests suits ..... ", func() {

	var amino Aminochain
	BeforeEach(func() {
		When("Amino initilaized ..... ", func() {
			Context("Empty amino object is not accepted", func() {
				It("should ammino object is valid ", func() {
					Expect(amino).Should(BeEmpty())
				})
			})
			Context("Codon pair order is valid", func() {
				It("should be ordered", func() {
					Expect(Class("ATT", 0, 2)).ShouldNot(BeEmpty())
				})
			})
			Context("Codon pair mass must not be zero", func() {
				It("pair should have mass", func() {
					mass, sym := GetMolarMass("ATT", 0, 2)
					Expect(mass).ShouldNot(BeZero())
					Expect(sym).ShouldNot(BeEmpty())
				})
			})
			Context("Codon pair must have acidiity value", func() {
				It("should have acidiity value", func() {
					_, a, b := GetPKa("ATT", 0, 2)
					Expect(a).ShouldNot(BeEmpty())
					Expect(b).ShouldNot(BeEmpty())
				})
			})
			Context("marcomolecules object tarits", func() {
				It("should have a marcomolecule physical properties", func() {
					Expect(GetAmino("ATT", 0, 2)).ShouldNot(BeEmpty())
				})
			})
		})
	})
})
