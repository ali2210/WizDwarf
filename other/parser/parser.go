package parser

import (
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/goombaio/namegenerator"
)

const INTERNAL_PATH string = "app_data/"

// Read File and extracts type format. File supported these formats "png, tif, jpeg, gif"
func ParseTags(s string) []string {
	var tags = make([]string, len(s))

	if strings.Contains(s, ".png") {
		tags = strings.Split(s, ".png")
	} else if strings.Contains(s, ".jpeg") {
		tags = strings.Split(s, ".jpeg")
	} else if strings.Contains(s, ".img") {
		tags = strings.Split(s, ".img")
	} else if strings.Contains(s, ".gif") {
		tags = strings.Split(s, ".gif")
	} else if strings.Contains(s, ".tif") {
		tags = strings.Split(s, ".tif")
	} else {
		tags = append(tags, " ")
	}
	return tags
}

// Generator returns File name
func Generator() string {

	left, err := rand.Int(rand.Reader, big.NewInt(55))
	if err != nil {
		panic(err.Error())
	}

	right, err := rand.Int(rand.Reader, big.NewInt(52))
	if err != nil {
		panic(err.Error())
	}

	leftGen := namegenerator.NewNameGenerator(left.Int64())
	rightGen := namegenerator.NewNameGenerator(right.Int64())

	return rightGen.Generate() + "_" + leftGen.Generate()
}

func CreateFile(filename ...string) (*os.File, error) {

	if _, err := os.Stat(filename[0]); os.IsExist(err) {
		return &os.File{}, err
	}

	paths, err := os.Stat(INTERNAL_PATH)
	if err != nil {
		return &os.File{}, err
	}

	// Application storage path
	if !paths.IsDir() {
		return &os.File{}, err
	}

	src_image := "Avatar-*-" + filename[0]

	// Store user-picture file in the storage directory
	_contentFile, err := ioutil.TempFile(filepath.Dir(INTERNAL_PATH), src_image)
	// _contentFile, err := os.OpenFile(src_image, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		return &os.File{}, err
	}

	// defer _contentFile.Close()

	return _contentFile, nil
}
