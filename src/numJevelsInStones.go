package main

import "strings"

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
