package structs


import(
  "github.com/kelvins/geocoder"
)

type Address struct{
  Country string
  City string
  PostalCode string
  State string
}

type YourGeoStationPoint interface{
  CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)
}


func (*Address)CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)  {

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
