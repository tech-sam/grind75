package main

import "fmt"

func main() {
	s := "abab"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {

	res := []int{}

	window, len := len(p), len(s)

	if window > len {
		return res
	}

	pHash, hash := [26]int{}, [26]int{}

	right, left := 0, 0

	for right < window {
		pHash[p[right]-'a']++
		hash[s[right]-'a']++
		right++
	}
	right--

	for right < len {
		if pHash == hash {
			res = append(res, left)
		}
		right++
		if right != len {
			hash[s[right]-'a']++
		}
		hash[s[left]-'a']--
		left++
	}

	return res

}
