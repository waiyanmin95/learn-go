package main

import (
	"fmt"

	"github.com/ivan-amity/learn-go/exercises/chapter-11-interfaces/Tape-Type-Assert/bull"
)

func TryOut(player bull.Player) {
	player.Play("Bullshit")
	player.Stop()
	recorder, ok := player.(bull.TapeRecorder)
	//fmt.Println(ok)
	if ok {
		recorder.Record()
	}
	fmt.Printf("%v\n", bull.TapeRecorder{})
}

func main() {
	TryOut(bull.TapeRecorder{})
	TryOut(bull.TapePlayer{})
}
