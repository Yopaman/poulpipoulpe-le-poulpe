package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
	sort    int8
	defense int8
	attack  int8
}

type Enemy struct {
	sort       int8
	health     int8
	damage     int8
	aggroRange int8
	pos        rl.Vector2
}

type Player struct {
	health          int8
	pos             rl.Vector2
	orientation     int
	inventory       map[Item]int
	keys            map[int32]bool
	texture         rl.Texture2D
	nextKeysRemoved []int
}

func NewPlayer(file string) Player {
	t := rl.LoadTexture(file)
	var p Player
	p.keys = map[int32]bool{
		rl.KeyUp:    true,
		rl.KeyLeft:  true,
		rl.KeyDown:  true,
		rl.KeyRight: true,
	}
	p.pos = rl.NewVector2(6, 3)
	p.texture = t
	p.orientation = 2
	return p
}

/*
Si le joueur est dans la range, on avance vers lui
Si il est à porté, on l'attaque
Sinon on ne fait rien.
*/
func (e *Enemy) Action(l *Level, p *Player) bool {
	d := rl.Vector2Distance(p.pos, e.pos)
	if d > float32(e.aggroRange) {
		return false
	}
	for _, offset := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nv := rl.NewVector2(e.pos.X+float32(offset[0]), e.pos.Y+float32(offset[1]))
		if n, ok := l.cases[int(nv.X)][int(nv.Y)]; ok && n.kind != KindWall && rl.Vector2Distance(nv, p.pos) < d {
			e.pos = nv
			return true
		}
	}
	return false
}

func (p *Player) Action(l *Level) bool {
	if p.keys[rl.KeyUp] && rl.IsKeyPressed(rl.KeyUp) {
		v := l.cases[int(p.pos.X)][int(p.pos.Y)-1]
		if v.kind != KindWall {
			p.pos.Y -= 1
			p.orientation = 0
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyDown] && rl.IsKeyPressed(rl.KeyDown) {
		v := l.cases[int(p.pos.X)][int(p.pos.Y)+1]
		if v.kind != KindWall {
			p.pos.Y += 1
			p.orientation = 2
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyLeft] && rl.IsKeyPressed(rl.KeyLeft) {
		v := l.cases[int(p.pos.X)-1][int(p.pos.Y)]
		if v.kind != KindWall {
			p.pos.X -= 1
			p.orientation = 3
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyRight] && rl.IsKeyPressed(rl.KeyRight) {
		v := l.cases[int(p.pos.X)+1][int(p.pos.Y)]
		if v.kind != KindWall {
			p.pos.X += 1
			p.orientation = 1
			desactivateNextKey(p)
			return true
		}
	}

	return false
}

func desactivateNextKey(p *Player) {
	for k := range p.keys {
		p.keys[k] = true
	}

	toRemove := int32(p.nextKeysRemoved[0])
	p.nextKeysRemoved = p.nextKeysRemoved[1:]

	p.keys[toRemove] = false

	possibleKeys := [4]int{rl.KeyRight, rl.KeyDown, rl.KeyLeft, rl.KeyUp}
	randomIndex := rand.Intn(4)
	p.nextKeysRemoved = append(p.nextKeysRemoved, possibleKeys[randomIndex])
}

func generateKeysRemoved(p *Player) {
	possibleKeys := [4]int{rl.KeyRight, rl.KeyDown, rl.KeyLeft, rl.KeyUp}
	keys := make([]int, 0)
	for i := 0; i < 4; i++ {
		randomIndex := rand.Intn(4)
		keys = append(keys, possibleKeys[randomIndex])
	}
	p.nextKeysRemoved = keys
}
