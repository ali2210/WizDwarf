package parser

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/goombaio/namegenerator"
)

// get picture file name without extension
// png, gif, tif , img, jpeg
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
