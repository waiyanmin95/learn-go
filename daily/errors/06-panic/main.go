package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err) // no timestamp
	}
}

// http://godoc.org/builtin#panic
