package main

import "fmt"

func main() {
	s := "AAa"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {

	count := make([]int, 128)
	for _, v := range s {
		count[v]++
	}
	ans := 0
	odd_exists := false
	for _, v := range count {

		if v%2 == 0 {
			ans += v
		} else {
			odd_exists = true
			ans += v - 1
		}
	}
	if odd_exists {
		ans++
	}

	return ans
}
