package handler

import (
	"testing"

	paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func RunPayPalTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paypal Test cases run")
}

var abstract PaypalClientLevel = PaypalClientGo()
var credit paypalSdk.CreditCard = paypalSdk.CreditCard{
	ID:                 "",
	PayerID:            "",
	ExternalCustomerID: "",
	Number:             "",
	Type:               "",
	ExpireMonth:        "",
	ExpireYear:         "",
	CVV2:               "",
	FirstName:          "",
	LastName:           "",
	BillingAddress:     &paypalSdk.Address{},
	State:              "",
	ValidUntil:         "",
}
var _ = Describe("Paypal API-Client Test", func() {
	var id string = "paypal-test"
	Context("Paypal abstraction defintionn", func() {
		client, _ := abstract.NewClient()
		It("Paypal Client Test", func() {
			Expect(abstract.NewClient()).ShouldNot(BeEmpty())
		})

		It("Paypal Client Token", func() {

			Expect(abstract.Token(client))
		})
		It("Paypal Client Remove Card", func() {
			Expect(abstract.RemoveCard(id, client)).ShouldNot(BeEmpty())
		})
		It("Paypal Client Retreive card info", func() {
			Expect(abstract.RetrieveCreditCardInfo(id, client)).ShouldNot(BeEmpty())
		})
		It("Paypal Client info storage", func() {
			Expect(abstract.StoreCreditCardInfo(credit, client)).ShouldNot(BeEmpty())
		})
	})

})
