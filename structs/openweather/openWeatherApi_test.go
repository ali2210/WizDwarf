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
	RunSpecs(t, "Open-Weather")
}

var _ = Describe("Open weather", func() {

	Context("Open-weather ", func ()  {
		data := DataVisualization{}
		client, _ := data.OpenWeather(apikey)
		coordinates := data.GetCoordinates(&MyCoordinates{
			Longitude: 31.5204,
			Latitude : 74.3587,
		})
		err := data.UVCoodinates(coordinates, client)
		uvindex , _ := data.UVCompleteInfo(client)
		It(" Return Open-weather client object ", func ()  {
			Specify(" open weather client object", func ()  {
				Expect(client).Should(BeNil())
			})
		})
		It(" Return city coordinates", func ()  {
			Specify(" City coordinates status", func(){
				Expect(coordinates).Should(BeAssignableToTypeOf(struct{
					Longitude float64 
					Latitude float64  
				}{}))
			})
		})
		It(" Calculate uv index of your city", func ()  {
			Specify(" uv predicition ", func ()  {
				Expect(err).ShouldNot(Panic())
			})
		})
		It(" Uv city index of a day", func ()  {
			Specify(" uv city index", func ()  {
				Expect(uvindex).Should(BeZero())
			})
		})
	})

})
