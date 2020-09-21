package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{Val: 0}
	result := sum
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		sum.Next = &ListNode{Val: tmp % 10}
		tmp = tmp / 10
		sum = sum.Next
	}
	return result.Next
}

func main() {
	node1 := &ListNode{Val: 7}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}

	node4 := &ListNode{Val: 8}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3

	node4.Next = node5
	node5.Next = node6

	l1 := node1
	l2 := node4

	node1 = addTwoNumbers(l1, l2)

	fmt.Println("Output: ", *node1, *node1.Next, *node1.Next.Next)
}
