package game

import (
	tl "github.com/JoelOtter/termloop"
)

// PortRow represents a row within the PortMatrix.
type PortRow struct {
	frags      []int   // represents port address fragments
	selectable bool    // flag for determining if the player can select the row
	status     tl.Attr // represents the state the row is in to the player
}

// NewPortRow instantiates a PortRow.
func NewPortRow(f []int) *PortRow {
	return &PortRow{frags: f, selectable: true, status: Active}
}

// Get returns the port fragment found at the given index.
func (r PortRow) Get(i int) int {
	if !(i < r.Len()) {
		panic("PortRow: overflow error")
	}
	return r.frags[i]
}

// Len returns the number of columns in the row.
func (r PortRow) Len() int { return len(r.frags) }

func (r PortRow) Selectable() bool      { return r.selectable }
func (r *PortRow) SetSelectable(s bool) { r.selectable = s }
func (r PortRow) Status() tl.Attr       { return r.status }
func (r *PortRow) SetStatus(s tl.Attr)  { r.status = s }

// PortMatrix represents the collection of port numbers.
type PortMatrix struct {
	rows  []PortRow
	pwner *Pwner
}

// NewPortMatrix instantiates a PortMatrix
func NewPortMatrix(r []PortRow, p *Pwner) *PortMatrix {
	return &PortMatrix{rows: r, pwner: p}
}

// Get returns the row found at the given index.
func (m PortMatrix) Get(i int) *PortRow {
	return &m.rows[i]
}

// Len returns the number of rows in the PortMatrix.
func (m PortMatrix) Len() int {
	return len(m.rows)
}

// Update updates the state of the rows in the PortMatrix.
func (m *PortMatrix) Update() {
	for i := 0; i < m.Len(); i++ {
		row := m.Get(i)
		if m.IsSelectable(i) {
			row.SetSelectable(true)
			row.SetStatus(Active)
		} else {
			row.SetSelectable(false)
			row.SetStatus(Inactive)
		}
	}
}

// IsSelectable determines if the row found with the given index can be selected
// by the player. It is not selectable only if we have an element in the
// PortRow where its index matches with the Pwner's element, the Pwner's
// element has been chosen, and the two elements' fragment values do not equal.
func (m PortMatrix) IsSelectable(i int) bool {
	row := m.Get(i)
	colLen := row.Len()

	for j := 0; j < colLen; j++ {
		pwnerEle := m.pwner.Get(j)

		if pwnerEle.Status() == Chosen && pwnerEle.Frag() != row.Get(j) {
			return false
		}
	}
	return true
}
