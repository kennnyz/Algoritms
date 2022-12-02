// LeetCode 704. Binary Search
package leetcode

//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
//If target exists, then return its index. Otherwise, return -1.
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target == nums[len(nums)/2] {
		return len(nums) / 2
	} else if target > nums[len(nums)/2] {
		for i := len(nums) / 2; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := len(nums) / 2; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
		}
	}
	return -1
}
