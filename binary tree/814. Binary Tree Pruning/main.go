package main

import (
	"fmt"
)

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func main() {
	// [1,0,1,0,0,0,1]
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 0}
	node5 := &TreeNode{Val: 0}
	node6 := &TreeNode{Val: 1}

	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6

	root = pruneTree(root)
	fmt.Println("Output: ", *root)
}
