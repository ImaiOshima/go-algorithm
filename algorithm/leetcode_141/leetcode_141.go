package leetcode_141

import (
	"go-algorithm/algorithm/structs"
)

/*
*
141. 环形链表
*/
func hasCycle(head *structs.ListNode) bool {
	l, r := head, head
	for r != nil && r.Next != nil {
		l, r = l.Next, r.Next.Next
		if l == r {
			return true
		}
	}
	return false
}
