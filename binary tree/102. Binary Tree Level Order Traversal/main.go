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

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		results = append(results, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			results[i] = append(results[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	return results
}

func main() {
	var node0, node1, node2, node3, node4 TreeNode
	// [3,9,20,null,null,15,7]
	node0.Val = 3
	node0.Left = &node1
	node0.Right = &node2
	node1.Val = 9
	node1.Left = nil
	node1.Right = nil
	node2.Val = 20
	node2.Left = &node3
	node2.Right = &node4
	node3.Val = 15
	node3.Left = nil
	node3.Right = nil
	node4.Val = 7
	node4.Left = nil
	node4.Right = nil

	fmt.Println("input: ", node0, node1, node2, node3, node4)

	results := levelOrder(&node0)

	fmt.Println("output: ", results)
}
