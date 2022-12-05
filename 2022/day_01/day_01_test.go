package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want int
	}{
		{"empty list", []int{}, 0},
		{"non-empty list", []int{1, 2, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.ints); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumGroups(t *testing.T) {
	type args struct {
		lines [][]int
	}
	tests := []struct {
		name   string
		groups [][]int
		want   []int
	}{
		{"empty list", [][]int{}, []int{}},
		{"non-empty list", [][]int{[]int{1, 2}, []int{3, 4}}, []int{3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumGroups(tt.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
