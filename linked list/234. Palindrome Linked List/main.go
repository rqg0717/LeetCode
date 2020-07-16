package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	reverse, temp := fast, slow

	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next

		temp.Next = reverse
		reverse = temp
	}

	if nil != fast { // length is odd
		slow = slow.Next
	}

	for ; slow != nil && reverse != nil; slow, reverse = slow.Next, reverse.Next {
		if slow.Val != reverse.Val {
			return false
		}
	}

	return true
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 2}
	node5 := &ListNode{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	fmt.Println("input: ", *node1, *node2, *node3, *node4)

	fmt.Println("Output: ", isPalindrome(node1))
}
