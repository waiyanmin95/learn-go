package main

import "fmt"

func main() {
	defer CalmDown()
	//FreakOut()
	//fmt.Println("Exiting Normally")
	err := fmt.Errorf("there's an error")
	panic(err)
}

func CalmDown() {
	//fmt.Println(recover()) // fmt.Println call "recover" and return nil
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

//func FreakOut() {
//	defer CalmDown()
//	panic("oh no")
//}
