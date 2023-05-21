package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range num {
		if v%3 == 0 || v%5 == 0 {
			fmt.Println("FizzBuzz =>", v)
		} else {
			fmt.Println(v)
		}
	}
}
