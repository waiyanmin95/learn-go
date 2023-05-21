package main

import "fmt"

func main() {

	// channel block
	c := make(chan int)

	c <- 42
	fmt.Println(<-c)
}
