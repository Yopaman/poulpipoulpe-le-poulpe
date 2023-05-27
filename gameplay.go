package main

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Item struct {
  sort int8
  defense int8
  attack int8
}

type Player struct {
  health int8
  pos rl.Vector2
  inventory map[Item]int
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
