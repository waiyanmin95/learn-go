package main

import "fmt"

type Part struct {
	info  string
	count int
}

func main() {
	var x Part
	x.info = "Hex Bolts"
	x.count = 22
	showInfo(x)

	gg := Order("Glitch", 101)
	fmt.Println(gg.info, gg.count)
}

func showInfo(p Part) {
	fmt.Println("Information:", p.info)
	fmt.Println("Count:", p.count)
}

func Order(i string, c int) Part {
	var y Part
	y.info = i
	y.count = c
	return y
}
