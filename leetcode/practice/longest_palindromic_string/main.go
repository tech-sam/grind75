package main

import "fmt"

func main() {
	fmt.Println("--testing---")
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {

	result := ""

	for i := 0; i < len(s); i++ {
		result = palindromeSubString(i, i, s)
		fmt.Println(result)

		result = palindromeSubString(i, i+1, s)

	}

	return result

}

func palindromeSubString(left, right int, s string) string {

	tmp_str := ""
	tmp_len := 0

	for left >= 0 && right < len(s) && s[left] == s[right] {
		if (right - left + 1) > tmp_len {
			tmp_str = s[left : right+1]
			tmp_len = right - left + 1
		}
		left--
		right++
	}

	return tmp_str
}
