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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := make(map[string]string)
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		data[strings.Split(scanner.Text(), " ")[0]] = strings.Split(scanner.Text(), " ")[1]
	}
	for scanner.Scan() {
		if data[scanner.Text()] == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", scanner.Text(), data[scanner.Text()])
		}
	}
}
