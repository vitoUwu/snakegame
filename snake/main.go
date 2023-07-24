package snake

import (
	"snakegame/enums"
	"snakegame/pos"
)

const (
	Up = iota
	Down
	Left
	Right
)

type Snake struct {
	Head        pos.Pos
	Body        []pos.Pos
	BoardWidth  int
	BoardHeight int
	Direction   int
}

func (s *Snake) Update() {
	if (s.Head.X == s.BoardWidth-1 && s.Direction == enums.Right) ||
		(s.Head.X == 0 && s.Direction == enums.Left) ||
		(s.Head.Y == s.BoardHeight-1 && s.Direction == enums.Down) ||
		(s.Head.Y == 0 && s.Direction == enums.Up) {
		panic("Game Over")
	}

	for i := len(s.Body) - 1; i >= 0; i-- {
		if i == 0 {
			s.Body[i].Set(s.Head.X, s.Head.Y)
		} else {
			s.Body[i].Set(s.Body[i-1].X, s.Body[i-1].Y)
		}
	}

	s.Head.Move(s.Direction)
}

func (s *Snake) BodyIn(x int, y int) bool {
	for _, body := range s.Body {
		if body.X == x && body.Y == y {
			return true
		}
	}

	return false
}

func (s *Snake) AppendBody() {
	if len(s.Body) == 0 {
		s.Body = append(s.Body, pos.Pos{X: s.Head.X, Y: s.Head.Y})
	} else {
		s.Body = append(s.Body, pos.Pos{X: s.Body[len(s.Body)-1].X, Y: s.Body[len(s.Body)-1].Y})
	}
}
