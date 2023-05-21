package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)
	fmt.Println("GOROUTINEs:", runtime.NumGoroutine())

	for v := range fanin {
		fmt.Println(v)
		fmt.Println("GOROUTINEs:", runtime.NumGoroutine())
	}

	fmt.Println("ABOUT to exit")
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
			//fmt.Println("EVEN CHANNEL: ", even)
		} else {
			odd <- i
			//fmt.Println("ODD CHANNEL: ", odd)
		}
	}
	close(even)
	close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
