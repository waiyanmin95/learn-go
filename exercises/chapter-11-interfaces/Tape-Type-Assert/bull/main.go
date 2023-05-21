package bull

import "fmt"

type Player interface {
	Play(s string)
	Stop()
}

// TapePlayer Start Here
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

// TapePlayer End Here

// TapeRecorder Start Here
type TapeRecorder struct {
	Microphone bool
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
