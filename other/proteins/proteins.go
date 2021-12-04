/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package proteins

import (
	"log"
	"strconv"
	"strings"
)

type Aminochain struct {
	Symbol    string
	Mass      float64
	Acidity_a string
	Acidity_b string
	Carbon    int64
	Hydrogen  int64
	Oxygen    int64
	Sulphur   int64
	Nitrogen  int64
	Magnetic  string
}

// aminochain object
var aminochain Aminochain = Aminochain{}
var total = 0

// ******************************* "0" indicate universal constant which means zero level ****************
// Magnetic = "0" no presence
// Element {c, n, s} = 0 atom
// acidic_a , acidic_b = 0 no acidic_a or no acidic_b

// proteins chains symbols
// Example Class
// symbols := class("TTT", 0, 3)

func Class(s string, i, j int) string {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		break
	case "TTC":
		aminochain.Symbol = "F"
		break
	case "TTA":
		aminochain.Symbol = "L"
		break
	case "TTG":
		aminochain.Symbol = "L"
		break
	case "CTT":
		aminochain.Symbol = "L"
		break
	case "CTC":
		aminochain.Symbol = "L"
		break
	case "CTA":
		aminochain.Symbol = "L"
		break
	case "CTG":
		aminochain.Symbol = "L"
		break
	case "ATT":
		aminochain.Symbol = "I"
		break
	case "ATC":
		aminochain.Symbol = "I"
		break
	case "ATA":
		aminochain.Symbol = "M"
		break
	case "ATG":
		aminochain.Symbol = "M"
		break
	case "GTT":
		aminochain.Symbol = "V"
		break
	case "GTC":
		aminochain.Symbol = "V"
		break
	case "GTA":
		aminochain.Symbol = "V"
		break
	case "GTG":
		aminochain.Symbol = "V"
		break
	case "TCT":
		aminochain.Symbol = "S"
		break
	case "TCC":
		aminochain.Symbol = "S"
		break
	case "TCA":
		aminochain.Symbol = "S"
		break
	case "TCG":
		aminochain.Symbol = "S"
		break
	case "CCT":
		aminochain.Symbol = "P"
		break
	case "CCC":
		aminochain.Symbol = "P"
		break
	case "CCA":
		aminochain.Symbol = "P"
		break
	case "CCG":
		aminochain.Symbol = "P"
		break
	case "ACT":
		aminochain.Symbol = "T"
		break
	case "ACC":
		aminochain.Symbol = "T"
		break
	case "ACA":
		aminochain.Symbol = "T"
		break
	case "ACG":
		aminochain.Symbol = "T"
		break
	case "GCT":
		aminochain.Symbol = "A"
		break
	case "GCC":
		aminochain.Symbol = "A"
		break
	case "GCA":
		aminochain.Symbol = "A"
		break
	case "GCG":
		aminochain.Symbol = "A"
		break
	case "TAT":
		aminochain.Symbol = "Y"
		break
	case "TAC":
		aminochain.Symbol = "Y"
		break
	case "TAA":
		aminochain.Symbol = "!"
		break
	case "TAG":
		aminochain.Symbol = "!*"
		break
	case "CAT":
		aminochain.Symbol = "H"
		break
	case "CAC":
		aminochain.Symbol = "H"
		break
	case "CAA":
		aminochain.Symbol = "Q"
		break
	case "CAG":
		aminochain.Symbol = "Q"
		break
	case "AAT":
		aminochain.Symbol = "N"
		break
	case "AAC":
		aminochain.Symbol = "N"
		break
	case "AAA":
		aminochain.Symbol = "K"
		break
	case "AAG":
		aminochain.Symbol = "K"
		break
	case "GAT":
		aminochain.Symbol = "D"
		break
	case "GAC":
		aminochain.Symbol = "D"
		break
	case "GAA":
		aminochain.Symbol = "E"
		break
	case "GAG":
		aminochain.Symbol = "E"
		break
	case "TGT":
		aminochain.Symbol = "C"
		break
	case "TGC":
		aminochain.Symbol = "C"
		break
	case "TGA":
		aminochain.Symbol = "!**"
		break
	case "TGG":
		aminochain.Symbol = "!**"
		break
	case "CGT":
		aminochain.Symbol = "R"
		break
	case "CGC":
		aminochain.Symbol = "R"
		break
	case "CGA":
		aminochain.Symbol = "R"
		break
	case "CGG":
		aminochain.Symbol = "R"
		break
	case "AGT":
		aminochain.Symbol = "S"
		break
	case "AGC":
		aminochain.Symbol = "S"
		break
	case "AGA":
		aminochain.Symbol = "R"
		break
	case "AGG":
		aminochain.Symbol = "R"
		break
	case "GGT":
		aminochain.Symbol = "G"
		break
	case "GGC":
		aminochain.Symbol = "G"
		break
	case "GGA":
		aminochain.Symbol = "G"
		break
	case "GGG":
		aminochain.Symbol = "G"
		break
	default:
		return " "
	}
	return aminochain.Symbol
}

//Potein chain molar mass

func GetMolarMass(s string, i, j int) (float64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191
		break
	case "TAA":
		aminochain.Symbol = "!"
		break
	case "TAG":
		aminochain.Symbol = "!*"
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15
		break
	case "TGA":
		aminochain.Symbol = "!**"
		break
	case "TGG":
		aminochain.Symbol = "!**"
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		break
	}
	return aminochain.Mass, aminochain.Symbol
}

// proteins chain acidiity level
func GetPKa(str string, i, j int) (string, string, string) {
	switch str[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		break
	default:
		return " ", " ", " "
	}
	return aminochain.Acidity_a, aminochain.Acidity_b, aminochain.Symbol
}

// proteins chain contains carbon details
func GetCarbon(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Carbon = 9
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Carbon = 9
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Carbon = 6
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Carbon = 6
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Carbon = 5
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Carbon = 5
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Carbon = 9
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Carbon = 9
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Carbon = 0
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Carbon = 0
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Carbon = 6
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Carbon = 6
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Carbon = 5
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Carbon = 5
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Carbon = 4
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Carbon = 4
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Carbon = 6
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Carbon = 6
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Carbon = 4
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Carbon = 4
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Carbon = 5
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Carbon = 5
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Carbon = 3
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Carbon = 3
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Carbon = 0
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Carbon = 0
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2
		break
	default:
		return 0, " "
	}
	return aminochain.Carbon, aminochain.Symbol
}

// proteins chain contains "h" atoms
func GetHydrogen(str string, i, j int) (int64, string) {
	switch str[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Hydrogen = 11
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Hydrogen = 11
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Hydrogen = 13
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Hydrogen = 13
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Hydrogen = 11
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Hydrogen = 11
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Hydrogen = 11
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Hydrogen = 11
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Hydrogen = 0
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Hydrogen = 0
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Hydrogen = 9
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Hydrogen = 9
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Hydrogen = 10
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Hydrogen = 10
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Hydrogen = 8
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Hydrogen = 8
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Hydrogen = 14
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Hydrogen = 14
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Hydrogen = 7
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Hydrogen = 7
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Hydrogen = 9
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Hydrogen = 9
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Hydrogen = 7
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Hydrogen = 7
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Hydrogen = 0
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Hydrogen = 0
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5
		break
	default:
		return 0, " "
	}
	return aminochain.Hydrogen, aminochain.Symbol
}

// protein chain contains "o" atoms
func GetOxgygen(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Oxygen = 2
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Oxygen = 2
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Oxygen = 2
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Oxygen = 2
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Oxygen = 2
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Oxygen = 2
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Oxygen = 3
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Oxygen = 3
		break
	case "TAA":
		aminochain.Symbol = "!"
		break
	case "TAG":
		aminochain.Symbol = "!*"
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Oxygen = 2
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Oxygen = 2
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Oxygen = 3
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Oxygen = 3
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Oxygen = 3
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Oxygen = 3
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Oxygen = 2
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Oxygen = 2
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Oxygen = 4
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Oxygen = 4
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Oxygen = 4
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Oxygen = 4
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Oxygen = 2
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Oxygen = 2
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Oxygen = 0
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Oxygen = 0
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2
		break
	default:
		return 0, " "
	}
	return aminochain.Oxygen, aminochain.Symbol
}

// proteins chain contains "n" atoms
func GetNitrogen(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Nitrogen = 1
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Nitrogen = 1
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Nitrogen = 1
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Nitrogen = 1
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Nitrogen = 1
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Nitrogen = 1
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Nitrogen = 1
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Nitrogen = 1
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Nitrogen = 0
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Nitrogen = 0
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Nitrogen = 3
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Nitrogen = 3
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Nitrogen = 2
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Nitrogen = 2
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Nitrogen = 2
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Nitrogen = 2
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Nitrogen = 2
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Nitrogen = 2
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Nitrogen = 1
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Nitrogen = 1
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Nitrogen = 1
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Nitrogen = 1
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Nitrogen = 1
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Nitrogen = 1
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Nitrogen = 0
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Nitrogen = 0
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1
		break
	default:
		return 0, " "
	}
	return aminochain.Nitrogen, aminochain.Symbol
}

// proteins chain contains "s" atoms
func GetSulphur(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Sulphur = 0
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Sulphur = 0
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Sulphur = 0
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Sulphur = 0
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Sulphur = 1
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Sulphur = 1
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Sulphur = 0
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Sulphur = 0
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Sulphur = 0
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Sulphur = 0
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Sulphur = 0
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Sulphur = 0
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Sulphur = 0
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Sulphur = 0
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Sulphur = 0
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Sulphur = 0
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Sulphur = 0
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Sulphur = 0
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Sulphur = 0
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Sulphur = 0
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Sulphur = 0
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Sulphur = 0
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Sulphur = 1
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Sulphur = 1
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Sulphur = 0
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Sulphur = 0
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0
		break
	default:
		return 0, " "
	}
	return aminochain.Sulphur, aminochain.Symbol
}

// proteins behaves wired in the presence of magnetic field.

func GetMagnetism(s string, i, j int) (string, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Magnetic = "0.0"
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Magnetic = "0.0"
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Magnetic = "−84.9·10−6"
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Magnetic = "−84.9·10−6"
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Magnetic = "0"
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Magnetic = "0"
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Magnetic = "-105.3·10−6"
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Magnetic = "-105.3·10−6"
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Magnetic = "0"
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Magnetic = "0"
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Magnetic = "0"
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Magnetic = "0"
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Magnetic = "0"
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Magnetic = "0"
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Magnetic = "-69.5·10−6"
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Magnetic = "-69.5·10−6"
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Magnetic = "0"
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Magnetic = "0"
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Magnetic = "-64.2·10−6"
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Magnetic = "-64.2·10−6"
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Magnetic = "−78.5·10−6"
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Magnetic = "−78.5·10−6"
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Magnetic = "0"
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Magnetic = "0"
		break
	case "TGA":
		aminochain.Symbol = "!**"
		break
	case "TGG":
		aminochain.Symbol = "!**"
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"
		break
	default:
		return " ", " "
	}
	return aminochain.Magnetic, aminochain.Symbol
}

func GetAmino(s string, i, j int) Aminochain {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"
		aminochain.Carbon = 9
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"
		aminochain.Carbon = 9
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-84.9·10−6"
		break
	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175
		aminochain.Magnetic = "−84.9·10−6"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175
		aminochain.Magnetic = "−84.9·10−6"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 13
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 1
		aminochain.Magnetic = "0.00"
		break
	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 1
		aminochain.Magnetic = "0.00"
		break
	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-74.3·10−6"
		break
	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-50.5·10−6"
		break
	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 9
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Magnetic = "-105.3·10−6"
		break
	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 9
		aminochain.Hydrogen = 11
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-105.3·10−6"
		break
	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Mass = 0.00
		aminochain.Acidity_a = "undefined"
		aminochain.Acidity_b = "undefined"
		aminochain.Carbon = 0
		aminochain.Hydrogen = 0
		aminochain.Nitrogen = 0
		aminochain.Oxygen = 0
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Mass = 0.00
		aminochain.Acidity_a = "undefined"
		aminochain.Acidity_b = "undefined"
		aminochain.Carbon = 0
		aminochain.Hydrogen = 0
		aminochain.Nitrogen = 0
		aminochain.Oxygen = 0
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 3
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 3
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 10
		aminochain.Oxygen = 3
		aminochain.Nitrogen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 10
		aminochain.Oxygen = 3
		aminochain.Nitrogen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 8
		aminochain.Oxygen = 3
		aminochain.Nitrogen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-69.5·10−6"
		break
	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"
		aminochain.Carbon = 4
		aminochain.Hydrogen = 8
		aminochain.Oxygen = 3
		aminochain.Nitrogen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-69.5·10−6"
		break
	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 2
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 2
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A
		aminochain.Carbon = 4
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 4
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-64.2·10−6"
		break
	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A
		aminochain.Carbon = 4
		aminochain.Hydrogen = 7
		aminochain.Oxygen = 4
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-64.2·10−6"
		break
	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 4
		aminochain.Sulphur = 0
		aminochain.Magnetic = "−78.5·10−6"
		break
	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"
		aminochain.Carbon = 5
		aminochain.Hydrogen = 9
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 4
		aminochain.Sulphur = 0
		aminochain.Magnetic = "−78.5·10−6"
		break
	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 1
		aminochain.Magnetic = "0.00"
		break
	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15
		aminochain.Acidity_a = "0.00"
		aminochain.Acidity_b = "0.00"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 2
		aminochain.Sulphur = 1
		aminochain.Magnetic = "0.00"
		break
	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Mass = 0.00
		aminochain.Acidity_a = "undefined"
		aminochain.Acidity_b = "undefined"
		aminochain.Carbon = 0
		aminochain.Hydrogen = 0
		aminochain.Nitrogen = 0
		aminochain.Oxygen = 0
		aminochain.Sulphur = 0
		aminochain.Magnetic = "undefined"
		break
	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Mass = 0.00
		aminochain.Acidity_a = "undefined"
		aminochain.Acidity_b = "undefined"
		aminochain.Carbon = 0
		aminochain.Hydrogen = 0
		aminochain.Nitrogen = 0
		aminochain.Oxygen = 0
		aminochain.Sulphur = 0
		aminochain.Magnetic = "undefined"
		break
	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"
		aminochain.Carbon = 3
		aminochain.Hydrogen = 7
		aminochain.Nitrogen = 1
		aminochain.Oxygen = 3
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino
		aminochain.Carbon = 6
		aminochain.Hydrogen = 14
		aminochain.Nitrogen = 4
		aminochain.Oxygen = 2
		aminochain.Sulphur = 0
		aminochain.Magnetic = "0.00"
		break
	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		aminochain.Carbon = 2
		aminochain.Hydrogen = 5
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		aminochain.Carbon = 2
		aminochain.Hydrogen = 5
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		aminochain.Carbon = 2
		aminochain.Hydrogen = 5
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-40.3·10−6"
		break
	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"
		aminochain.Carbon = 2
		aminochain.Hydrogen = 5
		aminochain.Oxygen = 2
		aminochain.Nitrogen = 1
		aminochain.Sulphur = 0
		aminochain.Magnetic = "-40.3·10−6"
		break
	default:
		return Aminochain{}
	}
	return aminochain
}

// total mass of a proteins chain, data exist in string format. For Add up we used default library
func TotalChainMass(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Printf(" Error converting :%v", err.Error())
		return 0
	}
	total += num
	return total
}

// UQProteins perform two operations. First is to calculate unique protein in a chain. (int)
// Second calculate proteins chain length
func UQProteins(s string, i, j int) int {
	return strings.Count(s, s[i:j])
}

// An healthy person have 20 amino acid in their celluar bodies but we used 21 because there is some research 21 amino acid also exist
func HealthCheck(a int) float64 {
	return float64(a) / 21
}
