package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	//log.Printf("Start time: %s", start)

	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {

		fmt.Println(v)
	}

	//time.Sleep(1 * time.Millisecond)

	elapsed := time.Since(start)

	log.Printf("This program took %s", elapsed)

}
