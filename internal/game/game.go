package game

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

type Game struct {
	Engine   *tl.Game
	Options  Options
	Controls *Controls
	Key      *Key
	Pwner    *Pwner
	Matrix   *PortMatrix
	UI       *UI
	isOver   bool
	hasWon   bool
	ticker   time.Ticker
}

func (g *Game) IsOver() bool {
	return g.isOver
}

func (g *Game) HasWon() bool {
	return g.hasWon
}

type Options struct {
	// width of the game world in terms of console pixels (cells)
	width int
	// height of the game world in terms of console pixels (cells)
	height int
	// fps sets the framerate for tl.Screen
	fps int
}

func NewGame(m *[]PortRow) *Game {
	game := &Game{}
	seed := rand.NewSource(time.Now().UnixNano())

	game.Engine = tl.NewGame()
	tw, th := game.Engine.Screen().Size()
	game.Options = Options{tw, th, 30}
	game.Key = NewKey(&seed)
	game.Pwner = NewPwner(game.Key, &seed)
	game.Matrix = NewPortMatrix(*m, game.Pwner)
	game.UI = NewUI(game)
	game.Controls = NewControls(game)
	game.ticker = *time.NewTicker(Duration)
	// TODO: Solve cyclic dependency problem with PortMatrix
	game.Pwner.SetMatrix(game.Matrix)
	return game
}

func (g *Game) Tick(e tl.Event) {}

func (g *Game) Draw(s *tl.Screen) {
	select {
	case <-g.ticker.C:
		g.Pwner.Update()
		g.Matrix.Update()
	default:
	}
}

func (g *Game) Start() {
	g.Key.RandomizeKey(g.Matrix)
	g.Pwner.Init()
	g.Matrix.Update()

	g.Engine.Screen().AddEntity(g.UI)
	g.Engine.Screen().AddEntity(g.Controls)
	g.Engine.Screen().AddEntity(g)
	g.Engine.Start()
}

func (g *Game) DeterminePortMatrixChoice() {
	if g.UI.MatrixInd == g.Key.ChosenIndex() {
		g.hasWon = true
		g.isOver = true
	} else if g.Matrix.Get(g.UI.MatrixInd).Status() == Inactive {
		return
	} else { // An incorrect active status PortRow was chosen
		g.hasWon = false
		g.isOver = true
	}
}

func (g *Game) DeterminePwnerChoice() {
	chosenEle := g.Pwner.Get(g.UI.PwnerInd)

	if !chosenEle.Selectable() {
		return
	}
	if chosenEle.Status() == Active {
		chosenEle.SetStatus(Chosen)
	}
	g.Pwner.UpdateWithoutRandomization()
	g.Matrix.Update()

	// See if the player has chosen all the correct port fragments in the Pwner
	// device
	for i := 0; i < g.Pwner.Len(); i++ {
		if g.Pwner.Get(i).Status() != Chosen {
			return
		}
	}
	g.hasWon = true
	g.isOver = true
}
