package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("mydir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Mode:", fileInfo.Mode())
}
