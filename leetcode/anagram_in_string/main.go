package main

import (
	"fmt"
)

func main() {

	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

// func findAnagrams(s string, p string) []int {
// 	m := [26]int{}
// 	for _, r := range p {
// 		v := r - 'a'
// 		m[v]++
// 	}
// 	res := []int{}
// 	for i := 0; i < len(s); i++ {
// 		l := i + len(p)
// 		if l <= len(s) {
// 			if isAnagram(s[i:l], m) {
// 				res = append(res, i)
// 			}
// 		}

// 	}

// 	return res
// }

// func isAnagram(s string, m [26]int) bool {
// 	n := [26]int{}
// 	for _, r := range s {
// 		v := r - 'a'
// 		n[v]++
// 	}
// 	return n == m
// }

// func findAnagrams(s string, p string) []int {
// 	if len(p) > len(s) {
// 		return []int{}
// 	}
// 	var res []int
// 	pMap := make(map[byte]int)
// 	for i := 0; i < len(p); i++ {
// 		pMap[p[i]]++
// 	}
// 	var left, right int
// 	var pCount int
// 	for right < len(s) {
// 		if pMap[s[right]] > 0 {
// 			pCount++
// 		}
// 		pMap[s[right]]--
// 		right++
// 		for pCount == len(p) {
// 			if right-left == len(p) {
// 				res = append(res, left)
// 			}
// 			if pMap[s[left]] == 0 {
// 				pCount--
// 			}
// 			pMap[s[left]]++
// 			left++
// 		}
// 	}
// 	return res
// }

// func findAnagrams(s string, p string) []int {
// 	if len(p) > len(s) {
// 		return []int{}
// 	}
// 	var res []int

// 	pCount, sCount := make(map[byte]int), make(map[byte]int)

// 	for i := 0; i < len(p); i++ {
// 		pCount[p[i]]++
// 		sCount[s[i]]++
// 	}

// 	if reflect.DeepEqual(sCount, pCount) {
// 		res = append(res, 0)
// 	}

// 	l := 0
// 	for r := len(p); r < len(s); r++ {
// 		sCount[s[r]]++
// 		sCount[s[l]]--
// 		if sCount[s[l]] == 0 {
// 			delete(sCount, s[l])
// 		}
// 		l++
// 		if reflect.DeepEqual(sCount, pCount) {
// 			res = append(res, l)
// 		}
// 	}

// 	return res
// }

func findAnagrams(s string, p string) []int {

	res := []int{}

	hash, pHash := [26]int{}, [26]int{}
	window, len := len(p), len(s)
	if len < window {
		return res
	}
	left, right := 0, 0

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
