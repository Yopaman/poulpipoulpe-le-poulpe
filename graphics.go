package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var tilesetCoords = map[int]rl.Vector2{
	// Floor
	0: rl.NewVector2(20, 16),
	// Poison Trap
	1: rl.NewVector2(168, 32),
	// Poison Trap activated
	2: rl.NewVector2(168, 23),
	// Movement Trap
	3: rl.NewVector2(168, 32),
	// Movement Trap activated
	4: rl.NewVector2(168, 23),
	// Wall
	5: rl.NewVector2(16, 32),
}

func drawWorld(level Level, x int, y int, tileset rl.Texture2D) {

	for X, yMap := range level.cases {
		for Y, tile := range yMap {
			coordX := tilesetCoords[int(tile)].X
			coordY := tilesetCoords[int(tile)].Y
			currX := float32(8*X + x)
			currY := float32(8*Y + y)
			tile := rl.NewRectangle(coordX, coordY, 8, 8)
			rl.DrawTextureRec(tileset, tile, rl.NewVector2(currX, currY), rl.White)
		}
	}

}

// Orientation :
// 0 : top
// 1 : right
// 2 : bottom
// 3 : left
func drawPlayer(texture rl.Texture2D, x int, y int, xMap int, yMap int, orientation int) {
	position := rl.NewVector2(float32(x*8+xMap), float32(y*8+yMap))

	var tile rl.Rectangle

	switch orientation {
	case 0:
		tile = rl.NewRectangle(8, 0, 8, 8)
	case 1:
		tile = rl.NewRectangle(16, 0, 8, 8)
	case 2:
		tile = rl.NewRectangle(0, 0, 8, 8)
	case 3:
		tile = rl.NewRectangle(24, 0, 8, 8)
	}

	rl.DrawTextureRec(texture, tile, position, rl.White)
}
