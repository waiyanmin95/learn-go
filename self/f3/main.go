package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 2
	var myFloat float64 = 3.74

	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(myFloat))
	fmt.Println(reflect.TypeOf(float64(myInt)))

	myInt = int(myFloat)
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(myInt)

}
