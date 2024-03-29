package game

import (
	"strconv"

	tl "github.com/JoelOtter/termloop"
)

const (
	xOffset = 6
	yOffset = 0
)

type UI struct {
	game *Game
	// PwnerInd is the current index of Pwner that UI refers to.
	PwnerInd int
	// MatrixInd is the current index of PortMatrix that UI refers to.
	MatrixInd int
}

func (ui *UI) Tick(e tl.Event) {}

func NewUI(game *Game) *UI {
	return &UI{game: game}
}

func (ui *UI) Draw(s *tl.Screen) {
	var bg, fg tl.Attr

	if ui.game.Status() == IsActive {
		// y represents the current y value we should be drawing to, as well as an
		// index into PortMatrix.
		y := 0

		// Draw the port Matrix rows.
		for ; y < ui.game.Matrix.Len(); y++ {
			bg = tl.ColorDefault
			row := ui.game.Matrix.Get(y)
			if y == ui.MatrixInd {
				bg = Hover
			}
			if row.Status() == Active {
				fg = Active
			} else {
				fg = Inactive
			}
			for x := 0; x < ui.game.Matrix.Get(y).Len(); x++ {
				frag := strconv.FormatInt(int64(ui.game.Matrix.Get(y).Get(x)), 10)
				if x != ui.game.Matrix.Get(y).Len()-1 {
					frag += " . "
				}
				tl.NewText(xOffset*x, yOffset+y, frag, fg, bg).Draw(s)
			}
		}
		// Draw the Pwner elements.
		for x := 0; x < ui.game.Pwner.Len(); x++ {
			bg = tl.ColorDefault
			ele := ui.game.Pwner.Get(x)
			frag := strconv.FormatInt(int64(ui.game.Pwner.Get(x).Frag()), 10)

			if x == ui.PwnerInd {
				bg = Hover
			}
			if ele.Status() == Active {
				fg = Active
			} else if ele.Status() == Inactive {
				fg = Inactive
			} else {
				fg = Chosen
			}
			if x != ui.game.Pwner.Len()-1 {
				frag += " . "
			}
			tl.NewText(xOffset*x, yOffset+1+y, frag, fg, bg).Draw(s)
		}
	} else if ui.game.Status() == HasWon {
		tl.NewText(xOffset, yOffset, "You have won", fg, bg).Draw(s)
	} else {
		tl.NewText(xOffset, yOffset, "You have lost", fg, bg).Draw(s)
	}
}

// MoveUp moves the PortMatrix cursor up. If it hits the topmost port address,
// it will not move.
func (ui *UI) MoveUp() {
	if !(ui.MatrixInd == 0) {
		ui.MatrixInd--
	}
}

// MoveDown moves the PortMatrix cursor down. If it hits the bottommost
// port address, it will not move.
func (ui *UI) MoveDown() {
	if !(ui.MatrixInd == ui.game.Matrix.Len()-1) {
		ui.MatrixInd++
	}
}

// MoveRight moves the Pwner cursor to the right. If it hits the rightmost
// port fragment, it will not move.
func (ui *UI) MoveRight() {
	if !(ui.PwnerInd == ui.game.Pwner.Len()-1) {
		ui.PwnerInd++
	}
}

// MoveLeft moves the Pwner cursor to the left. If it hits the rightmost
// port fragment, it will not move.
func (ui *UI) MoveLeft() {
	if !(ui.PwnerInd == 0) {
		ui.PwnerInd--
	}
}
