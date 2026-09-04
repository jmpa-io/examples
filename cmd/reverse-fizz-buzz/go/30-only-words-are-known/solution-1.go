package main

import (
	"unicode"
)

// splitPascalCase splits a string into PascalCase words.
func splitPascalCase(s string) (out []string) {

	// Loop over each character and determine the starting position of any
	// uppercase characters to split on.
	var position int
	for i, character := range s {

		// Check if the current character is uppercase, and append the substring
		// before the current uppercase letter as a new word.
		if unicode.IsUpper(character) && i > 0 {
			out = append(out, s[position:i])
			position = i
		}
	}

	// Append any characters after the last uppercase letter in the string.
	out = append(out, s[position:])
	return out
}

// SolutionV1 is a placeholder — implementation not yet complete.
func SolutionV1(_ []string) []string {
	return nil
}

// generateFizzBuzz is a helper to generate a fizz buzz slice!
func generateFizzBuzz(divisors []int, words []string, n int) (out []string) {

	for i := 1; i <= n; i++ {

		var line string

		// determine if the current index is divisible by any known divisor.
		// if divisible, append the corresponding word to the line.
		for j, d := range divisors {
			if i%d == 0 {
				line += words[j]
			}
		}

		// skip numbers from being added to the output.
		if line == "" {
			continue
		}

		out = append(out, line)
	}

	return out
}

