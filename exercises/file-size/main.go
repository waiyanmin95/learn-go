package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File size in byte:", fileinfo.Size(), "Bytes")
}
