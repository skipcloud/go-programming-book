package main

import "fmt"

/*
	Write viradic functions max and min, analogous to sum. What should these
	functions do when called with no arguments? Write variants that require at
	least one argument.
*/

func main() {

	fmt.Printf("max with arguments: %d\n", max(1, 2, 3, 4, 5, 6, 7, 100, 4))
	fmt.Printf("max with arguments lower than 0: %d\n", max(-10, -3, -2, -9))
	fmt.Printf("max with no arguments: %d\n", max())

	fmt.Printf("min with arguments: %d\n", min(1, 2, -3, 4, 5, 6, -7, 100, 4))
	fmt.Printf("min with arguments higher than 0: %d\n", min(1, 2, 3, 4, 5, 6, 7, 100, 4))
	fmt.Printf("min with no arguments: %d\n", min())

	fmt.Printf("maxOneRequiredArg with args: %d\n", maxOneRequiredArg(2, []int{-10, -5, 300, 4}...))
	fmt.Printf("minOneRequiredArg with args: %d\n", minOneRequiredArg(3, []int{-10, -5, 300, 4}...))
}

func max(nums ...int) int {
	var max int
	if len(nums) == 0 {
		return max

	}
	// just in case all numbers are lower than 0 set it
	// to the first number
	max = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func min(nums ...int) int {
	var min int
	if len(nums) == 0 {
		return min
	}
	// just in case all numbers are higher than 0 set it
	// to the first number
	min = nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func maxOneRequiredArg(first int, nums ...int) int {
	var max int
	if len(nums) == 0 {
		return max

	}
	// just in case all numbers are lower than 0 set it
	// to the first number
	max = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func minOneRequiredArg(first int, nums ...int) int {
	var min int
	if len(nums) == 0 {
		return min
	}
	// just in case all numbers are higher than 0 set it
	// to the first number
	min = nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
