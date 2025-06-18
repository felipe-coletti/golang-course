package main

import "fmt"

func sum1(nums ...int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func sum2(nums []int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func main() {
	x := []int{1, 2, 3, 4}

	fmt.Println(sum1(x...))
	fmt.Println(sum2(x))
}
