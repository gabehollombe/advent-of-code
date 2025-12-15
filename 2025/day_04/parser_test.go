package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseInput(t *testing.T) {
	// Create a test input file with the example from the problem
	testInput := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	// Write test input to a temporary file
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Parse the input
	room, err := ParseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	// Verify we have the correct number of cells (10 rows x 10 cols = 100)
	expectedCellCount := 100
	if len(room.Cells) != expectedCellCount {
		t.Errorf("ParseInput() returned %d cells, want %d", len(room.Cells), expectedCellCount)
	}

	// Count empty cells (.) and paper rolls (@)
	emptyCount := 0
	paperCount := 0
	for _, cell := range room.Cells {
		if cell.IsEmpty {
			emptyCount++
		} else {
			paperCount++
		}
	}

	// From the example: count the @ symbols
	expectedPaperCount := 71 // counted from the example
	if paperCount != expectedPaperCount {
		t.Errorf("ParseInput() found %d paper rolls, want %d", paperCount, expectedPaperCount)
	}

	expectedEmptyCount := 29 // 100 - 71
	if emptyCount != expectedEmptyCount {
		t.Errorf("ParseInput() found %d empty cells, want %d", emptyCount, expectedEmptyCount)
	}
}

func TestCellNeighbors(t *testing.T) {
	// Create a test input file
	testInput := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	tmpfile, err := os.CreateTemp("", "test_neighbors_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	room, err := ParseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	// Test corner cell (0,0) - should have 3 neighbors: (1,0), (0,1), (1,1)
	cornerCell := room.FindCell(0, 0)
	if cornerCell == nil {
		t.Fatal("Could not find corner cell at (0,0)")
	}
	if len(cornerCell.Neighbors) != 3 {
		t.Errorf("Corner cell (0,0) has %d neighbors, want 3", len(cornerCell.Neighbors))
	}
	expectedCornerNeighbors := map[string]bool{
		"1,0": false, "0,1": false, "1,1": false,
	}
	for _, neighbor := range cornerCell.Neighbors {
		key := fmt.Sprintf("%d,%d", neighbor.X, neighbor.Y)
		if _, exists := expectedCornerNeighbors[key]; exists {
			expectedCornerNeighbors[key] = true
		} else {
			t.Errorf("Corner cell (0,0) has unexpected neighbor at (%d,%d)", neighbor.X, neighbor.Y)
		}
	}
	for pos, found := range expectedCornerNeighbors {
		if !found {
			t.Errorf("Corner cell (0,0) missing expected neighbor at %s", pos)
		}
	}

	// Test edge cell (5,0) - should have 5 neighbors: (4,0), (6,0), (4,1), (5,1), (6,1)
	edgeCell := room.FindCell(5, 0)
	if edgeCell == nil {
		t.Fatal("Could not find edge cell at (5,0)")
	}
	if len(edgeCell.Neighbors) != 5 {
		t.Errorf("Edge cell (5,0) has %d neighbors, want 5", len(edgeCell.Neighbors))
	}
	expectedEdgeNeighbors := map[string]bool{
		"4,0": false, "6,0": false, "4,1": false, "5,1": false, "6,1": false,
	}
	for _, neighbor := range edgeCell.Neighbors {
		key := fmt.Sprintf("%d,%d", neighbor.X, neighbor.Y)
		if _, exists := expectedEdgeNeighbors[key]; exists {
			expectedEdgeNeighbors[key] = true
		} else {
			t.Errorf("Edge cell (5,0) has unexpected neighbor at (%d,%d)", neighbor.X, neighbor.Y)
		}
	}
	for pos, found := range expectedEdgeNeighbors {
		if !found {
			t.Errorf("Edge cell (5,0) missing expected neighbor at %s", pos)
		}
	}

	// Test center cell (5,5) - should have 8 neighbors in all directions
	centerCell := room.FindCell(5, 5)
	if centerCell == nil {
		t.Fatal("Could not find center cell at (5,5)")
	}
	if len(centerCell.Neighbors) != 8 {
		t.Errorf("Center cell (5,5) has %d neighbors, want 8", len(centerCell.Neighbors))
	}
	expectedCenterNeighbors := map[string]bool{
		"4,4": false, "5,4": false, "6,4": false, // top row
		"4,5": false, "6,5": false, // middle row (left and right)
		"4,6": false, "5,6": false, "6,6": false, // bottom row
	}
	for _, neighbor := range centerCell.Neighbors {
		key := fmt.Sprintf("%d,%d", neighbor.X, neighbor.Y)
		if _, exists := expectedCenterNeighbors[key]; exists {
			expectedCenterNeighbors[key] = true
		} else {
			t.Errorf("Center cell (5,5) has unexpected neighbor at (%d,%d)", neighbor.X, neighbor.Y)
		}
	}
	for pos, found := range expectedCenterNeighbors {
		if !found {
			t.Errorf("Center cell (5,5) missing expected neighbor at %s", pos)
		}
	}
}
