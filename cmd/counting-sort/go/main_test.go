package main

import (
	"reflect"
	"testing"
)

func makeFreq(pairs map[int]int32) []int32 {
	freq := make([]int32, 100)
	for i, v := range pairs {
		freq[i] = v
	}
	return freq
}

func TestCountingSort(t *testing.T) {
	tests := map[string]struct {
		input []int32
		want  []int32
	}{
		"multiple values": {
			input: []int32{1, 1, 3, 2, 1},
			want:  makeFreq(map[int]int32{1: 3, 2: 1, 3: 1}),
		},
		"single zero": {
			input: []int32{0},
			want:  makeFreq(map[int]int32{0: 1}),
		},
		"single max value": {
			input: []int32{99},
			want:  makeFreq(map[int]int32{99: 1}),
		},
		"all same zero": {
			input: []int32{0, 0, 0},
			want:  makeFreq(map[int]int32{0: 3}),
		},
		"two distinct values": {
			input: []int32{5, 5, 7},
			want:  makeFreq(map[int]int32{5: 2, 7: 1}),
		},
		"all distinct": {
			input: []int32{10, 20, 30},
			want:  makeFreq(map[int]int32{10: 1, 20: 1, 30: 1}),
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := countingSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countingSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
