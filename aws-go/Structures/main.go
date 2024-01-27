package main

import "fmt"

type SNS struct {
	Records []Entry `json:"Records"`
}

type Entry struct {
	EventSource  string `json:"EventSource"`
	EventVersion string `json:"EventVersion"`
	Sns          struct {
		Message string `json:"Message"`
	}
}

func main() {
	var message SNS
	var entry Entry

	entry.EventSource = "aws:sns"
	entry.EventVersion = "1.0.0"
	entry.Sns.Message = "Hello from SNS"

	message.Records = make([]Entry, 2)
	message.Records[0] = entry
	v2 := entry
	v2.EventVersion = "2.0.0-rc1"
	v2.Sns.Message = "This is v2"
	message.Records[1] = v2

	fmt.Println(entry)

	fmt.Println(message)

}
