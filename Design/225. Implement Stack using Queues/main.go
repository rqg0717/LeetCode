package main

import "fmt"

// MyStack  is a struct
type MyStack struct {
	queue []int
}

// Constructor init the stack
func Constructor() MyStack {
	return MyStack{queue: []int{}}
}

// Pop removes the last item
func (s *MyStack) Pop() int {
	item := s.queue[0]
	s.queue = s.queue[1:]
	return item
}

// Push adds an item
func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for n > 0 {
		s.queue = append(s.queue, s.Pop())
		n--
	}
}

// Top returns the first tiem
func (s *MyStack) Top() int {
	return s.queue[0]
}

// Empty returns if the stack is empty
func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.Top())   // returns -3
	fmt.Println(stack.Pop())   // returns -3
	fmt.Println(stack.Empty()) // returns false
}
