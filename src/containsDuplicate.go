package main

func containsDuplicate(nums []int) bool {

	mapOfValues := make(map[int]int)

	for _, currentNum := range nums {
		_, found := mapOfValues[currentNum]
		if found {
			return true
		}
		mapOfValues[currentNum] = currentNum
	}
	return false
}
