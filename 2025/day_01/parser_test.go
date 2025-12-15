package main

import (
	"testing"
)

func TestParseRotation(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		wantRotation  Rotation
		wantErr       bool
	}{
		{
			name:  "Valid Left Rotation",
			input: "L68",
			wantRotation: Rotation{
				Direction: Left,
				Distance:  68,
			},
			wantErr: false,
		},
		{
			name:  "Valid Right Rotation",
			input: "R48",
			wantRotation: Rotation{
				Direction: Right,
				Distance:  48,
			},
			wantErr: false,
		},
		{
			name:    "Invalid Direction",
			input:   "X10",
			wantErr: true,
		},
		{
			name:    "Invalid Short Input",
			input:   "L",
			wantErr: true,
		},
		{
			name:    "Invalid Number",
			input:   "Rabc",
			wantErr: true,
		},
		{
			name:  "Edge Case Zero Distance",
			input: "L0",
			wantRotation: Rotation{
				Direction: Left,
				Distance:  0,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotRotation, err := ParseRotation(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseRotation() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !tc.wantErr && gotRotation != tc.wantRotation {
				t.Errorf("ParseRotation() = %v, want %v", gotRotation, tc.wantRotation)
			}
		})
	}
}