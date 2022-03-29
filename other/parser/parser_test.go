package parser

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ParserBDDTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Tests ")
}

//  this test ensures that picture store correctly
var _ = Describe("parser unit tests started", func() {
	Context("Run tests on images ", func() {
		s := "mickeymouse.jpeg"
		It("Should image have been created", func() {
			Expect(ParseTags(s)).ShouldNot(BeEmpty())
		})
	})
	Context("Name Generate:", func() {
		It("Should have a name", func() {
			Expect(Generator()).ShouldNot(BeEmpty())
		})
	})
})
