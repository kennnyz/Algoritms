package leetcode

// Given a sorted array of distinct integers and a target value, return the index
// if the target is found. If not, return the index where it would be if it were inserted in order.

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for right > left {
		midpoint := (left + right) / 2
		if target < nums[midpoint] {
			right = midpoint - 1
		} else if target > nums[midpoint] {
			left = midpoint + 1
		} else {
			return midpoint
		}
	}

	if target <= nums[left] {
		return left
	}
	return left + 1
}
