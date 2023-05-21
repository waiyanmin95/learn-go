package main

import "fmt"

type cars struct {
	color  string
	wheels int
}

func main() {
	c := cars{
		color:  "white",
		wheels: 4,
	}
	fmt.Println(c)
	fmt.Println(c.color, c.wheels)
	changeme(&c)
	fmt.Println(c.color, c.wheels)
}

func changeme(ch *cars) {
	ch.color = "black"
	ch.wheels = 88
}
