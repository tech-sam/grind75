package main

import "fmt"

func main() {

	list := []int{1, 3, 5, 7, 9, 13}
	item := 13
	fmt.Println("--garbage repeat---")
	fmt.Println(binary_search(list, item))
}

func binary_search(list []int, item int) int {

	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + high
		guess := list[mid]

		if guess == item {
			return mid
		} else if item < guess {
			high = mid - 1

		} else {
			low = mid + 1
		}

	}

	return 0
}
