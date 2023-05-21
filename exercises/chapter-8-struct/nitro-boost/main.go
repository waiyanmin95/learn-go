package main

import "fmt"

type car struct {
	Name     string
	TopSpeed float64
}

func main() {
	var toyota car
	toyota.Name = ""
	toyota.TopSpeed = 443
	NitroBoost(&toyota)
	fmt.Println("Car Name:", toyota.Name)
	fmt.Println("Top Speed:", toyota.TopSpeed)
}

func NitroBoost(c *car) {
	c.TopSpeed += 50
}
