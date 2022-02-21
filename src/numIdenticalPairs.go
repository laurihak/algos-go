package main

func numIdenticalPairs(nums []int) int {

	sizeOfArray := len(nums)

	count := 0
	for i := 0; i < sizeOfArray; i++ {
		for j := i + 1; j < sizeOfArray; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}
	return count
}
