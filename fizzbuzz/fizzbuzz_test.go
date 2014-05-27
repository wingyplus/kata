package fizzbuzz_test

import (
	"github.com/wingyplus/kata/fizzbuzz"
	"testing"
)

func TestDoFizzBuzz(t *testing.T) {
	var testCases = []struct {
		actual   int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{6, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
	}

	for _, example := range testCases {
		var result string = fizzbuzz.DoFizzBuzz(example.actual)

		if result != example.expected {
			t.Errorf("expect %s but was %s", example.expected, result)
		}
	}
}
