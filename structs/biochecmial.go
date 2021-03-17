package structs

import (
	"github.com/ali2210/wizdwarf/structs/amino"
	"github.com/fogleman/ribbon/pdb"
)

type (
	MolecularBio struct {
		AtomLevel  *pdb.Atom
		HetAtom    *pdb.Atom
		AlphaHelix *pdb.Helix
		BetaSheets *pdb.Strand
		LinkBy     []*pdb.Connection
		Chains     *pdb.Chain
		BioChem    *pdb.Matrix
		Resuide    *pdb.Residue
	}
)

func SequenceStructure(st2 string) MolecularBio {

	molecule := MolecularBio{
		AtomLevel: &pdb.Atom{Serial: 0, Name: "", AltLoc: "", ResName: "", ChainID: "", ResSeq: 0, ICode: "", X: 0, Y: 0, Z: 0, Occupancy: 0, TempFactor: 0, Element: "", Charge: ""},
		HetAtom: &pdb.Atom{
			Serial:     0,
			Name:       "",
			AltLoc:     "",
			ResName:    "",
			ChainID:    "",
			ResSeq:     0,
			ICode:      "",
			X:          0,
			Y:          0,
			Z:          0,
			Occupancy:  0,
			TempFactor: 0,
			Element:    "",
			Charge:     "",
		},
		AlphaHelix: &pdb.Helix{Serial: 0, HelixID: "", InitResName: "", InitChainID: "", InitSeqNum: 0, InitICode: "", EndResName: "", EndChainID: "", EndSeqNum: 0, EndICode: "", HelixClass: 0, Length: 0},
		BetaSheets: &pdb.Strand{Strand: 0, SheetID: "", NumStrands: 0, InitResName: "", InitChainID: "", InitSeqNum: 0, InitICode: "", EndResName: "", EndChainID: "", EndSeqNum: 0, EndICode: "", Sense: 0, CurAtom: "", CurResName: "", CurChainId: "", CurResSeq: 0, CurICode: "", PrevAtom: "", PrevResName: "", PrevChainId: "", PrevResSeq: 0, PrevICode: ""},
		LinkBy:     []*pdb.Connection{},
		Chains:     &pdb.Chain{ChainID: "", Residues: []*pdb.Residue{}},
		BioChem:    &pdb.Matrix{},
		Resuide:    &pdb.Residue{ResName: "", ChainID: "", ResSeq: 0, Atoms: []*pdb.Atom{}, AtomsByName: map[string]*pdb.Atom{}, Type: 0},
	}

	molecule.AlphaHelix = pdb.ParseHelix(st2)
	molecule.BetaSheets = pdb.ParseStrand(st2)
	molecule.AtomLevel = pdb.ParseAtom(st2)
	molecule.BioChem = &pdb.Matrix{}
	molecule.Chains = &pdb.Chain{
		ChainID:  "",
		Residues: []*pdb.Residue{},
	}
	molecule.HetAtom = molecule.AtomLevel

	molecule.Resuide = &pdb.Residue{
		ResName:     "",
		ChainID:     "",
		ResSeq:      0,
		Atoms:       []*pdb.Atom{},
		AtomsByName: map[string]*pdb.Atom{},
		Type:        0,
	}
	molecule.LinkBy = pdb.ParseConnections(st2)

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
