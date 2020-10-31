package main

import "fmt"

// ListNode is a structure
type ListNode struct {
	Value int
	Next  *ListNode
}

// MyLinkedList is a double-pointer linked list
type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

// Constructor initializes the linked list
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (mll *MyLinkedList) Get(index int) int {
	if index < 0 || index >= mll.size {
		return -1
	}
	head := mll.head
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head.Value
}

// AddAtHead adds a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (mll *MyLinkedList) AddAtHead(val int) {
	mll.head = &ListNode{
		Value: val,
		Next:  mll.head,
	}
	if mll.size == 0 {
		mll.tail = mll.head
	}
	mll.size++
}

// AddAtTail appends a node of value val to the last element of the linked list.
func (mll *MyLinkedList) AddAtTail(val int) {
	if mll.size == 0 {
		mll.tail = &ListNode{
			Value: val,
			Next:  nil,
		}
		mll.head = mll.tail
	} else {
		mll.tail.Next = &ListNode{
			Value: val,
			Next:  nil,
		}
		mll.tail = mll.tail.Next
	}
	mll.size++
}

// AddAtIndex adds a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (mll *MyLinkedList) AddAtIndex(index int, val int) {
	if index > mll.size {
		return
	}
	if index <= 0 {
		mll.AddAtHead(val)
		return
	}
	if index == mll.size {
		mll.AddAtTail(val)
		return
	}
	head := mll.head
	for i := 1; i < index; i++ {
		head = head.Next
	}
	head.Next = &ListNode{
		Value: val,
		Next:  head.Next,
	}
	mll.size++
}

// DeleteAtHead removes the node from head.
func (mll *MyLinkedList) DeleteAtHead() {
	if mll.size == 0 {
		return
	}
	mll.head = mll.head.Next
	if mll.size == 1 {
		mll.tail = mll.head
	}
	mll.size--
}

// DeleteAtTail removes the node from tail.
func (mll *MyLinkedList) DeleteAtTail() {
	if mll.size == 0 {
		return
	}
	head := mll.head
	for head.Next != mll.tail {
		head = head.Next
	}
	mll.tail = head
	if mll.size == 1 {
		mll.head = mll.tail
	}
	mll.size--
}

// DeleteAtIndex removes the index-th node in the linked list, if the index is valid.
func (mll *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= mll.size {
		return
	}
	if index == 0 {
		mll.DeleteAtHead()
		return
	}
	if index == mll.size-1 {
		mll.DeleteAtTail()
		return
	}
	head := mll.head
	for i := 1; i < index; i++ {
		head = head.Next
	}
	head.Next = head.Next.Next
	mll.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2)               // linked list becomes 1->2->3
	fmt.Println("return ", myLinkedList.Get(1)) // return 2
	myLinkedList.DeleteAtIndex(1)               // now the linked list is 1->3
	fmt.Println("return ", myLinkedList.Get(1)) // return 3
}
