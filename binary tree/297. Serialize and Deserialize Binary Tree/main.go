package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode is definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec is not in use
type Codec struct {
	s []string
}

// Constructor is not in use
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (s *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "NULL"
	}
	return strconv.Itoa(root.Val) + "," + s.serialize(root.Left) + "," + s.serialize(root.Right)
}

// depth-first search
func dfs(s *[]string) *TreeNode {
	str := (*s)[0]
	*s = (*s)[1:]
	if str == "NULL" {
		return nil
	}
	val, _ := strconv.Atoi(str)
	node := &TreeNode{Val: val}
	node.Left = dfs(s)
	node.Right = dfs(s)
	return node
}

// Deserializes your encoded data to tree.
func (s *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	return dfs(&str)
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
func main() {
	// [4,2,5,1,3]
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 4}
	node4 := &TreeNode{Val: 5}

	root.Left = node1
	root.Right = node2
	root.Right.Left = node3
	root.Right.Right = node4

	obj := Constructor()
	s := obj.serialize(root)
	fmt.Println("serialize: ", s)
	d := obj.deserialize(s)
	fmt.Println("deserialize: ", d)
}
