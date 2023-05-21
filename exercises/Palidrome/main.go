package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a word: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // read OS STDIN with new line
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // start and end remove whitespaces

	if input == "" {

	}
	p := strings.Split(input, "")

	//reverse := []string{}
	//for i := len(p) - 1; i >= 0; i-- {
	//	reverse = append(reverse, p[i])
	//}
	pali := true
	i := len(p) - 1

	for j := 0; j < len(p)-1; j++ {
		if p[i] == p[j] {
			pali = true
		} else {
			pali = false
		}
		i--
	}

	if pali {
		fmt.Println("This is Palidrome.")
		return
	}
	fmt.Println("This is not Palidrome.")
}
