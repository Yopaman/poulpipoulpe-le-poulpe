package main

import (
	"math/rand"
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
	'.': {KindBase, 0},
	'P': {KindPoison, 0},
	'M': {KindMovement, 0},
	'#': {KindWall, 0},
}

type Case struct {
  kind int
  tile int
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
			caseKind, ok := RuneKinds[raw]
			if ok {
				if w.cases[x] == nil {
					w.cases[x] = make(map[int]Case)
				}
        if caseKind.kind == KindWall {
          caseKind.tile = rand.Intn(3)
        }
        w.cases[x][y] = caseKind
			}
		}
	}
	return w
}
