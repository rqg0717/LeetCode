package main

import (
	"fmt"
)

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

func main() {
	// [3,5,1,6,2,0,8,null,null,7,4]
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 6}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 0}
	node6 := &TreeNode{Val: 8}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 4}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	root.Right.Left = node5
	root.Right.Right = node6
	root.Left.Right.Left = node7
	root.Left.Right.Right = node8

	fmt.Println("Output: ", lowestCommonAncestor(root, node1, node8).Val)
}
