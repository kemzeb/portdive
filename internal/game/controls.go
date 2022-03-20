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
	if !c.g.IsOver() {
		// Handle in-game Controls
		switch e.Key {
		case tl.KeyCtrlJ:
			c.MoveUp()
		case tl.KeyCtrlK:
			c.MoveDown()
		case tl.KeyCtrlH:
			c.MoveLeft()
		case tl.KeyCtrlL:
			c.MoveRight()
		case tl.KeyCtrlM:
			c.g.DeterminePortMatrixChoice()
		case tl.KeyCtrlP:
			c.g.DeterminePwnerChoice()
		}
	}
}

// Draw is not implemented in Controls
func (c *Controls) Draw(screen *tl.Screen) {}

// MoveUp moves the PortMatrix cursor up. If it hits the topmost port address,
// it will not move.
func (c *Controls) MoveUp() {
	if !(c.g.UI.MatrixInd == 0) {
		c.g.UI.MatrixInd--
	}
}

// MoveDown moves the PortMatrix cursor down. If it hits the bottommost
// port address, it will not move.
func (c *Controls) MoveDown() {
	if !(c.g.UI.MatrixInd == c.g.Matrix.Len()-1) {
		c.g.UI.MatrixInd++
	}
}

// MoveRight moves the Pwner cursor to the right. If it hits the rightmost
// port fragment, it will not move.
func (c *Controls) MoveRight() {
	if !(c.g.UI.PwnerInd == c.g.Pwner.Len()-1) {
		c.g.UI.PwnerInd++
	}
}

// MoveLeft moves the Pwner cursor to the left. If it hits the rightmost
// port fragment, it will not move.
func (c *Controls) MoveLeft() {
	if !(c.g.UI.PwnerInd == 0) {
		c.g.UI.PwnerInd--
	}
}
