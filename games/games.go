package games

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayerWalkingWithAnimation() {
	screenWidth := int32(800)
	screenHeight := int32(800)

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(screenWidth, screenHeight, "Player Walking Animation")

	player := rl.NewRectangle(400, 400, 50, 50)

	rl.SetTargetFPS(60)

	for rl.WindowShouldClose() == false {
		if rl.IsKeyDown(rl.KeyRight) {
			player.X += 5
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			player.X -= 5
		}

		if rl.IsKeyDown(rl.KeyDown) {
			player.Y += 5
		}

		if rl.IsKeyDown(rl.KeyUp) {
			player.Y -= 5
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleRec(player, rl.Blue)

		rl.EndDrawing()
	}
}
