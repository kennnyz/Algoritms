package leetcode

// 33. Find First and Last Position of Element in Sorted Array

func searchRange(nums []int, target int) []int {
	begin, end := findingBegin(nums, target), findingEnd(nums, target)

	return []int{begin, end}
}

func findingBegin(nums []int, target int) int {
	index := -1

	start := 0
	end := len(nums) - 1

	for start <= end {
		midpoint := start + (end-start)/2
		if nums[midpoint] >= target {
			end = midpoint - 1
		} else {
			start = midpoint + 1
		}

		if nums[midpoint] == target {
			index = midpoint
		}

	}

	return index
}

func findingEnd(nums []int, target int) int {
	index := -1
	start := 0
	end := len(nums) - 1

	for start <= end {
		midpoint := start + (end-start)/2
		if nums[midpoint] <= target {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}
		if nums[midpoint] == target {
			index = midpoint
		}
	}

	return index
}
