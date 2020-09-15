package structs



import(
  openmap "github.com/briandowns/openweathermap"
)





type OpenWeatherApi interface{
  OpenWeather(units, lang , apikey string)(*openmap.CurrentWeatherData,error)
}

func OpenWeather(units, lang, apikey string)(*openmap.CurrentWeatherData,error)  {
  weather , err := openmap.NewCurrent(units,lang,apikey); if err != nil {
    return nil , err
  }
  return weather, nil
}
