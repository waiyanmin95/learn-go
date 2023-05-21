package main

import "fmt"

type human struct {
}

type cat struct {
}

type dog struct {
}

type devops struct {
}

type Animal interface {
	Speak() string
}

func main() {
	animals := []Animal{dog{}, cat{}, human{}, devops{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
func (c cat) Speak() string {
	return "MEOW!"
}

func (d dog) Speak() string {
	return "Wof!"
}

func (h human) Speak() string {
	return "HELLO"
}

func (dv devops) Speak() string {
	return "I've no idea what i'm doing"
}
