package main

import "algos/utils"

func maxSubArray(nums []int) int {
	subArrays := make([]int, len(nums))

	for i, value := range nums {
		if i == 0 {
			subArrays[i] = value
			continue
		}

		lastSum := subArrays[i-1] + value

		if lastSum > value {
			subArrays[i] = lastSum
			continue
		}
		subArrays[i] = value
	}
	maxValue := utils.Max(subArrays)
	return maxValue
}
