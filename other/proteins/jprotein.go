/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package proteins

import (
	"context"
	"encoding/json"
	"log"

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
var Client *firestore.Client
var Size_Bond int
var Ckk string

// service name
const hottopic = "peptideBond_Topic"

// error description
const error_desc = "data stream is not accepted because rules voliation"

// Biomolecules new object or receiver object
func NewPeptideTopic() AbstractBiomolecules { return &Biomolecules{} }

func (biomolecules *Biomolecules) AddPDB(context context.Context, list *binary.Micromolecule_List) *binary.MolecularState {

	doc, write, err := Client.Collection(hottopic).Add(context, map[string]interface{}{
		"chains": list.Peplide,
		"ckk":    Ckk,
	})
	if err != nil {
		log.Println("Data generation error:", err.Error())
		return &binary.MolecularState{State: false, Error: error_desc}
	}
	log.Println("document created successfully", doc, write)
	return &binary.MolecularState{State: true, Error: " "}
}

func (biomolecules *Biomolecules) DisplayPDB(context context.Context, req *binary.Request) *binary.Micromolecule_List {

	var list binary.Micromolecule_List
	query := Client.Collection(hottopic).Where("ckk", "==", Ckk).Documents(context)
	for {
		doc, err := query.Next()
		if err != iterator.Done {
			break
		}
		result := doc.Data()
		marshaldata, err := json.Marshal(result)
		if err != nil {
			log.Printf(" Error marshaling result%v ", err.Error())
			return &binary.Micromolecule_List{}
		}
		err = json.Unmarshal(marshaldata, &list)
		if err != nil {
			log.Printf(" Error unmarshaling list%v ", err.Error())
			return &binary.Micromolecule_List{}
		}
	}
	return &list
}
