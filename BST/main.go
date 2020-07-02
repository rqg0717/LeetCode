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

var last, first, second *TreeNode

func recoverTree(root *TreeNode) {
	last, first, second = nil, nil, nil
	inOrderTraversal(root)
	first.Val, second.Val = second.Val, first.Val
}

func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left)
	if last != nil && root.Val <= last.Val {
		if first != nil {
			second = root
			return
		}
		first, second = last, root
	}
	last = root
	inOrderTraversal(root.Right)
}

func main() {
	var node0, node1, node2, node3 TreeNode
	// [3,1,4,null,null,2]
	node0.Val = 3
	node0.Left = &node1
	node0.Right = &node2
	node1.Val = 1
	node1.Left = nil
	node1.Right = nil
	node2.Val = 4
	node2.Left = &node3
	node2.Right = nil
	node3.Val = 2
	node3.Left = nil
	node3.Right = nil

	fmt.Println("input: ", node0, node1, node2, node3)

	recoverTree(&node0)

	fmt.Println("output: ", node0, node1, node2, node3)

}
