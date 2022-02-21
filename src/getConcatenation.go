package main

func getConcatenation(nums []int) []int {
	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray*2)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[i]
		slice[i+sizeOfArray] = nums[i]
	}
	return slice
}
