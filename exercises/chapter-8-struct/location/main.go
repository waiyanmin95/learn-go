package main

import (
	"fmt"

	"github.com/ivan-amity/learn-go/exercises/chapter-8-struct/location/geo"
)

func main() {
	location := geo.Coordinates{
		Lat:  37.43,
		Long: -122.08,
	}
	fmt.Println("Latitude:", location.Lat)
	fmt.Println("Longitude:", location.Long)
}
