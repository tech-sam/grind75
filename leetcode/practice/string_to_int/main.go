package main

import (
	"fmt"
	"math"
)

func main() {

	s := "42"
	fmt.Println(myAtoi(s))

}

func myAtoi(s string) int {

	start := false
	res := 0
	sign := 1

	for _, v := range s {

		if v == ' ' && !start {
			continue
		}
		if v == '-' && !start {
			sign = -1
			start = true
			continue
		}
		if v == '+' && !start {
			start = true
			continue
		}

		if v >= '0' && v <= '9' {
			start = true
			digit := sign * int(v-'0')
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
				return math.MaxInt32
			}

			if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < -8) {
				return math.MinInt32
			}

			res = res*10 + digit

		} else {
			break
		}

	}

	return res
}
