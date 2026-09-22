package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortedInts(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)
	sort.Ints(out)
	return out
}

func TestSolutionV1(t *testing.T) {
	tests := map[string]struct {
		n     int
		input []string
		want  []int
	}{
		"rock-star-2-20": {
			n:     20,
			input: []string{"rock", "rock", "rock", "star", "rock", "rock", "rock", "rockstar", "rock", "rock", "rock"},
			want:  []int{2, 20, 20},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortedInts(SolutionV1(tc.n, tc.input))
			want := sortedInts(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
