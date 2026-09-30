package main

import (
	"strings"
)



// findWords is used to filter out numbers from FizzBuzz input & determine the
// sequence of words that appear in-order.
func findWordSequence(A []string) (out []string) {
	for i := 0; i < len(A); i++ {
		line := A[i]
		// ignore numbers.
		if !wordRgx.MatchString(line) {
			continue
		}
		out = append(out, line)
	}
	return out
}

// contains check if a given string value is found in the given array.
func contains(slice []string, value string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}
	return false
}

func allScoresEqualOne(m map[string]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}

func allScoresGreaterThanOne(m map[string]int) bool {
	for _, v := range m {
		if v <= 1 {
			return false
		}
	}
	return true
}

func SolutionV1(A []string) (out []int) {

	// determine sequence of words.
	sequence := findWordSequence(A)

	// find words in sequence, in-order.
	var words []string
	for i := 0; i < len(sequence); i++ {
		line := sequence[i]
		split := splitPascalCase(line)
		for j := 0; j < len(split); j++ {

			// check if the word is already in the word list.
			if contains(words, split[j]) {
				continue
			}
			words = append(words, split[j])
		}
	}

	// determine when words intersect.
	intersectWord := strings.Join(words, "")

	// setup scores.
	scores := make(map[string]int)
	for _, w := range words {
		scores[w] = 1
	}

	// score words.
	for i := 1; i < len(A); i++ {
		line := A[i-1]

		if line == intersectWord && allScoresEqualOne(scores) {
			for word := range scores {
				scores[word] = i
			}
		} else {

			for j := 0; j < len(words); j++ {
				word := words[j]
				if (line == word || line == intersectWord) && scores[word] == 0 {
					scores[word] = i
				}
			}
		}

		if allScoresGreaterThanOne(scores) {
			break
		}
	}

	// build output.
	for i := 0; i < len(words); i++ {
		out = append(out, scores[words[i]])
	}
	return out
}
