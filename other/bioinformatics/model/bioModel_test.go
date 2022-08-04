package model

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var result = Levenshtein{
	Probablity: 1250,
	Percentage: 0.0,
	Name:       "",
}

func ResultsTestCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Results Objects State")
}

var _ = Describe("Levenshtein Object State", func() {

	Context("Genetics tracebility", func() {
		It("Corelation between virus genome and human genome in probability", func() {
			Expect(result.Result(int(result.Probablity), 700)).Should(BeZero())
		})
	})

	Context("Identitical genome", func() {
		It("Corelation between virus genome and human genome through percentage ", func() {
			Expect(result.CalcualtePercentage(result.Probablity / 700)).Should(BeZero())
		})
	})

})
