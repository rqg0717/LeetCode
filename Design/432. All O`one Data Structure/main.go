package main

// DoublyLinkedNode ...
type DoublyLinkedNode struct {
	previous *DoublyLinkedNode
	next     *DoublyLinkedNode
	key      string
	value    int
}

// DoublyLinkedList ...
type DoublyLinkedList struct {
	head   *DoublyLinkedNode
	tail   *DoublyLinkedNode
	length int
}

// AllOne ...
type AllOne struct {
	dmap map[string]*DoublyLinkedNode
	list DoublyLinkedList
}

// Constructor Initializes data structure here.
func Constructor() AllOne {
	head := &DoublyLinkedNode{
		key:   "",
		value: 0,
	}
	tail := &DoublyLinkedNode{
		key:   "",
		value: 0,
	}
	head.next = tail
	tail.previous = head
	return AllOne{
		dmap: make(map[string]*DoublyLinkedNode),
		list: DoublyLinkedList{
			head:   head,
			tail:   tail,
			length: 0,
		},
	}
}

// Inc Inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (ao *AllOne) Inc(key string) {

}

// Dec Decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (ao *AllOne) Dec(key string) {

}

// GetMaxKey Returns one of the keys with maximal value.
func (ao *AllOne) GetMaxKey() string {

}

// GetMinKey Returns one of the keys with Minimal value.
func (ao *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
