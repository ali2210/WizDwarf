package structs

import (
	"github.com/ali2210/wizdwarf/structs/amino"
	"github.com/fogleman/ribbon/pdb"
)

type MolecularBio struct {

	// Helix Object Reference
	HelixA pdb.Helix
	// Strand Object Reference
	StrandB pdb.Strand
}

func SequenceStructure(st2 string) MolecularBio {

	molecule := MolecularBio{}

	hlix := *pdb.ParseHelix(st2)
	stand := *pdb.ParseStrand(st2)

	molecule.HelixA = hlix
	molecule.StrandB = stand

	return molecule
}

func RNA(sq []string) []string {

	var k []string

	for i := range sq {

		if sq[i] == "T" {
			sq[i] = "U"
		}
		k = append(k, sq[i])
	}

	return k

}

func Proteins(s []string) []amino.AminoClass {

	bases := []string{}
	for i := range s {
		bases = append(bases, s[i])
	}

	proteins := amino.AminoClass{}

	ls := proteins.Bases(bases)

	return ls
}
