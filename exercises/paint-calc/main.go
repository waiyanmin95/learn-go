package main

import (
	"fmt"
	"log"
)

func main() {
	amount, err := Paint(4.9, 3.4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", amount)
}

func Paint(w, h float64) (float64, error) {
	if w < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", w)
	}
	if h < 0 {
		return 0, fmt.Errorf("a high of %.2f is invalid", h)
	}

	area := w * h

	liter := area / 10.0

	return liter, nil
}
