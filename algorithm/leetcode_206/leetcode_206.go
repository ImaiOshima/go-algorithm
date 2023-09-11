package leetcode_206

import (
	"go-algorithm/algorithm/structs"
)

/*
*
206. 反转链表
*/

// 递归
func reverseList(head *structs.ListNode) *structs.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}

// 迭代
func reverseList2(head *structs.ListNode) *structs.ListNode {
	var l *structs.ListNode = nil
	r := head
	for r != nil {
		next := r.Next
		r.Next = l
		l = r
		r = next
	}
	return l
}
