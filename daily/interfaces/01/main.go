package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("You can call me", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am secret agent", sa.first)
}

type human interface {
	speak()
}

func SaySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Tom",
	}
	p2 := person{
		first: "Jerry",
	}

	sa := secretAgent{
		person: person{
			first: "James",
		},
		ltk: false,
	}

	fmt.Println(p1.first, p2.first)
	p1.speak()
	p2.speak()

	SaySomething(p1)
	SaySomething(p2)
	SaySomething(sa)
}
