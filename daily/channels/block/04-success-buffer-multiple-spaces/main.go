package main

import "fmt"

func main() {
	gg := make(chan int, 3)

	gg <- 99
	gg <- 11
	gg <- 222

	fmt.Println(<-gg)
	fmt.Println(<-gg)
	fmt.Println(<-gg)
}
