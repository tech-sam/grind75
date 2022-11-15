package main

import "fmt"

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2

		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1

}
