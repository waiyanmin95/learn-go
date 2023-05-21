package myinterface

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(f float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with:", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hello from MethodWithReturnValue"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
