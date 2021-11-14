package structs

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/sessions"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var routeParam RouteParameter = RouteParameter{
	Response: nil,
	Request:  &http.Request{},
	cookies:  &sessions.CookieStore{},
}

func TokenizationBddTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tokenization Test handle")
}

var _ = Describe("Tokenization handle", func() {

	Context("No Behaviour", func() {
		It("No record report", func() {
			value := reflect.ValueOf(&routeParam).Elem()
			for i := 0; i < value.NumField(); i++ {
				state := value.Field(i).IsNil()
				Expect(state).Should(BeTrue())
			}

		})
		It("Retreive Record", func() {
			// value := reflect.ValueOf(&routeParam).Elem()
			Expect(routeParam.GetContextSession()).Should(BeNil())
		})

		It("Have Data but incomplete and inconsist state", func() {

			Expect(routeParam.NewToken()).Should(BeNil())
		})
	})

})
