package main

import "fmt"

type Switch string

type Toggleable interface {
	toggle()
}

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
}

func main() {
	s := Switch("off")
	fmt.Println(s)
	var t Toggleable = &s
	t.toggle()
	fmt.Println(s)
	t.toggle()
	fmt.Println(s)
}
