package users

import (
	"context"
	"fmt"
	//"log"
	//firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collection_name string = "ProfileVisitors"
)

type (
	
	DBFirestore interface {
		AddUser(client *firestore.Client, member Visitors) (*firestore.DocumentRef,*firestore.WriteResult, error)
		GetDocumentById(client *firestore.Client, member Visitors) (map[string]interface{}, error)
		SearchUser(client *firestore.Client, member Visitors) (map[string]interface{}, error)
		
	}
	
	FirestoreClient struct{}
)

func NewCloudInstance() DBFirestore {
	return &FirestoreClient{}
}

func (*FirestoreClient) AddUser(client *firestore.Client, member Visitors) (*firestore.DocumentRef,*firestore.WriteResult, error) {
	
	doc, result, err := client.Collection(collection_name).Add(context.Background(), map[string]interface{}{
		"id" : member.Id,
		"name" : member.Name,
		"lastname" : member.LastName,
		"email" : member.Email,
		"password" : member.Password,
		"city" : member.City,
		"zip" : member.Zip,
		"address" : member.Address,
		"apparment" : member.Apparment,
		"country" : member.Country,
		"eve" : member.Eve,
		"phone" : member.PhoneNo,
		"twitter" : member.Twitter,
		// "remember" : member.Remember,
	})

	if err!= nil{
		fmt.Println("collection busy:", err.Error())
		return doc, result, err
	}
	return doc, result, nil
}

func (*FirestoreClient) GetDocumentById(client *firestore.Client, member Visitors) (map[string]interface{}, error) {
	
	var result_profile map[string]interface{}
	query := client.Collection(collection_name).Where("id", "==", member.Id).Documents(context.Background())
	for{
		doc , err := query.Next()
		if err == iterator.Done{break}
		result_profile = doc.Data()
	}
	return result_profile, nil
}

func (*FirestoreClient) SearchUser(client *firestore.Client, member Visitors) (map[string]interface{}, error) {

	var profile_search map[string]interface{}
	query := client.Collection(collection_name).Where("email", "==", member.Email).Where("password", "==", member.Password).Documents(context.Background())
	for{
		doc , err := query.Next()
		if err == iterator.Done{break}
		profile_search = doc.Data() 
				
	}
	return profile_search, nil
}


