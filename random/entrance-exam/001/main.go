package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(paried(num))
}
func paried(xi []int) int {
	for i := 0; i < len(xi); i++ {
		if i%2 == 0 && xi[i]%2 == 0 {
			return 0
		}
		if i%2 == 1 && xi[i]%2 == 1 {
			return 0
		}
	}
	return 1
}
