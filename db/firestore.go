package db


import (
	"context"
	"log"
	firebase "firebase.google.com/go"
	// "cloud.google.com/go/firestore"
	// "../record"
)


type Vistors struct{
	Id string 			`json:"Id"`
	Name string 		`json:"Name"`
	Email string    	`json:"Email"`
	Password string  	`json:"Password"`
}

const (
	// projectId string = "htickets-cb4d0"
	collectionName string = "ProfileVistors"
)


type DBFirestore interface{
	SaveData(visitor *Vistors, app *firebase.App)(*Vistors, error)
	// FindData(user *Create_User, visitor *profile.ProfileVistors)(*profile.ProfileVistors, error)
	FindAllData(app *firebase.App)([]Vistors, error)
}

type cloud_data struct{}


func NewCloudInstance()  DBFirestore{
	return &cloud_data{}
}

func (*cloud_data)SaveData(visitor *Vistors, app *firebase.App)(*Vistors, error){
	ctx := context.Background()
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	println("Client info", client)
	defer client.Close()
	println("Records added in collection")
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id" :	visitor.Id,
		"Name" : visitor.Name,
		"Email" : visitor.Email,
		"Password": visitor.Password,
	}); if err != nil{
		log.Fatal("Failed to retrive Vistor Record:", err)
		return nil, err
	}

	return visitor, nil
}

func (*cloud_data)FindAllData(app *firebase.App)([]Vistors,error){
	ctx := context.Background()
	println(ctx)
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}

	defer client.Close()

	var visit []Vistors
	iterator := client.Collection(collectionName).Documents(ctx)
	for{
		doc, err := iterator.Next(); if err != nil{
			log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}

		visitor := Vistors {
			Id : doc.Data()["Id"].(string),
			Name : doc.Data()["Name"].(string),
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
		} 
		// println("Data_id:", visitor.Id)
		// println("Data_name:", visitor.Name)
		// println("Data_email:", visitor.Email)
		// println("Data_password:", visitor.Password)
		visit = append(visit, visitor)
	}
	return visit, nil

}
