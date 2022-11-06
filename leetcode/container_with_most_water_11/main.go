package main

import "fmt"

func main() {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

// func maxArea(height []int) int {

// 	max_area := 1

// 	for i := 0; i < len(height); i++ {
// 		for j := 0; j < len(height); j++ {
// 			width := j - i
// 			area := min(height[i], height[j]) * width
// 			if area > max_area {
// 				max_area = area
// 			}
// 		}
// 	}
// 	return max_area
// }

func maxArea(height []int) int {

	max_area, a_pointer, b_pointer := 0, 0, len(height)-1
	for a_pointer < b_pointer {

		if height[a_pointer] < height[b_pointer] {
			area := height[a_pointer] * (b_pointer - a_pointer)
			max_area = max(max_area, area)
			a_pointer++
		} else {
			area := height[b_pointer] * (b_pointer - a_pointer)
			max_area = max(max_area, area)
			b_pointer--
		}
	}

	return max_area

}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
