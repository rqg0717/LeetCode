package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Val: 0}
	result.Next = head
	current, previous := result, result
	i := 1
	for head != nil {
		if i >= n {
			previous = current
			current = current.Next
		}
		head = head.Next
		i++
	}
	// current is the node to be removed
	previous.Next = current.Next
	return result.Next

}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	n := 2

	fmt.Println("input: ", *node1, *node2, *node3, *node4, *node5, *node6, *node7)

	node1 = removeNthFromEnd(node1, n)

	fmt.Println("Output: ", *node1, *node2, *node3, *node4, *node5, *node6, *node7)
}
