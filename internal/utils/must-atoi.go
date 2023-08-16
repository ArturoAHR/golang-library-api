package utils

import "strconv"

func MustAtoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return number
}
