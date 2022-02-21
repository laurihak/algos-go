package main

import "regexp"

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
