package main

import "fmt"

func main() {
	foo() // pass in arguments
	bar("SHEET")
	w1 := woo("BLUU HOO")
	fmt.Println(w1)
	x1, x2 := boo("WAI YAN", "MIN")
	fmt.Println(x2, x1)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("HELLO FROM FOO")
}

func bar(s string) {
	fmt.Println("HELLO FROM BAR, ", s)
}

func woo(s string) string {
	return fmt.Sprint("HELLO FROM WOO, ", s)
}

func boo(fn, ln string) (string, int) {
	a := fmt.Sprint(fn, " ", ln, " says HELLO")
	b := 33
	return a, b
}
