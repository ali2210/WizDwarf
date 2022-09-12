package cloudmedia

import (
	"context"
	"errors"
	"reflect"

	"cloud.google.com/go/firestore"
	"github.com/ali2210/wizdwarf/other/cloudmedia/media"
	"google.golang.org/api/iterator"
)

// Dc_1 means data center -1; our project have lots of data centers;
type DC_1 interface {
	PutData(data *media.IMAGE_METADATA, user_token ...string) error
	GetData(data *media.IMAGE_METADATA, code ...string) (map[string]interface{}, error)
	SearchData(data *media.IMAGE_METADATA, code ...string) (bool, *firestore.DocumentIterator)
}

const COLLECTION_NAME string = "_ADescriptor"

type Dc_1 struct {
	Ctx    context.Context
	Client *firestore.Client
}

func NewDc_1(ctx context.Context, client *firestore.Client) DC_1 {
	return &Dc_1{Ctx: ctx, Client: client}
}

func (d *Dc_1) PutData(data *media.IMAGE_METADATA, user_token ...string) error {

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
func (d *Dc_1) GetData(data *media.IMAGE_METADATA, code ...string) (map[string]interface{}, error) {

	var empty map[string]interface{}
	ok, query := d.SearchData(data, code...)
	if !ok {
		return empty, errors.New("no record found")
	}

	var queryiterator map[string]interface{}

	for {

		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		queryiterator = doc.Data()
	}

	return queryiterator, nil
}

func (d *Dc_1) SearchData(data *media.IMAGE_METADATA, code ...string) (bool, *firestore.DocumentIterator) {

	query := d.Client.Collection(COLLECTION_NAME).Where("Usercode", "==", code[0]).Where("Token_Category", "==", data.Tokens).Documents(d.Ctx)

	if reflect.DeepEqual(query, &firestore.DocumentIterator{}) {
		return false, &firestore.DocumentIterator{}
	}

	return true, query
}
