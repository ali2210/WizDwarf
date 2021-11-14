/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package proteins

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	"google.golang.org/api/iterator"
)

// protocol buffer messages services
type AbstractBiomolecules interface {
	AddPDB(context context.Context, list *binary.Micromolecule_List) *binary.MolecularState
	DisplayPDB(context context.Context, req *binary.Request) *binary.Micromolecule_List
}

// biomolecules data object
type Biomolecules struct{}

// variables which are initailize at time of calling
var client *firestore.Client
var Size_Bond int
var ckk string

// service name
const hottopic = "peptideBond_Topic"

// error description
const error_desc = "data stream is not accepted because rules voliation"

// Biomolecules new object or receiver object
func NewPeptideTopic() AbstractBiomolecules { return &Biomolecules{} }

func (biomolecules *Biomolecules) AddPDB(context context.Context, list *binary.Micromolecule_List) *binary.MolecularState {

	i := 0
	var empty_rec map[string]interface{}
	var docx *firestore.DocumentSnapshot

	// initial i declare at 0 which means i have enough lifespan
	if i != Size_Bond {

		// check whether data exists
		query := client.Collection(hottopic).Where("ckk", "==", ckk).Documents(context)
		for {
			// iterator point to the document if exist
			doc, err := query.Next()
			if err != iterator.Done {
				break
			}

			check := doc.Data()
			docx = doc

			// if the iterator point empty document
			if reflect.DeepEqual(check, empty_rec) {

				// create new document in the database
				doc, result, err := client.Collection(hottopic).Add(context, map[string]interface{}{
					"symbol":    list.Peplide[i].Symbol,
					"mass":      list.Peplide[i].Mass,
					"carbon":    list.Peplide[i].Composition.C,
					"hydrogen":  list.Peplide[i].Composition.H,
					"sulfpur":   list.Peplide[i].Composition.S,
					"nitrogen":  list.Peplide[i].Composition.N,
					"oxygen":    list.Peplide[i].Composition.O,
					"acid_a":    list.Peplide[i].Molecule.A,
					"acid_b":    list.Peplide[i].Molecule.B,
					"magnetism": list.Peplide[i].Molecule.Magnetic_Field,
					"ckk":       ckk,
				})
				if err != nil {
					log.Printf(" Error add your record in the topic:%v", err.Error())
					return &binary.MolecularState{State: false, Error: error_desc}
				}
				log.Println("Document created", doc, "Result:", result)
			}

			// update offset of a list
			i = i + 1

			// document already have created and have some data; list still have data
			docUpdate, err := docx.Ref.Set(context, map[string]interface{}{
				"symbol":    list.Peplide[i].Symbol,
				"mass":      list.Peplide[i].Mass,
				"carbon":    list.Peplide[i].Composition.C,
				"hydrogen":  list.Peplide[i].Composition.H,
				"sulfpur":   list.Peplide[i].Composition.S,
				"nitrogen":  list.Peplide[i].Composition.N,
				"oxygen":    list.Peplide[i].Composition.O,
				"acid_a":    list.Peplide[i].Molecule.A,
				"acid_b":    list.Peplide[i].Molecule.B,
				"magnetism": list.Peplide[i].Molecule.Magnetic_Field,
				"ckk":       ckk,
			})
			if err != nil {
				log.Printf(" Error document updating molecule%v", err.Error())
				return &binary.MolecularState{State: false, Error: error_desc}
			}
			log.Println("document updates:", docUpdate, "epoch:", i)
		}
	}
	return &binary.MolecularState{State: true, Error: ""}
}

func (biomolecules *Biomolecules) DisplayPDB(context context.Context, req *binary.Request) *binary.Micromolecule_List {

	var list binary.Micromolecule_List
	query := client.Collection(hottopic).Where("ckk", "==", ckk).Documents(context)
	for {
		doc, err := query.Next()
		if err != iterator.Done {
			break
		}
		result := doc.Data()
		marshaldata, err := json.Marshal(result)
		if err != nil {
			log.Printf(" Error marshaling result%v", err.Error())
			return &binary.Micromolecule_List{}
		}
		err = json.Unmarshal(marshaldata, &list)
		if err != nil {
			log.Printf(" Error unmarshaling list%v", err.Error())
			return &binary.Micromolecule_List{}
		}
	}
	return &list
}
