package main

import (
	"fmt"

	"github.com/ivan-amity/learn-go/exercises/chapter-11-interfaces/Interface/myinterface"
)

func main() {
	var value myinterface.MyInterface
	value = myinterface.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(3.3)
	fmt.Println(value.MethodWithReturnValue())
}
