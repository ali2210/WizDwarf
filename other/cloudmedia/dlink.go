/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package cloudmedia

// Libraries
import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/cloudmedia/dlink"
)

// Global Variables

var options skynet.UploadOptions

// Constants
const APIKEY string = "skynetdwarfs"
const USER_AGENT string = "Sia-Agent"
const ENDPOINT string = "/"

// Interface or Services
type DLINK_DC interface {
	Generate(data string, datatype ...string) (dlink.Errors, string)
	Get(filename string, link ...string) dlink.Errors
	Preview_Get(filename string, link ...string) (dlink.PreviewLink, dlink.Errors)
}

// Dlink Struct Object
type DLINK_Object struct {
	client *skynet.SkynetClient
	direc  string
}

// New Object
// @param client and directory path
// @return Dlink struct
func NewDlinkObject(Client *skynet.SkynetClient, Direc string) DLINK_DC {
	return &DLINK_Object{client: Client, direc: Direc}
}

// Upload content on skynet
// @param data string message , datatype string message as final
// @return Dlink Error & Content Url
// @receiver dlink
func (o *DLINK_Object) Generate(data string, datatype ...string) (dlink.Errors, string) {

	// ensure data validation
	if strings.Contains(data, " ") {
		return dlink.Errors_ERR_EMPTY_URL, " "
	}

	// ensure valid directory
	path, err := os.Stat(o.direc)
	if err != nil && !path.IsDir() {
		return dlink.Errors_ERR_PATH_MOVED, " "
	}

	// skynet credentials
	options = skynet.DefaultUploadOptions
	options.APIKey = APIKEY
	options.CustomUserAgent = USER_AGENT

	// Upload Content
	url, err := o.client.UploadFile(data, options)
	if err != nil {
		log.Println("Error upload file  :", err)
		return dlink.Errors_ERR_UNKNOWN_, " "
	}

	return dlink.Errors_NONE, url
}

// This function allow user to get copy of content which he had shared before
// @parm filename string message & final link vector of string
// @return error
// @receiver dlink
func (o *DLINK_Object) Get(filename string, link ...string) dlink.Errors {

	// ensure data validation
	if strings.Contains(filename, " ") && strings.Contains(link[0], " ") {
		return dlink.Errors_ERR_UNKNOWN_
	}

	// download content copy in your directory
	err := o.client.DownloadFile(filename, link[0], skynet.DefaultDownloadOptions)
	if err != nil {
		return dlink.Errors_ERR_INVALID_URL
	}

	return dlink.Errors_NONE
}

func (o *DLINK_Object) Preview_Get(filename string, link ...string) (dlink.PreviewLink, dlink.Errors) {

	if errs := o.Get(filename, link...); errs != dlink.Errors_NONE {
		return dlink.PreviewLink{}, errs
	}

	read, err := os.ReadDir("app_data/")
	if err != nil {
		return dlink.PreviewLink{}, dlink.Errors_ERR_PATH_MOVED
	}

	for dir, _ := range read {
		if strings.Contains(filename, read[dir].Name()) {

			content, err := os.ReadFile(filename)
			if err != nil {
				return dlink.PreviewLink{}, dlink.Errors_ERR_PATH_MOVED
			}

			return dlink.PreviewLink{
				Filename: filename,
				Path:     "app_data/" + filename,
				Size:     GetFileSize(),
				Readable: string(content[:]),
				DateTime: GetFileCreationTime().String(),
			}, dlink.Errors_NONE
		}
	}

	return dlink.PreviewLink{}, dlink.Errors_ERR_UNKNOWN_
}

func GetFileSize(filename ...string) string {

	_, err := os.ReadDir("app_data/")
	if err != nil {
		return " "
	}

	properties, err := os.Stat(filename[0])
	if err != nil {
		return " "
	}

	return fmt.Sprintf("%d", properties.Size()/1024) + "KiB"
}

func GetFileCreationTime(filename ...string) time.Time {

	_, err := os.ReadDir("app_data/")
	if err != nil {
		return time.Time{}
	}

	properties, err := os.Stat(filename[0])
	if err != nil {
		return time.Time{}
	}

	return properties.ModTime()
}
