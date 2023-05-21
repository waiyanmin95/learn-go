package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
	fmt.Println("Opening Refrigerator")
}

func (r Refrigerator) Close() {
	fmt.Println("Closing Refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
	r.Open()
	if find(food, r) {
		fmt.Println("Found: ", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	r.Close()
	return nil
}

func main() {
	fridge := Refrigerator{"Milk", "Coffee", "Pizza"}
	for _, food := range []string{"Milk", "Butter"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
