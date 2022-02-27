package main

func maxProfit(prices []int) int {
	var lowestPrice int = MaxInt
	var maxProfit int = 0
	var profitNow int = 0

	for _, priceNow := range prices {
		if priceNow < lowestPrice {
			lowestPrice = priceNow
		}

		profitNow = priceNow - lowestPrice

		if maxProfit < profitNow {
			maxProfit = profitNow
		}

	}
	return maxProfit
}
