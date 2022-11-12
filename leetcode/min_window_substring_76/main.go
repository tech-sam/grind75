package main

import "fmt"

func main() {
	s := "cabwefgewcwaefgcf"
	t := "cae"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {

	if len(s) < len(t) {
		return ""
	}
	mt := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}

	i, j, count := 0, 0, len(mt)
	left, right, min := 0, len(s)-1, len(s)
	found := false

	for j < len(s) {
		endChar := s[j]
		j++
		if _, ok := mt[endChar]; ok {
			mt[endChar]--
			if mt[endChar] == 0 {
				count--
			}
		}
		if count > 0 {
			continue
		}
		for count == 0 {
			startChar := s[i]
			i++
			if _, ok := mt[startChar]; ok {
				mt[startChar]++
				if mt[startChar] > 0 {
					count++
				}
			}
		}

		if (j - i) < min {
			left = i
			right = j
			min = j - i
			found = true
		}
	}
	if !found {
		return ""
	}
	return s[left-1 : right]
}
