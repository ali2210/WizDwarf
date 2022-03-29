package molecules

import (
	"reflect"
	"strings"

	"github.com/ali2210/wizdwarf/other/proteins"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
)

// genome function return  maromolecules props .
//  this function written in fp style . In a single line of code dozen of functions depend on one another
func Genome_Extract(m map[string]map[string]proteins.Aminochain, n map[string]string, key string) *binary.Micromolecule {

	// *************** declaration of props *************
	molecules := binary.Micromolecule{}
	var molecules_traits_a string = ""
	var molecules_traits_b string = ""
	var molecule_magnetic string = ""
	var carbonAtom int64 = 0
	var hydroAtom int64 = 0
	var sulphurAtom int64 = 0
	var oxygenAtom int64 = 0
	var nitrogenAtom int64 = 0
	// ***************************************************

	// iterate over map of map with specialized keys
	iterate := reflect.ValueOf(m[n[key]]).MapRange()
	for iterate.Next() {

		// hold molecules symbol
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") {
			molecules.Symbol = iterate.Value().FieldByName("Symbol").String()
		}

		// hold molecules mass
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			molecules.Mass = iterate.Value().FieldByName("Mass").Float()
		}

		// hold molecules acidity level
		if !strings.Contains(iterate.Value().FieldByName("Acidity_a").String(), "undefined") && !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			molecules_traits_a = iterate.Value().FieldByName("Acidity_a").String()
		}

		if !strings.Contains(iterate.Value().FieldByName("Acidity_b").String(), "undefined") && !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			molecules_traits_b = iterate.Value().FieldByName("Acidity_b").String()
		}

		// hold molecule magnetic fields level
		if !strings.Contains(iterate.Value().FieldByName("Magnetic").String(), "undefined") && !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			molecule_magnetic = iterate.Value().FieldByName("Magnetic").String()
		}

		// hold molecules carbon state
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			carbonAtom = iterate.Value().FieldByName("Carbon").Int()
		}

		// hold molecules hydrogen state
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			hydroAtom = iterate.Value().FieldByName("Hydrogen").Int()
		}

		// hold molecules nitrogen state
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			nitrogenAtom = iterate.Value().FieldByName("Nitrogen").Int()
		}

		// hold molecules sulphur state
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			sulphurAtom = iterate.Value().FieldByName("Sulphur").Int()
		}

		// hold molecule oxygen state
		if !strings.Contains(iterate.Value().FieldByName("Symbol").String(), " ") && special_proteins(iterate.Value().FieldByName("Symbol").String()) {
			oxygenAtom = iterate.Value().FieldByName("Oxygen").Int()
		}

		molecules.Composition = &binary.Element{C: carbonAtom, H: hydroAtom, N: nitrogenAtom, O: oxygenAtom, S: sulphurAtom}
		molecules.Molecule = &binary.Traits{A: molecules_traits_a, B: molecules_traits_b, Magnetic_Field: molecule_magnetic}
		//log.Println("molecules:", molecules)
	}
	return &molecules
}

// codon special properties
func GetMoleculesState(molecule *binary.Micromolecule) bool {

	// check protocol message traits .
	return !strings.Contains(molecule.Symbol, " ") && !reflect.DeepEqual(molecule.Molecule, &binary.Traits{})
}

// codon checmial composition
func GetCompositionState(molecule *binary.Micromolecule) bool {

	// check element traits
	return !strings.Contains(molecule.Symbol, " ") && !reflect.DeepEqual(molecule.Composition, &binary.Element{})
}

// create molecule
func Molecular(p *binary.Micromolecule) bool {

	// check whether chain have proper type and the molecule have some information
	if reflect.ValueOf(p).Elem().Kind() == reflect.Struct {
		return !reflect.DeepEqual(reflect.ValueOf(p).Elem(), &binary.Micromolecule{})
	}
	// return false state
	return false
}

// start codon or end codon & message codon for trscription
func special_proteins(str string) bool {

	// check molecule codon have start or end codon
	if strings.Contains(str, "!") {
		return false
	} else if strings.Contains(str, "!*") {
		return false
	} else if strings.Contains(str, "!**") {
		return false
	}
	// this is a valid codon
	return true
}

// count unique codon family members
func Abundance(molecule *binary.Micromolecule, str string, iter int) (int, string) {

	count := 0
	// if molecule have valid type and that molecule exist more than once
	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {
		count = strings.Count(str, reflect.ValueOf(molecule).Elem().Field(3).String())
	}
	return count, reflect.ValueOf(molecule).Elem().Field(3).String()
}

// codon properties analysis
func Predictwithvaribles(molecule *binary.Micromolecule) float64 {

	productOf := 1.0

	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {
		productOf = ((reflect.ValueOf(molecule).Elem().Field(4).Float()) * float64(reflect.ValueOf(molecule).Elem().Field(7).Int()))
	}

	return productOf
}

// Molecule symbols
func Symbol(molecule *binary.Micromolecule) string {

	symbol := ""
	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {

		symbol = reflect.ValueOf(molecule).Elem().Field(3).String()
	}
	return symbol
}

// Molecule mass
func Mass(molecule *binary.Micromolecule) float64 {

	mass := 0.00
	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {
		mass = reflect.ValueOf(molecule).Elem().Field(4).Float()
	}
	return mass
}

// Molecule uniquness
func Occurance(molecule *binary.Micromolecule) int64 {

	var occ int64 = 0

	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {
		occ = reflect.ValueOf(molecule).Elem().Field(7).Int()
	}
	return occ
}

// Molecules sub structs
func GetStructsFields(molecule *binary.Micromolecule, fields int, name string) reflect.Value {

	if reflect.ValueOf(molecule).Elem().Kind() == reflect.Struct {

		switch name {
		case "Symbol":
			return reflect.ValueOf(molecule).Elem().Field(fields)
		case "Mass":
			return reflect.ValueOf(molecule).Elem().Field(fields)
		case "Abundance":
			return reflect.ValueOf(molecule).Elem().Field(fields)
		}
	}
	return reflect.ValueOf(molecule)
}

// create molecules analytics
func DashboardAnalytics(amino []*binary.Micromolecule, sum int64) []map[string]interface{} {

	model := make([]map[string]interface{}, len(amino))
	//predict := make([]float64, len(amino))

	for i := range amino {

		refVal := GetStructsFields(amino[i], 3, "Symbol")
		refMass := GetStructsFields(amino[i], 4, "Mass")
		refOccu := GetStructsFields(amino[i], 7, "Abundance")

		if !(strings.Contains(refVal.String(), " ")) && !(reflect.DeepEqual(refVal, reflect.ValueOf(amino[i]))) {
			if !(reflect.DeepEqual(refMass, reflect.ValueOf(amino[i]))) {
				if !(reflect.DeepEqual(refOccu, reflect.ValueOf(amino[i]))) {

					model = append(model, map[string]interface{}{
						"Symbol":    refVal.String(),
						"Mass":      refMass.Float(),
						"Occurance": refOccu.Int(),
						// "Predict" : predict[i],
						"Max": sum,
					})
				}
			}
		}
	}

	return model
}

// Get name of the molecules
func Make_string(chain *binary.Micromolecule) string {

	symbols := ""
	if reflect.ValueOf(chain).Elem().Kind() == reflect.Struct {
		symbols = reflect.ValueOf(chain).Elem().Field(3).String()
	}
	return symbols
}
