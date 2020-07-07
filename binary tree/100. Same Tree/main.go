package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// both empty, same
	if p == nil && q == nil {
		return true
	}
	// one is empty, not the same
	if p == nil || q == nil {
		return false
	}
	// different values, not the same
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil || p == nil && q != nil || p.Val != q.Val {
		return false
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, p)
	queue = append(queue, q)
	for len(queue) > 0 {
		pNode := queue[0]
		qNode := queue[1]

		if pNode.Val != qNode.Val {
			return false
		}
		if pNode.Left != nil && qNode.Left != nil {
			queue = append(queue, pNode.Left)
			queue = append(queue, qNode.Left)
		} else if pNode.Left == nil && qNode.Left != nil {
			return false
		} else if pNode.Left != nil && qNode.Left == nil {
			return false
		}
		if pNode.Right != nil && qNode.Right != nil {
			queue = append(queue, pNode.Right)
			queue = append(queue, qNode.Right)
		} else if pNode.Right == nil && qNode.Right != nil {
			return false
		} else if pNode.Right != nil && qNode.Right == nil {
			return false
		}
		queue = queue[2:]
	}
	return true
}

func main() {
	var node0, node1, node2 TreeNode
	// [1, 2, 3]
	node0.Val = 1
	node0.Left = &node1
	node0.Right = &node2
	node1.Val = 2
	node1.Left = nil
	node1.Right = nil
	node2.Val = 3
	node2.Left = nil
	node2.Right = nil

	fmt.Println("same tree? ", isSameTree(&node0, &node0))

}
