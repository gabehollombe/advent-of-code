package main

const DialSize = 100

type Dial struct {
	Position int
}

func NewDial() *Dial {
	return &Dial{Position: 50}
}

func (d *Dial) Rotate(r Rotation) {
	switch r.Direction {
	case Right:
		d.Position = (d.Position + r.Distance) % DialSize
	case Left:
		// Handle negative modulo correctly for any distance
		d.Position = ((d.Position - r.Distance) % DialSize + DialSize) % DialSize
	}
}

func (d *Dial) IsAtZero() bool {
	return d.Position == 0
}

// CountZeroClicks counts how many times the dial clicks through position 0
// during a rotation, including the final position if it's 0
func (d *Dial) CountZeroClicks(r Rotation) int {
	count := 0
	startPos := d.Position
	
	switch r.Direction {
	case Right:
		// Count complete wraps around the dial
		count += r.Distance / DialSize
		
		// Check if partial rotation crosses 0 (from 99 to 0)
		if startPos+(r.Distance%DialSize) >= DialSize {
			count++
		}
		
	case Left:
		// Count complete wraps around the dial
		count += r.Distance / DialSize
		
		// Check if partial rotation crosses 0
		// We cross 0 when: startPos - remainder < 0, i.e., remainder >= startPos
		// Special case: starting at 0 doesn't count (handled by wraps only)
		remainder := r.Distance % DialSize
		if remainder > 0 && startPos > 0 && remainder >= startPos {
			count++
		}
	}
	
	// Update position
	d.Rotate(r)
	
	return count
}