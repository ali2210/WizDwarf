package users

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/structs/users/model"
)

const (
	collection string = "ProfileVistors"
)

type ProfileinJSON struct {
	ID            string `json:"Id"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	PhoneNo       string `json:"PhoneNo"`
	HouseAddress1 string `json:"HouseAddress1"`
	HouseAddress2 string `json:"HouseAddress2"`
	Country       string `json:"Country"`
	Zip           string `json:"Zip"`
	City          string `json:"City"`
	Eve           bool   `json:"Eve"`
	Email         string `json:"Email"`
	Twitter       string `json:"Twitter"`
}

type DBFirestore interface {
	SaveData(visitor *model.Vistors, app *firebase.App) (*model.Vistors, error)
	ToFindByGroupSet(id, email string, app *firebase.App) (*model.Vistors, error)
	FindAllData(app *firebase.App, email, password string) (*model.Vistors, error)
	UpdateProfiles(clientID *firebase.App, profile *model.UpdateProfile) (*model.UpdateProfile, error)
	GetProfile(clientID *firebase.App, ID, email string) (*ProfileinJSON, error)
}

type cloudData struct{}

func NewCloudInstance() DBFirestore {
	return &cloudData{}
}

func (*cloudData) SaveData(visitor *model.Vistors, app *firebase.App) (*model.Vistors, error) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"ID":       visitor.Id,
		"Name":     visitor.Name,
		"Email":    visitor.Email,
		"Password": visitor.Password,
		"FName":    visitor.FName,
		"Eve":      visitor.Eve,
		"Address":  visitor.Address,
		"LAddress": visitor.LAddress,
		"City":     visitor.City,
		"Zip":      visitor.Zip,
		"Country":  visitor.Country,
	})
	if err != nil {
		// log.Fatal("Failed to retrive Vistor Record:", err)
		return nil, err
	}
	return visitor, nil

}

func (*cloudData) FindAllData(app *firebase.App, email, password string) (*model.Vistors, error) {
	ctx := context.Background()
	var visit model.Vistors
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("Client Instance Failed to start", err)
		return &visit, err
	}

	defer client.Close()

	iterator := client.Collection(collection).Where("Email", "==", email).Where("Password", "==", password).Documents(ctx)
	defer iterator.Stop()
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatal("Iterator Failed on Vistor: ", err)
			return &visit, err
		}

		visit = model.Vistors{
			Id:       doc.Data()["ID"].(string),
			Name:     doc.Data()["Name"].(string),
			Email:    doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
			FName:    doc.Data()["FName"].(string),
			City:     doc.Data()["City"].(string),
			Country:  doc.Data()["Country"].(string),
			Zip:      doc.Data()["Zip"].(string),
			Address:  doc.Data()["Address"].(string),
			LAddress: doc.Data()["LAddress"].(string),
			Eve:      doc.Data()["Eve"].(bool),
		}
		break
	}
	return &visit, err

}

func (*cloudData) ToFindByGroupSet(id, email string, app *firebase.App) (*model.Vistors, error) {

	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()

	var visits model.Vistors
	it := client.Collection(collection).Where("Email", "==", email).Where("Id", "==", id).Documents(ctx)

	defer it.Stop()
	for {
		doc, err := it.Next()
		if err != nil {
			log.Fatal("Iterator Failed on Vistor: ", err)
			return &visits, err
		}
		if doc.Data() == nil {
			log.Fatal("Iterator Failed on Vistor: ", err, doc.Data())
			continue
		}
		visits = model.Vistors{
			Id:      doc.Data()["ID"].(string),
			Email:   doc.Data()["Email"].(string),
			Country: doc.Data()["Country"].(string),
			Zip:     doc.Data()["Zip"].(string),
			Eve:     doc.Data()["Eve"].(bool),
		}
		break
	}
	return &visits, err
}

func (*cloudData) UpdateProfiles(clientID *firebase.App, profile *model.UpdateProfile) (*model.UpdateProfile, error) {
	ctx := context.Background()
	client, err := clientID.Firestore(ctx)
	if err != nil {
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()

	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"ID":            profile.Id,
		"FirstName":     profile.FirstName,
		"Email":         profile.Email,
		"Eve":           profile.Male,
		"HouseAddress1": profile.HouseAddress,
		"HouseAddress2": profile.SubAddress,
		"Zip":           profile.Zip,
		"Country":       profile.Country,
		"LastName":      profile.LastName,
		"PhoneNo":       profile.Phone,
		"Twitter":       profile.Twitter,
		"City":          profile.City,
	})
	if err != nil {
		// log.Fatal("Failed to retrive Vistor Record:", err)
		return nil, err
	}
	return profile, nil
}

func (*cloudData) GetProfile(clientID *firebase.App, ID, email string) (*ProfileinJSON, error) {
	ctx := context.Background()
	var visits ProfileinJSON
	client, err := clientID.Firestore(ctx)
	if err != nil {
		log.Fatal("Client Instance Failed to start", err)
		return &visits, err
	}

	defer client.Close()
	iterator := client.Collection(collection).Where("Id", "==", ID).Where("Email", "==", email).Documents(ctx)

	defer iterator.Stop()
	for {
		doc, err := iterator.Next()
		if err != nil {
			return &visits, err
		}
		visits = ProfileinJSON{
			ID:            doc.Data()["ID"].(string),
			FirstName:     doc.Data()["FirstName"].(string),
			LastName:      doc.Data()["LastName"].(string),
			PhoneNo:       doc.Data()["PhoneNo"].(string),
			HouseAddress1: doc.Data()["HouseAddress1"].(string),
			HouseAddress2: doc.Data()["Address"].(string),
			Country:       doc.Data()["Country"].(string),
			Zip:           doc.Data()["Zip"].(string),
			City:          doc.Data()["City"].(string),
			Eve:           doc.Data()["Eve"].(bool),
			Email:         doc.Data()["Email"].(string),
			Twitter:       doc.Data()["Twitter"].(string),
		}
		break
	}
	return &visits, err
}
