/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package logformatter

// libraries
import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func LogsTestSuits(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Test suits starting .... ")
}

var _ = Describe("Logging Tests running ......", func() {

	BeforeEach(func() {
		When("Loging Configuration default test ", func() {
			Context("Logging Configuration setting test", func() {
				It("Should Logging configuration throw exception", func() {
					Expect(New()).ShouldNot(BeEmpty())
				})
			})
		})
	})
})
