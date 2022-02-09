package main

import "fmt"

func main() {
	fmt.Print(buildArray([]int{0, 1, 2, 3, 4, 5}))

}

func buildArray(nums []int) []int {

	var sizeOfArray = len(nums)
	var slice = make([]int, sizeOfArray) // or slice := make([]int, elems)

	for i := 0; i < sizeOfArray; i++ {
		slice[i] = nums[nums[i]]
	}
	return slice
}
