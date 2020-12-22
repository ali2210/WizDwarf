package users

import (
	firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/structs/users/model"
	"context"
	"log"
)

const (
	collection string = "ProfileVistors"
)



type ProfileinJSON struct{
	
	Id string `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	PhoneNo string `json:"PhoneNo"`
	HouseAddress1 string `json:"HouseAddress1"`
	HouseAddress2 string `json:"HouseAddress2"`
	Country string `json:"Country"`
	Zip string `json:"Zip"`
	Eve bool `json:"Eve"`
	Email string `json:"Email"`
}




type DBFirestore interface{
	
	SaveData(visitor *model.Vistors, app *firebase.App)(*model.Vistors, error)
	FindDataByID(id string , app *firebase.App)(*model.Vistors, error)
	FindAllData(app *firebase.App, email , password string)(*model.Vistors, error)
	UpdateProfiles(clientId *firebase.App, profile *model.UpdateProfile)(*model.UpdateProfile, error)
	GetProfile(clientId *firebase.App, Id string)(*model.UpdateProfile, error)
}

type cloud_data struct{}


func NewCloudInstance()  DBFirestore{
	return &cloud_data{}
}

func (*cloud_data) SaveData(visitor *model.Vistors, app *firebase.App)(*model.Vistors, error){
	ctx := context.Background()
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()
	
	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"Id" :	visitor.Id,
		"Name" : visitor.Name,
		"Email" : visitor.Email,
		"Password": visitor.Password,
		"FName": visitor.FName,
		"Eve": visitor.Eve,
		"Address":visitor.Address,
		"LAddress":visitor.LAddress,
		"City" : visitor.City,
		"Zip": visitor.Zip,
		"Country": visitor.Country,
		}); if err != nil{
			log.Fatal("Failed to retrive Vistor Record:", err)
			return nil, err
		}
		return visitor, nil

}

func (*cloud_data) FindAllData(app *firebase.App, email, password string)(*model.Vistors,error){
	ctx := context.Background()
	var visit model.Vistors
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return &visit, err
	}

	defer client.Close()

	iterator := client.Collection(collection).Where("Email", "==", email).Where("Password", "==", password).Documents(ctx)
	// defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			log.Fatal("Iterator Failed on Vistor: ", err)
			return &visit, err
		}

			visit = model.Vistors {
				Id : doc.Data()["Id"].(string),
			Name : doc.Data()["Name"].(string),
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
			FName: doc.Data()["FName"].(string),
			City : doc.Data()["City"].(string),
			Country: doc.Data()["Country"].(string),
			Zip: doc.Data()["Zip"].(string),
			Address: doc.Data()["Address"].(string),
			LAddress: doc.Data()["LAddress"].(string),
			Eve: doc.Data()["Eve"].(bool),
			} 
			break
	}
	return &visit, err

}

func (*cloud_data) FindDataByID(id string, app *firebase.App)(*model.Vistors, error){
	ctx := context.Background()
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	
	log.Println("AccountID:", id)

	defer client.Close()
	var visits model.Vistors
	iterator := client.Collection(collection).Where("Id", "==", id).Documents(ctx)
	
	//defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			return nil, err
		}
		visits = model.Vistors {
			Id : doc.Data()["Id"].(string),
			Name : doc.Data()["Name"].(string),
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
			FName: doc.Data()["FName"].(string),
			City : doc.Data()["City"].(string),
			Country: doc.Data()["Country"].(string),
			Zip: doc.Data()["Zip"].(string),
			Address: doc.Data()["Address"].(string),
			LAddress: doc.Data()["LAddress"].(string),
			Eve: doc.Data()["Eve"].(bool),
		}
		break
	}
	return &visits, nil
}

func (*cloud_data) UpdateProfiles(clientId *firebase.App, profile *model.UpdateProfile)(*model.UpdateProfile, error){
	ctx := context.Background()
	client , err := clientId.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()
	
	_, _, err = client.Collection(collection).Add(ctx, map[string]interface{}{
		"Id" :	profile.Id,
		"FirstName" : profile.FirstName,
		"Email" : profile.Email,
		"Eve": profile.Male,
		"HouseAddress1":profile.HouseAddress,
		"HouseAddress2":profile.SubAddress,
		"Zip": profile.Zip,
		"Country": profile.Country,
		"LastName": profile.LastName,
		"PhoneNo" : profile.Phone,
		}); if err != nil{
			log.Fatal("Failed to retrive Vistor Record:", err)
			return nil, err
		}
		return profile, nil
}

func (*cloud_data) GetProfile(clientId *firebase.App , Id string)(*model.UpdateProfile, error)   {
	ctx := context.Background()
	var visits *model.UpdateProfile
	client , err := clientId.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return visits, err
	}
	

	defer client.Close()
	iterator := client.Collection(collection).Where("Id", "==", Id).Documents(ctx)

	//defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			return visits, err
		}
		visits = &model.UpdateProfile {
			Id : doc.Data()["Id"].(string),
			FirstName : doc.Data()["FirstName"].(string),
			Email : doc.Data()["Email"].(string),
			Phone: doc.Data()["PhoneNo"].(string),
			LastName: doc.Data()["LastName"].(string),
			Country: doc.Data()["Country"].(string),
			Zip: doc.Data()["Zip"].(string),
			HouseAddress: doc.Data()["HouseAddress1"].(string),
			SubAddress: doc.Data()["Address"].(string),
			Male: doc.Data()["Eve"].(bool),
		}
		break
	}
	return visits, nil	
}

