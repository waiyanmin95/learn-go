package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatalf("Can't create a file %s", err)
	}
	defer f.Close()

	xs := []byte("Hello Gophers!")
	_, err = f.Write(xs)
	if err != nil {
		log.Fatalln("Can't write to the file", err)
	}
}
