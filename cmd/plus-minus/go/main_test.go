package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(fn func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestDetermineRatios(t *testing.T) {
	tests := map[string]struct {
		numbers []int32
		want    string
	}{
		"all-positive":  {[]int32{1, 2, 3, 4, 5}, "1.000000\n0.000000\n0.000000\n"},
		"all-negative":  {[]int32{-1, -2, -3}, "0.000000\n1.000000\n0.000000\n"},
		"all-neutral":   {[]int32{0, 0, 0}, "0.000000\n0.000000\n1.000000\n"},
		"mixed":         {[]int32{-4, 3, -9, 0, 3, 7, 7, -3, -2, -9}, "0.400000\n0.500000\n0.100000\n"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := captureOutput(func() { determineRatios(tc.numbers) })
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
