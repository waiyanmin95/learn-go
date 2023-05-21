// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatalf("can't calculate more than two numbers")
	}

	no1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("first input should be number: %w", err)
	}
	//operator := os.Args[2]
	no2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("second input should be number: %w", err)
	}

	switch os.Args[2] {
	case "+":
		fmt.Println(no1 + no2)
	case "-":
		fmt.Println(no1 - no2)
	case "x":
		fmt.Println(no1 * no2)
	case "/":
		if no2 == 0 {
			log.Fatalln("error: you tried to divide by zero")
		}
		fmt.Println(no1 / no2)
	default:
		fmt.Println("Invalid Operator")
		os.Exit(1)
	}
}
