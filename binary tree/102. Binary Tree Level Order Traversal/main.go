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

func levelOrder1(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, results [][]int) [][]int {
	if root == nil {
		return results
	}
	if len(results) == level {
		results = append(results, []int{root.Val})
	} else {
		results[level] = append(results[level], root.Val)
	}
	results = dfs(root.Left, level+1, results)
	results = dfs(root.Right, level+1, results)
	return results
}

func main() {
	// [3,9,20,null,null,15,7]
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

	fmt.Println("Input: ", *root, *node1, *node2, *node3, *node4)

	fmt.Println("output: ", levelOrder1(root))
}
