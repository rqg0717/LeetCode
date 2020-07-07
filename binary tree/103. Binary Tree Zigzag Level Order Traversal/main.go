package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	if root == nil {
		return results
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeftStart := true
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeftStart {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}
		}
		results = append(results, ans)
		isLeftStart = !isLeftStart
		queue = queue[l:]
	}
	return results
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
	root.Right.Left = node3
	root.Right.Right = node4

	fmt.Println("input: ", *root, *node1, *node2, *node3, *node4)

	results := zigzagLevelOrder(root)

	fmt.Println("output: ", results)
}
