package simpleTest

import (
  "testing"
  "github.com/ali2210/wizdwarf/structs"
  "fmt"
  "strings"
)

var(
  editTest structs.Levenshtein = structs.Levenshtein{}
)


func StringLengthTest(a, b[]string, t*testing.T)int{
  fmt.Println("Test {Status}")
  // []string{"Quantum Computer idea is very vague, for now!"}
  s1 := a
  s_array_col := strings.Join(s1, " ")
  array_split := strings.Split(s_array_col, "")
  s2 := b
  s_array_row := strings.Join(s2, " ")
  array_split2 := strings.Split(s_array_row, "")
  input_test := structs.EditDistanceStrings(array_split,array_split2)
  // if input_test < 0 {
  //   t.Errorf("Test {Failed}%d -- [BECAUSE THERE IS NO DATA EXISTS] ", input_test)
  // }
  return input_test
}

type StringLevenstein struct{
  inp0 []string
  inp1 []string
  out int
}

var testLeven = []StringLevenstein{
  {[]string{""},[]string{""}, -1},
  {[]string{""},[]string{""}, 0},
  {[]string{"a"},[]string{""}, -1},
  {[]string{""},[]string{"a"}, -1},
  {[]string{"a"},[]string{"a"}, 0},
  {[]string{"a"}, []string{"b"}, -1},
  {[]string{"aa"}, []string{"ab"}, -1},
  {[]string{"aa"}, []string{"aa"}, 0},
  {[]string{"aaa"},[]string{ "aaa"}, 0},
  {[]string{"abc"},[]string{"aabb"}, -1},
  {[]string{"aaaa"}, []string{"aaaa"}, 0},

}

func StringsTableTest(t *testing.T)  {
  for _, tt := range testLeven{
    testfunc := fmt.Sprintf("%s%s",tt.inp0,tt.inp1)
    t.Run(testfunc, func(t *testing.T) {
      expected := StringLengthTest(tt.inp0,tt.inp1,t)
      if expected != tt.out{
        t.Errorf("value1: %s, value_2%s, expected%d  ",  tt.inp0, tt.inp1, expected)
      }
    })
  }
}

type ResultTestcases struct{
  in int
  out float32
}
var testRes = []ResultTestcases{
  {0,0.0},
  {0,1.0},
}

func EditDistanceResultTest(a int ,t *testing.T)float32{
  output := editTest.Result(a)
  return output
}

func ResultTableDriven(t *testing.T) {
  for _, tt := range testRes {
    actual := EditDistanceResultTest(tt.in,t)
    if actual != tt.out{
      t.Errorf("expected output :%f, input: %d", actual, tt.in)
    }
  }
}

func CalcualtePercentageTest(a float32, t*testing.T)float32{
    expect := editTest.CalcualtePercentage(a)
    if expect != 0.0{
      t.Errorf("test failed output:%f, input: %f", a, expect)
    }
    return expect
}
type PercentageTest struct{
  in float32
  out float32
}


var test = []PercentageTest{
  {0.0,1.0},
  {100.0, 0.0},
}
func CalcualtePercentageTestDriven(t *testing.T){
  for _, tt := range test {
    actual := CalcualtePercentageTest(tt.in,t)
    if actual != tt.out{
      t.Errorf("expected output%f, input: %f", actual, tt.in)
    }
  }
}
