package main

import "fmt"

type part struct {
	info  string
	count int
}

func main() {
	var fuses part
	fuses.info = "Fuses"
	fuses.count = 2
	DoublePack(&fuses)
	fmt.Println("Info:", fuses.info)
	fmt.Println("Counts:", fuses.count)
}

func DoublePack(p *part) {
	p.count *= 3
}
