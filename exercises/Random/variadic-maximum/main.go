package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Maximum Number is:", max(-4.5, -6.7, -8.8, -9.2))
}

func max(n ...float64) float64 {
	max := math.Inf(-1)
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}
