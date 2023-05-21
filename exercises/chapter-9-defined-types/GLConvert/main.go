package main

import "fmt"

type Gallons float64
type Liters float64
type Milli float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milli) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%.3f Liters equal to %.3f Gallons\n", soda, soda.ToGallons())
	beer := Milli(4809)
	fmt.Printf("%.3f Milli equal to %.3f Gallons\n", beer, beer.ToGallons())
}
