package main

import "fmt"

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}

	if len(matrix) == 0 {
		return res
	}

	rowBegin, rowEnd, colBegin, colEnd := 0, len(matrix)-1, 0, len(matrix[0])-1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
		}
		rowBegin++
		for i := rowBegin; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--
		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				res = append(res, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			for i := rowEnd; i >= rowBegin; i-- {
				res = append(res, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return res
}
