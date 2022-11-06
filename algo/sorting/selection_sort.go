package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	//var sortArray []int
	fmt.Println(selectionSort(nums))
}

func selectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		min_index := i

		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		tmp := nums[min_index]
		nums[min_index] = nums[i]
		nums[i] = tmp

	}

	return nums
}
