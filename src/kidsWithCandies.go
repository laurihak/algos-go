package main

func kidsWithCandies(candies []int, extraCandies int) []bool {

	sizeOfArray := len(candies)

	maxValue := max(candies)

	result := make([]bool, sizeOfArray)

	for i := 0; i < sizeOfArray; i++ {
		if candies[i]+extraCandies >= maxValue {
			result[i] = true
			continue
		}
		result[i] = false
	}
	return result
}
