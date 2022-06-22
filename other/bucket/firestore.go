/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package bucket

import (
	"context"
	"image"
	"log"
	"reflect"

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
	_, _, err := client.Collection(COLLECTION_NAME).Add(ctx, map[string]interface{}{
		"Value":     db.Value,
		"Composite": db.Composite,
		"Key":       db.Key,
	})
	if err != nil {
		log.Printf(" Error creating document %v", err.Error())
		return Err
	}
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

type __Firestore struct {

	// Client allow to connect with firestore server
	Client *firestore.Client

	// With Ctx how much require to process information. Afterthat connection will disclose the connection
	Ctx context.Context

	// UserAgent specified which user agent want to use data; what is the access level of the user
	User_agent_id string

	// Hash_Color specified user avatars is not replicated in the database
	Hash_Color string

	// Image decode allow to get image binary level representation
	Image_Deoder image.Image
}

type __Firestore_Datasource interface {
	GenerateFingersprints(img image.Image) int
	AnalyzeFingersprints(img image.Image) (map[string]interface{}, int)
}

func NewClient(client *firestore.Client, ctx context.Context, agent, hcolor string) __Firestore_Datasource {
	return &__Firestore{Client: client, Ctx: ctx, User_agent_id: agent, Hash_Color: hcolor}
}

const schema_name string = "Avatarsprints"

func (h *__Firestore) GenerateFingersprints(img image.Image) int {

	h.Image_Deoder = img

	// add information in the database
	_, _, err := h.Client.Collection(schema_name).Add(h.Ctx, map[string]interface{}{
		"OwnerID":         h.User_agent_id,
		"Image_Signature": h.Hash_Color[64:128],
		"Decoder":         h.Image_Deoder,
	})

	// there may be case where infometion is not store due to any error then discard the request and throw exception
	if err != nil {
		log.Fatalln(" Error generate signature ", err)
		return Err
	}

	// return valid token
	return Ok
}

func (h *__Firestore) AnalyzeFingersprints(img image.Image) (map[string]interface{}, int) {

	var result map[string]interface{}
	h.Image_Deoder = img

	// query execute user commands on firestore
	query := h.Client.Collection(schema_name).Where("OwnerID", "==", h.User_agent_id).Documents(h.Ctx)
	for {

		// query get documentsnap ; if there exist; otherwise throw exception
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		// read the information
		result = doc.Data()
	}

	// get information should not be empty
	isvalid := func(query map[string]interface{}) bool {

		var nulify map[string]interface{}
		return !reflect.DeepEqual(query, nulify)
	}

	var nulify map[string]interface{}

	// analysis the document snapshot ; if the document is empty then throw exception
	if !isvalid(result) {
		return nulify, Err
	}

	// return result
	return result, Ok

}
