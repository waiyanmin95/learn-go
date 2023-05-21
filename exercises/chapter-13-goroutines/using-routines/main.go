package main

import (
	"fmt"
	"time"
)

func one() {
	for i := 0; i < 50; i++ {
		fmt.Print("A")
	}
}

func two() {
	for i := 0; i < 10; i++ {
		fmt.Print("B")
	}
}

func main() {
	go one()
	go two()
	time.Sleep(time.Millisecond)
	fmt.Println("END HERE main()")
}
