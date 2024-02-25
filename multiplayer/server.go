package multiplayer

import (
	"fmt"
	"log"
	"net"
)

func ListenForConnections() {
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	var players []string
	fmt.Println(players)

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Print(err)
		}

		go udpResponse(udpServer, addr, buf)

		fmt.Println(addr)
		fmt.Println("nothing...")
	}
}

func udpResponse(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	if string(buf)[:2] == "id" {
	}

	responseStr := fmt.Sprintf("%v", string(buf))

	_, err := udpServer.WriteTo([]byte("hello from server, first"), addr)
	if err != nil {
		fmt.Println(err)
	}

	_, err = udpServer.WriteTo([]byte("hello from server, second"), addr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("server: " + responseStr)
}
