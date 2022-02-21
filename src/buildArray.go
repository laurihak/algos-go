package main

func buildArray(nums []int) []int {

	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray) // or slice := make([]int, elems)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[nums[i]]
	}
	return slice
}
