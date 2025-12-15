package main

// CountAccessibleRolls counts how many rolls of paper can be accessed by a forklift
func CountAccessibleRolls(room *Room) int {
	numAccessibleCells := 0
	for _, cell := range room.Cells {
		if cell.IsAccessible() {
			numAccessibleCells++
		}
	}
	return numAccessibleCells
}
