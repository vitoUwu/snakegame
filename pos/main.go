package pos

import (
	"snakegame/enums"
)

type Pos struct {
	X int
	Y int
}

func (p *Pos) Move(direction int) {
	switch direction {
	case enums.Up:
		p.Y -= 1
	case enums.Down:
		p.Y += 1
	case enums.Left:
		p.X -= 1
	case enums.Right:
		p.X += 1
	}
}

func (p *Pos) Set(x int, y int) {
	p.X = x
	p.Y = y
}
