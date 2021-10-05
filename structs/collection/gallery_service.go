package collection

import (
	"context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"log"
	"encoding/json"
)

var(
	Firestore_Picture_Client *firestore.Client
	Size int64 = 1
)

const (
	collection_name = "pictures"
)

type Gallery_Stream_Server struct{}

func (s *Gallery_Stream_Server) NewPictures(ctx context.Context, images *Pictures)(*Collection){
	doc ,result, err := Firestore_Picture_Client.Collection(collection_name).Add(ctx, map[string]interface{}{
		"pic_id" : images.PicId,
		"pic_src": images.PicSrc,
		"pic_date" : images.PicDate,
		"userId" : images.UserAgentId,
		"pic_time" : images.PicTime,
		"pic_tags" : images.PicTags,
		// "picRef" : images.Pic_Reference,
	})
	if err != nil {
		log.Fatal(" images insertion error: ", err.Error())
		return &Collection{}
	}

	log.Println("Doc:", doc, "Result:", result)
	
	var pic_collection map[string]interface{}
	schema := Firestore_Picture_Client.Collection(collection_name).Where("userId", "==", images.UserAgentId).Where("pic_id", "==", images.PicId).Documents(ctx)
	for{
		doc, err := schema.Next()
		if err == iterator.Done {
			break
		}
		pic_collection = doc.Data()
	}

	convert , err := json.Marshal(pic_collection); if err != nil {
		log.Printf("error marshaling %v", err.Error())
		return &Collection{}
	}

	var pic Pictures 
	err = json.Unmarshal(convert, &pic); if err != nil {
		log.Printf("error unmarshaling %v", err.Error())
		return &Collection{}
	}
	
	collection := make([]*Pictures, Size-1)
	collection = append(collection, &pic)
	mycollection := &Collection{}
	mycollection.Gallery = collection
	
	//log.Println("Collection :", mycollection.Gallery)
	
	return mycollection
}

// func (s *Gallery_Stream_Server) SearchPictures(ctx context.Context, images *Pictures)(*Collection){
// 	var pic_collection map[string]interface{}
// 	schema := Firestore_Picture_Client.Collection(collection_name).Where("userId", "==", images.UserAgentId).Where("pic_id", "==", images.PicId).Documents(ctx)
// 	for{
// 		doc, err := schema.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		pic_collection = doc.Data()
// 	}

// 	convert , err := json.Marshal(pic_collection); if err != nil {
// 		log.Printf("error marshaling %v", err.Error())
// 		return &Collection{}
// 	}

// 	var pic Pictures 
// 	err = json.Unmarshal(convert, &pic); if err != nil {
// 		log.Printf("error unmarshaling %v", err.Error())
// 		return &Collection{}
// 	}
	
// 	collection := make([]*Pictures, Size)
// 	collection = append(collection, &pic)
// 	mycollection :=  &Collection{}
// 	mycollection.Gallery = collection
// 	return mycollection

// }

// func (s *Gallery_Stream_Server) DisplayPictures(ctx context.Context, images *Pictures)(*Collection){
// 	var pic_collection map[string]interface{}
// 	schema := Firestore_Picture_Client.Collection(collection_name).Where("userId", "==", images.UserAgentId).Where("pic_id", "==", images.PicId).Documents(ctx)
// 	for{
// 		doc, err := schema.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		pic_collection = doc.Data()
// 	}

// 	convert , err := json.Marshal(pic_collection); if err != nil {
// 		log.Printf("error marshaling %v", err.Error())
// 		return &Collection{}
// 	}

// 	var pic Pictures 
// 	err = json.Unmarshal(convert, &pic); if err != nil {
// 		log.Printf("error unmarshaling %v", err.Error())
// 		return &Collection{}
// 	}
	
// 	collection := make([]*Pictures, Size)
// 	collection = append(collection, &pic)
// 	mycollection := &Collection{}
// 	mycollection.Gallery = collection
// 	return mycollection

// }