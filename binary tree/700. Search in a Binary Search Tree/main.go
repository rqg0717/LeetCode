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

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func main() {
	// [4,2,5,1,3]
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4

	val := 5

	fmt.Println("Output: ", searchBST(root, val))
}
