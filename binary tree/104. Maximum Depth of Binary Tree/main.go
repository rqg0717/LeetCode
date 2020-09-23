package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 20}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	root.Left = node1
	root.Right = node2
	root.Left.Left = nil
	root.Left.Right = nil
	root.Right.Left = node3
	root.Right.Right = node4
	root.Right.Left.Left = nil
	root.Right.Left.Right = nil
	root.Right.Right.Left = nil
	root.Right.Right.Right = nil

	fmt.Println("input: ", *root, *node1, *node2, *node3, *node4)

}
