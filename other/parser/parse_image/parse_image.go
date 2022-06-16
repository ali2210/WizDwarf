package parse_image

import (
	"context"
	"image"
	"image/color"
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
	"github.com/ali2210/wizdwarf/other/bucket"
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

func Metadata(filename, key, ownership string, client *firestore.Client) int {

	if strings.Contains(filename, ".png") {

		datasource := bucket.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(decodeRawImage)

	} else if strings.Contains(filename, ".jpeg") {

		datasource := bucket.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(decodeRawImage)

	} else if strings.Contains(filename, ".gif") {

		datasource := bucket.NewClient(client, context.TODO(), ownership, key)

		return datasource.GenerateFingersprints(decodeRawImage)

	} else {
		log.Fatalln("FORMAT NOT SUPORTED")
		return bucket.Err
	}

}

func GetImageDecoder() image.Image {
	return decodeRawImage
}

func GetMetadata(key, ownership string, client *firestore.Client) (interface{}, int) {

	return bucket.NewClient(client, context.TODO(), ownership, key).AnalyzeFingersprints(decodeRawImage)
}

type PixelsValue struct {
	R, G, B, A uint32
}

var width, height int = 200, 200
var count int64 = 0

func EncodePixels(value []PixelsValue) *image.Paletted {

	avatar := make([]color.Color, decodeRawImage.Bounds().Max.X)
	for i := range value {

		// if reflect.DeepEqual(value[i].R, 0) && reflect.DeepEqual(value[i].G, 0) && reflect.DeepEqual(value[i].B, 0) && reflect.DeepEqual(value[i].A, 0) {
		// 	continue
		// }

		if reflect.DeepEqual(value[i], nil) {
			continue
		}
		// pixels_vec :=

		avatar = append(avatar, []color.Color{
			color.RGBA64{uint16(value[i].R), uint16(value[i].G), uint16(value[i].B), uint16(value[i].A)},
		}...)

		if reflect.DeepEqual(avatar[i], nil) {
			count += 1
			continue
		}

	}

	return image.NewPaletted(image.Rect(0, 0, width, height), avatar[count:])
}

func RGBA_Vec() []PixelsValue {

	encodeAvatar := make([]PixelsValue, decodeRawImage.Bounds().Max.X)

	for i := 0; i < GetImageDecoder().Bounds().Max.X; i++ {
		for j := 0; j < GetImageDecoder().Bounds().Max.Y; j++ {

			r, g, b, a := GetImageDecoder().At(i, j).RGBA()

			if reflect.DeepEqual(r, uint32(0)) && reflect.DeepEqual(g, uint32(0)) && reflect.DeepEqual(b, uint32(0)) && reflect.DeepEqual(a, uint32(0)) {
				continue
			}

			encodeAvatar[i].R, encodeAvatar[i].G, encodeAvatar[i].B, encodeAvatar[i].A = decodeRawImage.At(i, j).RGBA()

		}
	}

	return encodeAvatar
}

func Pixels_Vec(value []PixelsValue) *image.Paletted {

	return EncodePixels(value)
}
