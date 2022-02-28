package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	// fmt.Println("buildArray: ", buildArray([]int{0, 1, 2, 3, 4, 5}))
	// fmt.Println("getConcatenation: ", getConcatenation([]int{0, 1, 2, 3, 4, 5}))
	// fmt.Println("runningSum: ", runningSum([]int{0, 1, 1, 1, 1, 1}))
	// fmt.Println("maximumWealth: ", maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	// fmt.Println("finalValueAfterOperations: ", finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	// fmt.Println("mostWordsFound: ", mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
	// fmt.Println("defangIPaddr: ", defangIPaddr("1.1.1.1"))
	// fmt.Println("shuffle: ", shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	// fmt.Println("numIdenticalPairs: ", numIdenticalPairs([]int{1, 1, 1, 1}))
	// fmt.Println("kidsWithCandies: ", kidsWithCandies([]int{1, 3, 4, 2}, 2))
	// fmt.Println("numJewelsInStones: ", numJewelsInStones("aA", "aAAbbbb"))
	// fmt.Println("twoSum: ", twoSum([]int{3, 2, 4}, 6))
	// fmt.Println("maxProfit: ", maxProfit([]int{7, 6, 4, 3, 2, 1}))
	// fmt.Println("containsDuplicate: ", containsDuplicate([]int{1, 5, -2, -4, 0}))
	// fmt.Println("productExceptSelf: ", productExceptSelf([]int{1, 2, 3, 4}))
	// fmt.Println("maxSubArray: ", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println("maxSubArrayProduct: ", maxProduct([]int{2, -5, -2, -4, 3}))
	fmt.Println("findMin: ", findMin([]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2}))
}
