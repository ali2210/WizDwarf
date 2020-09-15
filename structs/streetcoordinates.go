package structs


import(
  "github.com/kelvins/geocoder"
)

type Address struct{
  Country string
  City string
  PostalCode string
  State string
  Long float64
  Lati float64
}

type YourGeoStationPoint interface{
  CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)
}


func CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)  {

  address := geocoder.Address{
    Country: a.Country,
    City : a.City,
    PostalCode : a.PostalCode,
  }
  loc , err :=  geocoder.Geocoding(address); if err != nil {
    return loc, err
  }
  return loc, nil
}
