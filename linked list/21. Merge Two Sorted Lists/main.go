package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	previous := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			previous.Next = l1
			l1 = l1.Next
		} else {
			previous.Next = l2
			l2 = l2.Next
		}
		previous = previous.Next
	}
	if l1 != nil {
		previous.Next = l1
	} else if l2 != nil {
		previous.Next = l2
	}
	return result.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3

	node4.Next = node5
	node5.Next = node6

	l1 := node1
	l2 := node4

	fmt.Println("input: ", *node1, *node2, *node3, *node4, *node5, *node6)

	node1 = mergeTwoLists(l1, l2)

	fmt.Println("Output: ", *node1, *node1.Next, *node1.Next.Next, *node1.Next.Next.Next, *node1.Next.Next.Next.Next, *node1.Next.Next.Next.Next.Next)
}
