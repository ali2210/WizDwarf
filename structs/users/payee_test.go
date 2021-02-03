package users

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var customers CreditCardInfo

func PayeeTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payee Test - cases")
}

var _ = Describe("Payee Interface", func() {
	Context("Payee Interface Client", func() {
		It("Payee Client Object:", func() {
			Expect(NewClient()).Should(BeNil())
		})
	})
	Context("Payee Interface :", func() {
		It("Payee Client ID:", func() {
			id := "payee"
			customers.SetAuthorizeStoreID(id)
			Expect(customers.GetAuthorizeStoreID()).Should(BeEmpty())
		})
	})
	Context("Payee Interface Client", func() {
		It("Payee Client", func() {
			Expect(customers.VoidStruct()).Should(BeNil())
		})
	})
})
