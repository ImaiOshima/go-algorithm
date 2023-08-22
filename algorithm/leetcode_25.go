package main

/*
*25. K 个一组翻转链表
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dump := &ListNode{}
	dump.Next = head
	lNode := dump
	rNode := dump
	for rNode.Next != nil {
		for i := 0; i < k && rNode.Next != nil; i++ {
			rNode = rNode.Next
		}
		if rNode == nil {
			break
		}
		pre := lNode.Next
		next := rNode.Next
		rNode.Next = nil
		lNode.Next = reverseNode(pre)
		pre.Next = next
		lNode = pre
		rNode = pre
	}
	return dump.Next
}

func reverseNode(target *ListNode) *ListNode {
	if target == nil || target.Next == nil {
		return target
	}
	curr := reverseNode(target.Next)
	target.Next.Next = target
	target.Next = nil
	return curr
}
