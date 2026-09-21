package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	dirs, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf("failed to read testdata: %v", err)
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		t.Run(d.Name(), func(t *testing.T) {
			// read input paths.
			inputBytes, err := os.ReadFile(filepath.Join("testdata", d.Name(), "input"))
			if err != nil {
				t.Fatalf("failed to read input: %v", err)
			}
			var paths []string
			for _, line := range strings.Split(strings.TrimSpace(string(inputBytes)), "\n") {
				if line != "" {
					paths = append(paths, strings.TrimRight(line, "\r"))
				}
			}

			// read expected answer.
			answerBytes, err := os.ReadFile(filepath.Join("testdata", d.Name(), "answer"))
			if err != nil {
				t.Fatalf("failed to read answer: %v", err)
			}
			lines := strings.Split(strings.TrimSpace(string(answerBytes)), "\n")
			if len(lines) != 2 {
				t.Fatalf("answer must have 2 lines, got %d", len(lines))
			}
			wantSubfolders, err := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
			if err != nil {
				t.Fatalf("invalid answer line 1: %v", err)
			}
			wantFiles, err := strconv.ParseInt(strings.TrimSpace(lines[1]), 10, 32)
			if err != nil {
				t.Fatalf("invalid answer line 2: %v", err)
			}

			got := solve(paths)
			want := []int32{int32(wantSubfolders), int32(wantFiles)}
			if got[0] != want[0] || got[1] != want[1] {
				t.Errorf("got %v, want %v\ninput:\n%s", got, want, fmt.Sprintf("  %s", strings.Join(paths, "\n  ")))
			}
		})
	}
}
