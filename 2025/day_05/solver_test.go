package main

import (
	"os"
	"testing"
)

func TestDatabaseIsFresh(t *testing.T) {
	db := &Database{
		Ranges: []FreshRange{
			{Start: 3, End: 5},
			{Start: 10, End: 14},
			{Start: 16, End: 20},
			{Start: 12, End: 18},
		},
	}

	testCases := []struct {
		name      string
		id        int
		wantFresh bool
		reason    string
	}{
		{
			name:      "ID 1 is spoiled",
			id:        1,
			wantFresh: false,
			reason:    "does not fall into any range",
		},
		{
			name:      "ID 5 is fresh",
			id:        5,
			wantFresh: true,
			reason:    "falls into range 3-5",
		},
		{
			name:      "ID 8 is spoiled",
			id:        8,
			wantFresh: false,
			reason:    "does not fall into any range",
		},
		{
			name:      "ID 11 is fresh",
			id:        11,
			wantFresh: true,
			reason:    "falls into range 10-14",
		},
		{
			name:      "ID 17 is fresh",
			id:        17,
			wantFresh: true,
			reason:    "falls into range 16-20 as well as range 12-18",
		},
		{
			name:      "ID 32 is spoiled",
			id:        32,
			wantFresh: false,
			reason:    "does not fall into any range",
		},
		{
			name:      "ID 3 is fresh (start of range)",
			id:        3,
			wantFresh: true,
			reason:    "falls into range 3-5 (inclusive)",
		},
		{
			name:      "ID 4 is fresh (middle of range)",
			id:        4,
			wantFresh: true,
			reason:    "falls into range 3-5",
		},
		{
			name:      "ID 10 is fresh (start of range)",
			id:        10,
			wantFresh: true,
			reason:    "falls into range 10-14 (inclusive)",
		},
		{
			name:      "ID 14 is fresh (end of range)",
			id:        14,
			wantFresh: true,
			reason:    "falls into range 10-14 (inclusive)",
		},
		{
			name:      "ID 20 is fresh (end of range)",
			id:        20,
			wantFresh: true,
			reason:    "falls into range 16-20 (inclusive)",
		},
		{
			name:      "ID 15 is fresh (overlapping ranges)",
			id:        15,
			wantFresh: true,
			reason:    "falls into range 12-18",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := db.IsFresh(tc.id)
			if got != tc.wantFresh {
				t.Errorf("IsFresh(%d) = %v, want %v (reason: %s)", tc.id, got, tc.wantFresh, tc.reason)
			}
		})
	}
}

func TestCountFreshIngredients(t *testing.T) {
	testCases := []struct {
		name      string
		db        *Database
		wantCount int
	}{
		{
			name: "Example from problem",
			db: &Database{
				Ranges: []FreshRange{
					{Start: 3, End: 5},
					{Start: 10, End: 14},
					{Start: 16, End: 20},
					{Start: 12, End: 18},
				},
				Available: []int{1, 5, 8, 11, 17, 32},
			},
			wantCount: 3, // IDs 5, 11, and 17 are fresh
		},
		{
			name: "All fresh",
			db: &Database{
				Ranges: []FreshRange{
					{Start: 1, End: 100},
				},
				Available: []int{1, 50, 100},
			},
			wantCount: 3,
		},
		{
			name: "None fresh",
			db: &Database{
				Ranges: []FreshRange{
					{Start: 10, End: 20},
				},
				Available: []int{1, 5, 25, 30},
			},
			wantCount: 0,
		},
		{
			name: "Empty available list",
			db: &Database{
				Ranges: []FreshRange{
					{Start: 1, End: 10},
				},
				Available: []int{},
			},
			wantCount: 0,
		},
		{
			name: "Empty ranges",
			db: &Database{
				Ranges:    []FreshRange{},
				Available: []int{1, 2, 3},
			},
			wantCount: 0,
		},
		{
			name: "Overlapping ranges",
			db: &Database{
				Ranges: []FreshRange{
					{Start: 5, End: 10},
					{Start: 8, End: 15},
					{Start: 12, End: 20},
				},
				Available: []int{4, 7, 9, 13, 21},
			},
			wantCount: 3, // IDs 7, 9, and 13 are fresh
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountFreshIngredients(tc.db)
			if got != tc.wantCount {
				t.Errorf("CountFreshIngredients() = %d, want %d", got, tc.wantCount)
			}
		})
	}
}

func TestCountAllPossibleFreshIngredients(t *testing.T) {
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

	// Count the number of all the IDs contained in all ranges
	if db.CountTotalPossibleFresh() != 14 {
		t.Errorf("CountTotalPossibleFresh() = %d, want 14", db.CountTotalPossibleFresh())
	}
}
