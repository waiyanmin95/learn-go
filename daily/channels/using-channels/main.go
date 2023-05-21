package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

// send
func send(c chan<- int) {
	c <- 74
}

// receive
func receive(c <-chan int) {
	fmt.Println("the value received from the channel:", <-c)
}
