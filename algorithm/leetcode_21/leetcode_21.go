package leetcode_21

import "go-algorithm/algorithm/structs"

/**
21. 合并两个有序链表
*/

func mergeTwoLists(list1 *structs.ListNode, list2 *structs.ListNode) *structs.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
