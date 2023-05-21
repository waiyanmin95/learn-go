// Package guu asks what she wanted to eat
package guu

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	y := 0
	for _, v := range xi {
		y += v
	}
	return y
}
