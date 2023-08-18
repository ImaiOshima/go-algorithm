package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return curr
}

// 迭代
func reverseList2(head *ListNode) *ListNode {
	var l *ListNode = nil
	r := head
	for r != nil {
		next := r.Next
		r.Next = l
		l = r
		r = next
	}
	return l
}
