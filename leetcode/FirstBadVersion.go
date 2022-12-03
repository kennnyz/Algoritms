package leetcode

// LeetCode 278. First Bad Version

var BV int = 4

func isBadversion(n int) bool {
	return n >= BV
}
func FirstBadVersion(n int) int {
	left := 0
	right := n

	for left < right {
		midpoint := left + (right-left)/2
		if isBadversion(midpoint) {
			right = midpoint
		} else {
			left = midpoint + 1
		}
	}
	if left == right && isBadversion(left) {
		return left
	}
	return -1
}
