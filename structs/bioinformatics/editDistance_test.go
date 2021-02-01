package bioinformatics

import (
	// "errors"
	"fmt"
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
			Expect(editTest.Result(algoTest.EditDistanceStrings(a, b))).Should(BeZero())
		})
		It("Calculate Percentage", func() {
			Expect(editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))).Should(BeZero())
		})
		It("Get Parameters", func() {
			parm := editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))
			editTest.SetProbParameter(parm)
			Expect(editTest.GetProbParameter())
		})
	})
})

type StringLevenstein struct {
	inp0 []string
	inp1 []string
	out  int
}

var testLeven = []StringLevenstein{
	{[]string{""}, []string{""}, -1},
	{[]string{""}, []string{""}, 0},
	{[]string{"a"}, []string{""}, -1},
	{[]string{""}, []string{"a"}, -1},
	{[]string{"a"}, []string{"a"}, 0},
	{[]string{"a"}, []string{"b"}, -1},
	{[]string{"aa"}, []string{"ab"}, -1},
	{[]string{"aa"}, []string{"aa"}, 0},
	{[]string{"aaa"}, []string{"aaa"}, 0},
	{[]string{"abc"}, []string{"aabb"}, -1},
	{[]string{"aaaa"}, []string{"aaaa"}, 0},
}

func StringsTableTest(t *testing.T) {
	for _, tt := range testLeven {
		testfunc := fmt.Sprintf("%s%s", tt.inp0, tt.inp1)
		t.Run(testfunc, func(t *testing.T) {
			var _ = Describe("Levenshtein Test", func() {
				a := []string{"a"}
				b := []string{"aa"}
				Context("Objects Handler", func() {
					It("Object !", func() {
						Expect(algoTest.EditDistanceStrings(a, b)).Should(BeZero())
					})
					It("Calculate Similarity", func() {
						Expect(editTest.Result(algoTest.EditDistanceStrings(a, b))).Should(BeZero())
					})
					It("Calculate Percentage", func() {
						Expect(editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))).Should(BeZero())
					})
					It("Get Parameters", func() {
						parm := editTest.CalcualtePercentage(editTest.Result(algoTest.EditDistanceStrings(a, b)))
						editTest.SetProbParameter(parm)
						Expect(editTest.GetProbParameter())
					})
				})
			})
		})
	}
}

type ResultTestcases struct {
	in  int
	out float32
}

var testRes = []ResultTestcases{
	{0, 0.0},
	{0, 1.0},
}

func EditDistanceResultTest(a int, t *testing.T) float32 {
	output := editTest.Result(a)
	return output
}

func ResultTableDriven(t *testing.T) {
	for _, tt := range testRes {
		actual := EditDistanceResultTest(tt.in, t)
		if actual != tt.out {
			t.Errorf("expected output :%f, input: %d", actual, tt.in)
		}
	}
}

func CalcualtePercentageTest(a float32, t *testing.T) float32 {
	expect := editTest.CalcualtePercentage(a)
	if expect != 0.0 {
		t.Errorf("test failed output:%f, input: %f", a, expect)
	}
	return expect
}

type PercentageTest struct {
	in  float32
	out float32
}

var test = []PercentageTest{
	{0.0, 1.0},
	{100.0, 0.0},
}

func CalcualtePercentageTestDriven(t *testing.T) {
	for _, tt := range test {
		actual := CalcualtePercentageTest(tt.in, t)
		if actual != tt.out {
			t.Errorf("expected output%f, input: %f", actual, tt.in)
		}
	}
}
