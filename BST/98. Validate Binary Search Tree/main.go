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

func isValidBST(root *TreeNode) bool {
	// min and max int64
	return dfs(root, -1<<63, 1<<63-1)
}

// dfs - Depth-First-Search
func dfs(root *TreeNode, min, max int) bool {
	// Pre-Order Traversal
	return root == nil || min < root.Val && root.Val < max &&
		dfs(root.Left, min, root.Val) &&
		dfs(root.Right, root.Val, max)
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

	fmt.Println("isValidBST? ", isValidBST(&node0))
}
