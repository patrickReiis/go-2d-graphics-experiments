package games

import (
	"fmt"
	"path/filepath"
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FPS        = 60
	MAX_FRAMES = 15
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

func PlayerWalkingWithAnimation() {

	screenWidth := int32(1500)
	screenHeight := int32(900)

	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)

	rl.InitWindow(screenWidth, screenHeight, "Player Walking Animation")

	var players []Player

	_, absolutePath, _, _ := runtime.Caller(1)
	currentDir := filepath.Dir(absolutePath)
	animationFile := currentDir + "/resources/animation/scarfy.png"

	scarfy := rl.LoadTexture(animationFile)

	for i := 0; i < 10; i++ {
		playerBot := Player{}
		playerBot.Position.X = 100
		playerBot.Position.Y = 100
		playerBot.frameRect = rl.Rectangle{X: 0, Y: 0, Width: float32(scarfy.Width / 6), Height: float32(scarfy.Height)}
		players = append(players, playerBot)
	}

	secondRect := rl.Rectangle{X: 0, Y: 0, Width: float32(scarfy.Width/6) * 2, Height: float32(scarfy.Height) * 2}
	fmt.Println(secondRect)

	currentFrame := 0
	frameCounter := 0

	rl.SetTargetFPS(FPS)

	for rl.WindowShouldClose() == false {
		if rl.IsKeyDown(rl.KeyRight) {
			for i := range players {
				player := &players[i]
				player.Position.X += float32(5 + (i))
			}
			frameCounter++
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			for i := range players {
				player := &players[i]
				player.Position.X -= float32(5 + (i))
			}
			frameCounter++
		}

		if rl.IsKeyDown(rl.KeyDown) {
			for i := range players {
				player := &players[i]
				player.Position.Y += float32(5 + (i * 2))
			}
		}

		if rl.IsKeyDown(rl.KeyUp) {
			for i := range players {
				player := &players[i]
				player.Position.Y -= float32(5 + (i * 2))
			}
		}

		if rl.IsKeyDown(rl.KeyLeft) == false && rl.IsKeyDown(rl.KeyRight) == false {
			frameCounter = 0
		}

		if frameCounter > FPS/MAX_FRAMES {
			currentFrame++
			frameCounter = 0
			if currentFrame > 5 {
				currentFrame = 0
			}
			for i := range players {
				player := &players[i]
				player.frameRect.X = player.frameRect.Width * float32(currentFrame)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for _, player := range players {
			//rl.DrawTexturePro(scarfy, player.frameRect, rl.Rectangle{Width: secondRect.Width, Height: secondRect.Height, X: player.Position.X, Y: player.Position.Y}, rl.Vector2{}, 0, rl.White)
			rl.DrawTextureRec(scarfy, player.frameRect, player.Position, rl.White)
			fmt.Println(player.Position)
		}

		rl.EndDrawing()
	}

	rl.UnloadTexture(scarfy)
	rl.CloseWindow()
}
