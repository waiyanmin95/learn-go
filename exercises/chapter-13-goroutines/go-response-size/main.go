package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://go.dev", "https://go.dev/doc/effective_go", "https://amity.co"}
	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d %v\n", page.URL, page.Size, "bytes")
	}
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting URL", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{
		URL:  url,
		Size: len(body),
	}
}
