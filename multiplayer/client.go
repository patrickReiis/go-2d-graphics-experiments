package multiplayer

import (
	"log"
	"net"

	"github.com/patrickReiis/go-2d-graphics-experiments/games"
)

func EstablishConn(playerId int) {
	udpServer, err := net.ResolveUDPAddr("udp", ":1053")

	if err != nil {
		log.Fatalf("ResolveUDPAddr failed: %s", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Fatalf("Listen failed: %s", err.Error())
	}
	// close the connection
	defer conn.Close()

	games.PlayerWalkingWithAnimation(conn, playerId)
}
