package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
	sort    int8
	defense int8
	attack  int8
}

type Enemy struct {
	health     int8
	damage     int8
	aggroRange int8
	pos        rl.Vector2
}

type Player struct {
	health    int8
	pos       rl.Vector2
	inventory map[Item]int
	keys      map[int32]bool
}

func NewPlayer() Player {
	var p Player
	p.keys = map[int32]bool{
		rl.KeyUp:    true,
		rl.KeyLeft:  true,
		rl.KeyDown:  true,
		rl.KeyRight: true,
	}
	p.pos = rl.NewVector2(6, 3)
	return p
}

/*
Si le joueur est dans la range, on avance vers lui
Si il est à porté, on l'attaque
Sinon on ne fait rien.
*/
func (e *Enemy) Action() {
}

func (p *Player) Action(l *Level) {
  done := false
  for !done {
    key := rl.GetKeyPressed()
    ok, err := p.keys[key]
    if !err && ok {
      switch key {
      case rl.KeyUp:
        v := l.cases[int(p.pos.X)][int(p.pos.Y)+1]
        if v.kind != KindWall {
          p.pos.Y += 1
          done = true
        }
      case rl.KeyDown:
        v := l.cases[int(p.pos.X)][int(p.pos.Y)-1]
        if v.kind != KindWall {
          p.pos.Y -= 1
          done = true
        }
      case rl.KeyLeft:
        v := l.cases[int(p.pos.X)-1][int(p.pos.Y)]
        if v.kind != KindWall {
          p.pos.X -= 1
          done = true
        }
      case rl.KeyRight:
        v := l.cases[int(p.pos.X)+1][int(p.pos.Y)]
        if v.kind != KindWall {
          p.pos.X += 1
          done = true
        }
      }
    }
  }
}
