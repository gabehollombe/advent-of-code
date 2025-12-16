package main

import (
	"os"
	"testing"
)

func TestParseRange(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantRange FreshRange
		wantErr   bool
	}{
		{
			name:  "Valid range 3-5",
			input: "3-5",
			wantRange: FreshRange{
				Start: 3,
				End:   5,
			},
			wantErr: false,
		},
		{
			name:  "Valid range 10-14",
			input: "10-14",
			wantRange: FreshRange{
				Start: 10,
				End:   14,
			},
			wantErr: false,
		},
		{
			name:  "Valid range 16-20",
			input: "16-20",
			wantRange: FreshRange{
				Start: 16,
				End:   20,
			},
			wantErr: false,
		},
		{
			name:    "Invalid format no dash",
			input:   "35",
			wantErr: true,
		},
		{
			name:    "Invalid format multiple dashes",
			input:   "3-5-7",
			wantErr: true,
		},
		{
			name:    "Invalid start value",
			input:   "abc-5",
			wantErr: true,
		},
		{
			name:    "Invalid end value",
			input:   "3-xyz",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRange, err := parseRange(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("parseRange() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr && gotRange != tc.wantRange {
				t.Errorf("parseRange() = %v, want %v", gotRange, tc.wantRange)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	// Create a temporary test file with the example input
	testInput := `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`
	tmpfile, err := os.CreateTemp("", "test_input_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testInput)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	db, err := ParseInput(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	// Check ranges - note that overlapping ranges are merged by AddRange
	// Input ranges: 3-5, 10-14, 16-20, 12-18
	// After merging: 3-5, and 10-20 (10-14 merged with 12-18 to become 10-18, then 16-20 merges with 10-18 to become 10-20)
	if len(db.Ranges) != 2 {
		t.Errorf("ParseInput() got %d ranges, want 2 (after merging overlaps)", len(db.Ranges))
	}

	// Verify that we have the expected ranges (order may vary)
	hasRange := func(r FreshRange) bool {
		for _, dbRange := range db.Ranges {
			if dbRange == r {
				return true
			}
		}
		return false
	}

	if !hasRange(FreshRange{Start: 3, End: 5}) {
		t.Errorf("ParseInput() missing range {3, 5}")
	}
	// Check for merged range - should be 10-20 after all merges
	if !hasRange(FreshRange{Start: 10, End: 20}) {
		t.Errorf("ParseInput() missing merged range {10, 20}, got ranges: %v", db.Ranges)
	}

	// Check available IDs
	expectedAvailable := []int{1, 5, 8, 11, 17, 32}

	if len(db.Available) != len(expectedAvailable) {
		t.Errorf("ParseInput() got %d available IDs, want %d", len(db.Available), len(expectedAvailable))
	}

	for i, expectedID := range expectedAvailable {
		if i >= len(db.Available) {
			break
		}
		if db.Available[i] != expectedID {
			t.Errorf("ParseInput() available[%d] = %d, want %d", i, db.Available[i], expectedID)
		}
	}
}

func TestAddRange(t *testing.T) {
	testCases := []struct {
		name           string
		initialRanges  []FreshRange
		rangeToAdd     FreshRange
		expectedRanges []FreshRange
	}{
		{
			name:          "Add to empty database",
			initialRanges: []FreshRange{},
			rangeToAdd:    FreshRange{Start: 5, End: 10},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
			},
		},
		{
			name: "Add non-overlapping range",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
			},
			rangeToAdd: FreshRange{Start: 15, End: 20},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 15, End: 20},
			},
		},
		{
			name: "Merge with overlapping range - extend end",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
			},
			rangeToAdd: FreshRange{Start: 8, End: 15},
			expectedRanges: []FreshRange{
				{Start: 5, End: 15},
			},
		},
		{
			name: "Merge with overlapping range - extend start",
			initialRanges: []FreshRange{
				{Start: 10, End: 15},
			},
			rangeToAdd: FreshRange{Start: 5, End: 12},
			expectedRanges: []FreshRange{
				{Start: 5, End: 15},
			},
		},
		{
			name: "Merge with overlapping range - extend both",
			initialRanges: []FreshRange{
				{Start: 10, End: 15},
			},
			rangeToAdd: FreshRange{Start: 5, End: 20},
			expectedRanges: []FreshRange{
				{Start: 5, End: 20},
			},
		},
		{
			name: "Merge with overlapping range - contained within existing",
			initialRanges: []FreshRange{
				{Start: 5, End: 20},
			},
			rangeToAdd: FreshRange{Start: 10, End: 15},
			expectedRanges: []FreshRange{
				{Start: 5, End: 20},
			},
		},
		{
			name: "Merge with adjacent range - touching at end",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
			},
			rangeToAdd: FreshRange{Start: 10, End: 15},
			expectedRanges: []FreshRange{
				{Start: 5, End: 15},
			},
		},
		{
			name: "Merge with adjacent range - touching at start",
			initialRanges: []FreshRange{
				{Start: 10, End: 15},
			},
			rangeToAdd: FreshRange{Start: 5, End: 10},
			expectedRanges: []FreshRange{
				{Start: 5, End: 15},
			},
		},
		{
			name: "Multiple ranges - merge with first",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 20, End: 25},
			},
			rangeToAdd: FreshRange{Start: 8, End: 12},
			expectedRanges: []FreshRange{
				{Start: 5, End: 12},
				{Start: 20, End: 25},
			},
		},
		{
			name: "Multiple ranges - merge with second",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 20, End: 25},
			},
			rangeToAdd: FreshRange{Start: 18, End: 22},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 18, End: 25},
			},
		},
		{
			name: "Multiple ranges - add between without overlap",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 20, End: 25},
			},
			rangeToAdd: FreshRange{Start: 12, End: 15},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 20, End: 25},
				{Start: 12, End: 15},
			},
		},
		{
			name: "Single point range - no overlap",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
			},
			rangeToAdd: FreshRange{Start: 15, End: 15},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
				{Start: 15, End: 15},
			},
		},
		{
			name: "Single point range - overlaps",
			initialRanges: []FreshRange{
				{Start: 5, End: 10},
			},
			rangeToAdd: FreshRange{Start: 7, End: 7},
			expectedRanges: []FreshRange{
				{Start: 5, End: 10},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := &Database{
				Ranges:    make([]FreshRange, len(tc.initialRanges)),
				Available: []int{},
			}
			copy(db.Ranges, tc.initialRanges)

			db.AddRange(tc.rangeToAdd)

			if len(db.Ranges) != len(tc.expectedRanges) {
				t.Errorf("AddRange() resulted in %d ranges, want %d", len(db.Ranges), len(tc.expectedRanges))
				t.Errorf("Got ranges: %v", db.Ranges)
				t.Errorf("Want ranges: %v", tc.expectedRanges)
				return
			}

			// Check if all expected ranges are present (order may vary for non-merged cases)
			for _, expected := range tc.expectedRanges {
				found := false
				for _, actual := range db.Ranges {
					if actual == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("AddRange() missing expected range %v in result %v", expected, db.Ranges)
				}
			}
		})
	}
}
