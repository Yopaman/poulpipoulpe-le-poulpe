package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var tilesetCoords = map[int](func(n int) rl.Vector2){
	// Floor
	0: func(tile int) rl.Vector2 {
		textures := [3]rl.Vector2{rl.NewVector2(16, 16), rl.NewVector2(24, 16), rl.NewVector2(32, 16)}
		return textures[tile]
	},
	// Poison trap
	1: func(tile int) rl.Vector2 {
		textures := [2]rl.Vector2{rl.NewVector2(168, 23), rl.NewVector2(168, 32)}
		return textures[tile]
	},
	// Movement trap
	2: func(tile int) rl.Vector2 {
		textures := [2]rl.Vector2{rl.NewVector2(176, 23), rl.NewVector2(176, 32)}
		return textures[tile]
	},
	// Wall
	3: func(tile int) rl.Vector2 {
		return rl.NewVector2(16, 32)
	},
	// Entry
	4: func(tile int) rl.Vector2 {
		return rl.NewVector2(168, 8)
	},
	// Exit
	5: func(tile int) rl.Vector2 {
		return rl.NewVector2(176, 8)
	},
}

func drawWorld(level Level, x int, y int, tileset rl.Texture2D) {

	for X, yMap := range level.cases {
		for Y, tile := range yMap {
			coordX := tilesetCoords[tile.kind](tile.tile).X
			coordY := tilesetCoords[tile.kind](tile.tile).Y
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

func drawArrows(texture rl.Texture2D, keys map[int32]bool, playerX float32, playerY float32) {
	var tile1 rl.Rectangle
	var tile2 rl.Rectangle
	var tile3 rl.Rectangle
	var tile4 rl.Rectangle
	var position1 rl.Vector2
	var position2 rl.Vector2
	var position3 rl.Vector2
	var position4 rl.Vector2

	if keys[rl.KeyDown] {
		tile1 = rl.NewRectangle(0, 0, 8, 8)
		position1 = rl.NewVector2(float32(playerX), float32(playerY+8))
	} else {
		tile1 = rl.NewRectangle(0, 8, 8, 8)
		position1 = rl.NewVector2(float32(playerX), float32(playerY+8))
	}
	if keys[rl.KeyLeft] {
		tile2 = rl.NewRectangle(8, 0, 8, 8)
		position2 = rl.NewVector2(float32(playerX-8), float32(playerY))
	} else {
		tile2 = rl.NewRectangle(8, 8, 8, 8)
		position2 = rl.NewVector2(float32(playerX-8), float32(playerY))
	}

	if keys[rl.KeyUp] {
		tile3 = rl.NewRectangle(16, 0, 8, 8)
		position3 = rl.NewVector2(float32(playerX), float32(playerY-8))
	} else {
		tile3 = rl.NewRectangle(16, 8, 8, 8)
		position3 = rl.NewVector2(float32(playerX), float32(playerY-8))
	}

	if keys[rl.KeyRight] {
		tile4 = rl.NewRectangle(24, 0, 8, 8)
		position4 = rl.NewVector2(float32(playerX+8), float32(playerY))
	} else {
		tile4 = rl.NewRectangle(24, 8, 8, 8)
		position4 = rl.NewVector2(float32(playerX+8), float32(playerY))
	}

	rl.DrawTextureRec(texture, tile1, position1, rl.ColorAlpha(rl.White, 0.5))
	rl.DrawTextureRec(texture, tile2, position2, rl.ColorAlpha(rl.White, 0.5))
	rl.DrawTextureRec(texture, tile3, position3, rl.ColorAlpha(rl.White, 0.5))
	rl.DrawTextureRec(texture, tile4, position4, rl.ColorAlpha(rl.White, 0.5))
}

func drawArrow(texture rl.Texture2D, x int, y int, key int) {
	var tile rl.Rectangle
	position := rl.NewVector2(float32(x), float32(y))

	switch key {
	case rl.KeyUp:
		tile = rl.NewRectangle(64, 0, 32, 32)
	case rl.KeyRight:
		tile = rl.NewRectangle(96, 0, 32, 32)
	case rl.KeyDown:
		tile = rl.NewRectangle(0, 0, 32, 32)
	case rl.KeyLeft:
		tile = rl.NewRectangle(32, 0, 32, 32)
	}

	rl.DrawTextureRec(texture, tile, position, rl.White)
}

func drawNextKeys(texture rl.Texture2D, p *Player) {
	xPos := 10
	yPos := 10
	for _, key := range p.nextKeysRemoved {
		drawArrow(texture, xPos, yPos, key)
		xPos += 40
	}
}

func drawGameOverScreen(screenWidth int32, screenHeight int32) {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Game Over !", screenWidth/2-rl.MeasureText("Game Over !", 50)/2, screenHeight/2-30, 50, rl.White)
	rl.DrawText("Appuyez sur R pour recommencer", screenWidth/2-rl.MeasureText("Appuyez sur R pour recommencer", 30)/2, screenHeight/2+100, 30, rl.White)
}

func drawEnnemy(texture rl.Texture2D, x int, y int) {
	tile := rl.NewRectangle(112, 8, 8, 8)
	position := rl.NewVector2(float32(x), float32(y))
	rl.DrawTextureRec(texture, tile, position, rl.White)
}

func drawEnemies(texture rl.Texture2D, level *Level) {
	for _, e := range level.enemies {
		drawEnnemy(texture, int(e.pos.X*8), int(e.pos.Y*8))
	}
}

func drawHealth(p *Player) {
	rl.DrawText("Sante : "+strconv.Itoa(int(p.health)), int32(rl.GetScreenWidth())-110, 10, 20, rl.White)
}

func drawEnd(screenWidth int32, screenHeight int32) {
	rl.DrawText("Bravo !", screenWidth/2-rl.MeasureText("Bravo !", 30)/2, screenHeight/2+100, 30, rl.White)
}
