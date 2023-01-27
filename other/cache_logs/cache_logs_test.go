package cache_logs

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Cache_test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logs stacked test suits ..... ")
}

var _ = Describe("Logs Tests suits should notbe failed during compliation", func() {

	var logstreams *Bigcached
	BeforeEach(func() {

		When("Logs Configuration setting  .....", func() {
			logstreams = New(GetBigcached_config())
			Expect(logstreams).ShouldNot(BeEmpty())
		})
	})

	AfterEach(func() {

		When("After Configuration test succeed, Logstream hold event value", func() {
			Context("Logstream with single test", func() {
				It("should log events are registered:", func() {
					Expect(logstreams.Set_Key("dump", "dumpy test suit:")).ShouldNot(BeNil())
				})
			})
		})

		When("Logstream retreive event value", func() {
			Context("Logstream with retreive test", func() {
				It("should log events retreive:", func() {
					value, err := logstreams.Get_Key("dump")
					Expect(value).ShouldNot(BeEmpty())
					Expect(err).ShouldNot(BeNil())
				})
			})
		})
	})
})
