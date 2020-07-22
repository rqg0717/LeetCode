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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// not rob = 0; rob = 1
// rob root: root:1 + left:0  + right:0
// not rob root: root:0 & max(left:0, left:1) + max(right:0, right:1)
func robNodes(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	left1, left0 := robNodes(root.Left)
	right1, right0 := robNodes(root.Right)
	return root.Val + left0 + right0,
		max(left0, left1) + max(right0, right1)
}

func rob(root *TreeNode) int {
	root1, root0 := robNodes(root)
	return max(root0, root1)
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

	//fmt.Println("input: ", *root, *node1, *node2, *node3, *node4)

	fmt.Println("Maximum amount of money the thief can rob: ", rob(root))
}
