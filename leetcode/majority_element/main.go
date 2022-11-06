package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {

	count, result := 0, 0

	for _, v := range nums {
		if count == 0 {
			result = v
			count++
		} else if result == v {
			count++
		} else {
			count--
		}
		fmt.Printf("-this is result and count %d - %d \n", result, count)
	}
	return result

}
