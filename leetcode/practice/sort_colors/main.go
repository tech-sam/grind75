package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}

	sortColors(nums)
	fmt.Println(nums)

}

func sortColors(nums []int) {

	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	start, end, index := 0, len(nums)-1, 0
	for start < end && index <= end {
		if nums[index] == 0 {
			nums[index] = nums[start]
			nums[start] = 0
			start++
			index++

		} else if nums[index] == 2 {
			nums[index] = nums[end]
			nums[end] = 2
			end--

		} else {
			index++
		}
	}
}
