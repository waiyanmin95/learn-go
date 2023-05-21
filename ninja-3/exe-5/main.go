package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("The Value is %v divided by 4 and the remainder is: %v\n", i, i%4)
	}
}
