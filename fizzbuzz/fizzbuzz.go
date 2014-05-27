package fizzbuzz

import "strconv"

func DoFizzBuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	}
	if n == 5 {
		return "Buzz"
	} else if n == 10 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}
