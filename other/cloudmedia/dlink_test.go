/*This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package cloudmedia

// Libraries

import (
	"testing"

	"github.com/SkynetLabs/go-skynet/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Decentralize_Link_Test_Suit(t *testing.T) {

	RegisterFailHandler(Fail)
	RunSpecs(t, "Dlink Test Suit starting ....")
}

var _ = Describe("Dlink Test Suit running ...", func() {
	BeforeEach(func() {
		When("Dlink object instance ....", func() {
			Context("Instaniate skynet client object ", func() {
				It("Should skynet client object created", func() {
					Expect(NewDlinkObject(&skynet.SkynetClient{}, " ")).ShouldNot(BeEmpty())
				})
			})
		})
	})

	AfterEach(func() {
		When("Dlink services test ....", func() {
			Context("User will upload content !", func() {
				It("should upload content", func() {
					Expect(NewDlinkObject(&skynet.SkynetClient{}, " ").Generate(" ", []string{""}...)).Should(BeEmpty())
				})
			})
			Context("User will get content copy", func() {
				It("should get content copy", func() {
					Expect(NewDlinkObject(&skynet.SkynetClient{}, " ").Get(" ", []string{" "}...)).Should(BeEmpty())
				})
			})
		})
	})
})
