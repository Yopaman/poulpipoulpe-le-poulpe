package main

import (
  "github.com/gen2brain/raylib-go/raylib"
)

/* trap :
    - 0 : none
    - 1 : poison
    - 2 : enlève un mouvement
*/
type Case struct {
  trap int8
  wall bool
  tile string
}

type Enemy struct {
  health int8
  damage int8
  aggroRange int8
  pos rl.Vector2
}
