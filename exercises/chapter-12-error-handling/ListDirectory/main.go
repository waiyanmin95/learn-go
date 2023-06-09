package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func scanDir(path string) {
	fmt.Println("Scanned Path:", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDir(filePath)
		} else {
			fmt.Println(filePath)

		}
	}
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	defer reportPanic()
	scanDir("go")
}
