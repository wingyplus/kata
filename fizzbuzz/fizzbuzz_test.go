package fizzbuzz_test

import (
	"github.com/wingyplus/kata/fizzbuzz"
	"testing"
)

func TestDoFizzBuzz(t *testing.T) {
	var result string = fizzbuzz.DoFizzBuzz(1)

	if result != "1" {
		t.Errorf("expect %s but was %s", "1", result)
	}
}
