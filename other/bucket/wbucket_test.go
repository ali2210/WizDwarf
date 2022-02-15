package bucket

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"context"
)

func BucketProtoTest(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "bucket proto started")
}

var _ = Describe("bucket proto provide additional services", func() {

	// bucket object help to access bucket proto serives
	Context("bucket proto services", func(){
		It("should with proto return bucket object", func() {
			ctx := context.Background()
			Expect(New_Client(&ctx)).Should(BeEmpty())
		})
	})

	// content address route validation
	Context("Content have content address router link", func() {
		It("should with valid content address", func() {
			SetContent("sia://example.com")
			Expect(GetContent()).Should(Equal("sia://example.com"))
		})
	})
})