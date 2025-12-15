package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	expected := []BatteryBank{
		{Batteries: []int32{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}},
		{Batteries: []int32{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}},
		{Batteries: []int32{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}},
		{Batteries: []int32{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1}},
	}

	banks, err := ParseInput("input_test.txt")
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	if !reflect.DeepEqual(banks, expected) {
		t.Errorf("ParseInput() = %v, want %v", banks, expected)
	}
}
