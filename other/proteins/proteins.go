/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package proteins

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
var total int64 = 0

// ******************************* "0" indicate universal constant which means zero level ****************
// Magnetic = "0" no presence
// Element {c, n, s} = 0 atom
// acidic_a , acidic_b = 0 no acidic_a or no acidic_b

// proteins chains symbols
// Example Class
// symbols := class("TTT", 0, 3)

// @param string message , first & last index as int
// @return string message

func Class(s string, i, j int) string {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"

	case "TTC":
		aminochain.Symbol = "F"

	case "TTA":
		aminochain.Symbol = "L"

	case "TTG":
		aminochain.Symbol = "L"

	case "CTT":
		aminochain.Symbol = "L"

	case "CTC":
		aminochain.Symbol = "L"

	case "CTA":
		aminochain.Symbol = "L"

	case "CTG":
		aminochain.Symbol = "L"

	case "ATT":
		aminochain.Symbol = "I"

	case "ATC":
		aminochain.Symbol = "I"

	case "ATA":
		aminochain.Symbol = "M"

	case "ATG":
		aminochain.Symbol = "M"

	case "GTT":
		aminochain.Symbol = "V"

	case "GTC":
		aminochain.Symbol = "V"

	case "GTA":
		aminochain.Symbol = "V"

	case "GTG":
		aminochain.Symbol = "V"

	case "TCT":
		aminochain.Symbol = "S"

	case "TCC":
		aminochain.Symbol = "S"

	case "TCA":
		aminochain.Symbol = "S"

	case "TCG":
		aminochain.Symbol = "S"

	case "CCT":
		aminochain.Symbol = "P"

	case "CCC":
		aminochain.Symbol = "P"

	case "CCA":
		aminochain.Symbol = "P"

	case "CCG":
		aminochain.Symbol = "P"

	case "ACT":
		aminochain.Symbol = "T"

	case "ACC":
		aminochain.Symbol = "T"

	case "ACA":
		aminochain.Symbol = "T"

	case "ACG":
		aminochain.Symbol = "T"

	case "GCT":
		aminochain.Symbol = "A"

	case "GCC":
		aminochain.Symbol = "A"

	case "GCA":
		aminochain.Symbol = "A"

	case "GCG":
		aminochain.Symbol = "A"

	case "TAT":
		aminochain.Symbol = "Y"

	case "TAC":
		aminochain.Symbol = "Y"

	case "TAA":
		aminochain.Symbol = "!"

	case "TAG":
		aminochain.Symbol = "!*"

	case "CAT":
		aminochain.Symbol = "H"

	case "CAC":
		aminochain.Symbol = "H"

	case "CAA":
		aminochain.Symbol = "Q"

	case "CAG":
		aminochain.Symbol = "Q"

	case "AAT":
		aminochain.Symbol = "N"

	case "AAC":
		aminochain.Symbol = "N"

	case "AAA":
		aminochain.Symbol = "K"

	case "AAG":
		aminochain.Symbol = "K"

	case "GAT":
		aminochain.Symbol = "D"

	case "GAC":
		aminochain.Symbol = "D"

	case "GAA":
		aminochain.Symbol = "E"

	case "GAG":
		aminochain.Symbol = "E"

	case "TGT":
		aminochain.Symbol = "C"

	case "TGC":
		aminochain.Symbol = "C"

	case "TGA":
		aminochain.Symbol = "!**"

	case "TGG":
		aminochain.Symbol = "!**"

	case "CGT":
		aminochain.Symbol = "R"

	case "CGC":
		aminochain.Symbol = "R"

	case "CGA":
		aminochain.Symbol = "R"

	case "CGG":
		aminochain.Symbol = "R"

	case "AGT":
		aminochain.Symbol = "S"

	case "AGC":
		aminochain.Symbol = "S"

	case "AGA":
		aminochain.Symbol = "R"

	case "AGG":
		aminochain.Symbol = "R"

	case "GGT":
		aminochain.Symbol = "G"

	case "GGC":
		aminochain.Symbol = "G"

	case "GGA":
		aminochain.Symbol = "G"

	case "GGG":
		aminochain.Symbol = "G"

	default:
		return " "
	}
	return aminochain.Symbol
}

//Potein chain molar mass
// @param string message , first & last index as int
// @return string message, int

func GetMolarMass(s string, i, j int) (float64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Mass = 165.192

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Mass = 131.175

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Mass = 131.175

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Mass = 149.21

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Mass = 117.148

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Mass = 115.132

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Mass = 119.120

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Mass = 89.094

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Mass = 181.191

	case "TAA":
		aminochain.Symbol = "!"

	case "TAG":
		aminochain.Symbol = "!*"

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Mass = 155.157

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Mass = 146.146

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Mass = 132.119

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Mass = 146.190

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Mass = 133.103

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Mass = 147.130

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Mass = 121.15

	case "TGA":
		aminochain.Symbol = "!**"

	case "TGG":
		aminochain.Symbol = "!**"

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Mass = 105.093

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Mass = 174.204

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Mass = 75.067

	}
	return aminochain.Mass, aminochain.Symbol
}

// proteins chain acidiity level
// @param string message , first & last index as int
// @return string messages
func GetPKa(str string, i, j int) (string, string, string) {
	switch str[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Acidity_a = "1.83"
		aminochain.Acidity_b = "9.13"

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Acidity_a = "2.36"
		aminochain.Acidity_b = "9.60"

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Acidity_a = "2.28"
		aminochain.Acidity_b = "9.21"

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Acidity_a = "2.32"
		aminochain.Acidity_b = "9.62"

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "10.96"

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Acidity_a = "2.63"
		aminochain.Acidity_b = "10.43"

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.87"

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Acidity_a = "2.2"
		aminochain.Acidity_b = "9.1"

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Acidity_a = "2.1"
		aminochain.Acidity_b = "8.80"

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Acidity_a = "1.99"
		aminochain.Acidity_b = "9.90" + "*/sc" + "3.90" //*/sc means side-chain pk_A

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Acidity_a = "2.10"
		aminochain.Acidity_b = "9.47" + "*/sc" + "4.07"

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Acidity_a = " "
		aminochain.Acidity_b = " "

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Acidity_a = "2.21"
		aminochain.Acidity_b = "9.15"

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Acidity_a = "2.18"
		aminochain.Acidity_b = "9.09" + "*/gc" + "13.2" //gc meeans guanidino

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Acidity_a = "2.34"
		aminochain.Acidity_b = "9.6"

	default:
		return " ", " ", " "
	}
	return aminochain.Acidity_a, aminochain.Acidity_b, aminochain.Symbol
}

// proteins chain contains carbon details
// @param string message , first & last index as int
// @return string message, int
func GetCarbon(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Carbon = 9

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Carbon = 9

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Carbon = 6

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Carbon = 6

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Carbon = 6

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Carbon = 5

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Carbon = 5

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Carbon = 5

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Carbon = 5

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Carbon = 4

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Carbon = 3

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Carbon = 9

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Carbon = 9

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Carbon = 0

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Carbon = 0

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Carbon = 6

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Carbon = 6

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Carbon = 5

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Carbon = 5

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Carbon = 4

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Carbon = 4

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Carbon = 6

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Carbon = 6

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Carbon = 4

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Carbon = 4

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Carbon = 5

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Carbon = 5

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Carbon = 3

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Carbon = 3

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Carbon = 0

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Carbon = 0

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Carbon = 3

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Carbon = 6

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Carbon = 2

	default:
		return 0, " "
	}
	return aminochain.Carbon, aminochain.Symbol
}

// proteins chain contains "h" atoms
// @param string message , first & last index as int
// @return string message, int
func GetHydrogen(str string, i, j int) (int64, string) {
	switch str[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Hydrogen = 11

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Hydrogen = 11

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Hydrogen = 13

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Hydrogen = 13

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Hydrogen = 13

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Hydrogen = 11

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Hydrogen = 11

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Hydrogen = 11

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Hydrogen = 9

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Hydrogen = 9

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Hydrogen = 7

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Hydrogen = 11

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Hydrogen = 11

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Hydrogen = 0

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Hydrogen = 0

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Hydrogen = 9

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Hydrogen = 9

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Hydrogen = 10

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Hydrogen = 10

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Hydrogen = 8

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Hydrogen = 8

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Hydrogen = 14

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Hydrogen = 14

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Hydrogen = 7

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Hydrogen = 7

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Hydrogen = 9

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Hydrogen = 9

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Hydrogen = 7

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Hydrogen = 7

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Hydrogen = 0

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Hydrogen = 0

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Hydrogen = 7

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Hydrogen = 14

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Hydrogen = 5

	default:
		return 0, " "
	}
	return aminochain.Hydrogen, aminochain.Symbol
}

// protein chain contains "o" atoms
// @param string message , first & last index as int
// @return string message, int
func GetOxgygen(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Oxygen = 2

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Oxygen = 2

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Oxygen = 2

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Oxygen = 2

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Oxygen = 2

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Oxygen = 2

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Oxygen = 2

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Oxygen = 2

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Oxygen = 2

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Oxygen = 3

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Oxygen = 2

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Oxygen = 3

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Oxygen = 3

	case "TAA":
		aminochain.Symbol = "!"

	case "TAG":
		aminochain.Symbol = "!*"

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Oxygen = 2

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Oxygen = 2

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Oxygen = 3

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Oxygen = 3

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Oxygen = 3

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Oxygen = 3

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Oxygen = 2

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Oxygen = 2

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Oxygen = 4

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Oxygen = 4

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Oxygen = 4

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Oxygen = 4

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Oxygen = 2

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Oxygen = 2

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Oxygen = 0

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Oxygen = 0

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Oxygen = 3

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Oxygen = 2

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Oxygen = 2

	default:
		return 0, " "
	}
	return aminochain.Oxygen, aminochain.Symbol
}

// proteins chain contains "n" atoms
// @param string message , first & last index as int
// @return string message, int64
func GetNitrogen(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Nitrogen = 1

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Nitrogen = 1

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Nitrogen = 1

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Nitrogen = 1

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Nitrogen = 1

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Nitrogen = 1

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Nitrogen = 1

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Nitrogen = 1

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Nitrogen = 1

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Nitrogen = 1

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Nitrogen = 1

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Nitrogen = 1

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Nitrogen = 1

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Nitrogen = 0

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Nitrogen = 0

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Nitrogen = 3

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Nitrogen = 3

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Nitrogen = 2

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Nitrogen = 2

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Nitrogen = 2

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Nitrogen = 2

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Nitrogen = 2

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Nitrogen = 2

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Nitrogen = 1

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Nitrogen = 1

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Nitrogen = 1

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Nitrogen = 1

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Nitrogen = 1

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Nitrogen = 1

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Nitrogen = 0

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Nitrogen = 0

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Nitrogen = 1

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Nitrogen = 4

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Nitrogen = 1

	default:
		return 0, " "
	}
	return aminochain.Nitrogen, aminochain.Symbol
}

// proteins chain contains "s" atoms
// @param string message , first & last index as int
// @return string message and int
func GetSulphur(s string, i, j int) (int64, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Sulphur = 0

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Sulphur = 0

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Sulphur = 0

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Sulphur = 0

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Sulphur = 0

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Sulphur = 1

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Sulphur = 1

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Sulphur = 0

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Sulphur = 0

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Sulphur = 0

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Sulphur = 0

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Sulphur = 0

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Sulphur = 0

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Sulphur = 0

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Sulphur = 0

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Sulphur = 0

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Sulphur = 0

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Sulphur = 0

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Sulphur = 0

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Sulphur = 0

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Sulphur = 0

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Sulphur = 0

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Sulphur = 0

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Sulphur = 0

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Sulphur = 0

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Sulphur = 0

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Sulphur = 0

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Sulphur = 1

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Sulphur = 1

	case "TGA":
		aminochain.Symbol = "!**"
		aminochain.Sulphur = 0

	case "TGG":
		aminochain.Symbol = "!**"
		aminochain.Sulphur = 0

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Sulphur = 0

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Sulphur = 0

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Sulphur = 0

	default:
		return 0, " "
	}
	return aminochain.Sulphur, aminochain.Symbol
}

// proteins behaves wired in the presence of magnetic field.
// @param string message , first & last index as int
// @return string message

func GetMagnetism(s string, i, j int) (string, string) {
	switch s[i:j] {
	case "TTT":
		aminochain.Symbol = "F"
		aminochain.Magnetic = "0.0"

	case "TTC":
		aminochain.Symbol = "F"
		aminochain.Magnetic = "0.0"

	case "TTA":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "TTG":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "CTT":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "CTC":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "CTA":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "CTG":
		aminochain.Symbol = "L"
		aminochain.Magnetic = "-84.9·10−6"

	case "ATT":
		aminochain.Symbol = "I"
		aminochain.Magnetic = "−84.9·10−6"

	case "ATC":
		aminochain.Symbol = "I"
		aminochain.Magnetic = "−84.9·10−6"

	case "ATA":
		aminochain.Symbol = "M"
		aminochain.Magnetic = "0"

	case "ATG":
		aminochain.Symbol = "M"
		aminochain.Magnetic = "0"

	case "GTT":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"

	case "GTC":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"

	case "GTA":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"

	case "GTG":
		aminochain.Symbol = "V"
		aminochain.Magnetic = "-74.3·10−6"

	case "TCT":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "TCC":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "TCA":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "TCG":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "CCT":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"

	case "CCC":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"

	case "CCA":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"

	case "CCG":
		aminochain.Symbol = "P"
		aminochain.Magnetic = "0"

	case "ACT":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"

	case "ACC":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"

	case "ACA":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"

	case "ACG":
		aminochain.Symbol = "T"
		aminochain.Magnetic = "0"

	case "GCT":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"

	case "GCC":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"

	case "GCA":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"

	case "GCG":
		aminochain.Symbol = "A"
		aminochain.Magnetic = "-50.5·10−6"

	case "TAT":
		aminochain.Symbol = "Y"
		aminochain.Magnetic = "-105.3·10−6"

	case "TAC":
		aminochain.Symbol = "Y"
		aminochain.Magnetic = "-105.3·10−6"

	case "TAA":
		aminochain.Symbol = "!"
		aminochain.Magnetic = "0"

	case "TAG":
		aminochain.Symbol = "!*"
		aminochain.Magnetic = "0"

	case "CAT":
		aminochain.Symbol = "H"
		aminochain.Magnetic = "0"

	case "CAC":
		aminochain.Symbol = "H"
		aminochain.Magnetic = "0"

	case "CAA":
		aminochain.Symbol = "Q"
		aminochain.Magnetic = "0"

	case "CAG":
		aminochain.Symbol = "Q"
		aminochain.Magnetic = "0"

	case "AAT":
		aminochain.Symbol = "N"
		aminochain.Magnetic = "-69.5·10−6"

	case "AAC":
		aminochain.Symbol = "N"
		aminochain.Magnetic = "-69.5·10−6"

	case "AAA":
		aminochain.Symbol = "K"
		aminochain.Magnetic = "0"

	case "AAG":
		aminochain.Symbol = "K"
		aminochain.Magnetic = "0"

	case "GAT":
		aminochain.Symbol = "D"
		aminochain.Magnetic = "-64.2·10−6"

	case "GAC":
		aminochain.Symbol = "D"
		aminochain.Magnetic = "-64.2·10−6"

	case "GAA":
		aminochain.Symbol = "E"
		aminochain.Magnetic = "−78.5·10−6"

	case "GAG":
		aminochain.Symbol = "E"
		aminochain.Magnetic = "−78.5·10−6"

	case "TGT":
		aminochain.Symbol = "C"
		aminochain.Magnetic = "0"

	case "TGC":
		aminochain.Symbol = "C"
		aminochain.Magnetic = "0"

	case "TGA":
		aminochain.Symbol = "!**"

	case "TGG":
		aminochain.Symbol = "!**"

	case "CGT":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "CGC":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "CGA":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "CGG":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "AGT":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "AGC":
		aminochain.Symbol = "S"
		aminochain.Magnetic = "0"

	case "AGA":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "AGG":
		aminochain.Symbol = "R"
		aminochain.Magnetic = "0"

	case "GGT":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"

	case "GGC":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"

	case "GGA":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"

	case "GGG":
		aminochain.Symbol = "G"
		aminochain.Magnetic = "-40.3·10−6"

	default:
		return " ", " "
	}
	return aminochain.Magnetic, aminochain.Symbol
}

// This function return physical propperties of codon pair
// @param string message , first & last index value as int
// @return aminochain
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

	default:
		return Aminochain{}
	}
	return aminochain
}

// @param radi int
// @return int
func Total_chain_filter(radi int64) int64 {

	return total + radi
}

// @param sum as int
// @return as float
func AminoHealth(sum int64) float64 {
	return float64(sum / 21)
}
