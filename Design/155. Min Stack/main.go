package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// MinStack is a struct
type MinStack struct {
	stack []int
	min   []int
}

// Constructor initializes MinStack.
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{math.MaxInt64},
	}
}

// Push adds an item
func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	top := m.min[len(m.min)-1]
	m.min = append(m.min, min(top, x))
}

// Pop removes the top item
func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.min = m.min[:len(m.min)-1]
}

// Top returns the top item
func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

// GetMin gets the minimum
func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin()) // returns -3
	stack.Pop()
	fmt.Println(stack.Top())    // returns 0
	fmt.Println(stack.GetMin()) // returns -2
}
