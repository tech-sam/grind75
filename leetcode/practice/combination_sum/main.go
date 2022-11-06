package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 9
	fmt.Println(combinationSum(candidates, target))

}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			var cur_copy []int
			cur_copy = append(cur_copy, cur...)
			result = append(result, cur_copy)
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		cur = append(cur, candidates[i])
		dfs(i, cur, total+candidates[i])
		cur = cur[:len(cur)-1]
		dfs(i+1, cur, total)

	}
	dfs(0, []int{}, 0)
	return result
}
