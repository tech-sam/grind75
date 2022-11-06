package main

import "fmt"

func main() {

	var a [6]int
	length := 0
	for i := 0; i < 3; i++ {
		a[i] = i + 1
		length++
	}
	fmt.Printf("Array capacity is %v \n", len(a))
	fmt.Printf("Array length is %v \n", length)

}
