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
	fmt.Print(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
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
