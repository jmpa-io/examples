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

func TestMiniMaxSum(t *testing.T) {
	tests := map[string]struct {
		arr  []int32
		want string
	}{
		"example-1": {[]int32{1, 2, 3, 4, 5}, "10 14\n"},
		"example-2": {[]int32{1, 3, 5, 7, 9}, "16 24\n"},
		"all-same":  {[]int32{5, 5, 5, 5, 5}, "20 20\n"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := captureOutput(func() { miniMaxSum(tc.arr) })
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
