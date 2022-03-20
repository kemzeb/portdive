package main

import (
	"portdive/internal/game"
)

func main() {
	// TODO: Generate winnable port matrix rather than manually adding port addresses
	rows := &[]game.PortRow{
		*game.NewPortRow([]int{193, 68, 30, 20}),
		*game.NewPortRow([]int{193, 69, 40, 20}),
		*game.NewPortRow([]int{194, 66, 20, 30}),
		*game.NewPortRow([]int{194, 66, 40, 40}),
		*game.NewPortRow([]int{194, 67, 10, 20}),
		*game.NewPortRow([]int{194, 67, 10, 30}),
		*game.NewPortRow([]int{194, 67, 20, 20}),
		*game.NewPortRow([]int{194, 68, 20, 20}),
	}
	game := game.NewGame(rows)

	game.Start()
}
