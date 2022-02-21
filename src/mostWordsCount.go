package main

import "strings"

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
