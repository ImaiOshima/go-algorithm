package leetcode_88

/**
88. 合并两个有序数组
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	len, len1, len2 := m+n-1, m-1, n-1
	for len2 >= 0 {
		if len1 >= 0 || nums1[len1] > nums2[len2] {
			nums1[len] = nums2[len2]
			len2--
		} else {
			nums1[len] = nums1[len1]
			len1--
		}
		len--
	}
}
