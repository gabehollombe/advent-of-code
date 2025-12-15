package main

import (
	"reflect"
	"testing"
)

func TestParseRanges(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		wantRanges []Range
		wantErr    bool
	}{
		{
			name:  "Single Range",
			input: "11-22",
			wantRanges: []Range{
				{Start: 11, End: 22},
			},
			wantErr: false,
		},
		{
			name:  "Multiple Ranges",
			input: "11-22,95-115,998-1012",
			wantRanges: []Range{
				{Start: 11, End: 22},
				{Start: 95, End: 115},
				{Start: 998, End: 1012},
			},
			wantErr: false,
		},
		{
			name:  "Large Numbers",
			input: "1188511880-1188511890",
			wantRanges: []Range{
				{Start: 1188511880, End: 1188511890},
			},
			wantErr: false,
		},
		{
			name:       "Empty Input",
			input:      "",
			wantRanges: []Range{},
			wantErr:    false,
		},
		{
			name:    "Invalid Format - No Dash",
			input:   "1122",
			wantErr: true,
		},
		{
			name:    "Invalid Format - Non-numeric",
			input:   "abc-def",
			wantErr: true,
		},
		{
			name:    "Invalid Format - Missing Start",
			input:   "-22",
			wantErr: true,
		},
		{
			name:    "Invalid Format - Missing End",
			input:   "11-",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRanges, err := ParseRanges(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseRanges() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr && !reflect.DeepEqual(gotRanges, tc.wantRanges) {
				t.Errorf("ParseRanges() = %v, want %v", gotRanges, tc.wantRanges)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	// This test assumes input.txt exists and contains valid data
	// We're just testing that it reads the file and calls ParseRanges
	ranges, err := ParseInput("input.txt")
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	if len(ranges) == 0 {
		t.Error("ParseInput() returned empty ranges")
	}
}
