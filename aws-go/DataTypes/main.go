package main

import "fmt"

type Super struct {
	name string
	age  int
}

func main() {
	xi := [...]string{"apple", "orange", "grapes"}
	for i := 0; i < len(xi); i++ {
		fmt.Println("Index", i, xi[i])
	}

	xxi := []int{1, 2, 3, 4, 5}
	fmt.Printf("Type %T\t Value: %v\n", xxi, xxi)

	mx := make(map[string]string)
	mx["A"] = "Apple"
	mx["B"] = "Ball"
	mx["C"] = "Cat"
	mx["D"] = "Doll"
	mx["F"] = "Football"

	something, ok := mx["Z"]
	fmt.Println(ok)
	if ok {
		fmt.Println(something, "is here")
	} else {
		fmt.Println("Not here!")
	}
	for i := range mx {
		fmt.Println(i)
	}

	p1 := Super{"Superman", 28}

	p2 := p1
	p1.name = "Batman"
	p2.age = 99
	fmt.Println("Bye Bye", p1.name, p1.age)
	fmt.Println(p1)
	fmt.Println(p2)

}
