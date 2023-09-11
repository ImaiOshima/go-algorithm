package leetcode_102

import "go-algorithm/algorithm/structs"

/**
102. 二叉树的层序遍历
*/

func levelOrder(root *structs.TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	list := []*structs.TreeNode{root}
	for len(list) != 0 {
		size := len(list)
		vals := make([]int, size)
		for i := range vals {
			// 实现队列 每次删除头结点
			node := list[0]
			list = list[1:]
			vals[i] = node.Val

			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
