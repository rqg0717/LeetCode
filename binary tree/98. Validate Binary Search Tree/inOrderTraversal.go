package main

var last *TreeNode

func isValidBSTinOrderTraversal(root *TreeNode) bool {
	last = &TreeNode{Val: -1 << 63}
	return dfs2(root)
}

func dfs2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs2(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root
	return dfs2(root.Right)
}
