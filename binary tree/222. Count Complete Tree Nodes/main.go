package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return 0
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

	fmt.Println("Output: ", countNodes(root))
}
