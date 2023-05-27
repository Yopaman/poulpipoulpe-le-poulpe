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

var possibleKeys = [4]int{rl.KeyRight, rl.KeyDown, rl.KeyLeft, rl.KeyUp}

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
	p.health = 10
	generateKeysRemoved(&p)
	return p
}

func desactivateNextKey(p *Player) {
	for k := range p.keys {
		p.keys[k] = true
	}

	toRemove := int32(p.nextKeysRemoved[0])
	p.nextKeysRemoved = p.nextKeysRemoved[1:]

	p.keys[toRemove] = false

	randomIndex := rand.Intn(4)
	p.nextKeysRemoved = append(p.nextKeysRemoved, possibleKeys[randomIndex])
}

func generateKeysRemoved(p *Player) {
	keys := make([]int, 0)
	for i := 0; i < 4; i++ {
		randomIndex := rand.Intn(4)
		keys = append(keys, possibleKeys[randomIndex])
	}
	p.nextKeysRemoved = keys
}

/*
Si le joueur est dans la range, on avance vers lui
Si il est à porté, on l'attaque
Sinon on ne fait rien.
*/
func (e *Enemy) Action(l *Level, p *Player) bool {
  if e.health <= 0 {
    e.pos.X = 300
    e.pos.Y = 300
  }
	d := rl.Vector2Distance(p.pos, e.pos)
	if d > float32(e.aggroRange) {
		return false
	}
  if rl.Vector2Distance(p.pos, e.pos) <= 1 {
    p.health--
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
    npx := int(p.pos.X)
    npy := int(p.pos.Y)-1
		v := l.cases[npx][npy]
    for i := range l.enemies {
      if int(l.enemies[i].pos.X) == npx && int(l.enemies[i].pos.Y) == npy {
        l.enemies[i].health--
        return true
      }
    }
		if v.kind != KindWall {
			p.pos.Y -= 1
			p.orientation = 0
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyDown] && rl.IsKeyPressed(rl.KeyDown) {
    npx := int(p.pos.X)
    npy := int(p.pos.Y)+1
		v := l.cases[npx][npy]
    for i := range l.enemies {
      if int(l.enemies[i].pos.X) == npx && int(l.enemies[i].pos.Y) == npy {
        l.enemies[i].health--
        return true
      }
    }
		if v.kind != KindWall {
			p.pos.Y += 1
			p.orientation = 2
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyLeft] && rl.IsKeyPressed(rl.KeyLeft) {
    npx := int(p.pos.X)-1
    npy := int(p.pos.Y)
		v := l.cases[npx][npy]
    for i := range l.enemies {
      if int(l.enemies[i].pos.X) == npx && int(l.enemies[i].pos.Y) == npy {
        l.enemies[i].health--
        return true
      }
    }
		if v.kind != KindWall {
			p.pos.X -= 1
			p.orientation = 3
			desactivateNextKey(p)
			return true
		}
	} else if p.keys[rl.KeyRight] && rl.IsKeyPressed(rl.KeyRight) {
    npx := int(p.pos.X)+1
    npy := int(p.pos.Y)
		v := l.cases[npx][npy]
    for i := range l.enemies {
      if int(l.enemies[i].pos.X) == npx && int(l.enemies[i].pos.Y) == npy {
        l.enemies[i].health--
        return true
      }
    }
		if v.kind != KindWall {
			p.pos.X += 1
			p.orientation = 1
			desactivateNextKey(p)
			return true
		}
	}

	return false
}

func (p *Player) CheckTrap(l *Level) bool {
	if c, ok := l.cases[int(p.pos.X)][int(p.pos.Y)]; ok && c.kind == KindPoison && c.tile == 0 {
		c.tile = 1
		l.cases[int(p.pos.X)][int(p.pos.Y)] = c
		p.health--
	} else if c, ok := l.cases[int(p.pos.X)][int(p.pos.Y)]; ok && c.kind == KindMovement && c.tile == 0 {
		c.tile = 1
		l.cases[int(p.pos.X)][int(p.pos.Y)] = c
		p.keys[int32(possibleKeys[rand.Intn(4)])] = false
	}
	return false
}

func (p *Player) CheckExit(l *Level) bool {
	return l.cases[int(p.pos.X)][int(p.pos.Y)].kind == KindExit
}

func (p *Player) CanMove(l *Level) bool {
	for _, offset := range [][]int{{-1, 0, rl.KeyLeft}, {1, 0, rl.KeyRight}, {0, -1, rl.KeyUp}, {0, 1, rl.KeyDown}} {
		if tile, ok := l.cases[int(p.pos.X)+offset[0]][int(p.pos.Y)+offset[1]]; ok && tile.kind != KindWall && p.keys[int32(offset[2])] {
			return true
		}
	}
	return false

}
