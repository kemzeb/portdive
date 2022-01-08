package game

import (
	"fmt"
	"github.com/pterm/pterm"
	"math/rand"
	"portdive/internal/input"
	"strings"
	"time"
)

type Game struct {
	m      *PortMatrix
	p      *Pwner
	k      *Key
	h      input.InHandler
	input  chan *input.InHandlerResult
	seed   *rand.Source
	hasWon bool
}

func NewGame(m *[]PortRow, s *rand.Source) *Game {
	key := NewKey(s)
	pwner := NewPwner(key, s)
	matrix := NewPortMatrix(*m, pwner)
	pwner.SetMatrix(matrix) // TODO: Solve cyclic dependency problem with PortMatrix
	in := make(chan *input.InHandlerResult)
	return &Game{
		m:     matrix,
		p:     pwner,
		k:     key,
		h:     *input.NewInHandler(in),
		input: in,
		seed:  s}
}

func (g *Game) StartGame() {
	var (
		message []string
		err     error
	)

	g.k.RandomizeKey(g.m)
	g.p.Init()
	g.p.UpdateWithoutRandomization()

	area, _ := pterm.DefaultArea.WithCenter().Start()

	finish := make(chan bool)
	go g.h.Handle(finish)

	g.printGame(area)
	for !g.IsGameOver() {
		message = nil

		select {
		case result := <-g.input:
			if result.Err != nil {
				message = append(message, result.Err.Error())
			} else {
				if result.Selection.Type == input.MatrixSelection {
					err = g.determinePortMatrixChoice(*result)
				} else {
					err = g.determinePwnerChoice(*result)
				}
				message = append(message, err.Error())
			}
			g.m.Update()
			g.p.UpdateWithoutRandomization()
		case <-time.After(Duration):
			g.m.Update()
			g.p.Update()

		}
		g.printGame(area, message...)
	}
	finish <- true // tell the input handler the game has finished

}

// IsGameOver determines if the game has ended. It will also set the hasWon
// field to true if the game was won. As of now, there are 3 ways
// of triggering a game over:
//
// 1. The player chooses a correct row in the PortMatrix.
//
// 2. The player chooses an incorrect row in the PortMatrix.
//
// 3. The player has the status of each PwnerElment in the Pwner as Chosen.
func (g Game) IsGameOver() bool {
	return g.hasWon
}

// determinePortMatrixChoice determines if a choice a player made is valid or
// not. It will then determine
func (g *Game) determinePortMatrixChoice(res input.InHandlerResult) error {
	return nil
}

func (g *Game) determinePwnerChoice(res input.InHandlerResult) error {
	return nil
}

func (g Game) printGame(area *pterm.AreaPrinter, messages ...string) {
	var sb strings.Builder

	for i := 0; i < g.m.Len(); i++ {
		row := g.m.Get(i)
		rowColor := row.Status()
		// Display the PortMatrix
		sb.WriteString(fmt.Sprintf("%-6d ", i))
		for j := 0; j < row.Len(); j++ {
			frag := row.Get(j)
			coloredFrag := rowColor.Sprint(frag)
			sb.WriteString(coloredFrag)
			if j+1 == row.Len() {
				break
			}
			sb.WriteString(" . ")
		}
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("\n%-7s", ""))
	// Display the Pwner
	for i := 0; i < g.p.Len(); i++ {
		element := g.p.Get(i)
		coloredFrag := element.Status().Sprint(element.Frag())
		sb.WriteString(coloredFrag)
		if i+1 == g.p.Len() {
			break
		}
		sb.WriteString(" . ")
	}
	sb.WriteString("\n")
	// Display any error messages
	for _, msg := range messages {
		sb.WriteString(pterm.FgYellow.Sprint(msg + "\n"))
	}
	area.Update(sb.String())
}
