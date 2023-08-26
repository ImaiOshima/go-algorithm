package main

import "math"

/*
*
53. 最大子数组和
*/
func maxSubArray(nums []int) int {
	length := len(nums)
	sum := 0
	max := math.MinInt
	for i := 0; i < length; i++ {
		sum = nums[i] + sum
		max = maxAbc(sum, max)
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

func maxAbc(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
