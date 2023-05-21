package gadget

import "fmt"

type Player interface {
	Play(s string)
	Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecoder struct {
	Microphone int
}

func (t TapeRecoder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecoder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecoder) Stop() {
	fmt.Println("Stopped")
}
