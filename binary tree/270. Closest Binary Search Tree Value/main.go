package main

import (
	"fmt"
	"math"
)

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

func main() {
	// [4,2,5,1,3]
	root := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 1}
	node4 := &TreeNode{Val: 3}

	root.Left = node1
	root.Right = node2
	root.Left.Left = node3
	root.Left.Right = node4

	target := 3.714286

	fmt.Println("Output: ", closestValue(root, target))
}
