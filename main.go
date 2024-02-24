package main

import (
	"time"

	"github.com/patrickReiis/go-2d-graphics-experiments/multiplayer"
)

func main() {

	go multiplayer.ListenForConnections()

	time.Sleep(2 * time.Second)

	go multiplayer.EstablishConn()

	select {}

	//games.PlayerWalkingWithAnimation()
}
