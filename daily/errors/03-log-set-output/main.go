package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	log.SetOutput(f) // Println calls Output to print to the standard logger.

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer func(f2 *os.File) {
		err := f2.Close()
		if err != nil {

		}
	}(f2)

	fmt.Println("check the log.txt file in the directory")
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
