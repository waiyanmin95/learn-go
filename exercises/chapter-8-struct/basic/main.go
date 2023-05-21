package main

import "fmt"

type basic struct {
	number float64
	word   string
	toggle bool
}

func main() {
	ssx := basic{
		number: 5.4,
		word:   "hello",
		toggle: true,
	}
	fmt.Printf("%v %v %#v\n", ssx.word, ssx.number, ssx.toggle)
}
