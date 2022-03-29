package date_time

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Datetime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Time unit tests")
}

// this test ensures that valid time and date generate
var _ = Describe("time format unit tests", func() {
	Context("Date tests", func() {
		s := "20"
		It("Should date returns empty data", func() {
			// empty date
			Expect(Date(s)).Should(BeEmpty())
		})
		It("Or should return data", func() {
			// date have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context("Month test suit", func() {

		s := "09"
		It("Should month returns empty", func() {
			// empty month
			Expect(Date(s)).Should(BeEmpty())
		})
		It("Or Should return data", func() {
			// month have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context("Year test suit", func() {

		s := "2021"
		It("Should this is a not valid year", func() {
			// empty year
			Expect(Date(s)).Should(BeEmpty())
		})
		It("Or Should this be a valid year", func() {
			// year have some numeric values
			Expect(Date(s)).ShouldNot(BeEmpty())
		})
	})
	Context("UTC time suit tests", func() {
		It("Should valid time format", func() {
			// time according utc format
			Expect(GetToday(2021, time.Month(9), 20)).Should(BeEmpty())
		})
	})
})
