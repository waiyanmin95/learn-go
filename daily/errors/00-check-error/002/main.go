package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string

	fmt.Println("Name :\t")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Age :\t")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Food :\t")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", ans2)
	fmt.Println(ans1, ans2, ans3)
}
