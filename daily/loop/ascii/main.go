package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Value: %v\t HEXA: %#x\t Unicode Full: %#U\t Unicode Character: %c\n", i, i, i, i)
	}
	fmt.Println("Runtime ARCHITECTURE: ", runtime.GOARCH, "\tRuntime OS: ", runtime.GOOS)
}
