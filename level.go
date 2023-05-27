package main

import (
	"strings"
)

const (
	// Kindase '.' is the kind for a base tile
	KindBase = iota
	// KindPoison 'P' is the kind for a poison trap
	KindPoison
	KindPoisonDeactivated
	// KindMovement 'M' is the kind of a movement trap
	KindMovement
	KindMovementDeactivated
	// KindWall '#' is the kind for a wall
	KindWall
)

var RuneKinds = map[rune]Case{
	'.': KindBase,
	'P': KindPoison,
	'M': KindMovement,
	'#': KindWall,
}

type Case int

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
