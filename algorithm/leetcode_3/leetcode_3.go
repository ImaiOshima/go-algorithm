package leetcode_3

/*
*3. 无重复字符的最长子串
 */
func lengthOfLongestSubstring(s string) int {
	var max, l int
	cnt := [128]int{}
	for r, c := range s {
		cnt[c]++
		// 注意 这个是for循环 不是if条件 大于2说明有重复元素
		for cnt[c] > 1 {
			// for循环 从老的左边界开始排查 到底是那个超过了2 因为重复元素是超过2的
			cnt[s[l]]--
			l++
		}
		max = Max(max, r-l+1)
	}
	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
