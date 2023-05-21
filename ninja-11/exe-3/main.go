package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is custom error %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "BLAHBLAHBLAH",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")
}
