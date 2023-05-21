package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // read OS STDIN with new line
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // start and end remove whitespaces
	grade, err := strconv.ParseFloat(input, 64) // convert string to float64
	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade > 100 {
		status = "You're Duu Yinn Thee"
	} else if grade == 100 {
		status = "Perfect."
	} else if grade >= 60 {
		status = "Passed and you did great."
	} else {
		status = "Failed."
	}

	fmt.Println("The grade of", grade, "is", status)
}
