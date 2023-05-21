package main

import "fmt"

func main() {
	Status("Anna")
	Status("Waiyan")
	Status("Tor")
	Status("Bulba")
}

func Status(s string) {
	grades := map[string]float64{"Anna": 0, "Bulba": 77.3, "Tor": 33.5}
	grade, ok := grades[s]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", s)
	} else if grade < 60 {
		fmt.Printf("%s is failling!\n", s)
	}
}
