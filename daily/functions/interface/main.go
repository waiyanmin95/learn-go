package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (r secretAgent) speak() { //method
	fmt.Println("I am", r.first, r.last, "and", r.age, "years old")
}
func (r person) speak() { //method
	fmt.Println("I am", r.first, r.last, "and", r.age, "years old")
}

type human interface {
	speak()
}

func boo(h human) {
	fmt.Println("This is passed into BOOOO", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			77,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
			99,
		},
		ltk: true,
	}

	p1 := person{
		"Wai Yan",
		"Min",
		27,
	}

	fmt.Println(sa1.first, sa1.last)
	sa1.speak()
	sa2.speak()

	p1.speak()
	boo(sa1)
	boo(sa2)
	boo(p1)
}
