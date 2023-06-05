package leetcode

import "fmt"

func RemoveDublicate80(nums []int) int {
	i := 0
	for _, v := range nums {
		if i < 2 || v > nums[i-2] { // here we are checking if the current element is greater than the element which is 2 positions before the current element to
			nums[i] = v
			i++
		}
	}

	// explanation:
	// nums = [1,1,1,2,2,3]
	// i = 0
	// v = 1
	// i < 2 || v > nums[i-2] => true
	// nums[i] = v => nums[0] = 1
	// i++ => i = 1
	//
	// i = 1
	// v = 1
	// i < 2 || v > nums[i-2] => true
	// nums[i] = v => nums[1] = 1
	// i++ => i = 2
	//
	// i = 2
	// v = 1
	// i < 2 || v > nums[i-2] => false
	//
	// i = 2
	// v = 2
	// i < 2 || v > nums[i-2] => true
	// nums[i] = v => nums[2] = 2
	// i++ => i = 3
	// nums = [1,1,2,2,2,3]
	// i = 3
	fmt.Println(nums)
	return i
}
