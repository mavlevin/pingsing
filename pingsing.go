package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/hajimehoshi/oto"

	"github.com/hajimehoshi/go-mp3"

	"github.com/tatsushid/go-fastping"
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

func initPinger() *fastping.Pinger {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		panic(err)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		go playPong()
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}

	return p
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: pingsing <host/ip>")
		return
	}

	initSound()
	pinger := initPinger()

	for {
		go playPing()
		err := pinger.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
