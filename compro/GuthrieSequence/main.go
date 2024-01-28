package main

import "fmt"

func IsGuthrieSequence(n int) int {
	if n < 0 {
		return 0
	}
	var xi []int
	var temp = n
	for temp != 1 {
		if temp%2 == 0 {
			temp = temp / 2
			xi = append(xi, temp)
		} else {
			temp = (temp * 3) + 1
			xi = append(xi, temp)
		}
	}
	fmt.Println("Slice: ", xi)
	return len(xi)
}

func main() {
	fmt.Println(IsGuthrieSequence(7))
}
