package amino

import (
	/*x*/
	"math/rand"
)

type AminoClass struct {
	Symbol      string
	Name        string
	Polar       bool
	Charge      bool
	ChargeType  string
	TypeChain   string
	Hydrophobic bool
	PKa         float64
	CodonStart  bool
	CodonEnd    bool
	ID          int
}

func (bioclass AminoClass) Bases(class []string) []AminoClass {

	size := (len(class) - 4)
	peptideBond := make([]AminoClass, size)

	for range class {
		peptideBond = append(peptideBond, AminoClass{})
	}

	bioclass.ID = rand.Intn((len(class) - 4)) * len(class)

	for i := 0; i < (len(class) - 4); i++ {

		if (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "U") || (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "C") {
			bioclass.Symbol = "F"
			bioclass.Name = "Phenylalanine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = true
			bioclass.Charge = false
			bioclass.ChargeType = " "
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "A") || (class[i+0] == "U" && class[i+1] == "U" && class[i+2] == "G") || (class[i+0] == "C" && class[i+1] == "U" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "U" && class[i+2] == "C") || (class[i+0] == "C" && class[i+1] == "U" && class[i+2] == "A") || (class[i+0] == "C" && class[i+1] == "U" && class[i+2] == "G") {
			bioclass.Symbol = "L"
			bioclass.Name = "Leucine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "U" && class[i+1] == "C" && class[i+2] == "U") || (class[i+0] == "U" && class[i+1] == "C" && class[i+2] == "C") || (class[i+0] == "U" && class[i+1] == "C" && class[i+2] == "A") || (class[i+0] == "U" && class[i+1] == "C" && class[i+2] == "G") || (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "U") || (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "C") {
			bioclass.Symbol = "S"
			bioclass.Name = "Serine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.TypeChain = " "
			bioclass.Hydrophobic = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "C" && class[i+1] == "C" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "C" && class[i+2] == "C") || (class[i+0] == "C" && class[i+1] == "C" && class[i+2] == "A") || (class[i+0] == "C" && class[i+1] == "C" && class[i+2] == "G") {
			bioclass.Symbol = "P"
			bioclass.Name = "Proline"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "A" && class[i+1] == "U" && class[i+2] == "C") || (class[i+0] == "A" && class[i+1] == "U" && class[i+2] == "A") {
			bioclass.Symbol = "I"
			bioclass.Name = "Isoleucine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if class[i+0] == "A" && class[i+1] == "U" && class[i+2] == "G" {
			bioclass.Symbol = "M"
			bioclass.Name = "Methionine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.CodonStart = true
			bioclass.PKa = 0.0
			bioclass.CodonEnd = false
			bioclass.ChargeType = " "
			bioclass.Hydrophobic = false
			bioclass.TypeChain = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "U") || (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "C") || (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "A") || (class[i+0] == "A" && class[i+1] == "C" && class[i+2] == "G") {
			bioclass.Symbol = "T"
			bioclass.Name = "Threonine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			bioclass.TypeChain = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "C") || (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "A") || (class[i+0] == "G" && class[i+1] == "U" && class[i+2] == "G") {
			bioclass.Symbol = "V"
			bioclass.Name = "Valine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "C") || (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "A") || (class[i+0] == "G" && class[i+1] == "C" && class[i+2] == "G") {
			bioclass.Symbol = "A"
			bioclass.Name = "Alanine"
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "U" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "U" && class[i+1] == "A" && class[i+2] == "C") {
			bioclass.Symbol = "Y"
			bioclass.Name = "Tyrosine"
			bioclass.Polar = true
			bioclass.TypeChain = "Aromatic"
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			bioclass.Hydrophobic = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "U" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] == "U" && class[i+1] == "A" && class[i+2] == "G") || (class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "A") {
			bioclass.Symbol = "X"
			bioclass.CodonEnd = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			bioclass.Name = " "
			bioclass.TypeChain = " "
			bioclass.Polar = false
			bioclass.Hydrophobic = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "C" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "A" && class[i+2] == "C") {
			bioclass.Symbol = "H"
			bioclass.Charge = true
			bioclass.ChargeType = "Positive"
			bioclass.Name = "Histidine"
			bioclass.Polar = true
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "C" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] == "C" && class[i+1] == "A" && class[i+2] == "G") {
			bioclass.Symbol = "Q"
			bioclass.Name = "Glutamine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			bioclass.Hydrophobic = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "U") || (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "C") || (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "A") || (class[i+0] == "C" && class[i+1] == "G" && class[i+2] == "G") || (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "A") || (class[i+0] == "A" && class[i+1] == "G" && class[i+2] == "G") {
			bioclass.Symbol = "R"
			bioclass.Name = "Arginine"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.ChargeType = "Positive"
			bioclass.Hydrophobic = false
			bioclass.PKa = 12.5
			bioclass.TypeChain = " "
			bioclass.CodonEnd = false
			bioclass.CodonStart = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "C") {
			bioclass.Symbol = "N"
			bioclass.Name = "Asparagine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] == "A" && class[i+1] == "A" && class[i+2] == "G") {
			bioclass.Symbol = "K"
			bioclass.Name = "Lysine"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.ChargeType = "Positive"
			bioclass.PKa = 10.5
			bioclass.TypeChain = " "
			bioclass.CodonEnd = false
			bioclass.CodonStart = false
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "G" && class[i+1] == "A" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "A" && class[i+2] == "C") {
			bioclass.Symbol = "D"
			bioclass.Name = "Aspartic Acid"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.PKa = 3.9
			bioclass.ChargeType = "Negative"
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.TypeChain = " "

			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "G" && class[i+1] == "A" && class[i+2] == "A") || (class[i+0] == "G" && class[i+1] == "A" && class[i+2] == "G") {
			bioclass.Symbol = "E"
			bioclass.Name = "Glutamate"
			bioclass.Polar = true
			bioclass.Charge = true
			bioclass.Hydrophobic = false
			bioclass.PKa = 4.2
			bioclass.ChargeType = "Negative"
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.TypeChain = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "U") || (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "C") || (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "A") || (class[i+0] == "G" && class[i+1] == "G" && class[i+2] == "G") {
			bioclass.Symbol = "G"
			bioclass.Name = "Glycine"
			bioclass.Polar = false
			bioclass.TypeChain = "Aliphatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)

		} else if (class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "U") || (class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "C") {
			bioclass.Symbol = "C"
			bioclass.Name = "Cysteine"
			bioclass.Polar = true
			bioclass.Charge = false
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.TypeChain = " "
			bioclass.ChargeType = " "
			bioclass.Hydrophobic = false
			peptideBond = append(peptideBond, bioclass)

		} else if class[i+0] == "U" && class[i+1] == "G" && class[i+2] == "G" {
			bioclass.Symbol = "W"
			bioclass.Name = "Tryptophan"
			bioclass.Polar = false
			bioclass.TypeChain = "Aromatic"
			bioclass.Hydrophobic = true
			bioclass.PKa = 0.0
			bioclass.CodonStart = false
			bioclass.CodonEnd = false
			bioclass.Charge = false
			bioclass.ChargeType = " "
			peptideBond = append(peptideBond, bioclass)
		}
	}

	return peptideBond

}
