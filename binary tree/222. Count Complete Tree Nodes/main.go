package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countLevel(root *TreeNode) int {
	level := 0
	for root != nil {
		root = root.Left
		level++
	}
	return level
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	} else {
		return countNodes(root.Left) + (1 << right)
	}
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 4}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 6}
	node6 := &TreeNode{Val: 7}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4
	node4.Left = node5
	node5.Right = node6

	fmt.Println("Output: ", countNodes(root))
}
