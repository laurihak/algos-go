package main

func twoSum(nums []int, target int) []int {
	sizeOfArray := len(nums)

	valueIndexMap := make(map[int]int)

	for i := 0; i < sizeOfArray; i++ {
		valueToFind := target - nums[i]
		value, found := valueIndexMap[valueToFind]

		if found && i != value {
			return []int{i, value}
		}
		valueIndexMap[nums[i]] = i
	}
	return []int{}
}
