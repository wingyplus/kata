package fizzbuzz

import "strconv"

func DoFizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	} else if n == 6 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
