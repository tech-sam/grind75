package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool)

	for _, v := range nums {
		if m[v] {
			return true
		}
		m[v] = true
	}

	return false
}
