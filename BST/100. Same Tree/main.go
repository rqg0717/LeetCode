package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// both empty, same
	if p == nil && q == nil {
		return true
	}
	// one is empty, not the same
	if p == nil || q == nil {
		return false
	}
	// different values, not the same
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	var node0, node1, node2 TreeNode
	// [1, 2, 3]
	node0.Val = 1
	node0.Left = &node1
	node0.Right = &node2
	node1.Val = 2
	node1.Left = nil
	node1.Right = nil
	node2.Val = 3
	node2.Left = nil
	node2.Right = nil

	fmt.Println("same tree? ", isSameTree(&node0, &node0))

}
