package game

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"portdive/internal"
	"sync"
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
	frags  []Element
	matrix *PortMatrix
	k      *Key
	seed   *rand.Source // used for randomizing port numbers
	lock   sync.Mutex
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
//columns in the PortMatrix. It chooses random port fragments in each column...
// for each element in the Pwner.
func (p *Pwner) Init() {
	rowLen := p.matrix.Len()
	colLen := p.matrix.Get(0).Len()

	for i := 0; i < colLen; i++ {
		randomRowIndex := internal.RandomizeInt(*p.seed, 0, rowLen-1)
		randomFrag := p.matrix.Get(randomRowIndex).Get(i)
		element := Element{randomFrag, false, internal.Inactive}
		p.frags = append(p.frags, element)
	}
}

// Update updates the state of the port fragments.
func (p *Pwner) Update() {
	p.lock.Lock()
	defer p.lock.Unlock()
	// Determine the random # of Pwner elements that should be replaced
	replaceLen := internal.RandomizeInt(*p.seed, 0, p.Len()/2)
	var randPwnerIndex int
	var excludedPwnerIndicies []int

	// Determine Pwner indices to exclude
	for i := 0; i < replaceLen; i++ {
		if p.Get(i).Status() == internal.Chosen {
			excludedPwnerIndicies = append(excludedPwnerIndicies, i)
		}
	}

	for i := 0; i < replaceLen; i++ {
		randPwnerIndex = internal.RandomizeInt(*p.seed, 0, p.Len()-1, excludedPwnerIndicies...)
		randMatrixRowIndex := internal.RandomizeInt(*p.seed, 0, p.matrix.Len()-1, i)
		randRow := p.matrix.Get(randMatrixRowIndex)
		p.frags[randPwnerIndex].frag = randRow.Get(randPwnerIndex)
	}

	// Update the select and status states of each Element. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := &p.frags[i]

		if !p.IsSelectable(i) {
			ele.SetSelectable(false)
			ele.SetStatus(internal.Inactive)
		} else {
			if ele.Status() == internal.Chosen {
				ele.SetSelectable(false)
			} else {
				ele.SetSelectable(true)
				ele.SetStatus(internal.Active)
			}
		}
	}
}

// UpdateWithoutRandomization updates the Pwner device without randomizing its
// port fragments
func (p *Pwner) UpdateWithoutRandomization() {
	p.lock.Lock()
	defer p.lock.Unlock()
	// Update the select and status states of each Element. Note that the
	// Chosen status is not determined here.
	for i := 0; i < p.Len(); i++ {
		ele := p.Get(i)

		if !p.IsSelectable(i) {
			if ele.Status() == internal.Active {
				ele.SetStatus(internal.Inactive)
			}
			ele.SetSelectable(false)

		} else {
			ele.SetSelectable(true)
			ele.SetStatus(internal.Active)
		}
	}
}

// IsSelectable determines if the port fragment found at the given index can
// be selected by the player. It is selectable only if the port fragment
// maps to the same port fragment in the Key.
func (p Pwner) IsSelectable(i int) bool {
	ele := p.Get(i)
	return ele.Frag() == p.k.Get(i) && ele.Status() != internal.Chosen
}

// TODO: Find a better solution to solving cyclic dependency with PortMatrix

func (p *Pwner) SetMatrix(m *PortMatrix) { p.matrix = m }
