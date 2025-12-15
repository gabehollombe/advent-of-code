package main

import (
	"testing"
)

func TestNewDial(t *testing.T) {
	dial := NewDial()
	if dial.Position != 50 {
		t.Errorf("NewDial() position = %d, want 50", dial.Position)
	}
}

func TestDialRotate(t *testing.T) {
	testCases := []struct {
		name         string
		startPos     int
		rotation     Rotation
		wantPosition int
	}{
		{
			name:     "Right rotation without wrap",
			startPos: 11,
			rotation: Rotation{Direction: Right, Distance: 8},
			wantPosition: 19,
		},
		{
			name:     "Left rotation without wrap",
			startPos: 11,
			rotation: Rotation{Direction: Left, Distance: 8},
			wantPosition: 3,
		},
		{
			name:     "Right rotation with wrap",
			startPos: 95,
			rotation: Rotation{Direction: Right, Distance: 10},
			wantPosition: 5,
		},
		{
			name:     "Left rotation with wrap",
			startPos: 5,
			rotation: Rotation{Direction: Left, Distance: 10},
			wantPosition: 95,
		},
		{
			name:     "Left from 0 wraps to 99",
			startPos: 0,
			rotation: Rotation{Direction: Left, Distance: 1},
			wantPosition: 99,
		},
		{
			name:     "Right from 99 wraps to 0",
			startPos: 99,
			rotation: Rotation{Direction: Right, Distance: 1},
			wantPosition: 0,
		},
		{
			name:     "Example: Start at 50, L68 -> 82",
			startPos: 50,
			rotation: Rotation{Direction: Left, Distance: 68},
			wantPosition: 82,
		},
		{
			name:     "Example: From 82, L30 -> 52",
			startPos: 82,
			rotation: Rotation{Direction: Left, Distance: 30},
			wantPosition: 52,
		},
		{
			name:     "Example: From 52, R48 -> 0",
			startPos: 52,
			rotation: Rotation{Direction: Right, Distance: 48},
			wantPosition: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dial := &Dial{Position: tc.startPos}
			dial.Rotate(tc.rotation)
			if dial.Position != tc.wantPosition {
				t.Errorf("After rotation, position = %d, want %d", dial.Position, tc.wantPosition)
			}
		})
	}
}

func TestDialIsAtZero(t *testing.T) {
	testCases := []struct {
		name     string
		position int
		want     bool
	}{
		{
			name:     "At zero",
			position: 0,
			want:     true,
		},
		{
			name:     "Not at zero",
			position: 50,
			want:     false,
		},
		{
			name:     "At 99",
			position: 99,
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dial := &Dial{Position: tc.position}
			if got := dial.IsAtZero(); got != tc.want {
				t.Errorf("IsAtZero() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDialCountZeroClicks(t *testing.T) {
	testCases := []struct {
		name      string
		startPos  int
		rotation  Rotation
		wantCount int
		wantPos   int
	}{
		{
			name:      "Right rotation crossing 0 once",
			startPos:  95,
			rotation:  Rotation{Direction: Right, Distance: 10},
			wantCount: 1,
			wantPos:   5,
		},
		{
			name:      "Left rotation crossing 0 once",
			startPos:  5,
			rotation:  Rotation{Direction: Left, Distance: 10},
			wantCount: 1,
			wantPos:   95,
		},
		{
			name:      "Right rotation not crossing 0",
			startPos:  10,
			rotation:  Rotation{Direction: Right, Distance: 20},
			wantCount: 0,
			wantPos:   30,
		},
		{
			name:      "Left rotation not crossing 0",
			startPos:  50,
			rotation:  Rotation{Direction: Left, Distance: 20},
			wantCount: 0,
			wantPos:   30,
		},
		{
			name:      "Example: 50 L68 crosses 0 once",
			startPos:  50,
			rotation:  Rotation{Direction: Left, Distance: 68},
			wantCount: 1,
			wantPos:   82,
		},
		{
			name:      "Example: 95 R60 crosses 0 once",
			startPos:  95,
			rotation:  Rotation{Direction: Right, Distance: 60},
			wantCount: 1,
			wantPos:   55,
		},
		{
			name:      "Example: 14 L82 crosses 0 once",
			startPos:  14,
			rotation:  Rotation{Direction: Left, Distance: 82},
			wantCount: 1,
			wantPos:   32,
		},
		{
			name:      "Multiple wraps: R1000 from 50",
			startPos:  50,
			rotation:  Rotation{Direction: Right, Distance: 1000},
			wantCount: 10,
			wantPos:   50,
		},
		{
			name:      "Right rotation ending at 0",
			startPos:  50,
			rotation:  Rotation{Direction: Right, Distance: 50},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Left rotation ending at 0",
			startPos:  50,
			rotation:  Rotation{Direction: Left, Distance: 50},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Right from 0 not crossing (stays in range)",
			startPos:  0,
			rotation:  Rotation{Direction: Right, Distance: 50},
			wantCount: 0,
			wantPos:   50,
		},
		{
			name:      "Left from 0 with distance < 100 (no wrap)",
			startPos:  0,
			rotation:  Rotation{Direction: Left, Distance: 50},
			wantCount: 0,
			wantPos:   50,
		},
		{
			name:      "Left from 0 with exact wrap (distance = 100)",
			startPos:  0,
			rotation:  Rotation{Direction: Left, Distance: 100},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Right from 0 with exact wrap (distance = 100)",
			startPos:  0,
			rotation:  Rotation{Direction: Right, Distance: 100},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Left rotation with multiple wraps",
			startPos:  25,
			rotation:  Rotation{Direction: Left, Distance: 250},
			wantCount: 3,
			wantPos:   75,
		},
		{
			name:      "Right rotation with multiple wraps",
			startPos:  25,
			rotation:  Rotation{Direction: Right, Distance: 250},
			wantCount: 2,
			wantPos:   75,
		},
		{
			name:      "Left from 1 wrapping to 0",
			startPos:  1,
			rotation:  Rotation{Direction: Left, Distance: 1},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Right exactly to boundary (99)",
			startPos:  50,
			rotation:  Rotation{Direction: Right, Distance: 49},
			wantCount: 0,
			wantPos:   99,
		},
		{
			name:      "Left with remainder exactly equal to startPos",
			startPos:  30,
			rotation:  Rotation{Direction: Left, Distance: 30},
			wantCount: 1,
			wantPos:   0,
		},
		{
			name:      "Left with remainder one less than startPos",
			startPos:  30,
			rotation:  Rotation{Direction: Left, Distance: 29},
			wantCount: 0,
			wantPos:   1,
		},
		{
			name:      "Right with distance 0",
			startPos:  50,
			rotation:  Rotation{Direction: Right, Distance: 0},
			wantCount: 0,
			wantPos:   50,
		},
		{
			name:      "Left with distance 0",
			startPos:  50,
			rotation:  Rotation{Direction: Left, Distance: 0},
			wantCount: 0,
			wantPos:   50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dial := &Dial{Position: tc.startPos}
			gotCount := dial.CountZeroClicks(tc.rotation)
			if gotCount != tc.wantCount {
				t.Errorf("CountZeroClicks() count = %d, want %d", gotCount, tc.wantCount)
			}
			if dial.Position != tc.wantPos {
				t.Errorf("After CountZeroClicks(), position = %d, want %d", dial.Position, tc.wantPos)
			}
		})
	}
}

func TestDialCountZeroClicks_CompleteExample(t *testing.T) {
	// This tests the complete example from problem_part_two.txt
	// The dial starts at 50 and goes through 10 rotations
	// Expected: 6 total zero clicks, ending at position 32
	
	rotations := []Rotation{
		{Direction: Left, Distance: 68},  // L68: 50 → 82, crosses 0 once
		{Direction: Left, Distance: 30},  // L30: 82 → 52, no crossing
		{Direction: Right, Distance: 48}, // R48: 52 → 0, ends at 0 (counts as 1)
		{Direction: Left, Distance: 5},   // L5: 0 → 95, no crossing
		{Direction: Right, Distance: 60}, // R60: 95 → 55, crosses 0 once
		{Direction: Left, Distance: 55},  // L55: 55 → 0, ends at 0 (counts as 1)
		{Direction: Left, Distance: 1},   // L1: 0 → 99, no crossing
		{Direction: Left, Distance: 99},  // L99: 99 → 0, ends at 0 (counts as 1)
		{Direction: Right, Distance: 14}, // R14: 0 → 14, no crossing
		{Direction: Left, Distance: 82},  // L82: 14 → 32, crosses 0 once
	}
	
	expectedCounts := []int{1, 0, 1, 0, 1, 1, 0, 1, 0, 1}
	expectedPositions := []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
	
	dial := NewDial() // Starts at 50
	totalZeroClicks := 0
	
	for i, rotation := range rotations {
		count := dial.CountZeroClicks(rotation)
		totalZeroClicks += count
		
		if count != expectedCounts[i] {
			t.Errorf("Rotation %d: got %d zero clicks, want %d", i+1, count, expectedCounts[i])
		}
		
		if dial.Position != expectedPositions[i] {
			t.Errorf("After rotation %d: position = %d, want %d", i+1, dial.Position, expectedPositions[i])
		}
	}
	
	if totalZeroClicks != 6 {
		t.Errorf("Total zero clicks = %d, want 6", totalZeroClicks)
	}
	
	if dial.Position != 32 {
		t.Errorf("Final position = %d, want 32", dial.Position)
	}
}