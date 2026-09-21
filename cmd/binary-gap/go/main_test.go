package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"9 (1001) gap of 2": {
			input: 9,
			want:  2,
		},
		"529 (1000010001) gap of 4": {
			input: 529,
			want:  4,
		},
		"20 (10100) gap of 1": {
			input: 20,
			want:  1,
		},
		"15 (1111) no gap": {
			input: 15,
			want:  0,
		},
		"32 (100000) trailing zeros no gap": {
			input: 32,
			want:  0,
		},
		"1041 (10000010001) gap of 5": {
			input: 1041,
			want:  5,
		},
		"1 (1) single bit no gap": {
			input: 1,
			want:  0,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Solution(tt.input)
			if got != tt.want {
				t.Errorf("Solution(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
