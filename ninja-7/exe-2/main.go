package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{name: "WYM"}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)
}

func changeme(p *person) {
	p.name = "GUU"
	//(*p).name = "Miss Moneyp"

}
