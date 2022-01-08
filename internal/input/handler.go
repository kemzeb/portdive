// Package input implements the logic necessary to parse user input for the
// game.
package input

import (
	"bufio"
	"os"
)

type InHandlerResult struct {
	Selection *SelectionStatement
	Err       error
}

type InHandler struct {
	parser Parser
	input  chan<- *InHandlerResult
}

func NewInHandler(in chan<- *InHandlerResult) *InHandler {
	return &InHandler{*NewParser(), in}
}

func (i *InHandler) Handle(finish chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		selection, err := i.parser.Parse(input)
		switch {
		case <-finish:
			return
		default:
			if err != nil {
				i.input <- &InHandlerResult{nil, err}
			} else {
				i.input <- &InHandlerResult{selection, nil}
			}
		}
	}
}
