/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package

package secrets

import (
	"context"
	"errors"
	"reflect"

	"cloud.google.com/go/firestore"
	error_codes "github.com/ali2210/wizdwarf/errors_codes"
	"google.golang.org/api/iterator"
)

type Secrets interface {
	SearchContent(collname ...string) (bool, *firestore.DocumentIterator)
	GetAllDocuments(collname ...string) ([]map[string]interface{}, error)
}

type Vault struct {
	Ctx       context.Context
	Firestore *firestore.Client
}

func NewSecretsInstance(ctx context.Context, client *firestore.Client) Secrets {
	return &Vault{Ctx: ctx, Firestore: client}
}

func (s *Vault) SearchContent(collname ...string) (bool, *firestore.DocumentIterator) {

	query := s.Firestore.Collection(collname[0]).Documents(s.Ctx)
	if reflect.DeepEqual(query, &firestore.DocumentIterator{}) {
		return false, &firestore.DocumentIterator{}
	}

	return true, query
}

func (s *Vault) GetAllDocuments(colname ...string) ([]map[string]interface{}, error) {

	list := make([]map[string]interface{}, 5)
	ok, query := s.SearchContent(colname[0])
	if !ok {
		return []map[string]interface{}{}, errors.New(error_codes.Operation_ERROR_CODE_EMPTY_OUTPUT.String())
	}

	for {

		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		list = append(list, doc.Data())
	}

	return list, nil
}
