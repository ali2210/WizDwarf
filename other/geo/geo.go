package geo

type Point struct {
	Latituide_Division string
	Longitude_Division string
}

func Location(str string) Point {
	current_nav := make(chan Point)
	go func() {
		current_nav <- Point{Longitude_Division: str[0:5], Latituide_Division: str[13:18]}
	}()
	location := <-current_nav
	return location
}
