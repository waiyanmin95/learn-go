package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 88
	c <- 9999

	fmt.Println(<-c)
}
