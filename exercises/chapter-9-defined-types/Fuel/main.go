package main

import "fmt"

type (
	Liters  float64
	Gallons float64
)

func main() {
	carFuel := Gallons(44.5)
	busFuel := Liters(882.9)

	//carFuel = 44.5
	//busFuel = 22.9

	fmt.Printf("Types: Car > %T Bus > %T\n", carFuel, busFuel)
	fmt.Println("Car:", carFuel)
	fmt.Println("Bus:", busFuel)

	fmt.Printf("Gallons to Liters: %.3f\n", GConvert(33.375))
	fmt.Printf("Liters to Gallons: %.3f\n", LConvert(126.42))
}

func GConvert(g Gallons) Liters {
	v := Liters(Gallons(g) * 3.785)
	return v
}

func LConvert(l Liters) Gallons {
	v := Gallons(Liters(l) * 0.264)
	return v
}
