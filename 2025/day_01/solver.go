package main

func CountZeroPositions(rotations []Rotation) int {
	dial := NewDial()
	count := 0

	for _, rotation := range rotations {
		dial.Rotate(rotation)
		if dial.IsAtZero() {
			count++
		}
	}

	return count
}

func CountZeroClicks(rotations []Rotation) int {
	dial := NewDial()
	count := 0

	for _, rotation := range rotations {
		count += dial.CountZeroClicks(rotation)
	}

	return count
}