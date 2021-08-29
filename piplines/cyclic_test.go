package piplines

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ali2210/wizdwarf/structs/users"
	bio "github.com/ali2210/wizdwarf/structs/bioinformatics"
	"fmt"
	"testing"
)


func Piplines_Cyclic_Testing(t *testing.T)  {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piplines cyclic testing")
}

var _ = Describe("Piplines cyclic testing", func ()  {
	
	Context(" Started...", func ()  {
		It(" Return Application Object", func ()  {
			Specify(" Empty Application Object", func ()  {
				assert := Expect(GetDBClientRef()).Should(BeNil())
				fmt.Println("Empty Application Object status:", assert)
			})
		})
		It(" Return Firestore Client Object", func ()  {
			var deps_object users.DBFirestore
			Specify(" Empty Firestore Object", func ()  {
				assert := Expect(GetDBCollect()).Should(BeEquivalentTo(deps_object))
				fmt.Println("Empty Firestore Object status:", assert)
			})
		})
		It(" Return Empty Credentials", func ()  {
			Specify(" Empty Credentials", func ()  {
				assert := Expect(GetKeyFile()).Should(BeEmpty())
				fmt.Println("Empty Credentials status:", assert)
			})
		})
		It(" Return Project-Credentials", func ()  {
			Specify(" Project Credentials", func ()  {
				assert := Expect(GetProjectID()).ShouldNot(BeEmpty())
				fmt.Println("Project Credentials status:", assert)
			})
		})
		It(" Return Levenshtein Algorithms parameters ", func ()  {
			var deps_object bio.LevenTable
			Specify(" Edit Distance parameters", func ()  {
				assert := Expect(GetDBCollect()).Should(BeEquivalentTo(deps_object))
				fmt.Println(" Levenshtein status:", assert)
			})
		})
		It(" Return User generated results ", func ()  {
			Specify(" have user generated results", func ()  {
				assert := Expect(GetBioAlgoParameters()).Should(BeAssignableToTypeOf(struct{
					Probablity float64
					Percentage float64
					Name       string
				}{}))
				fmt.Println("User generated status:", assert)
			})
		})
	})
})