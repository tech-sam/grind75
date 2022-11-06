package main

import "fmt"

func main() {
	s := "pwwkew"
	//s := "abcdefgklizk"
	//s := "abcabcbb"
	//s := " "
	//s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	a, b, max := 0, 0, 0

	m := make(map[byte]bool)
	for b < len(s) {
		if !m[s[b]] {
			m[s[b]] = true
			b++
			if len(m) > max {
				max = len(m)
			}
		} else {
			delete(m, s[a])
			a++
		}
	}

	return max

}
