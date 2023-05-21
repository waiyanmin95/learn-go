package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("THIS IS GOOS", runtime.GOOS)
	fmt.Println("THIS IS GOARCH", runtime.GOARCH)
	fmt.Println("THIS IS CPUs", runtime.NumCPU())
	fmt.Println("THIS IS GOROUTINES", runtime.NumGoroutine())
	wg.Add(1)
	go foo()

	bar()
	fmt.Println("THIS IS CPUs", runtime.NumCPU())
	fmt.Println("THIS IS GOROUTINES", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}

func bar() {
	for y := 0; y < 10; y++ {
		fmt.Println("Bar", y)
	}
}
