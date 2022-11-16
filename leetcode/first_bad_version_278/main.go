package main

import "fmt"

func main() {

	limit := 10
	fmt.Println(firstBadVersion(limit))

}

func firstBadVersion(n int) int {

	if n == 0 {
		return -1
	}

	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == right && isBadVersion(left) {
		return left
	}
	return -1
}

func isBadVersion(version int) bool {
	return version == 6
}
