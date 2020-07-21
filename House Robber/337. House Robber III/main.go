package main

import (
	"fmt"
)

//TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

}

func main() {
	// [3,2,3,null,3,null,1]
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 1}

	root.Left = node1
	root.Right = node2
	root.Left.Right = node3
	root.Right.Right = node4

	fmt.Println("input: ", *root, *node1, *node2, *node3, *node4)

	fmt.Println("max profit: ", rob(root))
}
