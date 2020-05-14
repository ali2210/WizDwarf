package db


import (
	"context"
	"log"
	"fmt"
	firebase "firebase.google.com/go"
	"time"
	// "cloud.google.com/go/firestore"
	// "../record"
)


type Vistors struct{
	Id string 			`json:"Id"`
	Name string 		`json:"Name"`
	Email string    	`json:"Email"`
	Password string  	`json:"Password"`
	FName string        `json:"FName"`
	City string         `json:"City"`
	Zip  string         `json:"Zip"`
	Address string      `json:"Address"`
	LAddress string     `json:"LAddress"`
	Country  string     `json:"Country"`
	Eve bool    		`json:"Eve"`
}

const (
	// projectId string = "htickets-cb4d0"
	collectionName string = "ProfileVistors"
)


type DBFirestore interface{
	SaveData(visitor *Vistors, app *firebase.App)(*Vistors, error)
	FindData(email string, pass string , app *firebase.App)(*Vistors, error)
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
	var visits Vistors
	timerStrt := time.Now()
	t := time.Now()
	fmt.Println("Please wait ...", t)
	iterator := client.Collection(collectionName).Where("Email", "==",visitor.Email).Documents(ctx)
	fmt.Printf("Iterator%v\n", iterator)
	defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			//log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}
		fmt.Printf("Data:%v\n", doc.Data())
		visits = Vistors {
			Email : doc.Data()["Email"].(string),
			Password: doc.Data()["Password"].(string),
		}
		fmt.Println("Process complete ...", t.Sub(timerStrt))
		break
	}
	fmt.Println("visits:", visits)
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
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
		fmt.Println("Process complete ...\n", t.Sub(timerStrt), "Record added")
		return visitor, nil

}

func (*cloud_data)FindAllData(app *firebase.App)([]Vistors,error){
	ctx := context.Background()
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	timerStrt := time.Now()
	t := time.Now()
	fmt.Println("Please wait ...", t)

	defer client.Close()

	var visits []Vistors
	iterator := client.Collection(collectionName).Documents(ctx)
	fmt.Printf("Iterator:%+v\n", iterator)
	// defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}
		 fmt.Printf("Data:%v", doc.Data())

			visit := Vistors {
				Id : doc.Data()["Id"].(string),
				Name : doc.Data()["Name"].(string),
				Email : doc.Data()["Email"].(string),
				Password: doc.Data()["Password"].(string),
			} 
			visits = append(visits, visit)
			fmt.Println("Process complete ...", t.Sub(timerStrt))
	}
	return visits, nil

}

func (*cloud_data)FindData(email string , pass string, app *firebase.App)(*Vistors, error){
	ctx := context.Background()
	client , err := app.Firestore(ctx); if err != nil{
		log.Fatal("Client Instance Failed to start", err)
		return nil, err
	}
	timerStrt := time.Now()
	t := time.Now()
	fmt.Println("Please wait ...", t)

	defer client.Close()
	var visits Vistors
	iterator := client.Collection(collectionName).Where("Email", "==", email).Where("Password", "==" , pass).Documents(ctx)
	fmt.Printf("Iterator%v\n", iterator)
	defer iterator.Stop()
	for{
		doc, err := iterator.Next();if err != nil{
			//log.Fatal("Iterator Failed on Vistor: ", err)
			return nil, err
		}
		fmt.Printf("Data:%v\n", doc.Data())
		visits = Vistors {
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
		fmt.Println("Process complete ...", t.Sub(timerStrt))
		break
	}
	return &visits, nil
}



