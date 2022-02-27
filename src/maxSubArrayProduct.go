package main

import (
	"algos/utils"
)

type data struct {
	min int
	max int
}

func maxProduct(nums []int) int {

	subArrays := make([]int, len(nums))
	previous := data{
		min: nums[0],
		max: nums[0],
	}

	for i, value := range nums {

		if value < 0 {
			// Current value is negative
			//  use negative from last round to get max result

			tempMin := previous.min
			previous.min = utils.Min([]int{previous.max * value, value})
			previous.max = utils.Max([]int{tempMin * value, value})
		} else {
			previous.max = utils.Max([]int{previous.max * value, value})
			previous.min = utils.Min([]int{previous.min * value, value})
		}

		if previous.max > value {
			subArrays[i] = previous.max
			continue
		}
		subArrays[i] = value
	}
	maxValue := utils.Max(subArrays)
	return maxValue
}
