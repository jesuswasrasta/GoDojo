package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{98, "98"},
		{99, "Fizz"},
		{100, "Buzz"},
	}

	for _, test := range tests {
		result := FizzBuzz(test.input)
		if result != test.expected {
			t.Errorf("FizzBuzz(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
