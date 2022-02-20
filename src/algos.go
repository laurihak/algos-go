package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Print(buildArray([]int{0, 1, 2, 3, 4, 5}))
	// fmt.Print(getConcatenation([]int{0, 1, 2, 3, 4, 5}))
	// fmt.Print(runningSum([]int{0, 1, 1, 1, 1, 1}))
	// fmt.Print(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	// fmt.Print(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	// fmt.Print(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
	// fmt.Print(defangIPaddr("1.1.1.1"))
	// fmt.Print(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	// fmt.Print(numIdenticalPairs([]int{1, 1, 1, 1}))
	// fmt.Print(kidsWithCandies([]int{1, 3, 4, 2}, 2))
	// fmt.Print(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Print(twoSum([]int{3, 2, 4}, 6))

}

func buildArray(nums []int) []int {

	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray) // or slice := make([]int, elems)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[nums[i]]
	}
	return slice
}

func getConcatenation(nums []int) []int {
	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray*2)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[i]
		slice[i+sizeOfArray] = nums[i]
	}
	return slice
}

func runningSum(nums []int) []int {
	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray)

	for i := 0; i < sizeOfArray; i++ {
		if i < 1 {
			slice[i] = nums[i]
			continue
		}
		slice[i] = nums[i] + slice[i-1]
	}
	return slice
}

func maximumWealth(accounts [][]int) int {

	var sizeOfAccounts = len(accounts)

	var res = 0

	for i := 0; i < sizeOfAccounts; i++ {
		var sizeOfAccount = len(accounts[i])
		var accountSum = 0
		for j := 0; j < sizeOfAccount; j++ {
			accountSum = accountSum + accounts[i][j]
		}

		if res < accountSum {
			res = accountSum
		}
	}
	return res
}

func finalValueAfterOperations(operations []string) int {

	var sizeOfArray = len(operations)
	var result = 0

	var add, _ = regexp.Compile(`\+`)

	for i := 0; i < sizeOfArray; i++ {
		var value = operations[i]
		if add.MatchString(value) {
			result++
			continue
		} else {
			result--
			continue
		}

	}
	return result
}

func mostWordsFound(sentences []string) int {
	sizeOfArray := len(sentences)

	largestWordCount := 0

	for i := 0; i < sizeOfArray; i++ {
		var result = sentences[i]

		var count = len(strings.Split(result, " "))

		if largestWordCount < count {
			largestWordCount = count
		}
	}
	return largestWordCount
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

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

func max(v []int) int {
	m := 0
	for i, e := range v {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func numJewelsInStones(jewels string, stones string) int {

	jewelsArray := strings.Split(jewels, "")
	stonesArray := strings.Split(stones, "")

	count := 0
	for _, v := range stonesArray {
		for _, v2 := range jewelsArray {
			if v == v2 {
				count++
			}
		}
	}
	return count
}

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
