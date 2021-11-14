package model

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var result = Levenshtein{
	Probablity: 0.0,
	Percentage: 0.0,
	Name:       "",
}

func ResultsTestCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Results Objects State")
}

var _ = Describe("Levenshtein Object State", func() {

	Context("Object Probability ", func() {
		It("Probability ", func() {
			Expect(result.Result(int(result.Probablity))).Should(BeZero())
		})
	})

	Context("Object Percentage", func() {
		It("Percentage ", func() {
			Expect(result.CalcualtePercentage(result.Percentage)).Should(BeZero())
		})
	})

})
