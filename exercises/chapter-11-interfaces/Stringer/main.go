package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	brand := CoffeePot("Premier")
	fmt.Println(brand.String())
	fmt.Print(brand, "\n")
	fmt.Printf("%s", brand)
}
