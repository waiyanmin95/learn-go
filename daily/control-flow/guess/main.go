package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number int
	random := rand.Intn(10) + 1

	fmt.Println("Random Number: ", random)

	found := false
	attempt := 0
	for !found {
		fmt.Printf("Enter your guess number: ")
		fmt.Scanln(&number)
		guess := number
		attempt++
		if guess == random {
			found = true
			fmt.Println("You guess it correctly at attempt: ", attempt)
		} else if guess > random {
			fmt.Println("Your guess is too high")
		} else {
			fmt.Println("Your guess is too low")
		}
	}
}
