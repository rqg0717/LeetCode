package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]int)
	current := head
	for current != nil {
		if _, exist := nodeMap[current]; exist {
			return current
		}
		nodeMap[current] = 1
		current = current.Next
	}
	return nil
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println("input: ", *node1, *node2, *node3, *node4)

	fmt.Println("Output: ", *detectCycle(node1))
}
