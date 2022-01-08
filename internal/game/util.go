// util.go provides constant values and functions that are used across package
// game.

package game

import (
	"github.com/pterm/pterm"
	"math/rand"
	"time"
)

const Duration time.Duration = 1500 * time.Millisecond

const (
	MaxHexDumpRows    = 4
	MaxHexDumpColumns = 6
)

// Colors used to represent the status of hex dump rows and of the p hex
// values
const (
	Inactive = pterm.FgDarkGray
	Active   = pterm.FgLightGreen
	Chosen   = pterm.FgLightWhite
)

// RandomizeInt generates a random integer value given minimum and maximum
// ranges.
func RandomizeInt(seed rand.Source, start int, end int, exclude ...int) int {
	if len(exclude) >= end {
		panic("game/util: exclude length >= end length")
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
