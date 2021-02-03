package bioinformatics

import (
	// "errors"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	bio "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
)

var (
	editTest bio.Levenshtein = bio.Levenshtein{
		Probablity: 0.0,
		Percentage: 0.0,
		Name:       "",
	}
	algoTest LevenTable = NewMatch()
)

func LevenshteinTestCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Levenshtein Test Cases")
}

var _ = Describe("Levenshtein Test", func() {
	a := []string{""}
	b := []string{""}
	Context("Empty Objects Handler", func() {
		It("Empty Object !", func() {
			Expect(algoTest.EditDistanceStrings(a, b)).Should(BeZero())
		})
		It("Calculate Similarity", func() {
			Expect(editTest.Result(algoTest.EditDistanceStrings(a, b))).ShouldNot(BeZero())
		})
		It("Calculate Percentage", func() {
			Expect(editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))).ShouldNot(BeZero())
		})
		It("Get Parameters", func() {
			parm := editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))
			editTest.SetProbParameter(parm)
			Expect(editTest.GetProbParameter()).ShouldNot(BeNil())
		})
	})
})
