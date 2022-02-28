package main

type Data struct {
	start int
	end   int
}

func findMin(nums []int) int {

	arraySize := len(nums) - 1
	temp := Data{
		start: 0,
		end:   arraySize,
	}

	for temp.start < temp.end {
		if nums[temp.start] < nums[temp.end] {
			return nums[temp.start]
		}
		midIndex := (temp.start + temp.end) / 2
		midValue := nums[midIndex]

		if midValue >= nums[temp.start] {
			// We know min is to the right side of midIndex, so we can forget left side of mid
			temp.start = midIndex + 1
		} else {
			// We know min is to the left side of midIndex, so we can forget right side of mid
			temp.end = midIndex
		}
	}
	return nums[temp.start]
}
