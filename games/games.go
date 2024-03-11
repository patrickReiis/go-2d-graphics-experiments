package games

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"path/filepath"
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FPS        = 60
	MAX_FRAMES = 10
)

type Power struct {
	countdown int
}

type LivingEntity struct {
	life      int
	Position  rl.Vector2
	frameRect rl.Rectangle
	speed     int
}

type Player struct {
	LivingEntity
	Animation
	money int
}

type Animation struct {
	counter      int
	currentFrame int
	framesSpeed  int
}

type PlayerJson struct {
	Id int
	X  int
	Y  int
}

func PlayerWalkingWithAnimation(conn *net.UDPConn, playerId int) {

	screenWidth := int32(1500)
	screenHeight := int32(900)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)

	rl.InitWindow(screenWidth, screenHeight, "Player Walking Animation")

	var players []Player

	_, absolutePath, _, _ := runtime.Caller(1)
	currentDir := filepath.Dir(absolutePath)
	animationFile := currentDir + "/../resources/animation/scarfy.png"

	scarfy := rl.LoadTexture(animationFile)

	player := Player{}
	player.frameRect = rl.Rectangle{X: 0, Y: 0, Width: float32(scarfy.Width / 6), Height: float32(scarfy.Height)}

	otherPlayers := Player{}
	otherPlayers.frameRect = rl.Rectangle{X: 0, Y: 0, Width: float32(scarfy.Width / 6), Height: float32(scarfy.Height)}

	players = append(players, player)

	currentFrame := 0
	frameCounter := 0

	playerJson := PlayerJson{Id: playerId, X: int(player.Position.X), Y: int(player.Position.Y)}
	playerJsoned, err := json.Marshal(playerJson)

	if err != nil {
		log.Fatal(err)
	}

	conn.Write(playerJsoned)

	playersFromServer := []PlayerJson{}

	received := make([]byte, 1024)
	go func() {
		for {

			n, err := conn.Read(received)
			if err != nil {
				log.Print(err)
				break
			}
			playersFromServer = []PlayerJson{}
			err = json.Unmarshal(received[:n], &playersFromServer)
			if err != nil {
				fmt.Println("CASE, trying to unmarshal", string(received))
				log.Fatal(err)
			}
		}

	}()

	rl.SetTargetFPS(FPS)
	for rl.WindowShouldClose() == false {
		_, err := conn.Write([]byte("[]"))
		if err != nil {
			log.Print(err)
			break
		}

		if rl.IsKeyDown(rl.KeyRight) {

			frameCounter++
			player.Position.X += 5

			playerJson := PlayerJson{Id: playerId, X: int(player.Position.X), Y: int(player.Position.Y)}
			playerJsoned, err := json.Marshal(playerJson)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(playerJsoned)
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			frameCounter++
			player.Position.X -= 5

			playerJson := PlayerJson{Id: playerId, X: int(player.Position.X), Y: int(player.Position.Y)}
			playerJsoned, err := json.Marshal(playerJson)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(playerJsoned)
		}

		if rl.IsKeyDown(rl.KeyDown) {
			player.Position.Y += 5
			playerJson := PlayerJson{Id: playerId, X: int(player.Position.X), Y: int(player.Position.Y)}
			playerJsoned, err := json.Marshal(playerJson)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(playerJsoned)
		}

		if rl.IsKeyDown(rl.KeyUp) {
			player.Position.Y -= 5
			playerJson := PlayerJson{Id: playerId, X: int(player.Position.X), Y: int(player.Position.Y)}
			playerJsoned, err := json.Marshal(playerJson)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(playerJsoned)
		}

		if rl.IsKeyDown(rl.KeyLeft) == false && rl.IsKeyDown(rl.KeyRight) == false {
			frameCounter = 0
			currentFrame = 0
			player.frameRect.X = float32(currentFrame)
		}

		if frameCounter > FPS/MAX_FRAMES {
			currentFrame++
			frameCounter = 0
			if currentFrame > 5 {
				currentFrame = 0
			}
			player.frameRect.X = player.frameRect.Width * float32(currentFrame)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for _, p := range playersFromServer {
			if p.Id == playerId {
				rl.DrawTextureRec(scarfy, player.frameRect, rl.Vector2{X: float32(p.X), Y: float32(p.Y)}, rl.White)
			} else {
				rl.DrawTextureRec(scarfy, otherPlayers.frameRect, rl.Vector2{X: float32(p.X), Y: float32(p.Y)}, rl.White)
			}
		}
		rl.EndDrawing()
	}

	rl.UnloadTexture(scarfy)
	rl.CloseWindow()
}
