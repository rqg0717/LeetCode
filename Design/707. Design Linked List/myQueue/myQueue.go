package main

import "fmt"

// ListNode is a structure
type ListNode struct {
	Value int
	Front *ListNode
	Back  *ListNode
}

// MyQueue is a double-pointer linked list
type MyQueue struct {
	head *ListNode
	tail *ListNode
}

// Constructor initializes the queue
func Constructor() MyQueue {
	head := &ListNode{0, nil, nil}
	tail := &ListNode{0, nil, nil}
	head.Back = tail
	tail.Front = head

	return MyQueue{head, tail}
}

// Enqueue pushes a node of value val into the back of the queue.
func (q *MyQueue) Enqueue(val int) {
	tail := q.tail
	head := q.tail.Front

	tmp := &ListNode{val, nil, nil}
	tmp.Front = head
	tmp.Back = tail
	head.Back = tmp
	tail.Front = tmp
}

// Dequeue removes the first node in the front of the queue
func (q *MyQueue) Dequeue() {
	if q.head.Back == q.tail {
		return
	}
	q.head.Back = q.head.Back.Back
}

func main() {
	myQueue := Constructor()
	myQueue.Dequeue()
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println("myQueue Front is: ", myQueue.head.Back.Value) // return 3
}
