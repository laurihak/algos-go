package main

import "fmt"

func main() {
	// fmt.Print(buildArray([]int{0, 1, 2, 3, 4, 5}))
	// fmt.Print(getConcatenation([]int{0, 1, 2, 3, 4, 5}))
	fmt.Print(runningSum([]int{0, 1, 1, 1, 1, 1}))
}

func buildArray(nums []int) []int {

	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray) // or slice := make([]int, elems)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[nums[i]]
	}
	return slice
}

func getConcatenation(nums []int) []int {
	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray*2)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[i]
		slice[i+sizeOfArray] = nums[i]
	}
	return slice
}

func runningSum(nums []int) []int {
	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray)

	for i := 0; i < sizeOfArray; i++ {
		if i < 1 {
			slice[i] = nums[i]
			continue
		}
		slice[i] = nums[i] + slice[i-1]
	}
	return slice
}
