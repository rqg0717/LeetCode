package main

import "fmt"

// Stack is a struct
type Stack struct {
	nums []int
}

// NewStack init the stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Push adds an item
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop removes the last item
func (s *Stack) Pop() int {
	last := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return last
}

// IsEmpty returns if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}

// MyQueue is a struct
type MyQueue struct {
	in, out *Stack
}

// Constructor initializes the queue
func Constructor() MyQueue {
	return MyQueue{in: NewStack(), out: NewStack()}
}

// Push an item in the back of the queue
func (q *MyQueue) Push(x int) {
	q.in.Push(x)
}

// Pop removes the item from the front of the queue
func (q *MyQueue) Pop() int {
	if len(q.out.nums) > 0 {
		return q.out.Pop()
	}
	for len(q.in.nums) > 0 {
		q.out.Push(q.in.Pop())
	}
	return q.out.Pop()
}

// Peek obtains the item from the front of the queue
func (q *MyQueue) Peek() int {
	item := q.Pop()
	q.out.Push(item)
	return item
}

// Empty returns if the queue is empty
func (q *MyQueue) Empty() bool {
	return (len(q.in.nums) + len(q.out.nums)) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor()
	queue.Push(-2)
	queue.Push(0)
	queue.Push(-3)
	fmt.Println(queue.Peek())  // returns -2
	fmt.Println(queue.Pop())   // returns -2
	fmt.Println(queue.Pop())   // returns 0
	fmt.Println(queue.Pop())   // returns -3
	fmt.Println(queue.Empty()) // returns true
}
