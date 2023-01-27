/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package session

// libraries
import (
	"net/http"
	"testing"

	"github.com/gorilla/sessions"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// test suit
func SessionTestsuit(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Session suits get started ....")
}

var _ = Describe("User sessions test suits running .....", func() {

	var cookies Cookies
	BeforeEach(func() {
		When("Sessions initilaized test", func() {
			Context("Sessions with empty fields", func() {
				It("Should session succeed with empty fields", func() {
					Expect(&Cookies{
						Response: nil,
						Request:  &http.Request{},
						cookies:  &sessions.CookieStore{},
					}).Should(BeEmpty())
				})
			})
			Context("Session created with empty parameters", func() {
				It("Should create a session", func() {
					cookies.SetContextSession(&sessions.CookieStore{}, nil, &http.Request{})
					Expect(cookies.GetContextSession()).Should(BeEmpty())
				})
			})
		})
	})

	AfterEach(func() {
		When("Token validation test suits started....", func() {
			Context("New Token is valid", func() {
				It("Should session is created or fail", func() {
					Expect(cookies.NewToken()).Should(BeNil())
				})
			})
			Context("Token expire before login", func() {
				It("Should session expire", func() {
					Expect(cookies.ExpireToken()).Should(BeNil())
				})
			})
		})
	})
})
