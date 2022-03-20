package game

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
)

// Element represents an element within the Pwner device.
type Element struct {
	frag       int
	selectable bool
	status     tl.Attr
}

func (e Element) Frag() int             { return e.frag }
func (e Element) Selectable() bool      { return e.selectable }
func (e *Element) SetSelectable(s bool) { e.selectable = s }
func (e Element) Status() tl.Attr       { return e.status }
func (e *Element) SetStatus(s tl.Attr)  { e.status = s }

// Pwner represents the Pwner device.
type Pwner struct {
	frags            []Element
	matrix           *PortMatrix
	k                *Key
	seed             *rand.Source // used for randomizing port numbers
	availableIndices []int
}

func NewPwner(k *Key, s *rand.Source) *Pwner {
	p := Pwner{k: k, seed: s}
	return &p
}

// Get returns the port fragment found at the given index.
func (p Pwner) Get(i int) *Element {
	if !(i < p.Len()) {
		panic("Pwner: overflow error")
	}
	return &p.frags[i]
}

// Len returns the number of port fragments.
func (p Pwner) Len() int { return len(p.frags) }

// Init instantiates a slice of Element types with a length = the # of
// columns in the PortMatrix. It chooses random port fragments in each column
// for each element in the Pwner.
func (p *Pwner) Init() {
	rowLen := p.matrix.Len()
	colLen := p.matrix.Get(0).Len()
	p.availableIndices = make([]int, colLen)

	for i := 0; i < colLen; i++ {
		randomRowIndex := RandomizeInt(*p.seed, 0, rowLen-1)
		randomFrag := p.matrix.Get(randomRowIndex).Get(i)
		element := Element{randomFrag, false, Inactive}
		p.frags = append(p.frags, element)

		p.availableIndices[i] = i
	}
}

// Update updates the state of the port fragments.
func (p *Pwner) Update() {
	for i := 0; i < len(p.availableIndices); i++ {
		// Remove elements from the available list that were chosen
		if p.Get(p.availableIndices[i]).Status() == Chosen {
			p.availableIndices = append(p.availableIndices[:i],
				p.availableIndices[i+1:]...)
		}
	}
	// Determine the random # of Pwner elements that should be replaced
	replaceLen := RandomizeInt(*p.seed, 0, len(p.availableIndices))
	for i := 0; i < replaceLen; i++ {
		randAvailInd := RandomizeInt(*p.seed, 0, len(p.availableIndices)-1)
		randMatrixRowIndex := RandomizeInt(*p.seed, 0, p.matrix.Len()-1, i)
		randRow := p.matrix.Get(randMatrixRowIndex)
		p.frags[p.availableIndices[randAvailInd]].frag = randRow.Get(p.availableIndices[randAvailInd])
	}

	// Update the select and status states of each Element. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := p.Get(i)

		if !p.IsSelectable(i) {
			if ele.Status() == Active {
				ele.SetStatus(Inactive)
			}
			ele.SetSelectable(false)
		} else {
			ele.SetSelectable(true)
			ele.SetStatus(Active)
		}
	}
}

// UpdateWithoutRandomization updates the Pwner device without randomizing its
// port fragments
func (p *Pwner) UpdateWithoutRandomization() {
	// Update the select and status states of each Element. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := p.Get(i)

		if !p.IsSelectable(i) {
			if ele.Status() == Active {
				ele.SetStatus(Inactive)
			}
			ele.SetSelectable(false)
		} else {
			ele.SetSelectable(true)
			ele.SetStatus(Active)
		}
	}
}

// IsSelectable determines if the port fragment found at the given index can
// be selected by the player. It is selectable only if the port fragment
// maps to the same port fragment in the Key.
func (p Pwner) IsSelectable(i int) bool {
	ele := p.Get(i)
	return ele.Frag() == p.k.Get(i) && ele.Status() != Chosen
}

// TODO: Find a better solution to solving cyclic dependency with PortMatrix

func (p *Pwner) SetMatrix(m *PortMatrix) { p.matrix = m }
