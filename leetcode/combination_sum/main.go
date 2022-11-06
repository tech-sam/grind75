package main

import "fmt"

func main() {

	
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))

}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var findCombination func(i int, total int, arr []int)

	findCombination = func(i int, total int, arr []int) {
		if total == target {

			combination := make([]int, len(arr))
			copy(combination, arr)
			result = append(result, combination)
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		arr = append(arr, candidates[i])
		findCombination(i, total+candidates[i], arr)
		arr = arr[:len(arr)-1]
		findCombination(i+1, total, arr)

	}
	findCombination(0, 0, []int{})

	return result
}
