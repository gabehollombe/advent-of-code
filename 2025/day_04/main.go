package main

import (
	"fmt"
	"log"
)

func main() {
	room, err := ParseInput("input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	accessibleCount := CountAccessibleRolls(room)
	fmt.Printf("Part One: Number of accessible rolls: %d\n", accessibleCount)

	totalRemovedRolls := 0
	for CountAccessibleRolls(room) > 0 {
		accessibleCells := []*Cell{}
		for _, cell := range room.Cells {
			if cell.IsAccessible() {
				accessibleCells = append(accessibleCells, cell)
			}
		}
		for _, cell := range accessibleCells {
			cell.IsEmpty = true
			totalRemovedRolls++
		}
	}

	fmt.Printf("Part Two: Number of total removable rolls: %d\n", totalRemovedRolls)

}
