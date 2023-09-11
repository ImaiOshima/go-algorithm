package leetcode_215

/*
*
215. 数组中的第K个最大元素
*/
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	l, r := 0, len(nums)-1
	tIndex := n - k
	for {
		index := Partition(nums, l, r)
		if index == tIndex {
			return nums[index]
		} else if index < tIndex {
			l = index + 1
		} else {
			r = index - 1
		}
	}
	return -1
}

func Partition(nums []int, start, end int) int {
	if start >= end {
		return start
	}
	pivot := nums[start]
	l, r := start, end
	// 需要遍历全部 按照pivot分成小于和大于pivot 最后还需要l和r替换
	for i := l + 1; i <= r; i++ {
		if nums[i] <= pivot {
			l++
			nums[l], nums[i] = nums[i], nums[l]
		}
	}
	nums[l], nums[start] = nums[start], nums[l]
	return l
}
