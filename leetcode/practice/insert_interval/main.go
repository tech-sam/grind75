package main

import "fmt"

func main() {

	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {

	l, r := [][]int{}, [][]int{}
	a, b := newInterval[0], newInterval[1]

	for _, interval := range intervals {

		if interval[1] < a {
			l = append(l, interval)
		} else if interval[0] > b {
			r = append(r, interval)
		} else {
			a = min(a, interval[0])
			b = max(b, interval[1])
		}

	}

	return append(append(l, []int{a, b}), r...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
