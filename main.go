package main

import (
	"io/ioutil"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

  player := NewPlayer()

	rl.InitWindow(screenWidth, screenHeight, "test")

	camera := rl.Camera2D{}
	// camera.Offset = rl.NewVector2(float32(screenWidth) / 2, float32(screenHeight) / 2)
	camera.Rotation = 0.0
	camera.Zoom = 5.0
  camera.Target = rl.NewVector2(player.pos.X + 20, player.pos.Y + 20)

	rl.SetTargetFPS(60)

	levelContent, err := ioutil.ReadFile("level.txt")
	if err != nil {
		return
	}
	level := ParseWorld(string(levelContent))
	tileset := rl.LoadTexture("tileset.png")
	player_texture := rl.LoadTexture("chars.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
    rl.ClearBackground(rl.Black)
		drawWorld(level, 20, 20, tileset)
		drawPlayer(player_texture, int(player.pos.X), int(player.pos.Y), 20, 20, 1)
		rl.EndMode2D()
		rl.EndDrawing()
    player.Action(&level)
    camera.Target = rl.NewVector2(player.pos.X + 20, player.pos.Y + 20)
	}

	rl.UnloadTexture(tileset)
  rl.UnloadTexture(player_texture)

	rl.CloseWindow()
}
