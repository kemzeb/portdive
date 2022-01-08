package input

// RuneStack provides typical stack operations for runes. It is used by the
// Scanner in order to maintain a history of runes when scanning its input
// string.
type RuneStack []rune

// IsEmpty returns true if our stack is not empty, else false.
func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push appends a rune to the top of our stack.
func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}

// Pop removes the rune at the top of our stack and returns it.
func (s *RuneStack) Pop() rune {
	if s.IsEmpty() {
		return EOFRune
	}
	index := len(*s) - 1
	r := (*s)[index]
	*s = (*s)[:index]
	return r
}

// Peek returns the rune at the top of our stack.
func (s RuneStack) Peek() rune {
	if s.IsEmpty() {
		return EOFRune
	}
	topIndex := len(s) - 1
	return (s)[topIndex]
}

// RemoveAllElements simply resets the value to a new RuneStack.
func (s *RuneStack) RemoveAllElements() {
	*s = RuneStack{}
}
