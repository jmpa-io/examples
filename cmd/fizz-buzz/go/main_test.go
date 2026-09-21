package main

import (
	"context"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := map[string]struct {
		input fizzBuzzInput
		want  []string
	}{
		"n=5 fizz and buzz": {
			input: fizzBuzzInput{
				n: 5,
				divisors: []divisor{
					{3, "Fizz"},
					{5, "Buzz"},
				},
			},
			want: []string{"1", "2", "3 - Fizz", "4", "5 - Buzz"},
		},
		"n=15 includes fizzbuzz at index 14": {
			input: fizzBuzzInput{
				n: 15,
				divisors: []divisor{
					{3, "Fizz"},
					{5, "Buzz"},
				},
			},
			want: []string{
				"1", "2", "3 - Fizz", "4", "5 - Buzz",
				"6 - Fizz", "7", "8", "9 - Fizz", "10 - Buzz",
				"11", "12 - Fizz", "13", "14", "15 - FizzBuzz",
			},
		},
		"n=1 no divisor matches": {
			input: fizzBuzzInput{
				n:        1,
				divisors: []divisor{{2, "hello"}},
			},
			want: []string{"1"},
		},
		"n=2 second number matches": {
			input: fizzBuzzInput{
				n:        2,
				divisors: []divisor{{2, "hello"}},
			},
			want: []string{"1", "2 - hello"},
		},
		"n=1 empty divisors": {
			input: fizzBuzzInput{
				n:        1,
				divisors: []divisor{},
			},
			want: []string{"1"},
		},
		"multiple divisors combine words": {
			input: fizzBuzzInput{
				n: 6,
				divisors: []divisor{
					{2, "hello"},
					{3, "world"},
				},
			},
			want: []string{"1", "2 - hello", "3 - world", "4 - hello", "5", "6 - helloworld"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := fizzBuzz(context.Background(), tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz(%+v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
