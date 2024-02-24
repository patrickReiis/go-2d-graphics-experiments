package multiplayer

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ListenForConnections() {
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Print(err)
		}

		go udpResponse(udpServer, addr, buf)

		fmt.Println(addr)
	}
}

func udpResponse(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

	udpServer.WriteTo([]byte(responseStr), addr)
}
