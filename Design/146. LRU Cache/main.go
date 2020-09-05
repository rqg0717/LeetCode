package main

import "fmt"

// LinkedNode is a doubly linked list
type LinkedNode struct {
	key, value     int
	previous, next *LinkedNode
}

// LRUCache is a Least Recently Used (LRU) cache
type LRUCache struct {
	nodes      map[int]*LinkedNode
	capacity   int
	head, tail *LinkedNode
}

// Constructor constructs LRUCache
func Constructor(capacity int) LRUCache {
	head := &LinkedNode{0, 0, nil, nil}
	tail := &LinkedNode{0, 0, nil, nil}
	head.next = tail
	tail.previous = head
	return LRUCache{make(map[int]*LinkedNode), capacity, head, tail}
}

// RemoveNode removes a node from the doubly linked list
func (lru *LRUCache) RemoveNode(node *LinkedNode) {
	node.previous.next = node.next
	node.next.previous = node.previous
}

// AddNode inserts a node into the doubly linked list
func (lru *LRUCache) AddNode(node *LinkedNode) {
	head := lru.head
	node.next = head.next
	head.next.previous = node
	node.previous = head
	head.next = node
}

// MoveToHead moves a node to head
func (lru *LRUCache) MoveToHead(node *LinkedNode) {
	lru.RemoveNode(node)
	lru.AddNode(node)
}

// Get obtains the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
func (lru *LRUCache) Get(key int) int {
	nodes := lru.nodes
	if node, existed := nodes[key]; existed {
		lru.MoveToHead(node)
		return node.value
	}
	return -1
}

// Put inserts the value if the key is not already present.
func (lru *LRUCache) Put(key int, value int) {
	head := lru.head
	tail := lru.tail
	nodes := lru.nodes

	if node, existed := nodes[key]; existed {
		node.value = value
		lru.MoveToHead(node)
	} else {
		node := &LinkedNode{key, value, nil, nil}
		if len(nodes) == lru.capacity { // it's full, evicts previous node
			delete(nodes, tail.previous.key) // built-in function deletes the element with the specified key (m[key]) from the map
			tail.previous.previous.next = tail
			tail.previous = tail.previous.previous
		}
		node.next = head.next
		node.previous = head
		head.next.previous = node
		head.next = node
		nodes[key] = node
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println("Get: ", cache.Get(1)) // returns 1
	cache.Put(3, 3)
	fmt.Println("Get: ", cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)
	fmt.Println("Get: ", cache.Get(1)) // returns -1 (not found)
	fmt.Println("Get: ", cache.Get(3)) // returns 3
	fmt.Println("Get: ", cache.Get(4)) // returns 4
}
