package main

import "fmt"

func main() {
	amount := 7
	// assign the address "&" of amount into total
	total := &amount

	// print out the address of amount
	fmt.Println("Print the address of the amount:", &amount)

	fmt.Println("Execute the Double Function:", Double(&amount))

	// assign a value into amount
	amount = 8
	// point "*" to and give you the value from the address "&"
	fmt.Println("Print the value from the address with pointer:", *total)
	fmt.Println("Print the address of the amount ( After ):", &amount)
	fmt.Println("Recheck the value from amount:", amount)
}

func Double(n *int) int {
	*n *= 2
	return *n
}
