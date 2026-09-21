package main

import (
	"testing"
)

func TestLonelyinteger(t *testing.T) {
	tests := map[string]struct {
		input []int32
		want  int32
	}{
		"single element": {
			input: []int32{1},
			want:  1,
		},
		"one pair and one unique": {
			input: []int32{1, 1, 2},
			want:  2,
		},
		"unique at end": {
			input: []int32{0, 0, 1, 2, 1},
			want:  2,
		},
		"three pairs one unique": {
			input: []int32{1, 2, 3, 4, 3, 2, 1},
			want:  4,
		},
		"unique zero": {
			input: []int32{0, 1, 1},
			want:  0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := lonelyinteger(tt.input)
			if got != tt.want {
				t.Errorf("lonelyinteger(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestLonelyintegerV2(t *testing.T) {
	tests := map[string]struct {
		input []int32
		want  int32
	}{
		"single element": {
			input: []int32{1},
			want:  1,
		},
		"one pair and one unique": {
			input: []int32{1, 1, 2},
			want:  2,
		},
		"unique at end": {
			input: []int32{0, 0, 1, 2, 1},
			want:  2,
		},
		"three pairs one unique": {
			input: []int32{1, 2, 3, 4, 3, 2, 1},
			want:  4,
		},
		"unique zero": {
			input: []int32{0, 1, 1},
			want:  0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := lonelyintegerV2(tt.input)
			if got != tt.want {
				t.Errorf("lonelyintegerV2(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
