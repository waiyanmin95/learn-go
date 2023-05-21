package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	r := strings.NewReader("James Bonddd")

	_, _ = io.Copy(f, r)
}
