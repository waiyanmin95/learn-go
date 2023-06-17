package main

import "fmt"

func factorial(n int) int {
	fmt.Println("The N value: ", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factLoop(n int) int {
	count := 1
	for n > 0 {
		count *= n
		n--
	}
	return count
}

func main() {
	fmt.Println("Factorial Function", factorial(4))
	fmt.Println("Factorial Loop", factLoop(4))
}
