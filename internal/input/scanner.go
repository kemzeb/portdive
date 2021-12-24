package input

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

// scanner.go provides the token and scanning logic for user inputs while the
//game is running.

// Token defines a token spawned from the Scanner
type Token uint

const (
	SELECTION Token = iota // alias to a keyword

	INDEX // an index value either for the Pwner or Hex Dump

	ILLEGAL
	EOS // end of a statement
)

// Scanner provides lexical analysis logic.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns an instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewReader(r)}
}

// Scan returns the next token and literal value, else an error.
func (s *Scanner) Scan() (tok Token, lit string) {
	ch, err := s.read()
	if err != nil {
		return EOS, ""
	}
	if unicode.IsSpace(ch) {
		s.scanSpace()
		ch, err = s.read()
		// Since we iterated the buffer ignoring space runes, check to make sure
		// no errors were found.
		if err != nil {
			return EOS, ""
		}
	}
	if ch == 'p' || ch == 'd' {
		return SELECTION, string(ch)
	} else if unicode.IsDigit(ch) {
		s.unread()
		return s.scanIndex()
	}
	return ILLEGAL, string(ch)

}

// scanSpace consumes the current rune and all adjacent space runes. Since
// space is irrelevant in this context, we do not return any state.
func (s *Scanner) scanSpace() {
	// Consume subsequent space-related runes until either we reached the end
	// or we get a non-whitespace character. If we reach either condition, break
	// from the loop.
	for {
		if ch, err := s.read(); err != nil {
			break
		} else if !unicode.IsSpace(ch) {
			s.unread()
			break
		}
	}
}

// scanIndex consumes the current rune and all adjacent integer runes.
func (s *Scanner) scanIndex() (tok Token, lit string) {
	// Create a buffer and read the current integer rune into it.
	var buf bytes.Buffer
	ch, _ := s.read()
	buf.WriteRune(ch)

	// Consume subsequent integer runes
	for {
		if ch, err := s.read(); err != nil {
			break
		} else if !unicode.IsDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return INDEX, buf.String()
}

// read returns the next rune from the buffered reader. Returns an error if an
// error occurs or the end to the buffer has been reached.
func (s *Scanner) read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return -1, err
	}
	return ch, nil
}

// unread places the previously-read rune back on to the reader.
func (s *Scanner) unread() {
	s.r.UnreadRune()
}
