package main

func productExceptSelf(nums []int) []int {
	sizeOfArray := len(nums)
	result := make([]int, sizeOfArray)

	left := 1
	for i := 0; i < sizeOfArray; i++ {
		if i > 0 {
			left = left * nums[i-1]
		}
		result[i] = left
	}

	right := 1

	for i := sizeOfArray - 1; i >= 0; i-- {
		if i < sizeOfArray-1 {
			right = right * nums[i+1]
		}
		result[i] = result[i] * right
	}
	return result
}
