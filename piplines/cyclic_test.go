package piplines

import (
	"testing"

	bio "github.com/ali2210/wizdwarf/other/bioinformatics"
	"github.com/ali2210/wizdwarf/other/users"
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
