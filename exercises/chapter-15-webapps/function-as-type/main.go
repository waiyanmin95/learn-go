package main

import "fmt"

func doMath(f func(int, int) float64) {
	result := f(8, 2)
	fmt.Println(result)
}

func divide(x, y int) float64 {
	return float64(x) / float64(y)
}

func multiply(x, y int) float64 {
	return float64(x) * float64(y)
}

func main() {
	doMath(divide)
	doMath(multiply)
}
