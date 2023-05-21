package main

import "fmt"

func main() {
	x := 11.3
	fmt.Printf("The Square Value: %#v", Square(&x))
}

func Square(x *float64) float64 {
	*x = *x * *x
	return *x
}
