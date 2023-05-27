package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var tilesetCoords = map[int]rl.Vector2{
	// Floor
	0: rl.NewVector2(20, 16),
	// Poison Trap
	1: rl.NewVector2(55, 40),
	// Movement Trap
	2: rl.NewVector2(55, 40),
	// Wall
	3: rl.NewVector2(16, 32),
}

func drawWorld(level Level, x int, y int, tileset rl.Texture2D) {

	for X, yMap := range level.cases {
		for Y, tile := range yMap {
			coordX := tilesetCoords[tile.kind].X
			coordY := tilesetCoords[tile.kind].Y
			currX := float32(8*X + x)
			currY := float32(8*Y + y)
			tile := rl.NewRectangle(coordX, coordY, 8, 8)
			rl.DrawTextureRec(tileset, tile, rl.NewVector2(currX, currY), rl.White)
		}
	}

}
