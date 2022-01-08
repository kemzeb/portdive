package game

import (
	"math/rand"

	"github.com/pterm/pterm"
)

// PwnerElement represents an element within the p device.
type PwnerElement struct {
	frag       int
	selectable bool
	status     pterm.Color
}

func (e PwnerElement) Frag() int                { return e.frag }
func (e PwnerElement) Selectable() bool         { return e.selectable }
func (e *PwnerElement) SetSelectable(s bool)    { e.selectable = s }
func (e PwnerElement) Status() pterm.Color      { return e.status }
func (e *PwnerElement) SetStatus(s pterm.Color) { e.status = s }

// Pwner represents the p device.
type Pwner struct {
	frags  []PwnerElement
	matrix *PortMatrix
	key    *Key
	seed   *rand.Source // used for randomizing port numbers
}

func NewPwner(k *Key, s *rand.Source) *Pwner {
	p := Pwner{key: k, seed: s}
	return &p
}

// Get returns the port fragment found at the given index.
func (p Pwner) Get(i int) PwnerElement {
	if !(i < p.Len()) {
		panic("p: overflow error")
	}
	return p.frags[i]
}

// Len returns the number of port fragments.
func (p Pwner) Len() int { return len(p.frags) }

// Init instantiates a slice of PwnerElement types with a length = the # of
//columns in the PortMatrix. It chooses random port fragments in each column...
// for each element in the Pwner.
func (p *Pwner) Init() {
	rowLen := p.matrix.Len()
	colLen := p.matrix.Get(0).Len()

	for i := 0; i < colLen; i++ {
		randomRowIndex := RandomizeInt(*p.seed, 0, rowLen-1)
		randomFrag := p.matrix.Get(randomRowIndex).Get(i)
		element := PwnerElement{randomFrag, false, Inactive}
		p.frags = append(p.frags, element)
	}
}

// Update updates the state of the port fragments.
func (p *Pwner) Update() {
	// Determine the random # of Pwner elements that should be replaced
	replaceLen := RandomizeInt(*p.seed, 0, p.Len()/2)
	var randPwnerIndex int
	var excludedPwnerIndicies []int

	// Determine Pwner indices to exclude
	for i := 0; i < replaceLen; i++ {
		if p.Get(i).Status() == Chosen {
			excludedPwnerIndicies = append(excludedPwnerIndicies, i)
		}
	}

	for i := 0; i < replaceLen; i++ {
		randPwnerIndex = RandomizeInt(*p.seed, 0, p.Len()-1, excludedPwnerIndicies...)
		randMatrixRowIndex := RandomizeInt(*p.seed, 0, p.matrix.Len()-1, i)
		randRow := p.matrix.Get(randMatrixRowIndex)
		p.frags[randPwnerIndex].frag = randRow.Get(randPwnerIndex)
	}

	// Update the select and status states of each PwnerElement. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := &p.frags[i]

		if !p.IsSelectable(i) {
			ele.SetSelectable(false)
			ele.SetStatus(Inactive)
		} else {
			if ele.Status() == Chosen {
				ele.SetSelectable(false)
			} else {
				ele.SetSelectable(true)
				ele.SetStatus(Active)
			}
		}
	}
}

// UpdateWithoutRandomization updates the p device without randomizing its
// port fragments
func (p *Pwner) UpdateWithoutRandomization() {
	// Update the select and status states of each PwnerElement. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := &p.frags[i]

		if !p.IsSelectable(i) {
			ele.SetSelectable(false)
			ele.SetStatus(Inactive)
		} else {
			if ele.Status() == Chosen {
				ele.SetSelectable(false)
			} else {
				ele.SetSelectable(true)
				ele.SetStatus(Active)
			}
		}
	}
}

// IsSelectable determines if the port fragment found at the given index can
// be selected by the player. It is selectable only if the port fragment
// maps to the same port fragment in the Key.
func (p Pwner) IsSelectable(i int) bool {
	element := p.Get(i)
	return element.Frag() == p.key.Get(i)
}

// TODO: Find a better solution to solving cyclic dependency with PortMatrix

func (p *Pwner) SetMatrix(m *PortMatrix) { p.matrix = m }
