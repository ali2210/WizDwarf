/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

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

	// create new document
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

	// if the document is not created then
	if err != nil {
		fmt.Println("collection busy:", err.Error())
		return doc, result, err
	}
	return doc, result, nil
}

func (*FirestoreClient) GetDocumentById(client *firestore.Client, member Visitors) (map[string]interface{}, error) {

	var result_profile map[string]interface{}
	// execute user search query
	query := client.Collection(collection_name).Where("email", "==", member.Email).Where("id", "==", member.Id).Documents(context.Background())
	for {

		//query return documents then go forward
		doc, err := query.Next()

		// if query return empty documednt then terminate
		if err == iterator.Done {
			break
		}

		// read document there may be possible while reading something unexpected
		docsnaps, err := doc.Ref.Get(context.Background())
		if err != nil {
			log.Printf(" Error searching ... %v", err.Error())
			return result_profile, err
		}

		// store document results
		result_profile = docsnaps.Data()
		log.Println("Results:", result_profile)
	}
	return result_profile, nil
}

func (*FirestoreClient) SearchUser(client *firestore.Client, member Visitors) (map[string]interface{}, error) {

	var profile_search map[string]interface{}

	// execute user search query
	query := client.Collection(collection_name).Where("email", "==", member.Email).Where("password", "==", member.Password).Documents(context.Background())
	for {

		//query return documents then go forward
		doc, err := query.Next()
		// if query return empty documednt then terminate
		if err == iterator.Done {
			break
		}

		// read document there may be possible while reading something unexpected
		docsnaps, err := doc.Ref.Get(context.Background())
		if err != nil {
			log.Printf(" Error searching ... %v", err.Error())
			return profile_search, err
		}
		// store document results
		profile_search = docsnaps.Data()
		log.Println("Results:", profile_search)
	}
	return profile_search, nil
}

func (*FirestoreClient) UpdateUserDetails(client *firestore.Client, member Visitors) error {

	// execute user search query
	query := client.Collection(collection_name).Where("email", "==", member.Email).Documents(context.Background())

	for {
		// if query return empty documednt then terminate
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		// write document there may be possible while reading something unexpected
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
