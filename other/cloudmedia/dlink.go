package cloudmedia

import (
	"log"
	"os"
	"strings"

	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ali2210/wizdwarf/other/cloudmedia/dlink"
)

var options skynet.UploadOptions

const APIKEY string = "skynetdwarfs"
const USER_AGENT string = "Sia-Agent"
const ENDPOINT string = "/"

type DLINK_DC interface {
	Generate(data string, datatype ...string) (dlink.Errors, string)
	Get(filename string, link ...string) dlink.Errors
}

type DLINK_Object struct {
	client *skynet.SkynetClient
	direc  string
}

func NewDlinkObject(Client *skynet.SkynetClient, Direc string) DLINK_DC {
	return &DLINK_Object{client: Client, direc: Direc}
}

func (o *DLINK_Object) Generate(data string, datatype ...string) (dlink.Errors, string) {

	if strings.Contains(data, " ") {
		return dlink.Errors_ERR_EMPTY_URL, " "
	}

	path, err := os.Stat(o.direc)
	if err != nil && !path.IsDir() {
		return dlink.Errors_ERR_PATH_MOVED, " "
	}

	options = skynet.DefaultUploadOptions
	options.APIKey = APIKEY
	options.CustomUserAgent = USER_AGENT

	url, err := o.client.UploadFile(data, options)
	if err != nil {
		log.Println("Error upload file  :", err)
		return dlink.Errors_ERR_UNKNOWN_, " "
	}

	return dlink.Errors_NONE, url
}

func (o *DLINK_Object) Get(filename string, link ...string) dlink.Errors {

	if strings.Contains(filename, " ") && strings.Contains(link[0], " ") {
		return dlink.Errors_ERR_UNKNOWN_
	}

	err := o.client.DownloadFile(filename, link[0], skynet.DefaultDownloadOptions)
	if err != nil {
		return dlink.Errors_ERR_INVALID_URL
	}

	return dlink.Errors_NONE
}
