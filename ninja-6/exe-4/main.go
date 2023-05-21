package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	gg := person{
		first: "WAI YAN",
		last:  "MIN",
		age:   27,
	}
	gg.speak()
}

func (r person) speak() {
	fmt.Println("I'm", r.first, r.last, r.age)
}
