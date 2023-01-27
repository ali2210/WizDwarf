package geo

// Our globe consist of both latitude and longitude. Point is an actual location that will represent in vector
type Point struct {
	Latituide_Division string // Latituide
	Longitude_Division string // Longitude
}

// @param string , @ return Point object
// @@ this function act as goroutine. This function transform value into vector. In most cases vector is a good choice against complex processing.
// l = [25.820, 25.820]
func Location(str string) Point {

	current_nav := make(chan Point)
	go func() {
		current_nav <- Point{Longitude_Division: str[0:5], Latituide_Division: str[13:18]}
	}()
	location := <-current_nav
	return location
}
