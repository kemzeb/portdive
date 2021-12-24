package input

import (
	"errors"
	"io"
	"strconv"
)

// parser.go provides parsing logic for user-defined inputs from a command
// prompt.

// SelectionType represents whether the user wishes to select the Pwner or the
// HexDump to interact with.
type SelectionType int

const (
	PWNR_SELECTION SelectionType = iota
	HEXDUMP_SELECTION
)

// SelectionStatement represents a sucessful parsing of user input, where its
// fields will be utilized by the game
type SelectionStatement struct {
	Selection SelectionType
	Index     int
}

// Parser represents a parser for user input
type Parser struct {
	s   *Scanner
	buf struct {
		tok  Token  // previously-read token
		lit  string // previously-read literal
		flag uint8  // true = read from this buffer, false = read next token & literal
	}
}

// NewParser instantiates a Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse enforces the following grammar rules:
//
// statement 	::= SELECTION INDEX EOS
//
// SELECTION 	::= 'p' | 'd'
//
// INDEX		::= an integer value
func (p *Parser) Parse() (*SelectionStatement, error) {
	statement := SelectionStatement{}
	var err error

	// The initial token should be a "SELECTION" token
	if tok, lit := p.s.Scan(); tok == SELECTION {
		if lit == "p" {
			statement.Selection = PWNR_SELECTION
		} else {
			statement.Selection = HEXDUMP_SELECTION
		}
	} else {
		return nil, errors.New("parser: invalid grammar; expected a \"SELECTION\"")
	}

	// The second token should be an "INDEX" token
	if tok, lit := p.s.Scan(); tok == INDEX {
		statement.Index, err = strconv.Atoi(lit)
		// Index value may be out of representable range
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("parser: invalid grammar; expected \"INDEX\"")
	}

	// The final token should be a "EOS" token
	if tok, _ := p.s.Scan(); tok != EOS {
		return nil, errors.New("parser: invalid grammar; expected \"EOS\"")
	}
	return &statement, nil
}
