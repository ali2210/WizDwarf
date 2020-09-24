package structs


import(
  "github.com/kelvins/geocoder"
  "fmt"
)


type Address struct{
    StreetAddress string
      RouteNum int
      City string
      State string
      PostalCode string
      Country string
}

type YourGeoStationPoint interface{
  CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)
}


func (*Address)CurrentLocationByPostalAddress(a Address)(geocoder.Location, error)  {

  address := geocoder.Address{
    Street : a.StreetAddress,
    Number : a.RouteNum,
    City : a.City,
    State : a.State,
    PostalCode : a.PostalCode,
    Country: a.Country,
  }
  fmt.Println("Address :", address)
  geocoder.ApiKey = "AIzaSyDDCKNvLe1Yo8N31ruebiYuGnmVvhYvgg"
  loc , err :=  geocoder.Geocoding(address); if err != nil {
    return loc, err
  }
  return loc, nil
}
