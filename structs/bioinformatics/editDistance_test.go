package bioinformatics

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

var (
	editTest Levenshtein = Levenshtein{}
)

func StringLengthTest(a, b []string, t *testing.T) (int, error) {
	fmt.Println("Test {Status}")
	err := errors.New("[x] Leveinstein algorithms ")

	s1 := a
	sArrayCol := strings.Join(s1, " ")
	arraySplit := strings.Split(sArrayCol, "")
	s2 := b
	sArrayRow := strings.Join(s2, " ")
	arraySplit2 := strings.Split(sArrayRow, "")
	inputTest := EditDistanceStrings(arraySplit, arraySplit2)
	if inputTest < 0 {
		t.Errorf("Test {Failed}%v -- [Leveinstein match failed] ", err)
		return inputTest, err
	}
	return inputTest, nil
}

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
			expected, err := StringLengthTest(tt.inp0, tt.inp1, t)
			if err != nil {
				return
			}
			if expected != tt.out || expected < 0 {
				t.Errorf("value1: %s, value_2%s, expected%d  ", tt.inp0, tt.inp1, expected)
			}
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
