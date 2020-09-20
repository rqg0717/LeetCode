package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result, previous := &ListNode{Val: 0}, &ListNode{Val: 0}
	result.Next = head
	current := result
	i := 1
	for head != nil {
		if i >= n {
			previous = current
			current = current.Next
		}
		head = head.Next
		i++
	}
	previous.Next = previous.Next.Next
	return result.Next

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

	n := 2

	fmt.Println("input: ", *node1, *node2, *node3, *node4, *node5)

	node1 = removeNthFromEnd(node1, n)

	fmt.Println("Output: ", *node1, *node2, *node3, *node5)
}
