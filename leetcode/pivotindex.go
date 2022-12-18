package leetcode

func PivotIndex(nums []int) int {
	totalSum := totalSum(nums)
	leftSum := 0

	for i := 0; i < len(nums); i++ {
		if i != 0 {
			leftSum += nums[i-1]
		}
		if totalSum-leftSum-nums[i] == leftSum {
			return i
		}
	}
	return -1

}

func totalSum(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans += v
	}
	return ans
}
