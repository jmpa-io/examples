package main

import (
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	tests := map[string]struct {
		input [][]int32
		want  int32
	}{
		"3x3 with negatives": {
			input: [][]int32{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			want: 15,
		},
		"2x2 equal diagonals": {
			input: [][]int32{
				{1, 0},
				{0, 1},
			},
			want: 2,
		},
		"3x3 sequential": {
			input: [][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 0,
		},
		"1x1 single element": {
			input: [][]int32{
				{5},
			},
			want: 0,
		},
		"3x3 all zeros": {
			input: [][]int32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		"2x2 asymmetric": {
			input: [][]int32{
				{3, 7},
				{2, 4},
			},
			want: 2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := diagonalDifference(tt.input)
			if got != tt.want {
				t.Errorf("diagonalDifference(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
