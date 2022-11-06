package main

import "fmt"

func main() {

	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
func productExceptSelf(nums []int) []int {

	n := len(nums)
	result := make([]int, n)

	result[0] = 1

	for i := 1; i < n; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * r
		r = nums[i] * r
	}

	return result
}
