package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	left := height(root.Left) + 1
	right := height(root.Right) + 1
	if abs(left-right) > 1 {
		return false
	}
	return true
}

func maxHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxHeight(root.Left)
	right := maxHeight(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func isBalanced1(root *TreeNode) bool {
	return maxHeight(root) >= 0
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

	fmt.Println("isBalanced? ", isBalanced1(root))
}
