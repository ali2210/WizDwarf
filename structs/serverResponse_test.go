package structs

import (
	"net/http"
	"regexp"
	"testing"
	"text/template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ServerTestBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Test-cases")
}

var _ = Describe("Serves Response against your request", func() {
	responseStruct := Response{
		Flag:      false,
		Message:   "",
		Links:     "",
		cResponse: nil,
		cRequest:  &http.Request{},
		t:         &template.Template{},
	}

	Context("User Html request ->:", func() {
		It("File Request must be exist in directory", func() {
			pattern := "(^[A-Z][a-z]|[0-9][+\\].[a-z]+\\$)"
			filename := "index.html"
			validPattern := regexp.MustCompile(pattern)
			ok := validPattern.MatchString(filename)
			if ok {
				Expect(responseStruct.ClientHTMLRequest(filename)).Should(BeAnExistingFile())
			}

		})
	})

	Context("Response generate:", func() {
		It("Response must be valid ", func() {
			param := Response{
				Flag:      false,
				Message:   "",
				Links:     "",
				cResponse: nil,
				cRequest:  &http.Request{},
				t:         &template.Template{},
			}
			Expect(responseStruct.ClientRequestHandle(param.Flag, param.Message, param.Links, param.cResponse, param.cRequest)).Should(BeNil())
		})
	})

	Context("Run command", func() {
		It("Parse File", func() {
			Expect(responseStruct.Run(responseStruct.t)).Should(MatchError("fail"))
		})
	})
})
