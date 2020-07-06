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

func levelOrder(root *TreeNode) [][]int {
	// Breadth-first search & O(n) & O(n)

	results := [][]int{}

	if root == nil {
		return results
	}
	// start with the first level
	first := []*TreeNode{root}
	for i := 0; len(first) > 0; i++ {
		results = append(results, []int{})
		next := []*TreeNode{}
		for j := 0; j < len(first); j++ {
			node := first[j]
			results[i] = append(results[i], node.Val)
			// build the next level
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		// go to the next level
		first = next
	}

	return results
}

func main() {
	// [3,9,20,null,null,15,7]
	root := TreeNode{}
	node1 := TreeNode{}
	node2 := TreeNode{}
	node3 := TreeNode{}
	node4 := TreeNode{}

	root.Val = 3
	root.Left = &node1
	root.Right = &node2
	root.Left.Val = 9
	root.Left.Left = nil
	root.Left.Right = nil
	root.Right.Val = 20
	root.Right.Left = &node3
	root.Right.Right = &node4
	root.Right.Left.Val = 15
	root.Right.Left.Left = nil
	root.Right.Left.Right = nil
	root.Right.Right.Val = 7
	root.Right.Right.Left = nil
	root.Right.Right.Right = nil

	fmt.Println("input: ", root, node1, node2, node3, node4)

	results := levelOrder(&root)

	fmt.Println("output: ", results)
}
