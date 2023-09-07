package main

/**
20. 有效的括号
*/

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	con := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if _, ok := con[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != con[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
