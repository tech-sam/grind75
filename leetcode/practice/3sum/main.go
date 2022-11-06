package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			threeSum := v + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				result = append(result, []int{v, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}

		}
	}
	return result

}
