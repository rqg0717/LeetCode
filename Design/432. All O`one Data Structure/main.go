package main

import "fmt"

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
		}}
}

func (ao *AllOne) removeNode(node *DoublyLinkedNode) {
	node.previous.next = node.next
	node.next.previous = node.previous
}

func (ao *AllOne) addNode(node, current *DoublyLinkedNode) {
	node.next = current
	node.previous = current.previous
	current.previous.next = node
	current.previous = node
}

// Inc Inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (ao *AllOne) Inc(key string) {
	if node, ok := ao.dmap[key]; ok {
		node.value++
		current := node.next
		for current != ao.list.tail {
			if current.value >= node.value {
				break
			}
			current = current.next
		}
		if node.next != current {
			ao.removeNode(node)
			ao.addNode(node, current)
		}
	} else {
		node := &DoublyLinkedNode{
			key:   key,
			value: 1,
		}
		node.next = ao.list.head.next
		node.previous = ao.list.head
		ao.list.head.next.previous = node
		ao.list.head.next = node

		ao.dmap[key] = node
		ao.list.length++
	}
}

// Dec Decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (ao *AllOne) Dec(key string) {
	if node, ok := ao.dmap[key]; ok {
		if node.value <= 1 {
			ao.removeNode(node)
			ao.list.length--
			delete(ao.dmap, key)
		} else {
			node.value--
			current := node.previous
			for current != ao.list.tail {
				if current.value <= node.value {
					break
				}
				current = current.previous
			}
			if node.previous != current {
				ao.removeNode(node)
				ao.addNode(node, current)
			}
		}
	}
}

// GetMaxKey Returns one of the keys with maximal value.
func (ao *AllOne) GetMaxKey() string {
	if ao.list.length == 0 {
		return ""
	}
	return ao.list.tail.previous.key
}

// GetMinKey Returns one of the keys with Minimal value.
func (ao *AllOne) GetMinKey() string {
	if ao.list.length == 0 {
		return ""
	}
	return ao.list.head.next.key
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
func main() {
	ao := Constructor()
	ao.Inc("1")
	ao.Inc("2")
	ao.Inc("3")
	ao.Dec("1")
	fmt.Println("Max Key: ", ao.GetMaxKey())
	fmt.Println("Min Key: ", ao.GetMinKey())
}
