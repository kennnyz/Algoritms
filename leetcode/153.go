package leetcode

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		midpoint := left + (right-left)/2

		if nums[midpoint] < nums[midpoint-1] {
			return nums[midpoint]
		} else if nums[left] <= nums[midpoint] && nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint - 1
		}
	}
	return nums[left]
}
