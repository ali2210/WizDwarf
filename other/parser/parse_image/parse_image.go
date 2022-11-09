package parse_image

import (
	"context"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	fingerprints "github.com/WisdomEnigma/urban-fiesta/fingerprint"
	"github.com/ali2210/wizdwarf/other/bucket/fireclient"
)

var decodeRawImage image.Image

// PNG COLOR HASH Decode image and return the corresponding unique color hash

func PNG_Color_Hash(file *multipart.File) string {

	// decode picture and translate pxiels into vector for processing
	_upload_content, err := png.Decode(*file)
	if err != nil && err != io.EOF {
		log.Fatalln("Error decode :", err)
		return ""
	}

	decodeRawImage = _upload_content

	// return image signature
	return image_signature(_upload_content)
}

// JPEG COLOR HASH Decode image and return the corresponding unique color hash
func JPEG_Color_Hash(file multipart.File) string {

	// decode picture and translate pxiels into vector for processing
	_upload_content, err := jpeg.Decode(file)
	if err != nil {
		return ""
	}

	decodeRawImage = _upload_content

	// return image signature
	return image_signature(_upload_content)

}

// GIF COLOR HASH Decode image and return the corresponding unique color hash
func GIF_Color_Hash(file multipart.File) string {

	// decode picture and translate pxiels into vector for processing
	_upload_content, err := gif.Decode(file)
	if err != nil {
		return ""
	}

	decodeRawImage = _upload_content
	// return image signature
	return image_signature(_upload_content)

}

// image signature translate pixels into scalar for futher processing
func image_signature(_upload_content image.Image) string {

	// image properties
	detector := &fingerprints.Image_Print{}
	scanner_prints := make([]string, _upload_content.Bounds().Max.X)

	// Get Pixels values in rgba and translate into unique colors hash
	for i := 0; i < _upload_content.Bounds().Max.X; i++ {

		red, green, blue, alpha := _upload_content.At(i, i).RGBA()
		scanner_prints[i] = detector.CalculateHashColor(red, green, blue, alpha, int64(i))
	}

	// memory allocation reduction and return

	return strings.Join(scanner_prints, " ")[:]
}

// Metadata function return file description information
func Metadata(filename, key, ownership string, client *firestore.Client) int {

	if reflect.DeepEqual(ownership, "") {

		log.Fatalln("Empty ownership:", ownership)
		return fireclient.Err
	}

	if strings.Contains(filename, ".png") {

		datasource := fireclient.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(GetImageDecoder())

	} else if strings.Contains(filename, ".jpeg") {

		datasource := fireclient.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(GetImageDecoder())

	} else if strings.Contains(filename, ".gif") {

		datasource := fireclient.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(GetImageDecoder())

	} else {
		log.Fatalln("FORMAT NOT SUPORTED")
		return fireclient.Err
	}

}

// Get Image Decode function is a special function that returns image pixels in vec form
func GetImageDecoder() image.Image {
	return decodeRawImage
}

// GetMetadata instance is an bridge object. If user provide certain parameters then it will return metadata of the shared data
func GetMetadata(key, ownership string, client *firestore.Client) (interface{}, int) {

	return fireclient.NewClient(client, context.TODO(), ownership, key).AnalyzeFingersprints(GetImageDecoder())
}

// Pixels Value hold RGBA value
type PixelsValue struct {
	R, G, B, A uint32
}
