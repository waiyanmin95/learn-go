package main

import (
	"github.com/ivan-amity/learn-go/exercises/chapter-11-interfaces/Tape/gadget"
)

func PlayList(device gadget.Player, songs []string) {
	for _, v := range songs {
		device.Play(v)
	}
	//device.Record()
	device.Stop()
}

func main() {
	mixtapes := []string{"I NEED YOU", "Shape of you", "Waka waka"}
	var player gadget.Player = gadget.TapePlayer{}
	PlayList(player, mixtapes)

	player = gadget.TapeRecoder{}
	PlayList(player, mixtapes)

}
