package main

import (
	"log"
	"os"
	"snakegame/board"
	"snakegame/enums"
	"snakegame/pos"
	"snakegame/snake"

	"github.com/gdamore/tcell"
)

var size = 50

func main() {
	snake := snake.Snake{
		Head:        pos.Pos{X: 5, Y: 5},
		Body:        []pos.Pos{{X: 4, Y: 5}, {X: 3, Y: 5}},
		BoardWidth:  size,
		BoardHeight: size,
		Direction:   enums.Right,
	}

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatal(err)
	}

	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	board := board.Board{
		Fruit:  pos.Pos{X: 8, Y: 5},
		Height: size,
		Width:  size,
		Screen: screen,
	}

	go update(&board, &snake)

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyUp:
				snake.Direction = enums.Up
			case tcell.KeyDown:
				snake.Direction = enums.Down
			case tcell.KeyLeft:
				snake.Direction = enums.Left
			case tcell.KeyRight:
				snake.Direction = enums.Right
			case tcell.KeyEscape:
			case tcell.KeyExit:
				screen.Fini()
				os.Exit(0)
			}
		}
	}

}

func update(b *board.Board, s *snake.Snake) {
	for {
		s.Update()
		b.Update(s)
	}
}
