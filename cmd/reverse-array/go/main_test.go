package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	tests := map[string]struct {
		input []int32
		want  []int32
	}{
		"five elements": {
			input: []int32{1, 2, 3, 4, 5},
			want:  []int32{5, 4, 3, 2, 1},
		},
		"single element": {
			input: []int32{1},
			want:  []int32{1},
		},
		"two elements": {
			input: []int32{1, 2},
			want:  []int32{2, 1},
		},
		"duplicates preserved": {
			input: []int32{3, 1, 4, 1, 5},
			want:  []int32{5, 1, 4, 1, 3},
		},
		"already reversed": {
			input: []int32{5, 4, 3, 2, 1},
			want:  []int32{1, 2, 3, 4, 5},
		},
		"all same values": {
			input: []int32{7, 7, 7},
			want:  []int32{7, 7, 7},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverseArray(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseArray(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
