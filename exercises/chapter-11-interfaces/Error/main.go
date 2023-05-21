package main

import (
	"fmt"
	"log"
)

type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
}

// OverHeatError Start Here
type OverHeatError float64

func (o OverHeatError) Error() string { // "Error() string" is the method that satisfied the "error" Interface
	return fmt.Sprintf("Overheating by %0.2f degrees", o)
}

func CheckTemperature(actual, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverHeatError(excess)
	}
	return nil
}

func main() {
	var reel error // error = error interface
	reel = ComedyError("What's a programmer's favourite beer?, Logger!")
	fmt.Println(reel)

	reel = OverHeatError(4.5)
	fmt.Println(reel)

	reel = CheckTemperature(134.2, 100.0)
	if reel != nil {
		log.Fatal(reel)
	}
}
