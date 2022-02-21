package main

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
