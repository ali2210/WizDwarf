package collection

import (
	"context"
	"crypto/sha256"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var collection_test = "pictures collection tests started"
var pictures_store_test = "pictures store as collection tests started"
var results = "should picture saved"

func ProtocolBuffers_CollectionTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, collection_test)
}

var _ = Describe(collection_test, func() {
	Context(pictures_store_test, func() {
		gallery := Gallery_Stream_Server{}
		pic := Pictures{
			PicId:   "micky0",
			PicSrc:  "mickymouse.jpeg",
			PicDate: "20-11-2021",
			PicTime: "20-11-2021,21:21,utc+0000",
			PicTags: "micky0",
		}
		pic.CDR = make(map[string]string, 1)
		hashvalue := sha256.Sum256([]byte(pic.PicId))
		pic.CDR["unecrypted"] = string(hashvalue[:])
		It(results, func() {
			Expect(gallery.NewPictures(context.Background(), &pic)).ShouldNot(BeTrue())
		})
	})
})
