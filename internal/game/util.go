// util.go provides constant values and functions that are used across package
// game.

package game

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

const Duration = 1700 * time.Millisecond

const (
	MaxHexDumpRows    = 4
	MaxHexDumpColumns = 6
)

// Colors used to represent the status of port dump rows and of the Pwner port
// values
const (
	Inactive = tl.ColorDefault
	Active   = tl.ColorGreen
	Chosen   = tl.ColorYellow
	Hover    = tl.ColorMagenta
)

// RandomizeInt generates a random integer value given minimum and maximum
// ranges.
//
// 3 preconditions must be met if exclude is used:
//
// 1. Must be sorted in ascending order.
//
// 2. Contains integers within the specified range.
//
// 3. Contains integers that are unique.
func RandomizeInt(seed rand.Source, start int, end int, exclude ...int) int {
	if len(exclude) > end {
		panic("game/util: exclude length > end")
	}
	rand := rand.New(seed)
	random := start + rand.Intn(end-start+1-len(exclude))
	for _, ex := range exclude {
		if random < ex {
			break
		}
		random++
	}
	return random
}
