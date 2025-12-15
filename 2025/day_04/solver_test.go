package main

import (
	"os"
	"testing"
)

func TestCountAccessibleRolls(t *testing.T) {
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

	tmpfile, err := os.CreateTemp("", "test_solver_*.txt")
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

	// According to the problem, there should be 13 accessible rolls
	expectedAccessible := 13
	got := CountAccessibleRolls(room)
	if got != expectedAccessible {
		t.Errorf("CountAccessibleRolls() = %d, want %d", got, expectedAccessible)
	}
}

func TestCellIsAccessible(t *testing.T) {
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

	tmpfile, err := os.CreateTemp("", "test_accessible_*.txt")
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

	// Test specific cells from the example
	// According to the problem output, these should be accessible (marked with x):
	// Row 0: positions 2,3,6 and 8
	// Row 1: position 0
	// Row 2: position 6
	// Row 4: positions 0,9
	// Row 7: position 0
	// Row 9: positions 0,2,8

	testCases := []struct {
		x                  int
		y                  int
		shouldBeAccessible bool
		description        string
	}{
		{2, 0, true, "Position (2,0) should be accessible"},
		{3, 0, true, "Position (3,0) should be accessible"},
		{6, 0, true, "Position (6,0) should be accessible"},
		{7, 0, false, "Position (7,0) should NOT be accessible"},
		{8, 0, true, "Position (8,0) should be accessible"},
		{0, 1, true, "Position (0,1) should be accessible"},
		{6, 2, true, "Position (6,2) should be accessible"},
		{0, 4, true, "Position (0,4) should be accessible"},
		{9, 4, true, "Position (9,4) should be accessible"},
		{0, 7, true, "Position (0,7) should be accessible"},
		{0, 9, true, "Position (0,9) should be accessible"},
		{2, 9, true, "Position (2,9) should be accessible"},
		{8, 9, true, "Position (8,9) should be accessible"},
		// Test some that should NOT be accessible (surrounded by too many rolls)
		{1, 1, false, "Position (1,1) should NOT be accessible (surrounded)"},
		{2, 2, false, "Position (2,2) should NOT be accessible (surrounded)"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			cell := room.FindCell(tc.x, tc.y)
			if cell == nil {
				t.Fatalf("Could not find cell at (%d,%d)", tc.x, tc.y)
			}

			// Only test paper rolls, not empty cells
			if cell.IsEmpty {
				t.Skip("Skipping empty cell")
			}

			got := cell.IsAccessible()
			if got != tc.shouldBeAccessible {
				t.Errorf("Cell at (%d,%d).IsAccessible() = %v, want %v", tc.x, tc.y, got, tc.shouldBeAccessible)
			}
		})
	}
}
