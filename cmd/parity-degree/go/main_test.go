package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"24 divides by 2 three times": {
			input: 24,
			want:  3,
		},
		"1 is odd no division": {
			input: 1,
			want:  0,
		},
		"8 is 2^3": {
			input: 8,
			want:  3,
		},
		"16 is 2^4": {
			input: 16,
			want:  4,
		},
		"6 divides by 2 once": {
			input: 6,
			want:  1,
		},
		"2 is 2^1": {
			input: 2,
			want:  1,
		},
		"odd number 15": {
			input: 15,
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
