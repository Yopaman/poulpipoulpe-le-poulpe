package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"io/ioutil"
	"math/rand"
	"strconv"
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
	// KindEntry 'E' is the kind for the level entry
	KindEntry
	// KindExit 'S' is the kind for the level exit
	KindExit
)

var RuneKinds = map[rune]Case{
	'.': {KindBase, 0},
	'P': {KindPoison, 0},
	'M': {KindMovement, 0},
	'#': {KindWall, 0},
	'E': {KindEntry, 0},
	'S': {KindExit, 0},
}

type Case struct {
	kind int
	tile int
}

type Level struct {
	cases   map[int](map[int]Case)
	enemies []Enemy
}

func ParseWorld(input string) map[int](map[int]Case) {
	cases := make(map[int](map[int]Case))
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			caseKind, ok := RuneKinds[raw]
			if ok {
				if cases[x] == nil {
					cases[x] = make(map[int]Case)
				}
				if caseKind.kind == KindBase {
					caseKind.tile = rand.Intn(3)
				}
				cases[x][y] = caseKind
			}
		}
	}
	return cases
}

func readInt(s string) int8 {
	n, _ := strconv.ParseInt(s, 10, 8)
	return int8(n)
}

func readFloat(s string) float32 {
	n, _ := strconv.ParseFloat(s, 32)
	return float32(n)
}

func ParseEnemies(input string) []Enemy {
	enemiesInput := strings.Split(strings.TrimSpace(input), "\n")
	enemies := make([]Enemy, len(enemiesInput))
	for y, row := range enemiesInput {
		enemies[y] = Enemy{}
		infos := strings.Split(strings.TrimSpace(row), " ")
		enemies[y].sort = readInt(infos[0])
		enemies[y].health = readInt(infos[1])
		enemies[y].damage = readInt(infos[2])
		enemies[y].aggroRange = readInt(infos[3])
		enemies[y].pos = rl.NewVector2(readFloat(infos[4]), readFloat(infos[5]))
	}
	return enemies
}

func ParseLevel(casesFile string, enemiesFile string) (l Level, err error) {
	levelContent, err := ioutil.ReadFile(casesFile)
	if err != nil {
		return l, err
	}
	l.cases = ParseWorld(string(levelContent))

	enemiesContent, err := ioutil.ReadFile(enemiesFile)
	if err != nil {
		return l, err
	}
	l.enemies = ParseEnemies(string(enemiesContent))
	return l, nil
}
