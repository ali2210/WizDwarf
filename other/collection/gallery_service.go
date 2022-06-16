package collection

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var (
	Firestore_Picture_Client *firestore.Client
	Size                     int64 = 1
)

const (
	collection_name = "pictures"
)

type Gallery_Stream_Server struct{}

type Gallery_Declaration interface {
	NewPictures(ctx context.Context, images *Pictures) *Collection
	SearchPictures(ctx context.Context, compress *Compressed) *ContentRoute
	GetMetadata(ctx context.Context, compress *Compressed) *Metadata_File
}

// Store Images allow you to store images data in database
func (s *Gallery_Stream_Server) NewPictures(ctx context.Context, images *Pictures) *Collection {

	// store images data
	_, _, err := Firestore_Picture_Client.Collection(collection_name).Add(ctx, map[string]interface{}{
		"pic_id":   images.PicId,
		"pic_src":  images.PicSrc,
		"pic_date": images.PicDate,
		"userId":   images.UserAgentId,
		"pic_time": images.PicTime,
		"pic_tags": images.PicTags,
		"picRef":   images.CDR,
	})
	if err != nil {
		log.Fatal(" images insertion error: ", err.Error())
		return &Collection{}
	}

	// add image in gallery
	var pic_collection map[string]interface{}
	schema := Firestore_Picture_Client.Collection(collection_name).Where("userId", "==", images.UserAgentId).Where("pic_id", "==", images.PicId).Documents(ctx)
	for {
		doc, err := schema.Next()
		if err == iterator.Done {
			break
		}
		pic_collection = doc.Data()
	}

	convert, err := json.Marshal(pic_collection)
	if err != nil {
		log.Printf("error marshaling %v", err.Error())
		return &Collection{}
	}

	var pic Pictures
	err = json.Unmarshal(convert, &pic)
	if err != nil {
		log.Printf("error unmarshaling %v", err.Error())
		return &Collection{}
	}

	collection := make([]*Pictures, Size-1)
	collection = append(collection, &pic)
	mycollection := &Collection{}
	mycollection.Gallery = collection

	return mycollection
}

// search images in the collection

func (c *Gallery_Stream_Server) SearchPictures(ctx context.Context, compress *Compressed) *ContentRoute {

	// if no data will provided by the client then throw error
	if reflect.DeepEqual(compress, &Compressed{}) {
		return &ContentRoute{UserAgentId: "", IsP2PAddress: false}
	}

	if reflect.DeepEqual(compress.GetPicId(), " ") {
		return &ContentRoute{UserAgentId: "", IsP2PAddress: false}
	}

	if reflect.DeepEqual(compress.GetUserAgentId(), " ") {
		return &ContentRoute{UserAgentId: "", IsP2PAddress: false}
	}

	// otherwise search data query
	var result map[string]interface{}
	query := Firestore_Picture_Client.Collection(collection_name).Where("pic_id", "==", compress.PicId).Where("userId", "==", compress.UserAgentId).Documents(ctx)

	for {

		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		result = doc.Data()
	}

	// there may be possible that nothing will be returned by database

	if reflect.DeepEqual(result, map[string]interface{}{}) {
		return &ContentRoute{UserAgentId: "", IsP2PAddress: false}
	}

	// return query result
	return &ContentRoute{UserAgentId: compress.UserAgentId, IsP2PAddress: true}
}

// Retreive image meta data that store in the data store
func (c *Gallery_Stream_Server) GetMetadata(ctx context.Context, compress *Compressed) *Metadata_File {

	// result attributes
	var result map[string]interface{}

	if reflect.DeepEqual(compress.GetUserAgentId(), "") {
		return &Metadata_File{PicSrc: "", CDR: "", UserAgentId: ""}
	}

	if reflect.DeepEqual(compress.GetPicId(), "") {
		return &Metadata_File{PicSrc: "", CDR: "", UserAgentId: ""}
	}

	query := Firestore_Picture_Client.Collection(collection_name).Where("pic_id", "==", compress.PicId).Where("userId", "==", compress.UserAgentId).Documents(ctx)

	for {

		doc, err := query.Next()
		if err == iterator.Done {
			break
		}

		result = doc.Data()
	}

	if reflect.DeepEqual(result, map[string]interface{}{}) {
		return &Metadata_File{PicSrc: "", UserAgentId: "", CDR: ""}
	}

	return &Metadata_File{PicSrc: reflect.ValueOf(result["pic_src"]).String(),
		UserAgentId: compress.UserAgentId,
		CDR:         reflect.ValueOf(result["user_agent_id"]).String()}

}
