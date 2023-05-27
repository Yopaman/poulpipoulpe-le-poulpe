package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "test")

	tileset := rl.LoadTexture("tileset.png")

	tile := rl.NewRectangle(4, 8, 8, 8)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Beige)

		rl.DrawTextureRec(tileset, tile, rl.NewVector2(50, 50), rl.White)

		rl.EndDrawing()
	}

	rl.UnloadTexture(tileset)

	rl.CloseWindow()
}
