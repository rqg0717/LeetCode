package main

import "fmt"

// Node is a structure
type Node struct {
	Value int
	Front *Node
	Back  *Node
}

// MyQueue is a struture contains two sentinel nodes
type MyQueue struct {
	head *Node
	tail *Node
}

// Constructor initializes the queue
func Constructor() MyQueue {
	head := &Node{0, nil, nil}
	tail := &Node{0, nil, nil}
	head.Back = tail
	tail.Front = head

	return MyQueue{head, tail}
}

// Enqueue pushes a node of value val into the back of the queue.
func (q *MyQueue) Enqueue(val int) {
	tail := q.tail
	head := q.tail.Front

	tmp := &Node{val, nil, nil}
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
