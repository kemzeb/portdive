package game

import (
	tl "github.com/JoelOtter/termloop"
	"portdive/internal"
	"strconv"
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

	if !ui.game.IsOver() {
		// represents the current y value we should be drawing to, as well as an
		// index for the PortMatrix
		y := 0

		// Draw the port Matrix rows
		for ; y < ui.game.Matrix.Len(); y++ {
			bg = tl.ColorDefault
			row := ui.game.Matrix.Get(y)
			if y == ui.MatrixInd {
				bg = internal.Hover
			}
			if row.Status() == internal.Active {
				fg = internal.Active
			} else {
				fg = internal.Inactive
			}
			for x := 0; x < ui.game.Matrix.Get(y).Len(); x++ {
				frag := strconv.FormatInt(int64(ui.game.Matrix.Get(y).Get(x)), 10)
				if x != ui.game.Matrix.Get(y).Len()-1 {
					frag += " . "
				}
				tl.NewText(xOffset*x, yOffset+y, frag, fg, bg).Draw(s)
			}
		}
		// Draw the Pwner elements
		for x := 0; x < ui.game.Pwner.Len(); x++ {
			bg = tl.ColorDefault
			ele := ui.game.Pwner.Get(x)
			frag := strconv.FormatInt(int64(ui.game.Pwner.Get(x).Frag()), 10)

			if x == ui.PwnerInd {
				bg = internal.Hover
			}
			if ele.Status() == internal.Active {
				fg = internal.Active
			} else if ele.Status() == internal.Inactive {
				fg = internal.Inactive
			} else {
				fg = internal.Chosen
			}
			if x != ui.game.Pwner.Len()-1 {
				frag += " . "
			}
			tl.NewText(xOffset*x, yOffset+1+y, frag, fg, bg).Draw(s)
		}
	} else {
		if ui.game.HasWon() {
			tl.NewText(xOffset, yOffset, "You have won", fg, bg).Draw(s)
		} else {
			tl.NewText(xOffset, yOffset, "You have lost", fg, bg).Draw(s)
		}
	}
}
