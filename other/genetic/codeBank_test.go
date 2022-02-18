package genetics

import(
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func GenesTests(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "genes tests suited")
}

var _ = Describe("Genetics provide data storage", func() {
	Context("Genes object created ", func() {
		It("should create a new Genetics object", func(){
			Expect(New()).ShouldNot(BeEmpty())
		})
	})
})