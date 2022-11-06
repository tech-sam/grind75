package main

import "fmt"

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {

	var result []int
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if val, ok := m[num]; ok {
			result = append(result, val, i)
			break
		}
		m[nums[i]] = i
	}

	return result

}
