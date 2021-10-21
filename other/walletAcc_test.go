package structs

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func WalletAccountTestCases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wallet Account information Test")
}

var _ = Describe(" Wallet Account Test ->", func() {
	account := Acc{
		Email:      "",
		Password:   "",
		Terms:      false,
		EthAddress: "",
		PubKey:     "",
		PrvteKey:   "",
		Allowed:    false,
	}
	Context("Wallet Account [1] Private Key ", func() {

		It("Account must have Private Key", func() {
			typeOf := reflect.ValueOf(&account).Elem()
			valueKey := typeOf.Index(5)
			Expect(valueKey).Should(BeEmpty())
		})
	})

	Context("Wallet Account Key Value Status", func() {
		It("Wallet Account Key", func() {
			Expect(account.GetPrivateKey()).Should(BeEmpty())
		})
	})

})
