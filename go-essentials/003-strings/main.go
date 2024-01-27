package main

import "fmt"

func main() {
	book := "The colour of magic"
	fmt.Println(book)
	fmt.Println(len(book))

	fmt.Println(book[2])
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
}
