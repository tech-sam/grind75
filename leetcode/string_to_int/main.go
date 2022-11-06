package main

import (
	"fmt"
	"math"
)

func main() {
	//s := "words and 987"
	//s := "   0032"
	//s := "+-12"
	//s := "-91283472332"
	//output = "-2147483648"

	s := "21474836460"

	// fmt.Println(math.MaxInt32)
	// fmt.Println(math.MinInt32)
	fmt.Println(myAtoi(s))

	//fmt.Println(21474836460 > math.MaxInt32)
}

func myAtoi(s string) int {

	start := false
	sign := 1
	res := 0

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
			digit := int(v-'0') * sign

			if res > math.MaxInt32 {
				return math.MaxInt32
			}
			if res < math.MinInt32 {
				return math.MinInt32
			}
			res = res*10 + digit
			fmt.Printf("---res---%d\n", res)

		} else {
			break
		}
	}
	return res
}
