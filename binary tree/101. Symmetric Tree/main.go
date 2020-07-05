package main

import "fmt"

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrorRecursive(root.Left, root.Right)
}

func isMirrorRecursive(p *TreeNode, q *TreeNode) bool {
	// both empty, Symmetric
	if p == nil && q == nil {
		return true
	}
	// one is empty, not Symmetric
	if p == nil || q == nil {
		return false
	}
	// compares left and right, right and left ...
	return p.Val == q.Val && isMirrorRecursive(p.Left, q.Right) && isMirrorRecursive(p.Right, q.Left)
}

func isMirrorIterate(left *TreeNode, right *TreeNode) bool {
	q := []*TreeNode{} // stack
	q = append(q, left, right)

	for len(q) > 0 {
		q1, q2 := q[0], q[1]
		q = q[2:] // pop

		if q1 == nil && q2 == nil {
			continue
		}
		if q1 == nil || q2 == nil {
			return false
		}
		if q1.Val != q2.Val {
			return false
		}

		q = append(q, q1.Left, q2.Right)
		q = append(q, q1.Right, q2.Left)
	}
	return true
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

	fmt.Println("isSymmetric? ", isSymmetric(&node0))
}
