package openweather

import (
	"testing"

	"github.com/briandowns/openweathermap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var data = DataVisualization{
	Percentage:  0,
	Process:     0,
	SeenBy:      "",
	coordinates: openweathermap.Coordinates{},
	UVinfo:      []openweathermap.UVIndexInfo{},
}

var coordinates = MyCoordinates{
	Longitude: 0,
	Latitude:  0,
}

var apikey = "7efdb33cK"

func OpenWeatherPackageTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Open-Weather testing")
}

var _ = Describe("Open weather testcase", func() {

	// Open-weather API VERSION TETST
	Context("Open weather API-KEY Request", func() {
		It("Open-weather API-KEY test", func() {

			Expect(data.OpenWeather(apikey)).ShouldNot(BeEmpty())
		})
	})

	// Open- weather track your geopointt

	Context("Open Weather Geopoint Testcase", func() {
		It("Open weather Geopoint Results", func() {
			Expect(data.GetCoordinates(&coordinates)).ShouldNot(BeEmpty())
		})
	})

	// Open -weather uv and coordinates
	Context("Open weather UV Coordinatees Testcase", func() {
		It("Open-weather UV Results based on coordinates ", func() {
			c := data.GetCoordinates(&coordinates)
			u, _ := data.OpenWeather(apikey)
			Expect(data.UVCoodinates(c, u)).ShouldNot(BeEmpty())
		})
	})

	// Open-weather uv archieve

	Context("Open weather Uv Archeive... ", func() {
		It("Open-weather Uv history Results", func() {
			u, _ := data.OpenWeather(apikey)
			Expect(data.UVCompleteInfo(u)).ShouldNot(BeEmpty())
		})
	})

})
