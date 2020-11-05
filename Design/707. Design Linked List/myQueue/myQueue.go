package main

import (
	"errors"
	"fmt"
)

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
func (q *MyQueue) Dequeue() (int, error) {
	if q.head.Back == q.tail {
		return -1, errors.New("empty queue")
	}
	ret := q.head.Back.Value
	q.head.Back = q.head.Back.Back
	return ret, nil
}

func main() {
	q := Constructor()
	data, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	data, err = q.Dequeue()
	if err == nil {
		fmt.Println("myQueue Front is: ", data) // return 1
	}
	q.Dequeue()
	fmt.Println("myQueue Front is: ", q.head.Back.Value) // return 3
}
