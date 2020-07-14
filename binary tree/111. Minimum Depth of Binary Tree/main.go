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

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil {
		return right + 1
	} else if root.Right == nil {
		return left + 1
	} else {
		return min(left, right) + 1
	}
}

func main() {
	// [3,9,20,1,null,15,7]
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}

	root.Left = node1
	root.Right = node2
	root.Right.Left = node3
	root.Right.Right = node4
	root.Left.Left = node5

	fmt.Println("input: ", *root, *node1, *node2, *node3, *node4, *node5)

	results := minDepth(root)

	fmt.Println("Output: ", results)
}
