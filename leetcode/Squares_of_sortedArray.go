package leetcode

// Given an integer array nums sorted in non-decreasing order,
//return an array of the squares of each number sorted in non-decreasing order.

func SortedSquares(arr []int) []int {
	res := make([]int, len(arr))
	j, k := 0, len(arr)-1
	for i := 0; i < len(arr); i++ {
		a, b := arr[j]*arr[j], arr[k]*arr[k]
		if a > b {
			res[len(arr)-1-i] = a
			j++
			continue
		}
		res[len(arr)-1-i] = b
		k--
	}
	return res
}
