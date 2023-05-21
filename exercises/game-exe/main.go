package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	score := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)

	success := false
	for g := 10; g > 0; g-- {
		fmt.Println("You have", g, "guesses left.")

		fmt.Print("Guess a score: ")
		gs, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		gs = strings.TrimSpace(gs) // start and end remove whitespaces

		c, err := strconv.Atoi(gs)
		if err != nil {
			log.Fatal(err)
		}

		if c == score {
			success = true
			fmt.Println("Good Job!, You guessed it!.")
			break
		} else if c < score {
			fmt.Println("Oops. Your guess was low.")
		} else {
			fmt.Println("Oops. Your guess was high.")
		}
	}

	if success != true {
		fmt.Println("Sorry, You didn't guess the score. The Score was:", score)
	}
}
