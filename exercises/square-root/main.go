package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	sqr, err := SquareRoot(36)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Squre Root is %.2f", sqr)
}

func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("can't get the square root of negative number")
	}

	return math.Sqrt(n), nil
}
