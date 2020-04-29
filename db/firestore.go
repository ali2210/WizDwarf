package db


import (
	"context"
	"log"
	"fmt"
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
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	fmt.Printf("Documents:%v\n", ctx)

	defer client.Close()

	var visits []Vistors
	iterator := client.Collection(collectionName).Documents(ctx)
	//fmt.Printf("Iterator:%+v\n", iterator)
	defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}
		// fmt.Printf("Data:%v", doc.Data())

		visit := Vistors {
			Id : doc.Data()["Id"].(string),
			Name : doc.Data()["Name"].(string),
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
		} 
		visits = append(visits, visit)
		if err == nil{
			break
		}
	}
	return visits, nil

}
