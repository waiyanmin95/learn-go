package main

import "fmt"

func main() {
	truth := true
	Negate(&truth)
	fmt.Println("Truth:", truth)

	lie := false
	Negate(&lie)
	fmt.Println("Lie:", lie)
}

func Negate(b *bool) {
	*b = !*b
}
