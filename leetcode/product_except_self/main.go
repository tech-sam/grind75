package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	//nums := []int{-1, 1, 0, -3, 3}
	//nums := []int{0, 0}
	fmt.Println(productExceptSelf(nums))
	// [-1,1,0,-3,3]

}

func productExceptSelf(nums []int) []int {
	n := len(nums)

	output_arr := make([]int, n)
	output_arr[0] = 1

	for i := 1; i < n; i++ {
		output_arr[i] = nums[i-1] * output_arr[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		output_arr[i] = output_arr[i] * r
		r = r * nums[i]
	}

	return output_arr

}

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)

// 	left_products, right_products, output_arr := make([]int, n), make([]int, n), make([]int, n)

// 	left_products[0] = 1
// 	right_products[n-1] = 1

// 	for i := 1; i < n; i++ {
// 		left_products[i] = nums[i-1] * left_products[i-1]
// 	}

// 	for i := n - 2; i >= 0; i-- {
// 		right_products[i] = nums[i+1] * right_products[i+1]
// 	}

// 	for i := 0; i < n; i++ {
// 		output_arr[i] = left_products[i] * right_products[i]
// 	}

// 	return output_arr

// }

// func productExceptSelf(nums []int) []int {

// 	l, r := make(map[int]int), make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		product := 1
// 		for j := i + 1; j < len(nums); j++ {

// 			product = nums[j] * product
// 			fmt.Printf("---coming inside %d \n", product)

// 		}
// 		r[nums[i]] = product
// 		fmt.Printf("---product outside %d \n", product)

// 	}
// 	fmt.Println(r)

// 	for i := len(nums) - 1; i >= 0; i-- {
// 		product := 1
// 		for j := i - 1; j >= 0; j-- {
// 			product = nums[j] * product

// 		}
// 		l[nums[i]] = product

// 	}
// 	fmt.Println(l)
// 	result := []int{}

// 	for _, v := range nums {
// 		product := l[v] * r[v]
// 		result = append(result, product)
// 	}
// 	return result
// }
