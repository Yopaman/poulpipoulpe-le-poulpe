package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(1500)
	screenHeight := int32(900)

	isGameOver := false

	rl.InitWindow(screenWidth, screenHeight, "test")

	player := NewPlayer("chars.png")

	camera := rl.Camera2D{}
	camera.Offset = rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)
	camera.Rotation = 0.0
	camera.Zoom = 5.0
	camera.Target = rl.NewVector2(player.pos.X*8, player.pos.Y*8)

	rl.SetTargetFPS(60)

	currentLevel := 1

	level, err := ParseLevel(fmt.Sprintf("level%d.txt", currentLevel), fmt.Sprintf("enemies%d.txt", currentLevel))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	tileset := rl.LoadTexture("tileset.png")
	arrows := rl.LoadTexture("arrows.png")
	arrowsBig := rl.LoadTexture("arrows_big.png")

	for !rl.WindowShouldClose() {
		if !isGameOver {
			camera.Zoom += rl.GetMouseWheelMove() * 0.05
			camera.Zoom = float32(math.Max(float64(camera.Zoom), 3.0))
			rl.BeginDrawing()
			rl.BeginMode2D(camera)
			rl.ClearBackground(rl.Black)
			drawWorld(level, 0, 0, tileset)
			drawPlayer(player.texture, int(player.pos.X), int(player.pos.Y), 0, 0, player.orientation)
			drawArrows(arrows, player.keys, player.pos.X*8, player.pos.Y*8)
			isGameOver = !player.CanMove(&level)
			if player.Action(&level) {
				player.CheckTrap(&level)
				for i := range level.enemies {
					level.enemies[i].Action(&level, &player)
				}
			}
			for _, e := range level.enemies {
				drawEnnemy(tileset, int(e.pos.X)*8, int(e.pos.Y)*8)
			}
			if player.CheckExit(&level) {
				currentLevel += 1
				level, err = ParseLevel(fmt.Sprintf("level%d.txt", currentLevel), fmt.Sprintf("enemies%d.txt", currentLevel))
				if err != nil {
					fmt.Printf("Error: %v", err)
					return
				}
				player.pos = rl.NewVector2(6, 3)
			}
			camera.Target = rl.NewVector2(player.pos.X*8, player.pos.Y*8)
			rl.EndMode2D()
			drawNextKeys(arrowsBig, &player)
			drawHealth(&player)
			rl.EndDrawing()
			if player.health <= 0 {
				isGameOver = true
			}
		} else {
			rl.BeginDrawing()
			drawGameOverScreen(screenWidth, screenHeight)
			if rl.IsKeyPressed(rl.KeyR) {
				player = NewPlayer("chars.png")
				currentLevel = 1
				level, err = ParseLevel(fmt.Sprintf("level%d.txt", currentLevel), fmt.Sprintf("enemies%d.txt", currentLevel))
				if err != nil {
					fmt.Printf("Error: %v", err)
					return
				}
				isGameOver = false
			}
			rl.EndDrawing()
		}
	}

	rl.UnloadTexture(tileset)
	rl.UnloadTexture(player.texture)
	rl.UnloadTexture(arrows)
	rl.UnloadTexture(arrowsBig)

	rl.CloseWindow()
}
