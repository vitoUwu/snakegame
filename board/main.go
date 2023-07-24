package board

import (
	"math/rand"
	"snakegame/pos"
	"snakegame/snake"
	"time"

	"github.com/gdamore/tcell"
)

type Board struct {
	Fruit  pos.Pos
	Height int
	Width  int
	Screen tcell.Screen
}

func (b *Board) GenerateFruit() {
	b.Fruit = pos.Pos{X: rand.Intn(b.Width), Y: rand.Intn(b.Height)}
}

func (b *Board) Update(s *snake.Snake) {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if s.BodyIn(s.Head.X, s.Head.Y) {
				panic("Game Over")
			}

			if b.Fruit.X == s.Head.X && b.Fruit.Y == s.Head.Y {
				s.AppendBody()
				b.GenerateFruit()
			}

			if b.Fruit.X == j && b.Fruit.Y == i {
				b.Screen.SetContent(j, i, 'X', nil, tcell.StyleDefault)
				continue
			}

			if s.Head.X == j && s.Head.Y == i {
				b.Screen.SetContent(j, i, 'O', nil, tcell.StyleDefault)
				continue
			}

			if s.BodyIn(j, i) {
				b.Screen.SetContent(j, i, '#', nil, tcell.StyleDefault)
				continue
			}

			if j == b.Width-1 || j == 0 {
				b.Screen.SetContent(j, i, '|', nil, tcell.StyleDefault)
				continue
			}

			if i == b.Height-1 || i == 0 {
				b.Screen.SetContent(j, i, '-', nil, tcell.StyleDefault)
				continue
			}

			b.Screen.SetContent(j, i, ' ', nil, tcell.StyleDefault)
		}
	}

	b.Screen.Show()
	time.Sleep(time.Millisecond * 250)
}
