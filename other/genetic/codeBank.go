/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package genetics

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ali2210/wizdwarf/other/genetic/binary"
	"google.golang.org/api/iterator"
)

type Lifecode_Plugin interface {
	AddCode(context context.Context, code *binary.Lifecode) *binary.Lifecode_Status
	DisplayCode(context context.Context, code *binary.Request) *binary.Lifecode
}

type LifeCode struct{}

var Client *firestore.Client
var Pkk string

const genetictopice = "genome"

func New() Lifecode_Plugin { return &LifeCode{} }
func (l *LifeCode) AddCode(context context.Context, code *binary.Lifecode) *binary.Lifecode_Status {

	_, _, err := Client.Collection(genetictopice).Add(context, map[string]interface{}{
		"genes": code.Genes,
		"pk":    code.Pkk,
	})
	if err != nil {
		log.Printf("Error creating genetictopice%v", err.Error())
		return &binary.Lifecode_Status{Status: false, Error: err.Error(), ErrorCode: binary.Errors_Error}
	}
	// log.Println("Document created @", doc, "Result:", result)
	return &binary.Lifecode_Status{Status: true, Error: "", ErrorCode: binary.Errors_OK}
}

func (l *LifeCode) DisplayCode(context context.Context, code *binary.Request) *binary.Lifecode {

	var codes *binary.Lifecode
	query := Client.Collection(genetictopice).Where("pk", "==", Pkk).Documents(context)
	for {
		doc, err := query.Next()
		if err != iterator.Done {
			break
		}
		data, err := json.Marshal(doc.Data)
		if err != nil {
			log.Printf("Error marshalling document%v", err.Error())
			return &binary.Lifecode{}
		}
		err = json.Unmarshal(data, &codes)
		if err != nil {
			log.Printf("Error unmarshaling document%v", err.Error())
			return &binary.Lifecode{}
		}
	}
	return codes
}
