package main

import "sort"

/**
15. 三数之和
*/

func threeSum(nums []int) [][]int {
	var ans [][]int
	length := len(nums)
	if nums != nil && length < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		l := i + 1
		r := length - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for nums[l] == nums[l-1] {
					l++
				}
				for nums[r] == nums[r-1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
