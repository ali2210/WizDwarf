/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package bucket

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/bucket/fireclient"
	"github.com/ali2210/wizdwarf/other/bucket/proto"
)

// general variables
// var content_address string = ""
// var Client *firestore.Client
var Key string = ""
var Composite string = ""
var instance fireclient.DBStorage

// constants
const APIKEY string = "skynetdwarfs"
const USER_AGENT string = "Sia-Agent"
const ENDPOINT string = "/"

type Bucket_Service interface {
	New_Bucket(o *proto.Object) *proto.IObject
	Preview(query *proto.Query, content string) *proto.QState
	Download(query *proto.Query, content string) *proto.QState
}

type GBucket struct {
	ctx    *context.Context
	Client *firestore.Client
}

func New_Client(ctx *context.Context, client *firestore.Client) Bucket_Service {
	return &GBucket{ctx: ctx, Client: client}
}

func (c *GBucket) New_Bucket(o *proto.Object) *proto.IObject {

	// ensure data validation
	if strings.Contains(o.Name, " ") {
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	if strings.Contains(o.Types, " ") {
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	data := "app_data/" + o.Name + o.Types

	// read docker apps directory
	path, err := os.Stat("app_data/")
	if err != nil {
		log.Println(" Error stat :", err.Error())
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	// if the path is not a directory then throw an exception
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err.Error())
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	client := skynet.New()

	// skynet param defined
	option := skynet.DefaultUploadOptions
	option.APIKey = APIKEY
	option.CustomUserAgent = USER_AGENT
	// option.EndpointPath = "/"

	log.Println("New File created on the disk ....")

	log.Println("Uploading ..... ")

	// upload file over skynet storage
	root_link, err := client.UploadFile(data, option)
	if err != nil {
		log.Println(" Error file upload: ", err.Error())
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	log.Println("Your content address:", root_link)

	// store content address in a cache
	// SetContent(root_link)

	// decentralize file storage client access
	// store cache data in our database for different purpose
	status := fireclient.New(context.Background(), c.Client).Store(Key, root_link, Composite)

	log.Println("Store User File Description : ....", Key, Composite)

	// if the status is "2" then there may be some reason code reject cache data
	if status == 2 {
		return &proto.IObject{Iobject: &proto.Object{}, Istatus: proto.Object_Status_ERROR}
	}

	return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_OK}
}

// content address is an machine generated address for a particular content.
// This is similarly to public key but the key is 32bit long.
// "Set" will hold new link
// func SetContent(link string) { content_address = link }

// // content address have another operation "Get"
// func GetContent() string { return content_address }

// "preview" doesn't mean user will read the file content.
// preview is a snapshot of a document. User will view document properties.

func (c *GBucket) Preview(query *proto.Query, content string) *proto.QState {

	// var client_Bucket fireclient.DBStorage
	var sia_client = skynet.New()

	// the store transaction utx00000000 refer to the content.
	// the main reason for "preview" is to verified ownership.
	// If the ownership verification fails, then no such transaction exists & vice versa.
	metadata_option := skynet.DefaultMetadataOptions
	// metadata_option.APIKey = APIKEY
	metadata_option.CustomUserAgent = USER_AGENT

	err := sia_client.Metadata(content, metadata_option)
	if err != nil {
		log.Println(" Error getting metadata", err.Error())
		return &proto.QState{Qstatus: proto.QStatus_Err}
	}
	return &proto.QState{Qstatus: proto.QStatus_Ok}
}

// "Download" function allow user to download document file on your local machine
func (c *GBucket) Download(query *proto.Query, content string) *proto.QState {

	// var client_Bucket fireclient.DBStorage
	var sia_client = skynet.New()
	download_option := skynet.DefaultDownloadOptions
	// download_option.APIKey = APIKEY
	download_option.CustomUserAgent = USER_AGENT

	// the verified transaction will be download on the local machine.
	// If the transaction is not verified, it will return an error and vice versa.
	err := sia_client.DownloadFile("app_data/"+query.ByName, content, download_option)
	if err != nil {
		log.Printf(" Error download file from %s: %v", query.ByName, err.Error())
		return &proto.QState{Qstatus: proto.QStatus_Err}
	}

	log.Println("Downloading succeed .... ", content)
	return &proto.QState{Qstatus: proto.QStatus_Ok}
}

// replicate copy of bucket
func NewObject(i fireclient.DBStorage) { instance = i }

// return copy of bucket
func GetObject() fireclient.DBStorage { return instance }
