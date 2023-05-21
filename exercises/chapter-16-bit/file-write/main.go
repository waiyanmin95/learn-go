package main

import (
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//func main() {
//	file, err := os.OpenFile("program.log", os.O_WRONLY, os.FileMode(0600))
//	check(err)
//	_, err = file.Write([]byte("this is shit"))
//	check(err)
//	err = file.Close()
//	check(err)
//}

func main() {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("program.log", options, os.FileMode(0644))
	check(err)
	_, err = file.Write([]byte("this is good\n"))
	check(err)
	err = file.Close()
	check(err)
}
