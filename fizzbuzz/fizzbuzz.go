package fizzbuzz

import "strconv"

func Say(n int) string {

	r := ""

	if n%3 == 0 {
		r += "Fizz"
	}
	if n%5 == 0 {
		r += "Buzz"
	}

	if r == "" {
		r = strconv.Itoa(n)
	}

	return r
}
