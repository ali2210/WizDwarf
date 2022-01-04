/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package bucket

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type DBStore struct {
	Key       string
	Composite string
	Value     string
}

const (
	Ok int = iota << 1
	Err
)
const COLLECTION_NAME string = "ContentAddress"

var client *firestore.Client

type DBStorage interface {
	Store(context.Context, *firestore.Client) int
	Get(context.Context, *firestore.Client) (interface{}, int)
}

func New(key, value, composite string) DBStorage {
	return &DBStore{Key: key, Value: value, Composite: composite}
}

func (db *DBStore) Store(ctx context.Context, client *firestore.Client) int {

	// user shared content during sessions. Each address in onces generated and reference of a document.
	// user will have reference or have public signature that unlocks the content
	doc, result, err := client.Collection(COLLECTION_NAME).Add(ctx, map[string]interface{}{
		"Value":     db.Value,
		"Composite": db.Composite,
		"Key":       db.Key,
	})
	if err != nil {
		log.Printf(" Error creating document %v", err.Error())
		return Err
	}

	log.Println("Doc creating ", doc, result)
	return Ok
}

func (db *DBStore) Get(ctx context.Context, client *firestore.Client) (interface{}, int) {

	// user cannot remember every machine genated address.
	// If the address exists in the database then user will do any operation such as view
	var query_result map[string]interface{}
	query := client.Collection(COLLECTION_NAME).Where("Key", "==", db.Key).Where("Composite", "==", db.Composite).Documents(ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		query_result = doc.Data()
	}

	return query_result[db.Key], Ok
}

func SetClient(c *firestore.Client) {
	client = c
}

func GetClient() *firestore.Client {
	return client
}
