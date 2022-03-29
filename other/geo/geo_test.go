package geo

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func GeoBDDTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geonavigation unit tests")
}

// this test ensures that geo navigation tests provide valid coordinates
var _ = Describe("Geo test started", func() {

	Context("Run tests on coordinates", func() {
		It("Should coordinates are valid", func() {
			Expect(Location("30.77, 70.77")).Should(BeEmpty())
		})
	})

})
