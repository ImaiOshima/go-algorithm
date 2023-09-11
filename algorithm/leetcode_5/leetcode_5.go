package leetcode_5

import "fmt"

/*
*
5. 最长回文子串
*/
func longestPalindrome(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	var left, right, naibu = 0, 0, 1
	length := len(s)
	maxLength := 0
	maxStart := 0
	for i := 0; i < length; i++ {
		left = i - 1
		right = i + 1
		for left >= 0 && s[i] == s[left] {
			naibu++
			left--
		}
		for right < length && s[right] == s[i] {
			naibu++
			right++
		}
		for left >= 0 && right < length && s[left] == s[right] {
			naibu += 2
			left--
			right++
		}
		if naibu > maxLength {
			maxLength = naibu
			maxStart = left
		}
		naibu = 1
		fmt.Println(s[maxStart+1 : maxStart+maxLength+1])
	}
	return s[maxStart+1 : maxStart+maxLength+1]
}

//func main() {
//	s := "cbbd"
//	longestPalindrome(s)
//}
