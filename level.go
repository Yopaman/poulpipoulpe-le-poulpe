package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"strings"
)

const (
	// Kindase '.' is the kind for a base tile
	KindBase = iota
	// KindPoison 'P' is the kind for a poison trap
	KindPoison
	// KindMovement 'M' is the kind of a movement trap
	KindMovement
	// KindWall '#' is the kind for a wall
	KindWall
)

var RuneKinds = map[rune]Case{
	'.': Case{KindBase, true},
	'P': Case{KindPoison, false},
	'M': Case{KindMovement, false},
	'#': Case{KindWall, true},
}

type Case struct {
	kind       int
	discovered bool
}

type Enemy struct {
	health     int8
	damage     int8
	aggroRange int8
	pos        rl.Vector2
}

type Level struct {
	cases   map[int](map[int]Case)
	enemies []Enemy
}

func ParseWorld(input string) Level {
	var w Level
	if w.cases == nil {
		w.cases = make(map[int](map[int]Case))
	}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, ok := RuneKinds[raw]
			if ok {
				if w.cases[x] == nil {
					w.cases[x] = make(map[int]Case)
				}
				w.cases[x][y] = kind
			}
		}
	}
	return w
}
