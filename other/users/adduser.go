/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package users

import (
	"context"
	"errors"
	"log"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
	user "github.com/ali2210/wizdwarf/other/users/register"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	collection_name string = "ProfileVisitors"
)

type Doc_Response struct {
	Profile    user.New_User
	Update     user.Updated_User
	Doc_status bool
}

type (
	DBFirestore interface {
		AddUser(client *firestore.Client, member user.New_User) (*firestore.DocumentRef, *firestore.WriteResult, error)
		GetDocumentById(client *firestore.Client, member user.New_User) (map[string]interface{}, error)
		SearchUser(client *firestore.Client, member user.New_User) (*Doc_Response, error)
		UpdateUserDetails(client *firestore.Client, member user.Updated_User) (*firestore.WriteResult, error)
		GetUpdatedRecord(client *firestore.Client, member user.Updated_User) (map[string]interface{}, error)
	}

	FirestoreClient struct{}
)

func NewCloudInstance() DBFirestore {
	return &FirestoreClient{}
}

func (*FirestoreClient) AddUser(client *firestore.Client, member user.New_User) (*firestore.DocumentRef, *firestore.WriteResult, error) {

	// create new document
	doc, result, err := client.Collection(collection_name).Add(context.Background(), map[string]interface{}{
		"Name":            member.Name,
		"Email":           member.Email,
		"Password":        member.Password,
		"Lastname":        member.Lastname,
		"Address":         member.Address,
		"Phone":           member.Phone,
		"Zip":             member.Zip,
		"City":            member.City,
		"State":           member.State,
		"Gender":          member.Gender,
		"ID":              member.ID,
		"Friends":         member.Friends,
		"Inspire":         member.Inspire,
		"Lead":            member.Lead,
		"SocialEvolution": member.SocialEvolution,
	})

	// if the document is not created then
	if err != nil {
		return &firestore.DocumentRef{}, result, err
	}
	return doc, result, nil
}

func (*FirestoreClient) GetDocumentById(client *firestore.Client, member user.New_User) (map[string]interface{}, error) {

	var result_profile map[string]interface{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// execute user search query
	query := client.Collection(collection_name).Where("Email", "==", member.Email).Where("ID", "==", member.ID).Documents(ctx)
	for {

		//query return documents then go forward
		doc, err := query.Next()

		// if query return empty document then terminate
		if err == iterator.Done {
			break
		}

		// read document there may be possible while reading something unexpected
		docsnaps, err := doc.Ref.Get(context.Background())
		if err != nil {
			return result_profile, err
		}

		// store document results
		result_profile = docsnaps.Data()

	}
	return result_profile, nil
}

// Search User will return dpcuments if exist... Documents either for new profile and updated profile
// Flag Value                               Description
//  true 									  	Document exist
//   false										No Document exist

func (fire *FirestoreClient) SearchUser(client *firestore.Client, member user.New_User) (*Doc_Response, error) {

	var profile_search map[string]interface{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// execute user search query
	query := client.Collection(collection_name).Where("Email", "==", member.Email).Where("Password", "==", member.Password).Documents(ctx)

	if reflect.DeepEqual(query, &firestore.DocumentIterator{}) {
		log.Fatalln("Doc iterator empty")
		return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, errors.New("no user found")
	}

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
			log.Fatalln("Error getting doc  :", err)
			return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, err
		}

		if reflect.DeepEqual(docsnaps, map[string]interface{}{}) {
			log.Fatalln("Error empty document")
			return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, err
		}

		// store document results
		profile_search = docsnaps.Data()
	}

	snaps, err := structpb.NewValue(profile_search)
	if err != nil {
		log.Fatalln("Error data formatting")
		return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, err
	}

	data, err := snaps.MarshalJSON()
	if err != nil {
		log.Fatalln("Error data marshalling")
		return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, err
	}

	if err = snaps.UnmarshalJSON(data); err != nil {
		log.Fatalln("Error data new unmarshaling: ")
		return &Doc_Response{Profile: user.New_User{}, Update: user.Updated_User{}, Doc_status: false}, err
	}

	Iterate := reflect.ValueOf(profile_search).MapRange()
	updated := user.Updated_User{}

	for Iterate.Next() {

		if reflect.DeepEqual(Iterate.Key().String(), "Address") {
			updated.Address = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Citation") {
			updated.Citation = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Achievements") {
			updated.Achievements = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Password") {
			updated.Password = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Company") {
			updated.Company = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Lastname") {
			updated.Lastname = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "ID") {
			updated.ID = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "State") {
			updated.State = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Inspire") {
			updated.Inspire = Iterate.Value().Interface().(int64)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "City") {
			updated.City = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Published") {
			updated.Published = Iterate.Value().Interface().(int64)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Date") {
			updated.Date = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Circulum") {
			updated.Circulum = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "University") {
			updated.University = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Zip") {
			updated.Zip = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "SocialEvolution") {
			updated.SocialEvolution = Iterate.Value().Interface().(int64)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Name") {
			updated.Name = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Phone") {
			updated.Phone = Iterate.Value().Interface().(string)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Lead") {
			updated.Lead = Iterate.Value().Interface().(int64)
		}

		if reflect.DeepEqual(Iterate.Key().String(), "Email") {
			updated.Email = Iterate.Value().Interface().(string)
		}

	}

	if updated.SocialEvolution == 2 {
		_new := user.New_User{}
		_new.Name = updated.Name
		_new.Email = updated.Email
		_new.Password = updated.Password
		_new.Address = updated.Address
		_new.City = updated.City
		_new.State = updated.State
		_new.ID = updated.ID
		_new.Lastname = updated.Lastname
		_new.Inspire = updated.Inspire
		_new.Friends = updated.Friends
		_new.Lead = updated.Lead
		_new.Zip = updated.Zip
		_new.SocialEvolution = updated.SocialEvolution
		return &Doc_Response{Profile: _new, Update: user.Updated_User{}, Doc_status: true}, nil
	}

	return &Doc_Response{Profile: user.New_User{}, Update: updated, Doc_status: true}, nil

}

func (*FirestoreClient) UpdateUserDetails(client *firestore.Client, member user.Updated_User) (*firestore.WriteResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// execute user search query
	query := client.Collection(collection_name).Where("Email", "==", member.Email).Where("ID", "==", member.ID).Documents(ctx)
	var result *firestore.WriteResult
	for {
		// if query return empty documednt then terminate
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		// write document there may be possible while reading something unexpected
		result, err = doc.Ref.Set(context.Background(), map[string]interface{}{
			"ID":              member.ID,
			"Name":            member.Name,
			"Email":           member.Email,
			"Phone":           member.Phone,
			"Lastname":        member.Lastname,
			"Address":         member.Address,
			"State":           member.State,
			"City":            member.City,
			"Zip":             member.Zip,
			"Gender":          member.Gender,
			"Status":          member.Status,
			"Suffix":          member.Suffix,
			"University":      member.University,
			"Circulum":        member.Circulum,
			"Company":         member.Company,
			"Date":            member.Date,
			"Published":       member.Published,
			"Achievements":    member.Achievements,
			"Citation":        member.Citation,
			"Friends":         member.Friends,
			"Inspire":         member.Inspire,
			"Lead":            member.Lead,
			"SocialEvolution": member.SocialEvolution,
			"Password":        member.Password,
		})
		if err != nil {
			return &firestore.WriteResult{}, err
		}

		log.Println("Document Updated:", *result)
	}

	return result, nil
}

func (*FirestoreClient) GetUpdatedRecord(client *firestore.Client, member user.Updated_User) (map[string]interface{}, error) {

	var result_profile map[string]interface{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// execute user search query
	query := client.Collection(collection_name).Where("Email", "==", member.Email).Where("ID", "==", member.ID).Documents(ctx)
	for {

		//query return documents then go forward
		doc, err := query.Next()

		// if query return empty document then terminate
		if err == iterator.Done {
			break
		}

		// read document there may be possible while reading something unexpected
		docsnaps, err := doc.Ref.Get(ctx)
		if err != nil {
			return result_profile, err
		}

		// store document results
		result_profile = docsnaps.Data()

	}
	return result_profile, nil
}
