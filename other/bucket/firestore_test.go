package bucket

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func BucketTest(t *testing.T){
	
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bucket services test started ")
}

var _ = Describe("Bucket offered services ", func() {
	
	// create a test buffer 
	Context("Bucket Service Object create", func() {
		It("should create a bucket", func(){
			Expect(New("init", "init-hello", "hello world")).ShouldNot(BeEmpty())
		})
	})
})