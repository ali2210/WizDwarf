package users

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	collection_name string = "ProfileVisitors"
)

type (
	DBFirestore interface {
		AddUser(client *firestore.Client, member Visitors) (*firestore.DocumentRef, *firestore.WriteResult, error)
		GetDocumentById(client *firestore.Client, member Visitors) (map[string]interface{}, error)
		SearchUser(client *firestore.Client, member Visitors) (map[string]interface{}, error)
		UpdateUserDetails(client *firestore.Client, member Visitors) error
	}

	FirestoreClient struct{}
)

func NewCloudInstance() DBFirestore {
	return &FirestoreClient{}
}

func (*FirestoreClient) AddUser(client *firestore.Client, member Visitors) (*firestore.DocumentRef, *firestore.WriteResult, error) {

	doc, result, err := client.Collection(collection_name).Add(context.Background(), map[string]interface{}{
		"id":         member.Id,
		"name":       member.Name,
		"lastname":   member.LastName,
		"email":      member.Email,
		"password":   member.Password,
		"city":       member.City,
		"zip":        member.Zip,
		"address":    member.Address,
		"appartment": member.Appartment,
		"country":    member.Country,
		"eve":        member.Eve,
		"phone":      member.PhoneNo,
		"twitter":    member.Twitter,
	})

	if err != nil {
		fmt.Println("collection busy:", err.Error())
		return doc, result, err
	}
	return doc, result, nil
}

func (*FirestoreClient) GetDocumentById(client *firestore.Client, member Visitors) (map[string]interface{}, error) {

	var result_profile map[string]interface{}
	query := client.Collection(collection_name).Where("email", "==", member.Email).Where("id", "==", member.Id).Documents(context.Background())
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		result_profile = doc.Data()
	}
	return result_profile, nil
}

func (*FirestoreClient) SearchUser(client *firestore.Client, member Visitors) (map[string]interface{}, error) {

	var profile_search map[string]interface{}
	query := client.Collection(collection_name).Where("email", "==", member.Email).Where("password", "==", member.Password).Documents(context.Background())
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		profile_search = doc.Data()

	}
	return profile_search, nil
}

func (*FirestoreClient) UpdateUserDetails(client *firestore.Client, member Visitors) error {

	query := client.Collection(collection_name).Where("email", "==", member.Email).Documents(context.Background())
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		result, err := doc.Ref.Set(context.Background(), map[string]interface{}{
			"id":         member.Id,
			"name":       member.Name,
			"email":      member.Email,
			"lastname":   member.LastName,
			"password":   member.Password,
			"address":    member.Address,
			"appartment": member.Appartment,
			"country":    member.Country,
			"city":       member.City,
			"zip":        member.Zip,
			"phone":      member.PhoneNo,
			"twitter":    member.Twitter,
			"eve":        member.Eve,
		})
		if err != nil {
			log.Println("Error updating member: ", err.Error())
			return err
		}
		log.Println("Result:", result)
	}
	return nil
}
