package main

func shuffle(nums []int, n int) []int {

	sizeOfArray := len(nums)

	result := make([]int, sizeOfArray)

	j := 0

	for i := 0; i < (sizeOfArray / 2); i++ {
		result[j] = nums[i]
		result[j+1] = nums[i+n]
		j = j + 2
	}

	return result
}
