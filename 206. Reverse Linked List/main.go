package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var temp *ListNode = nil
	current := head
	for nil != current {
		next := current.Next
		current.Next = temp
		temp = current
		current = next
	}
	return temp
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println("input: ", *node1, *node2, *node3, *node4, *node5)
	reverseList(node1)
	fmt.Println("Output: ", *node1, *node2, *node3, *node4, *node5)
}
