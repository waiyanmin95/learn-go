package main

import "fmt"

func main() {
	fmt.Println("Is it a prime number?: ", isPrime(99))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
