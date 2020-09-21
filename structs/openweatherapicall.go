package structs



import(
  openmap "github.com/briandowns/openweathermap"
  geocoder "github.com/kelvins/geocoder"
)



type DataVisualization struct{
  Percentage float32
  UVinfo []openmap.UVIndexInfo
}


type OpenWeatherApi interface{
  OpenWeather(apikey string)(*openmap.UV,error)
  GetCoordinates(loc geocoder.Location)(*openmap.Coordinates)
  // getPollutionState(api string)(*openmap.Pollution,error)
}

func (*DataVisualization)OpenWeather(apikey string)(*openmap.UV,error)  {
  weather , err := openmap.NewUV(apikey); if err != nil {
    return nil , err
  }
  return weather, nil
}

func (*DataVisualization)GetCoordinates(loc geocoder.Location)(*openmap.Coordinates)  {
  coo := &openmap.Coordinates{
    Longitude : loc.Longitude,
    Latitude : loc.Latitude,
  }
  return coo
}

// func getPollutionState(api string)(*openmap.Pollution,error)  {
//   Poll ,err := openmap.NewPollution(api); if err != nil {
//     return nil , err
//   }
//   return nil, Poll
// }
