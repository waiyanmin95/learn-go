package main

import "fmt"

func main() {
	c := make(chan int)    // channel ( Bidirectional )
	cs := make(chan<- int) //send channel
	cr := make(<-chan int) // receive channel

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)
}
