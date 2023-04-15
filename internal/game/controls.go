package game

import (
	tl "github.com/JoelOtter/termloop"
)

// Controls manages all the data and logic associated to controlling the game
// world.
type Controls struct {
	g *Game
}

func NewControls(game *Game) *Controls {
	return &Controls{g: game}
}

func (c *Controls) Tick(e tl.Event) {
	if c.g.Status() == IsActive {
		switch e.Key {
		case tl.KeyCtrlJ:
			c.g.MoveUp()
		case tl.KeyCtrlK:
			c.g.MoveDown()
		case tl.KeyCtrlH:
			c.g.MoveLeft()
		case tl.KeyCtrlL:
			c.g.MoveRight()
		case tl.KeyCtrlM:
			c.g.DeterminePortMatrixChoice()
		case tl.KeyCtrlP:
			c.g.DeterminePwnerChoice()
		}
	}
}

// Draw is not implemented in Controls.
func (c *Controls) Draw(screen *tl.Screen) {}
