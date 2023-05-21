package main

import "fmt"

type person struct {
	Name string
}

func (p *person) speak() {
	fmt.Println("HELLO")
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		Name: "WAI",
	}

	// this is not work
	//saySomething(&p1)
	saySomething(&p1)

	p1.speak()
}

func saySomething(h human) {
	h.speak()

}
