package main

import "fmt"

type (
	Car   string
	Truck string
)

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(s string)
}

func (c Car) Accelerate() {
	fmt.Println("Speeding Up")
}

func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turing ", direction)
}

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Stopping")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turing ", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

func TryVehicle(v Vehicle) {
	v.Steer("Left")
	v.Accelerate()
	v.Steer("Right")
	v.Brake()
	tt, ok := v.(Truck)
	if ok {
		tt.LoadCargo("Boxes")
	}
}
func main() {
	TryVehicle(Truck("BMW"))
}
