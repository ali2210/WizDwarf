package bucket

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func BucketTest(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Bucket services test started ")
}

var _ = Describe("bucket proto provide additional services", func() {

	// bucket object help to access bucket proto serives
	Context("bucket proto services", func() {
		It("should with proto return bucket object", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			Expect(New_Client(&ctx, &firestore.Client{})).Should(BeEmpty())
		})
	})

})
