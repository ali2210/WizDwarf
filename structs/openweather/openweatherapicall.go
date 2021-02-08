package openweather

import (
	"log"

	openmap "github.com/briandowns/openweathermap"
	// geocoder "github.com/kelvins/geocoder"
)

type DataVisualization struct {
	Percentage float32
	Process    int
	SeenBy     string

	// additionalParameters
	coordinates openmap.Coordinates
	UVinfo      []openmap.UVIndexInfo
}

type MyCoordinates struct {
	Longitude float64
	Latitude  float64
}

type OpenWeatherApi interface {
	OpenWeatherApiCall(apikey string) (*openmap.UV, error)
	GetCoordinates(loc *MyCoordinates) openmap.Coordinates
	UVCoodinates(c openmap.Coordinates, u *openmap.UV) error
	UVCompleteInfo(u *openmap.UV) ([]openmap.UVIndexInfo, error)
	PrintLogs()
}

func (*DataVisualization) OpenWeather(apikey string) (*openmap.UV, error) {

	weather, err := openmap.NewUV(apikey)
	return weather, err
}

func (*DataVisualization) GetCoordinates(loc *MyCoordinates) *openmap.Coordinates {
	coo := &openmap.Coordinates{
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
	}
	return coo
}

func (*DataVisualization) UVCoodinates(c *openmap.Coordinates, u *openmap.UV) error {

	err := u.Current(c)
	return err

}

func (*DataVisualization) UVCompleteInfo(u *openmap.UV) ([]openmap.UVIndexInfo, error) {

	i, err := u.UVInformation()
	return i, err
}

func (d *DataVisualization) PrintLogs() {
	log.Println("[Logs]... ", *d)
}
