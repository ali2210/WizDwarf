/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package fireclient

import (
	"context"
	"image"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"
	error_codes "github.com/ali2210/wizdwarf/errors_codes"
	"google.golang.org/api/iterator"
)

// Crdedntials structure
type DBStore struct {
	// Key       string
	// Composite string
	// Value     string

	Ctx context.Context
	Ref *firestore.Client
}

// Enum selectors
const (
	Ok int = iota << 1
	Err
)

// Collection name
const COLLECTION_NAME string = "ContentAddress"

// client object
var client *firestore.Client

// Credentials services
type DBStorage interface {
	Store(key, value string, composite ...string) int
	Get(key string, composite ...string) (map[string]interface{}, int)
	GetAll() ([]map[string]interface{}, error)
}

// Instanitation of Credentials Object
func New(ctx context.Context, client *firestore.Client) DBStorage {
	return &DBStore{Ctx: ctx, Ref: client}
}

// Store credentials object information
func (db *DBStore) Store(key, value string, composite ...string) int {

	// user shared content during sessions. Each address in onces generated and reference of a document.
	// user will have reference or have public signature that unlocks the content
	_, _, err := db.Ref.Collection(COLLECTION_NAME).Add(db.Ctx, map[string]interface{}{
		"Value":     value,
		"Composite": composite[0],
		"Key":       key,
	})
	if err != nil {
		log.Printf(" Error creating document %v", err.Error())
		return Err
	}

	return Ok
}

func (db *DBStore) GetAll() ([]map[string]interface{}, error) {

	exp := make([]map[string]interface{}, 10)

	query := db.Ref.CollectionGroup(COLLECTION_NAME).Limit(10).Documents(db.Ctx)

	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		exp = append(exp, doc.Data())
	}

	return exp, nil

	// for {

	// 	doc, err := query.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}

	// 	if err != nil {
	// 		return []map[string]interface{}{}, err
	// 	}
	// 	exp = append(exp, doc.Data())
	// }

	// if reflect.DeepEqual(exp, map[string]interface{}{}) {
	// 	log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_OUTPUT)
	// 	return
	// }

	// return exp, nil
}

// Credentials Information Retreive
func (db *DBStore) Get(key string, composite ...string) (map[string]interface{}, int) {

	// user cannot remember every machine genated address.
	// If the address exists in the database then user will do any operation such as view
	var query_result map[string]interface{}
	query := db.Ref.Collection(COLLECTION_NAME).Where("Key", "==", key).Where("Composite", "==", composite[0]).Documents(db.Ctx)
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		query_result = doc.Data()
	}

	if reflect.DeepEqual(query_result, map[string]interface{}{}) {
		log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_OUTPUT)
		return map[string]interface{}{}, Err
	}

	return query_result, Ok
}

// Client reference
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
