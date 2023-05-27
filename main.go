package main

import (
	"io/ioutil"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "test")

	camera := rl.Camera2D{}
	camera.Offset = rl.NewVector2(0, 0)
	camera.Rotation = 0.0
	camera.Zoom = 5.0

	rl.SetTargetFPS(60)

	levelContent, err := ioutil.ReadFile("level.txt")
	if err != nil {
		return
	}
	level := ParseWorld(string(levelContent))
	tileset := rl.LoadTexture("tileset.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		drawWorld(level, 20, 20, tileset)
		rl.EndMode2D()
		rl.EndDrawing()
	}

	rl.UnloadTexture(tileset)

	rl.CloseWindow()
}
