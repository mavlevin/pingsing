package main

import (
	"bytes"
	"io"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"
)

/*
to build:
	go-bindata resources/ && go build && pingsing
*/

var player *oto.Player = nil
var pongDecoder *mp3.Decoder = nil
var pingDecoder *mp3.Decoder = nil

func playPong() {
	pongDecoder.Seek(0, 0)
	if _, err := io.Copy(player, pongDecoder); err != nil {
		panic(err)
	}
}

func playPing() {
	pingDecoder.Seek(0, 0)
	if _, err := io.Copy(player, pingDecoder); err != nil {
		panic(err)
	}
}

func initSound() {
	context, err := oto.NewContext(44100, 2, 2, 8192)
	if err != nil {
		panic(err)
	}
	player = context.NewPlayer()

	pongMP3, err := Asset("resources/pong_concise.mp3")
	if err != nil {
		panic(err)
	}
	pongDecoder, err = mp3.NewDecoder(bytes.NewReader(pongMP3))

	pingMP3, err := Asset("resources/ping_concise.mp3")
	if err != nil {
		panic(err)
	}
	pingDecoder, err = mp3.NewDecoder(bytes.NewReader(pingMP3))
}

func main() {
	initSound()

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			playPing()
		} else {
			playPong()
		}
	}
}
