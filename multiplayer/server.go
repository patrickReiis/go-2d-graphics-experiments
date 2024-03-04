package multiplayer

import (
	"encoding/json"
	"log"
	"net"

	"github.com/patrickReiis/go-2d-graphics-experiments/games"
)

var Players []games.PlayerJson

func ListenForConnections() {
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Print(err)
		}

		go udpResponse(udpServer, addr, buf[:n])

	}
}

func udpResponse(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	if string(buf) == "[]" {
		err := sendAllPlayersToClient(udpServer, addr)
		if err != nil {
			log.Print(err)
		}
		return
	}

	var player games.PlayerJson
	err := json.Unmarshal(buf, &player)
	if err != nil {
		log.Fatal(err)
	}

	if isPlayerPresent(Players, player) == true {
		updatePlayerData(Players, player)

		err := sendAllPlayersToClient(udpServer, addr)
		if err != nil {
			log.Print(err)
		}

		return
	}
	err = sendAllPlayersToClient(udpServer, addr)
	if err != nil {
		log.Print(err)
	}

	Players = append(Players, player)
}

func updatePlayerData(players []games.PlayerJson, player games.PlayerJson) {
	for i, e := range players {
		if e.Id == player.Id {
			players[i].X = player.X
			players[i].Y = player.Y
		}
	}
}

func isPlayerPresent(players []games.PlayerJson, player games.PlayerJson) bool {
	for _, e := range players {
		if e.Id == player.Id {
			return true
		}
	}

	return false
}

func sendAllPlayersToClient(udpServer net.PacketConn, addr net.Addr) error {
	playersJson, err := json.Marshal(Players)
	if err != nil {
		return err
	}

	_, err = udpServer.WriteTo(playersJson, addr)
	if err != nil {
		return err
	}

	return nil
}
