// parser.go provides parsing logic for player-defined inputs from a command
// prompt.

package input

import (
	"errors"
	"strconv"
)

// SelectionType represents whether the player wishes to select the Pwner or the
// HexDump to interact with.
type SelectionType int

const (
	PwnrSelection SelectionType = iota
	MatrixSelection
)

// SelectionStatement represents a successful parsing of player input, where its
// fields will be utilized by the game
type SelectionStatement struct {
	Type  SelectionType
	Index int
}

// Parser represents a parser for player input
type Parser struct {
	scanner *Scanner
}

// NewParser instantiates a Parser.
func NewParser() *Parser {
	return &Parser{NewScanner()}
}

// Parse enforces the following grammar rules:
//
// statement 		::= SelectionToken IndexToken EndOfStatementToken
//
// SelectionToken 	::= 'p' | 'd'
//
// IndexToken		::= an integer value
func (p *Parser) Parse(input string) (*SelectionStatement, error) {
	statement := SelectionStatement{}
	var err error

	p.scanner.Reset(input) // parse a new string

	// The initial token should be a "SELECTION" token
	if tok, lit := p.scanner.NextToken(); tok == SelectionToken {
		if lit == "p" {
			statement.Type = PwnrSelection
		} else {
			statement.Type = MatrixSelection
		}
	} else {
		return nil, errors.New("parse: expected a \"selection\"")
	}

	// The second token should be an IndexToken
	if tok, lit := p.scanner.NextToken(); tok == IndexToken {
		statement.Index, err = strconv.Atoi(lit)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("parse: invalid grammar; expected an \"index\"")
	}

	// The final token should be an EndOfStatementToken
	if tok, _ := p.scanner.NextToken(); tok != EndOfStatementToken {
		return nil, errors.New("parse: invalid grammar; expected \"end of statement\"")
	}
	return &statement, nil
}
