package main

import (
	"fmt"
	"runtime"
)

func greet(g chan string) {
	g <- "Hello World"
	g <- "WOW"
}

func main() {
	xx := make(chan string)
	go greet(xx)
	fmt.Println(<-xx)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(<-xx)
}
