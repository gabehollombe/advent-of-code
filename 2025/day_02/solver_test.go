package main

import (
	"reflect"
	"testing"
)

func TestIsInvalidId(t *testing.T) {
	testCases := []struct {
		name string
		id   int
		want bool
	}{
		{name: "Example 11", id: 11, want: true},
		{name: "Example 55", id: 55, want: true},
		{name: "Example 6464", id: 6464, want: true},
		{name: "Example 123123", id: 123123, want: true},
		{name: "Example 1188511885", id: 1188511885, want: true},
		{name: "Example 222222", id: 222222, want: true},
		{name: "Example 446446", id: 446446, want: true},
		{name: "Example 38593859", id: 38593859, want: true},
		{name: "Example 1010", id: 1010, want: true},
		{name: "Example 99", id: 99, want: true},

		{name: "Valid 101", id: 101, want: false},
		{name: "Valid 1234", id: 1234, want: false},
		{name: "Valid 12123", id: 12123, want: false},
		{name: "Valid Single Digit", id: 5, want: false},
		{name: "Valid Odd Length", id: 12345, want: false},
		{name: "Valid Different Halves", id: 123456, want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsInvalidId(tc.id)
			if got != tc.want {
				t.Errorf("IsInvalidId(%d) = %v, want %v", tc.id, got, tc.want)
			}
		})
	}
}

func TestGetInvalidIds(t *testing.T) {
	testCases := []struct {
		name           string
		rng            Range
		wantInvalidIds []int
	}{
		{
			name:           "Range 11-22",
			rng:            Range{Start: 11, End: 22},
			wantInvalidIds: []int{11, 22},
		},
		{
			name:           "Range 95-115",
			rng:            Range{Start: 95, End: 115},
			wantInvalidIds: []int{99},
		},
		{
			name:           "Range 998-1012",
			rng:            Range{Start: 998, End: 1012},
			wantInvalidIds: []int{1010},
		},
		{
			name:           "Range 1188511880-1188511890",
			rng:            Range{Start: 1188511880, End: 1188511890},
			wantInvalidIds: []int{1188511885},
		},
		{
			name:           "Range 1698522-1698528",
			rng:            Range{Start: 1698522, End: 1698528},
			wantInvalidIds: []int{},
		},
		{
			name:           "Range 446443-446449",
			rng:            Range{Start: 446443, End: 446449},
			wantInvalidIds: []int{446446},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetInvalidIds(tc.rng)
			// Handle nil vs empty slice comparison
			if len(got) == 0 && len(tc.wantInvalidIds) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.wantInvalidIds) {
				t.Errorf("GetInvalidIds(%v) = %v, want %v", tc.rng, got, tc.wantInvalidIds)
			}
		})
	}
}

func TestSumInvalidIds(t *testing.T) {
	// Integration test using the full example from the problem
	ranges := []Range{
		{Start: 11, End: 22},
		{Start: 95, End: 115},
		{Start: 998, End: 1012},
		{Start: 1188511880, End: 1188511890},
		{Start: 222220, End: 222224},
		{Start: 1698522, End: 1698528},
		{Start: 446443, End: 446449},
		{Start: 38593856, End: 38593862},
		{Start: 565653, End: 565659},
		{Start: 824824821, End: 824824827},
		{Start: 2121212118, End: 2121212124},
	}

	expectedSum := 1227775554

	actualSum := SumInvalidIds(ranges)

	if actualSum != expectedSum {
		t.Errorf("SumInvalidIds() = %d, want %d", actualSum, expectedSum)
	}
}

func TestGetInvalidIdsPartTwo(t *testing.T) {
	testCases := []struct {
		name           string
		rng            Range
		wantInvalidIds []int
	}{
		{
			name:           "Range 11-22 (still has 11 and 22)",
			rng:            Range{Start: 11, End: 22},
			wantInvalidIds: []int{11, 22},
		},
		{
			name:           "Range 95-115 (now has 99 and 111)",
			rng:            Range{Start: 95, End: 115},
			wantInvalidIds: []int{99, 111},
		},
		{
			name:           "Range 998-1012 (now has 999 and 1010)",
			rng:            Range{Start: 998, End: 1012},
			wantInvalidIds: []int{999, 1010},
		},
		{
			name:           "Range 1188511880-1188511890 (still has 1188511885)",
			rng:            Range{Start: 1188511880, End: 1188511890},
			wantInvalidIds: []int{1188511885},
		},
		{
			name:           "Range 222220-222224 (still has 222222)",
			rng:            Range{Start: 222220, End: 222224},
			wantInvalidIds: []int{222222},
		},
		{
			name:           "Range 1698522-1698528 (still contains no invalid IDs)",
			rng:            Range{Start: 1698522, End: 1698528},
			wantInvalidIds: []int{},
		},
		{
			name:           "Range 446443-446449 (still has 446446)",
			rng:            Range{Start: 446443, End: 446449},
			wantInvalidIds: []int{446446},
		},
		{
			name:           "Range 38593856-38593862 (still has 38593859)",
			rng:            Range{Start: 38593856, End: 38593862},
			wantInvalidIds: []int{38593859},
		},
		{
			name:           "Range 565653-565659 (now has 565656)",
			rng:            Range{Start: 565653, End: 565659},
			wantInvalidIds: []int{565656},
		},
		{
			name:           "Range 824824821-824824827 (now has 824824824)",
			rng:            Range{Start: 824824821, End: 824824827},
			wantInvalidIds: []int{824824824},
		},
		{
			name:           "Range 2121212118-2121212124 (now has 2121212121)",
			rng:            Range{Start: 2121212118, End: 2121212124},
			wantInvalidIds: []int{2121212121},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetInvalidIdsPartTwo(tc.rng)
			// Handle nil vs empty slice comparison
			if len(got) == 0 && len(tc.wantInvalidIds) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.wantInvalidIds) {
				t.Errorf("GetInvalidIdsPartTwo(%v) = %v, want %v", tc.rng, got, tc.wantInvalidIds)
			}
		})
	}
}
