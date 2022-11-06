package main

import (
	"fmt"
	"sort"
)

func main() {

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals := [][]int{{0, 4}, {1, 4}}
	//intervals := [][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}
	//intervals := [][]int{{5, 7}, {4, 6}}

	fmt.Println(merge(intervals))
	// [[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]]

}

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{}
	current_interval := intervals[0]
	result = append(result, current_interval)

	for _, interval := range intervals {
		//current_begin := current_interval[0]
		current_end := current_interval[1]

		next_begin := interval[0]
		next_end := interval[1]

		if current_end >= next_begin {
			current_interval[1] = max(current_end, next_end)
		} else {
			current_interval = interval
			result = append(result, current_interval)
		}
	}

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
