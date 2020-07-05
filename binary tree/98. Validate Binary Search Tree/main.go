package main

import (
	"errors"
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

// Insert insets a node into a binary search tree
func (bst *TreeNode) Insert(value int) error {
	if bst == nil {
		return errors.New("Tree is null")
	}

	if bst.Val == value {
		return errors.New("This node already exists")
	}

	if bst.Val > value {
		if bst.Left == nil {
			bst.Left = &TreeNode{Val: value}
			return nil
		}
		return bst.Left.Insert(value)
	}

	if bst.Val < value {
		if bst.Right == nil {
			bst.Right = &TreeNode{Val: value}
			return nil
		}
		return bst.Right.Insert(value)
	}
	return nil
}

func main() {
	t := &TreeNode{Val: 3}

	err := t.Insert(6)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Insert(1)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Insert(5)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Insert(2)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Insert(7)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Insert(4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("isValidBST? ", isValidBST(t))
}
