package main

import "fmt"

func main() {
	a1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Wai Yan",
		last:  "Min",
		age:   27,
	}

	fmt.Println(a1)
	fmt.Println(a1.age)
}
