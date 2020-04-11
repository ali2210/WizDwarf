package db


import (
	"context"
	"log"
	"cloud.google.com/go/firestore"
	// "../record"
)


 type ProfileVistors struct{
	Id string 			`json:"Id"`
	Name string 		`json:"Name"`
	Email string    	`json:"Email"`
	Password string  	`json:"Password"`
}

const (
	projectId string = "htickets-cb4d0"
	collectionName string = "ProfileVistors"
)


type DBFirestore interface{
	SaveData(visitor *ProfileVistors)(*ProfileVistors, error)
	// FindData(user *Create_User, visitor *profile.ProfileVistors)(*profile.ProfileVistors, error)
	FindAllData()([]ProfileVistors, error)
}

type cloud_data struct{}


func NewCloudInstance()  DBFirestore{
	return &cloud_data{}
}

func (*cloud_data)SaveData(visitor *ProfileVistors)(*ProfileVistors, error){
	ctx := context.Background()
	client , err := firestore.NewClient(ctx, projectId); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	defer client.Close()

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

func (*cloud_data)FindAllData()([]ProfileVistors,error){
	ctx := context.Background()
	client , err := firestore.NewClient(ctx, projectId); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}

	defer client.Close()

	var visitors []ProfileVistors 
	iterator := client.Collection(collectionName).Documents(ctx)
	for{
		doc, err := iterator.Next(); if err != nil{
			log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}

		visitor := ProfileVistors{
			Id: doc.Data()["ID"].(string),
			Name : doc.Data()["Name"].(string),
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
		} 
		visitors = append(visitors, visitor)
	}
	return visitors, nil

}
