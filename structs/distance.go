package structs



// import libraries
import(
  "strings"
  "github.com/agnivade/levenshtein"
)


type Levenshtein struct{
  Probablity float64
}


type LevenTable interface{
  // Calcualte Distance
   EditDistanceStrings(s1 , s2 []string)(int)
    Result(d, l1 int)(float64)
}



func EditDistanceStrings(s1, s2 []string)(int)  {
  rowString := strings.Join(s1," ")
  colStrings := strings.Join(s2, " ")

  return levenshtein.ComputeDistance(rowString, colStrings)
}

func (*Levenshtein) Result(d, l1 int) (float64)  {
   return float64((d/l1)*100)
}
