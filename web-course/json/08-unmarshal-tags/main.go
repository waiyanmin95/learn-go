package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	A string  `json:"Postal"`
	B float64 `json:"Latitude"`
	C float64 `json:"Longitude"`
	D string  `json:"Address"`
	E string  `json:"City"`
	F string  `json:"State"`
	G string  `json:"Zip"`
	H string  `json:"Country"`
}

type cities []city

func main() {
	var data cities
	xx := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(xx), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling", err)
	}

	fmt.Println(data)

	fmt.Printf("%T\n", data)
	fmt.Println(data[0].E, data[1].H)
}
