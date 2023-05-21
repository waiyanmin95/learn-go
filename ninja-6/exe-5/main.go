package main

import (
	"fmt"
	"math"
)

type (
	square struct {
		length float64
	}
	circle struct { //π r 2
		radius float64
	}
)

func main() {
	cir := circle{radius: 3.44}
	info(shape(cir))

}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
