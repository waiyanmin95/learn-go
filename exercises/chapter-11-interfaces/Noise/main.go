package main

import "fmt"

type (
	Whistle string
	Horn    string
	Robot   string
)

type NoiseMaker interface {
	MakeSound()
}

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!", w)
}

func (w Horn) MakeSound() {
	fmt.Println("Honk!", w)
}

func (r Robot) MakeSound() {
	fmt.Println("Beep Bop", r)
}

func (r Robot) Walk() {
	fmt.Println("I'm walking", r)
}

func main() {
	var toy NoiseMaker = Robot("Bumblebee")
	//var toy NoiseMaker
	//toy = Whistle("Tee Tuu")
	//toy.MakeSound()
	//toy = Horn("Tee Tee Ti")
	//toy.MakeSound()
	//toy = Robot("Transformer")
	//toy.MakeSound()
	//toy = Robot("Transformer")
	toy.MakeSound()
	var robot = toy.(Robot)
	robot.Walk()
}
