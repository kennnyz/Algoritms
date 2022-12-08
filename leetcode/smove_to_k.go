package leetcode

func Rotate(nums []int, k int) []int {
	if k%len(nums) == 0 {
		return nums
	}
	for i := 0; i < k; i++ {
		for j := len(nums) - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}z
