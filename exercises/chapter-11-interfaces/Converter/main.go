package main

import "fmt"

type (
	Gallons     float64
	Liters      float64
	Milliliters float64
)

func (g Gallons) String() string {
	return fmt.Sprintf("%.2f gallons", g)
}

func (l Liters) String() string {
	return fmt.Sprintf("%.2f liters", l)
}
func (m Milliliters) String() string {
	return fmt.Sprintf("%.2f milliliters", m)
}

func main() {
	fmt.Println(Gallons(33.2))
	fmt.Println(Liters(98.44))
	fmt.Println(Milliliters(88.256))
}
