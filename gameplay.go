package main

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
  sort int8
  defense int8
  attack int8
}

type Enemy struct {
	health     int8
	damage     int8
	aggroRange int8
	pos        rl.Vector2
}

type Player struct {
  health int8
  pos rl.Vector2
  inventory map[Item]int
  keys map[int]bool
}

func NewPlayer() Player {
  var p Player
  p.keys = map[int]bool{
    rl.KeyUp : true,
    rl.KeyLeft : true,
    rl.KeyDown : true,
    rl.KeyRight : true,
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

func (p *Player) Action() {

}
