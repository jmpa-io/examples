package main

import (
	"testing"
)

func TestGridChallenge(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"5x5 sortable YES": {
			input: []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"},
			want:  "YES",
		},
		"already sorted rows and columns YES": {
			input: []string{"abc", "hjk", "mpq", "rtv", "xyw"},
			want:  "YES",
		},
		"3x4 sorted YES": {
			input: []string{"abcd", "efgh", "ijkl"},
			want:  "YES",
		},
		"column violation after sort NO": {
			input: []string{"zbc", "abc"},
			want:  "NO",
		},
		"single row always YES": {
			input: []string{"dcba"},
			want:  "YES",
		},
		"single character per row YES": {
			input: []string{"a", "b", "c"},
			want:  "YES",
		},
		"single character per row descending NO": {
			input: []string{"c", "b", "a"},
			want:  "NO",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// pass a copy so the function's in-place row sort doesn't affect other cases
			input := make([]string, len(tt.input))
			copy(input, tt.input)
			got := gridChallenge(input)
			if got != tt.want {
				t.Errorf("gridChallenge(%v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
