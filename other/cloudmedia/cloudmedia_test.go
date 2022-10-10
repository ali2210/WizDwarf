/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package

package cloudmedia

// Libraries

import (
	"context"
	"testing"

	"time"

	"cloud.google.com/go/firestore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func CloudMediaTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloudmedia Test suits starting ....")
}

var _ = Describe("Cloudmedia Test suites running .....", func() {
	BeforeEach(func() {
		When("Cloudmedia Instance created", func() {
			Context("Instance created ensure validation", func() {
				It("Should it is a valid instance", func() {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
					defer cancel()
					Expect(NewDc_1(ctx, &firestore.Client{})).Should(BeEmpty())
				})
			})
		})
	})
})
