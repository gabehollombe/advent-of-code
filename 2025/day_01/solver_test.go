package main

import (
	"testing"
)

func TestCountZeroPositions(t *testing.T) {
	testCases := []struct {
		name      string
		rotations []Rotation
		wantCount int
	}{
		{
			name: "Example from problem",
			rotations: []Rotation{
				{Direction: Left, Distance: 68},
				{Direction: Left, Distance: 30},
				{Direction: Right, Distance: 48},
				{Direction: Left, Distance: 5},
				{Direction: Right, Distance: 60},
				{Direction: Left, Distance: 55},
				{Direction: Left, Distance: 1},
				{Direction: Left, Distance: 99},
				{Direction: Right, Distance: 14},
				{Direction: Left, Distance: 82},
			},
			wantCount: 3,
		},
		{
			name:      "Empty rotations",
			rotations: []Rotation{},
			wantCount: 0,
		},
		{
			name: "Single rotation to zero",
			rotations: []Rotation{
				{Direction: Right, Distance: 50},
			},
			wantCount: 1,
		},
		{
			name: "Single rotation not to zero",
			rotations: []Rotation{
				{Direction: Right, Distance: 10},
			},
			wantCount: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountZeroPositions(tc.rotations)
			if got != tc.wantCount {
				t.Errorf("CountZeroPositions() = %d, want %d", got, tc.wantCount)
			}
		})
	}
}

func TestCountZeroPositionsDetailedExample(t *testing.T) {
	// This test verifies the step-by-step progression from the problem example
	rotations := []Rotation{
		{Direction: Left, Distance: 68},   // 50 -> 82
		{Direction: Left, Distance: 30},   // 82 -> 52
		{Direction: Right, Distance: 48},  // 52 -> 0 (count: 1)
		{Direction: Left, Distance: 5},    // 0 -> 95
		{Direction: Right, Distance: 60},  // 95 -> 55
		{Direction: Left, Distance: 55},   // 55 -> 0 (count: 2)
		{Direction: Left, Distance: 1},    // 0 -> 99
		{Direction: Left, Distance: 99},   // 99 -> 0 (count: 3)
		{Direction: Right, Distance: 14},  // 0 -> 14
		{Direction: Left, Distance: 82},   // 14 -> 32
	}

	dial := NewDial()
	expectedPositions := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
	
	for i, rotation := range rotations {
		dial.Rotate(rotation)
		if dial.Position != expectedPositions[i] {
			t.Errorf("After rotation %d, position = %d, want %d", i+1, dial.Position, expectedPositions[i])
		}
	}

	count := CountZeroPositions(rotations)
	if count != 3 {
		t.Errorf("CountZeroPositions() = %d, want 3", count)
	}
}

func TestCountZeroClicks(t *testing.T) {
	testCases := []struct {
		name      string
		rotations []Rotation
		wantCount int
	}{
		{
			name: "Part 2 Example from problem",
			rotations: []Rotation{
				{Direction: Left, Distance: 68},   // 50 -> 82, crosses 0 once
				{Direction: Left, Distance: 30},   // 82 -> 52, no cross
				{Direction: Right, Distance: 48},  // 52 -> 0, ends at 0 (1 click)
				{Direction: Left, Distance: 5},    // 0 -> 95, no cross
				{Direction: Right, Distance: 60},  // 95 -> 55, crosses 0 once
				{Direction: Left, Distance: 55},   // 55 -> 0, ends at 0 (1 click)
				{Direction: Left, Distance: 1},    // 0 -> 99, no cross
				{Direction: Left, Distance: 99},   // 99 -> 0, ends at 0 (1 click)
				{Direction: Right, Distance: 14},  // 0 -> 14, no cross
				{Direction: Left, Distance: 82},   // 14 -> 32, crosses 0 once
			},
			wantCount: 6, // 3 from crossing during rotation + 3 from ending at 0
		},
		{
			name:      "Empty rotations",
			rotations: []Rotation{},
			wantCount: 0,
		},
		{
			name: "Single large rotation",
			rotations: []Rotation{
				{Direction: Right, Distance: 1000},
			},
			wantCount: 10, // Crosses 0 ten times
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountZeroClicks(tc.rotations)
			if got != tc.wantCount {
				t.Errorf("CountZeroClicks() = %d, want %d", got, tc.wantCount)
			}
		})
	}
}