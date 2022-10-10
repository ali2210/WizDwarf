/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package

package cloudmedia

// Libraries
import (
	"context"
	"errors"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"
	error_codes "github.com/ali2210/wizdwarf/errors_codes"
	"github.com/ali2210/wizdwarf/other/cloudmedia/media"
	"google.golang.org/api/iterator"
)

// Interface or dc1 servcies
type DC_1 interface {
	PutData(data *media.IMAGE_METADATA, user_token ...string) error
	GetData(data *media.IMAGE_METADATA, code ...string) (map[string]interface{}, error)
	SearchData(data *media.IMAGE_METADATA, code ...string) (bool, *firestore.DocumentIterator)
}

// Constants
const COLLECTION_NAME string = "_ADescriptor"

// dc_1 struct object
type Dc_1 struct {
	Ctx    context.Context
	Client *firestore.Client
}

// @param context & client
// @return dc_1
func NewDc_1(ctx context.Context, client *firestore.Client) DC_1 {
	return &Dc_1{Ctx: ctx, Client: client}
}

// Store Image data
// @param data & final token string message
// @return error message
// @receiver dc_1
func (d *Dc_1) PutData(data *media.IMAGE_METADATA, user_token ...string) error {

	// Upload content information
	_, _, err := d.Client.Collection(COLLECTION_NAME).Add(d.Ctx, map[string]interface{}{
		"Name":           data.Name,
		"Type":           data.Type,
		"Timeline":       data.Created,
		"Token_Category": data.Tokens,
		"OnAccount":      data.MyProfile,
		"Usercode":       user_token[0],
		"CDR_LINK":       data.Cdr,
		"Date":           data.Timeline,
		"Tag":            data.Tags,
	})
	return err
}

// GetData is an retreival function which return content information
// @param data , final code as string message
// @return map set
// @receiver dc_1
func (d *Dc_1) GetData(data *media.IMAGE_METADATA, code ...string) (map[string]interface{}, error) {

	var empty map[string]interface{}

	// search data for data validation
	ok, query := d.SearchData(data, code...)

	// if validation failed then throw @exception
	if !ok {
		return empty, errors.New("no record found")
	}

	var queryiterator map[string]interface{}

	// read document and store in map set
	for {

		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		queryiterator = doc.Data()
	}

	return queryiterator, nil
}

// @param data , final code as string message
// @return boolean and document_info
// #receiver dc_1

func (d *Dc_1) SearchData(data *media.IMAGE_METADATA, code ...string) (bool, *firestore.DocumentIterator) {

	// query command is created
	query := d.Client.Collection(COLLECTION_NAME).Where("Usercode", "==", code[0]).Where("Token_Category", "==", data.Tokens).Documents(d.Ctx)

	// if the query found nothing then throw exception
	if reflect.DeepEqual(query, &firestore.DocumentIterator{}) {
		log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_OUTPUT)
		return false, &firestore.DocumentIterator{}
	}

	return true, query
}
