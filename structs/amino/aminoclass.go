package amino

import(
	"fmt"
)

type AminoClass struct{

	Symbol string
	Name string
	Polar bool
	Charge bool
	ChargeType string 
	TypeChain string
	Hydrophobic bool
	PK float64
	codon_Start bool
	codon_End bool
}


func (bioclass *AminoClass) Bases(class []string)[]*AminoClass{


	size := (len(class)- 4) 
	peptideBond := make([]*AminoClass, size)

	for i := 0; i < (len(class) - 4); i++{
		// fmt.Println("input_i:", class[i+0])
		// fmt.Println("input_j:", class[i+1])
		// fmt.Println("input_k:", class[i+2])
		  if (class[i+0] == "U"  && class[i+1] == "U" &&  class[i+2] == "U")|| (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "C"){
			bioclass.Symbol = "F"
			bioclass.Name = "Phenylalanine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:", bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "A")|| (class[i+0] == "U" && class[i+1] =="U" && class[i+2] =="G")|| (class[i+0] == "C" && class[i+1] =="U" && class[i+2] == "U") || (class[i+0] ==  "C" && class[i+1] =="U" && class[i+2] == "C")|| (class[i+0] ==  "C" &&  class[i+1] =="U"  && class[i+2] =="A") || (class[i+0] ==  "C" && class[i+1] =="U" && class[i+2] =="G") {
			bioclass.Symbol = "L"
			bioclass.Name = "Leucine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] ==  "U" && class[i+1] =="C" && class[i+2] =="U")|| (class[i+0] ==  "U" &&  class[i+1] =="C" && class[i+2] =="C")|| (class[i+0] ==  "U" && class[i+1] == "C" && class[i+2] == "A")|| (class[i+0] ==  "U" &&  class[i+1] == "C" && class[i+2] ==  "G")|| (class[i+0] ==  "A" && class[i+1] == "G" && class[i+2] == "U")|| (class[i+0] == "A" &&  class[i+1] == "G" && class[i+2] == "C") {
			bioclass.Symbol = "S"
			bioclass.Name = "Serine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.Hydrophobic = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==  "C" && class[i+1] == "C" && class[i+2] == "U")|| (class[i+0] ==  "C" && class[i+1] == "C" && class[i+2] == "C")|| (class[i+0] ==  "C" && class[i+1] == "C" && class[i+2] == "A")|| (class[i+0] == "C" && class[i+1] == "C" && class[i+2] == "G" ){
			bioclass.Symbol = "P"
			bioclass.Name = "Proline"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==  "A" && class[i+1] == "A" && class[i+2] == "U")|| (class[i+0] ==  "A" && class[i+1] == "U" && class[i+2] == "C")|| (class[i+0] == "A" && class[i+1] == "U" && class[i+2] == "A" ) {
			bioclass.Symbol = "I"
			bioclass.Name = "Isoleucine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if class[i+0] ==  "A" && class[i+1] == "U" && class[i+2] == "G"  {
			bioclass.Symbol = "M"
			bioclass.Name = "Methionine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.codon_Start = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==  "A" && class[i+1] == "C" && class[i+2] == "U")|| (class[i+0] ==  "A" && class[i+1] == "C" && class[i+2] == "C")|| (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "A" )|| (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "G") {
			bioclass.Symbol = "T"
			bioclass.Name = "Threonine"
			bioclass.Polar = true
			bioclass.Charge = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==  "G" && class[i+1] == "U" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "C" )|| (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "A" )|| (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "G" ){
			bioclass.Symbol = "V"
			bioclass.Name = "Valine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "G" && class[i+1] == "C" && class[i+2] == "U")|| (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "C")|| (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "A" )|| (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "G") {
			bioclass.Symbol = "A"
			bioclass.Name = "Alanine"
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "U" && class[i+1] == "A" && class[i+2] == "U" )||  (class[i+0] == "U" && class[i+1] == "A" && class[i+2] == "C")  {
			bioclass.Symbol = "Y"
			bioclass.Name = "Tyrosine"
			bioclass.Polar = true
			bioclass.TypeChain = "Aromatic"
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "U" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] ==  "U" && class[i+1] == "A" && class[i+2] == "G")|| (class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "A")  {
			bioclass.Symbol = "X"
			bioclass.codon_End = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			

		}else if (class[i+0] ==   "C" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "A" && class[i+2] == "C")  {
			bioclass.Symbol = "H"
			bioclass.Charge = true
			bioclass.ChargeType = "Positive"
			bioclass.Name = "Histidine"
			bioclass.Polar = true
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "C" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] ==  "C" && class[i+1] == "A" && class[i+2] == "G")  {
			bioclass.Symbol = "Q"
			bioclass.Name = "Glutamine"
			bioclass.Polar = true
			bioclass.Charge = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "C" && class[i+1] == "G" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "C" )|| (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "A")|| (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "G" )|| (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "A" )|| (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "G")  {
			bioclass.Symbol = "R"
			bioclass.Name = "Arginine"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.ChargeType = "Positive"
			bioclass.Hydrophobic= false
			bioclass.PK= 12.5
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i+0] ==   "A" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "C")  {
			bioclass.Symbol = "N"
			bioclass.Name = "Asparagine"
			bioclass.Polar = true
			bioclass.Charge = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] ==   "A" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "G")  {
			bioclass.Symbol = "K"
			bioclass.Name = "Lysine"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.ChargeType = "Positive"
			bioclass.PK = 10.5
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] ==   "G" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "A" && class[i+2] == "C")  {
			bioclass.Symbol = "D"
			bioclass.Name = "Aspartic Acid"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.PK = 3.9
			bioclass.ChargeType = "Negative"
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] ==   "G" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] ==  "G" && class[i+1] == "A" && class[i+2] == "G")  {
			bioclass.Symbol = "E"
			bioclass.Name = "Glutamate"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.PK = 4.2
			bioclass.ChargeType = "Negative"
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if (class[i] ==   "G" && class[i] == "G" && class[i] == "U" )|| (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "C" )|| (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "A" )|| (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "G")  {
			bioclass.Symbol = "G"
			bioclass.Name = "Glycine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
		}else if (class[i+0] ==   "U" && class[i+1] == "G" && class[i+2] == "U" )|| (class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "C")  {
			bioclass.Symbol = "C"
			bioclass.Name = "Cysteine"
			bioclass.Polar = true
			bioclass.Charge = false
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
			
			
		}else if class[i+0] ==   "U" && class[i+1] == "G" && class[i+2] == "G" {
			bioclass.Symbol = "W"
			bioclass.Name = "Tryptophan"
			bioclass.Polar = false
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = true
			fmt.Println("class:",bioclass)
			peptideBond = append(peptideBond, bioclass)
		}	
   	}
   	
	return peptideBond

}

