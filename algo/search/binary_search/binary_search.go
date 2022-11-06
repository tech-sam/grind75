package main

import "fmt"

func main() {
	list := []int{17, 25, 45, 67, 89, 123, 456, 1110}
	item := 1110
	fmt.Println(binary_search(list, item))
}

func binary_search(list []int, item int) int {

	low := 0
	high := len(list) - 1
	count := 0
	for low <= high {
		count++
		mid := low + high
		guess := list[mid]
		fmt.Printf("count %v \n", count)
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return 0
}
