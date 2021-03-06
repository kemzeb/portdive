package game

import (
	"math/rand"
)

// Key represents a row within the hex dump that represents the Key.
type Key struct {
	frags       []int        // the port fragment values that make up the Key
	chosenIndex int          // the index of a row in PortMatrix that was chosen
	seed        *rand.Source // used for choosing a random PortRow in PortMatrix
}

// NewKey instantiates a Key.
func NewKey(s *rand.Source) *Key {
	return &Key{chosenIndex: -1, seed: s}
}

// Get returns a hex value within the Key.
func (k Key) Get(i int) int { return k.frags[i] }

// RandomizeKey chooses a random row within the hex dump. This implementation is
// bound to change.
func (k *Key) RandomizeKey(dump *PortMatrix) {
	rowLen := dump.Len()
	colLen := dump.Get(0).Len()
	k.chosenIndex = RandomizeInt(*k.seed, 0, rowLen-1)

	for i := 0; i < colLen; i++ {
		row := dump.Get(k.chosenIndex)
		k.frags = append(k.frags, row.Get(i))
	}
}

// Reset resets the state of the Key.
func (k *Key) Reset() {
	k.frags = nil
	k.chosenIndex = -1
}

func (k Key) ChosenIndex() int { return k.chosenIndex }
