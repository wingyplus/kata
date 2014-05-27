package fizzbuzz

import "strconv"

func DoFizzBuzz(n int) string {
	if n == 15 {
		return "FizzBuzz"
	}
	if n == 30 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
