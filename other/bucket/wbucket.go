/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package bucket

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/bucket/proto"
)

//
var content_address string = ""
var Client *firestore.Client
var Key string = ""
var Composite string = ""
var instance DBStorage
var client_Bucket DBStorage

var sia_client = skynet.New()

const API_KEY_SIA = "wizStorage"
const APP_VENDOR_CODE = "Sia-Agent"

type Bucket_Service interface {
	New_Bucket(o *proto.Object) *proto.IObject
	Preview(query *proto.Query) *proto.QState
	Download(query *proto.Query) *proto.QState
}

type GBucket struct {
	ctx *context.Context
}

func New_Client(ctx *context.Context) Bucket_Service {
	return &GBucket{ctx: ctx}
}

func (c *GBucket) New_Bucket(o *proto.Object) *proto.IObject {

	// read docker apps directory
	path, err := os.Stat("app_data/")
	if err != nil {
		log.Println(" Error stat :", err.Error())
		return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_ERROR}
	}

	// if the path is not a directory then throw an exception
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err.Error())
		return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_ERROR}
	}

	// skynet param defined
	option := skynet.DefaultUploadOptions
	option.APIKey = API_KEY_SIA
	option.CustomUserAgent = APP_VENDOR_CODE

	// upload file over skynet storage
	root_link, err := sia_client.UploadFile("app_data/"+o.Name+o.Types, option)
	if err != nil {
		log.Println(" Error file upload: ", err.Error())
		return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_ERROR}
	}

	// store content address in a cache
	SetContent(root_link)

	// decentralize file storage client access
	client_Bucket = New(Key, root_link, Composite)

	// store cache data in our database for different purpose
	status := client_Bucket.Store(context.Background(), Client)

	// if the status is not "0" then there may be some reason code reject cache data
	if status != 0 {
		return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_ERROR}
	}

	return &proto.IObject{Iobject: o, Istatus: proto.Object_Status_OK}
}

// content address is an machine generated address for a particular content.
// This is similarly to public key but the key is 32bit long.
// "Set" will hold new link
func SetContent(link string) { content_address = link }

// content address have another operation "Get"
func GetContent() string { return content_address }

// "preview" doesn't mean user will read the file content.
// preview is a snapshot of a document. User will view document properties.

func (c *GBucket) Preview(query *proto.Query) *proto.QState {

	// the store transaction utx00000000 refer to the content.
	// the main reason for "preview" is to verified ownership.
	// If the ownership verification fails, then no such transaction exists & vice versa.

	metadata_option := skynet.DefaultMetadataOptions
	metadata_option.APIKey = API_KEY_SIA
	metadata_option.CustomUserAgent = APP_VENDOR_CODE

	err := sia_client.Metadata(content_address, metadata_option)
	if err != nil {
		log.Println(" Error getting metadata", err.Error())
		return &proto.QState{Qstatus: proto.QStatus_Err}
	}
	return &proto.QState{Qstatus: proto.QStatus_Ok}
}

// "Download" function allow user to download document file on your local machine
func (c *GBucket) Download(query *proto.Query) *proto.QState {

	download_option := skynet.DefaultDownloadOptions
	download_option.APIKey = API_KEY_SIA
	download_option.CustomUserAgent = APP_VENDOR_CODE
	download_option.EndpointPath = "app_data/" + query.ByName

	// the verified transaction will be download on the local machine.
	// If the transaction is not verified, it will return an error and vice versa.
	err := sia_client.DownloadFile("app_data/"+query.ByName, content_address, download_option)
	if err != nil {
		log.Printf(" Error download file from %s: %v", query.ByName, err.Error())
		return &proto.QState{Qstatus: proto.QStatus_Err}
	}
	return &proto.QState{Qstatus: proto.QStatus_Ok}
}

// replicate copy of bucket
func NewObject(i DBStorage) { instance = i }

// return copy of bucket
func GetObject() DBStorage { return instance }
