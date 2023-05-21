package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type city struct {
	Precision string  `json:"precision"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Address   string  `json:"Address"`
	City      string  `json:"City"`
	State     string  `json:"State"`
	Zip       string  `json:"Zip"`
	Country   string  `json:"Country"`
}

type cities []city

func main() {
	var data cities
	xx := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"SOMETHING","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"NOTHING","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(xx), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling", err)
	}

	fmt.Println(data[0].City, data[1].Country)
}
