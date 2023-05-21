package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
	fmt.Println(runtime.GOARCH)
}
