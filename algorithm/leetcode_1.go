package main

/*
*
1. 两数之和
*/
func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := targetMap[target-nums[i]]; ok {
			return []int{i, targetMap[target-nums[i]]}
		}
		targetMap[nums[i]] = i
	}
	return nil
}
