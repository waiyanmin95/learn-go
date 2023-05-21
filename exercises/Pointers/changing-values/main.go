package main

import "fmt"

func main() {
	n := 4
	myintp := &n // assign to "n"'s memory address
	fmt.Println(myintp)
	fmt.Println(*myintp)
	*myintp = 998834
	fmt.Println("After Address:", myintp)
	fmt.Println("After Pointer:", *myintp)
	fmt.Println(n)
}
