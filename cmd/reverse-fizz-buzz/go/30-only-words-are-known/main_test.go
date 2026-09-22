package main

import (
	"reflect"
	"testing"
)

func TestGenerateFizzBuzz(t *testing.T) {
	tests := map[string]struct {
		divisors []int
		words    []string
		n        int
		want     []string
	}{
		"fizz-buzz-3-5-n15": {
			divisors: []int{3, 5},
			words:    []string{"Fizz", "Buzz"},
			n:        15,
			want:     []string{"Fizz", "Buzz", "Fizz", "Fizz", "Buzz", "Fizz", "FizzBuzz"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := generateFizzBuzz(tc.divisors, tc.words, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
