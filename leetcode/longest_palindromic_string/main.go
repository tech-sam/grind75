package main

import "fmt"

func main() {
	//s := "babad"
	//s := "ab"
	s := "aa"

	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {

	res := ""
	res_len := 0

	for j := 0; j < len(s); j++ {
		// odd length palindrome
		l, r := j, j

		for l >= 0 && r < len(s) && s[l] == s[r] {

			if (r - l + 1) > res_len {
				res = s[l : r+1]
				res_len = r - l + 1
			}
			l--
			r++
		}
		// even length palindrome
		l, r = j, j+1
		for l >= 0 && r < len(s) && s[l] == s[r] {

			if (r - l + 1) > res_len {
				res = s[l : r+1]
				res_len = r - l + 1
			}
			l--
			r++
		}
	}

	return res

}

// func longestPalindrome(s string) string {

// 	if len(s) == 1 {
// 		return s
// 	}
// 	max_sub_string := ""
// 	for i := 0; i < len(s); i++ {

// 		for j := i + 1; j < len(s); j++ {
// 			sub_string := s[i : j+1]
// 			if isPalindrome(sub_string) {
// 				if len(max_sub_string) < len(sub_string) {
// 					max_sub_string = sub_string
// 				}
// 			}

// 		}
// 	}
// 	if len(max_sub_string) == 0 {
// 		return string(s[0])
// 	}
// 	return max_sub_string
// }

// func isPalindrome(s string) bool {
// 	i, j := 0, len(s)-1
// 	for i < j {
// 		if s[i] != s[j] {
// 			return false
// 		}

// 		i++
// 		j--
// 	}
// 	return true
// }
