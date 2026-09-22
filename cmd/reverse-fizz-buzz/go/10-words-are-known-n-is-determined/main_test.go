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

func TestSolutionV2(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []int
	}{
		"input-1": {
			input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "Fizz", "Buzz"},
			want:  []int{10, 10},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortedInts(SolutionV2(tc.input))
			want := sortedInts(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
