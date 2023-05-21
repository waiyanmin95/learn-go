package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
	fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
	fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
	fmt.Printf("%09b %s\n", 0755, os.FileMode(0755))
	fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))
}
