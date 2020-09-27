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

func deleteCandidate(root *TreeNode) *TreeNode {
	if root.Right == nil {
		left := root.Left
		root.Left = nil
		return left
	}
	root.Right = deleteCandidate(root.Right)
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)

		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	if root.Left == nil {
		return root.Right
	}
	candidate := root.Left
	for candidate.Right != nil {
		candidate = candidate.Right
	}
	root.Val = candidate.Val
	root.Left = deleteCandidate(root.Left)
	return root
}

func main() {
	root := &TreeNode{Val: 7}
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 5}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	node4.Left = node5
	node5.Right = node6

	key := 3
	deleteNode(root, key)

	fmt.Println("Output: ", *root, *root.Left, *root.Right)
}
