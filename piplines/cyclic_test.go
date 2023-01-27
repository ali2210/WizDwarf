package piplines

import (
	"testing"

	bio "github.com/ali2210/wizdwarf/other/bioinformatics"
	"github.com/ali2210/wizdwarf/other/users"
	"github.com/hashicorp/vault/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Piplines_Cyclic_Testing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piplines cyclic testing")
}

var _ = Describe("Piplines cyclic testing", func() {

	Context(" Started...", func() {
		It(" Return Application Object", func() {
			Specify(" Empty Application Object", func() {
				Expect(GetDBClientRef()).Should(BeNil())
			})
		})
		It(" Return Firestore Client Object", func() {
			var deps_object users.DBFirestore
			Specify(" Empty Firestore Object", func() {
				Expect(GetDBCollect()).Should(BeEquivalentTo(deps_object))
			})
		})
		It(" Return Empty Credentials", func() {
			Specify(" Empty Credentials", func() {
				Expect(GetKeyFile()).Should(BeEmpty())
			})
		})
		It(" Return Project-Credentials", func() {
			Specify(" Project Credentials", func() {
				Expect(GetProjectID()).ShouldNot(BeEmpty())
			})
		})
		It(" Return Levenshtein Algorithms parameters ", func() {
			var deps_object bio.LevenTable
			Specify(" Edit Distance parameters", func() {
				Expect(GetDBCollect()).Should(BeEquivalentTo(deps_object))
			})
		})
		It(" Return User generated results ", func() {
			Specify(" have user generated results", func() {
				Expect(GetBioAlgoParameters()).Should(BeAssignableToTypeOf(struct {
					Probablity float64
					Percentage float64
					Name       string
				}{}))
			})
		})
	})
})

var _ = Describe("Vault test suits running", func() {

	// Before running vault tests
	BeforeEach(func() {
		When("Parser code analysis tests", func() {
			Context("Parser returns empty credentials", func() {
				It("Should empty credentials succeed", func() {
					Expect(HCLDeclaration{
						Weatherapi:  "",
						Channel_key: "",
						Channel_id:  "",
						Secret:      "",
						Cluster_ID:  "",
						Token_Auth:  "",
					}).Should(BeEmpty())
				})
			})
		})
		When("Pepper extraction from salt message", func() {
			Context("Let's started on on some text", func() {
				It("Should extract salt message or function will throw exception", func() {
					Expect(Extractor("Hello, world!", "world")).ShouldNot(BeEmpty())
				})
			})
		})
	})
	AfterEach(func() {
		When("When data should be empty", func() {
			Context("data is empty", func() {
				It("should be empty", func() {
					Expect(PutKV(&HCLDeclaration{
						Weatherapi:  "",
						Channel_key: "",
						Channel_id:  "",
						Secret:      "",
						Cluster_ID:  "",
						Token_Auth:  "",
					}, " ", &api.Client{})).Should(BeNil())
				})
			})
		})
		When("When data is not provided", func() {
			Context("Does vault return no data", func() {
				It("returns no data", func() {
					value, err := GetKV("")
					Expect(err).Should(BeNil())
					Expect(value).To(BeEmpty())
				})
			})
		})
	})
})
