package multiplayer

import (
	"fmt"
	"log"
	"net"
)

func EstablishConn() {
	udpServer, err := net.ResolveUDPAddr("udp", ":1053")

	if err != nil {
		log.Fatalf("ResolveUDPAddr failed: %s", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Fatalf("Listen failed: %s", err.Error())
	}
	// close the connection
	//defer conn.Close()

	_, err = conn.Write([]byte("hello from the other side of the world"))
	if err != nil {
		log.Fatalf("Write data failed: %s", err.Error())
	}

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		log.Fatalf("Read data failed: %s", err.Error())
	}

	fmt.Println(string(received))
}
