// scanner.go provides the token and scanning logic for user inputs while the
//game is running.

package input

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Token defines a token spawned from the Scanner
type Token uint

const (
	EOFRune        rune  = -1
	SelectionToken Token = iota // alias to a keyword
	IndexToken                  // an index value either for the Pwner or Hex Dump
	IllegalToken
	EndOfStatementToken // end of a statement
)

// Scanner provides lexical analysis logic for user input.
type Scanner struct {
	input    string // the string to be lexed
	position int    // starting and position indices
	rewind   RuneStack
}

// NewScanner returns an instance of Scanner.
func NewScanner() *Scanner {
	return &Scanner{}
}

// Reset resets the fields of the lexer to their zero values with the exception
// to "input", where we replace it with "new".
func (s *Scanner) Reset(new string) {
	s.input = new
	s.position = 0
	s.rewind.RemoveAllElements()
}

// NextToken returns a type-value pair representing the next token scanned in
// the input string. It ignores any space characters and considers consecutive
// digits to be a number.
func (s *Scanner) NextToken() (tok Token, lit string) {
	r := s.scan()

	if r == EOFRune {
		return EndOfStatementToken, ""
	}
	if unicode.IsSpace(r) {
		s.scanSpace()
		r = s.scan()
		if r == EOFRune {
			return EndOfStatementToken, ""
		}
	}
	lit = string(r)
	if r == 'p' || r == 'd' {
		tok = SelectionToken
	} else if unicode.IsDigit(r) {
		tok = IndexToken
		s.unscan()
		lit = s.scanInteger()
	} else {
		tok = IllegalToken
	}
	return tok, lit

}

func (s *Scanner) scan() rune {
	var (
		r    rune
		size int // essentially the ending index byte for a rune
	)
	str := s.input[s.position:]
	if len(str) == 0 {
		return EOFRune
	}
	r, size = utf8.DecodeRuneInString(str)
	s.position += size
	s.rewind.Push(r)

	return r
}

func (s *Scanner) unscan() {
	r := s.rewind.Pop()
	if r > EOFRune {
		size := utf8.RuneLen(r)
		s.position -= size
	}
}

// scanSpace reads a rune and all subsequent runes until it cannot scan any
// further runes or if it reaches a non-space character.
func (s *Scanner) scanSpace() {
	r := s.scan()

	for r != EOFRune {
		if !unicode.IsSpace(r) {
			s.unscan()
			break
		}
		r = s.scan()
	}
}

// scanInteger reads a rune and all subsequent runes until it cannot scan
// any further runes or if we reach a non-digit character. It returns these
// subsequent digits as a string (i.e. it represents a number).
func (s *Scanner) scanInteger() string {
	builder := strings.Builder{}
	r := s.scan()

	for r != EOFRune {
		if !unicode.IsDigit(r) {
			s.unscan() // don't throw away a rune that may be scan later
			break
		}
		builder.WriteRune(r)
		r = s.scan()
	}
	return builder.String()
}
