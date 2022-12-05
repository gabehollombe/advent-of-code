package main

import "testing"

func Test_scoreRound(t *testing.T) {
	type args struct {
		r [2]shape
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"tie with rock", args{r: [2]shape{Rock, Rock}}, 4},
		{"win with rock", args{r: [2]shape{Scissors, Rock}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreRound(tt.args.r); got != tt.want {
				t.Errorf("scoreRound() = %v, want %v", got, tt.want)
			}
		})
	}
}
