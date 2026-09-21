package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"first unique is not first element": {
			input: []int{1, 3, 6, 4, 1, 2},
			want:  3,
		},
		"all unique returns first": {
			input: []int{1, 2, 3},
			want:  1,
		},
		"no unique returns -1": {
			input: []int{1, 1, 2, 2, 3, 3},
			want:  -1,
		},
		"unique buried in repeats": {
			input: []int{9, 3, 9, 3, 9, 7, 9},
			want:  7,
		},
		"single element": {
			input: []int{42},
			want:  42,
		},
		"all same returns -1": {
			input: []int{5, 5, 5, 5},
			want:  -1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Solution(tt.input)
			if got != tt.want {
				t.Errorf("Solution(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestSolutionV2(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"missing 5 in sequence": {
			input: []int{1, 3, 6, 4, 1, 2},
			want:  5,
		},
		"consecutive from 1 missing next": {
			input: []int{1, 2, 3},
			want:  4,
		},
		"all negative returns 1": {
			input: []int{-1, -3},
			want:  1,
		},
		"missing 2 in sequence": {
			input: []int{1, 3},
			want:  2,
		},
		"large gap returns 1": {
			input: []int{100, 200, 300},
			want:  1,
		},
		"single positive 1 missing 2": {
			input: []int{1},
			want:  2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := SolutionV2(tt.input)
			if got != tt.want {
				t.Errorf("SolutionV2(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
