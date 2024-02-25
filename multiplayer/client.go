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
	//	received := make([]byte, 1024)
	//	_, err = conn.Read(received)
	//	if err != nil {
	//		log.Fatal("error reading, cliet game loop")
	//	}

	games.PlayerWalkingWithAnimation(conn, playerId)

	//	_, err = conn.Write([]byte("hello first"))
	//	if err != nil {
	//		log.Fatalf("Write data failed: %s", err.Error())
	//	}
	//
	//	received := make([]byte, 1024)
	//	_, err = conn.Read(received)
	//	if err != nil {
	//		log.Fatalf("Read data failed: %s", err.Error())
	//	}
	//
	//	fmt.Println("client: " + string(received))

}
