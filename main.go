package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/patrickReiis/go-2d-graphics-experiments/multiplayer"
)

func main() {

	var playerId int
	flag.IntVar(&playerId, "id", 1, "set the player id")

	client := flag.Bool("client", false, "runs a new the client")
	server := flag.Bool("server", false, "creates an UDP server")

	flag.Parse()

	if *client == true && *server == true {
		log.Fatal("You can't run both a client and server")
	}

	if *client == true {
		fmt.Println("Running a new client")
		multiplayer.EstablishConn(playerId)
		return
	}

	if *server == true {
		fmt.Println("Server listening...")
		multiplayer.ListenForConnections()
		return
	}

	fmt.Println("You need to specify either a 'server' or 'client' flag.")
}
