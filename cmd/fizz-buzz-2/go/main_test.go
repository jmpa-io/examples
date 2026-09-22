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

func TestFizzBuzz(t *testing.T) {
	tests := map[string]struct {
		n    int32
		want string
	}{
		"n=1":  {1, "1\n"},
		"n=3":  {3, "1\n2\nFizz\n"},
		"n=5":  {5, "1\n2\nFizz\n4\nBuzz\n"},
		"n=15": {15, "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := captureOutput(func() { fizzBuzz(tc.n) })
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
