package main

import (
	"testing"
)

func TestTimeConversion(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"PM time": {
			input: "07:05:45PM",
			want:  "19:05:45",
		},
		"midnight AM": {
			input: "12:00:00AM",
			want:  "00:00:00",
		},
		"noon PM": {
			input: "12:00:00PM",
			want:  "12:00:00",
		},
		"early morning AM": {
			input: "01:30:00AM",
			want:  "01:30:00",
		},
		"last second PM": {
			input: "11:59:59PM",
			want:  "23:59:59",
		},
		"midnight with seconds AM": {
			input: "12:45:54AM",
			want:  "00:45:54",
		},
		"noon with seconds PM": {
			input: "12:01:01PM",
			want:  "12:01:01",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := timeConversion(tt.input)
			if got != tt.want {
				t.Errorf("timeConversion(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
