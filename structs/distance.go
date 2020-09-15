package structs



// import libraries
import(
  "strings"
  "github.com/agnivade/levenshtein"

)


type Levenshtein struct{
  Probablity float32
  Percentage float32
  Name string

}


type LevenTable interface{
  // Calcualte Distance
   EditDistanceStrings(s1 , s2 []string)(int)
   Result(d int)(float32)
   CalcualtePercentage(p float32)float32
}



func EditDistanceStrings(s1, s2 []string)(int)  {
  rowString := strings.Join(s1," ")
  colStrings := strings.Join(s2, " ")

  return levenshtein.ComputeDistance(rowString, colStrings)
}

func (*Levenshtein) Result(d int) (float32)  {
   return float32(d/1000)
}

func (*Levenshtein) CalcualtePercentage(p float32) float32  {
  return ((100-p)/100)
}
