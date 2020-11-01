package main

import "fmt"

// ListNode is a structure
type ListNode struct {
	Value    int
	Next     *ListNode
	Previous *ListNode
}

// MyLinkedList is a double-pointer linked list
type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

// Constructor initializes the linked list
func Constructor() MyLinkedList {
	head := &ListNode{0, nil, nil}
	tail := &ListNode{0, nil, nil}
	head.Next = tail
	tail.Previous = head

	return MyLinkedList{head, tail, 0}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (mll *MyLinkedList) Get(index int) int {
	// if index is invalid
	if index < 0 || index >= mll.size {
		return -1
	}

	current := mll.head
	if index+1 < mll.size-index {
		for i := 0; i < index+1; i++ {
			current = current.Next
		}
	} else {
		current = mll.tail
		for i := 0; i < mll.size-index; i++ {
			current = current.Previous
		}
	}

	return current.Value

}

// AddAtHead adds a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (mll *MyLinkedList) AddAtHead(val int) {
	predecessor := mll.head
	successor := mll.head.Next

	tmp := &ListNode{val, nil, nil}
	tmp.Previous = predecessor
	tmp.Next = successor
	predecessor.Next = tmp
	successor.Previous = tmp

	mll.size++

}

// AddAtTail appends a node of value val to the last element of the linked list.
func (mll *MyLinkedList) AddAtTail(val int) {
	successor := mll.tail
	predecessor := mll.tail.Previous

	tmp := &ListNode{val, nil, nil}
	tmp.Previous = predecessor
	tmp.Next = successor
	predecessor.Next = tmp
	successor.Previous = tmp

	mll.size++
}

// AddAtIndex adds a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (mll *MyLinkedList) AddAtIndex(index int, val int) {

	if index > mll.size {
		return
	}

	if index < 0 {
		index = 0
	}

	// find predecessor and successor of the node to be added
	predecessor := mll.head
	successor := mll.tail
	if index < mll.size-index {
		for i := 0; i < index; i++ {
			predecessor = predecessor.Next
		}
		successor = predecessor.Next
	} else {

		for i := 0; i < mll.size-index; i++ {
			successor = successor.Previous
		}
		predecessor = successor.Previous
	}

	tmp := &ListNode{val, nil, nil}
	tmp.Previous = predecessor
	tmp.Next = successor
	predecessor.Next = tmp
	successor.Previous = tmp

	mll.size++
}

// DeleteAtIndex removes the index-th node in the linked list, if the index is valid.
func (mll *MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || index >= mll.size {
		return
	}

	// find predecessor and successor of the node to be deleted
	predecessor := mll.head
	successor := mll.tail
	if index < mll.size-index {
		for i := 0; i < index; i++ {
			predecessor = predecessor.Next
		}

		successor = predecessor.Next.Next
	} else {
		for i := 0; i < mll.size-index-1; i++ {
			successor = successor.Previous
		}

		predecessor = successor.Previous.Previous
	}

	// delete predecessor.Next
	predecessor.Next = successor
	successor.Previous = predecessor

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
