package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for _, r := range s {
		v := r - 'a'
		m[v]++
	}

	n := [26]int{}
	for _, r := range t {
		v := r - 'a'
		n[v]++
	}
	return m == n
}

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	s_sort := []rune(s)
// 	sort.Slice(s_sort, func(i, j int) bool {
// 		return s_sort[i] < s_sort[j]
// 	})

// 	t_sort := []rune(t)
// 	sort.Slice(t_sort, func(i, j int) bool {
// 		return t_sort[i] < t_sort[j]
// 	})

// 	for i := 0; i < len(s_sort); i++ {
// 		if s_sort[i] != t_sort[i] {
// 			return false
// 		}
// 	}

// 	return true
// }
