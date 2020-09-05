package main

import "fmt"

// LinkedNode is a doubly linked list
type LinkedNode struct {
	key, value, frequency int
	previous, next        *LinkedNode
}

// DoublyLinkedNode is a doubly linked list
type DoublyLinkedNode struct {
	size       int
	head, tail *LinkedNode
}

// LFUCache is a Least Frequently Used (LFU) cache
type LFUCache struct {
	capacity    int
	current     int // current the least used
	size        int
	cache       map[int]*LinkedNode
	frequencies map[int]*DoublyLinkedNode
}

// Constructor creates LFUCache
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:    capacity,
		current:     0,
		size:        0,
		cache:       make(map[int]*LinkedNode),
		frequencies: make(map[int]*DoublyLinkedNode),
	}
}

func (dln *DoublyLinkedNode) remove(node *LinkedNode) {
	previous := node.previous
	previous.next = node.next
	previous.next.previous = previous
}

// RemoveNode removes a node from the doubly linked list
func (lfu *LFUCache) RemoveNode(node *LinkedNode) {
	dln := lfu.frequencies[node.frequency]
	dln.remove(node)
	dln.size--
}

func (dln *DoublyLinkedNode) add(node *LinkedNode) {
	next := dln.head.next
	next.previous = node
	dln.head.next = node
	node.next = next
	node.previous = dln.head
}

// AddNode inserts a node into the doubly linked list
func (lfu *LFUCache) AddNode(node *LinkedNode) {
	if val, ok := lfu.frequencies[node.frequency]; ok {
		val.add(node)
		val.size++
	} else {
		dln := &DoublyLinkedNode{0, &LinkedNode{}, &LinkedNode{}}
		dln.head.next = dln.tail
		dln.tail.previous = dln.head
		dln.add(node)
		dln.size++
		lfu.frequencies[node.frequency] = dln
	}
}

// Get obtains the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
func (lfu *LFUCache) Get(key int) int {
	if node, ok := lfu.cache[key]; ok {
		lfu.RemoveNode(node)
		if node.frequency == lfu.current && lfu.frequencies[node.frequency].size == 0 {
			lfu.current++
		}
		node.frequency++
		lfu.AddNode(node)
		return node.value
	}
	return -1
}

func (dln *DoublyLinkedNode) removeTail() *LinkedNode {
	previous := dln.tail.previous
	previous.previous.next = dln.tail
	dln.tail.previous = previous.previous
	return previous
}

func (lfu *LFUCache) removeCurrent() {
	dln := lfu.frequencies[lfu.current]
	key := dln.removeTail().key
	dln.size--
	delete(lfu.cache, key)
}

// Put inserts the value if the key is not already present.
func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	newnode := &LinkedNode{key, value, 1, nil, nil}
	if node, ok := lfu.cache[key]; !ok {
		if lfu.size < lfu.capacity {
			lfu.AddNode(newnode)
			lfu.size++
			lfu.current = 1
		} else {
			lfu.removeCurrent()
			lfu.current = 1
			lfu.AddNode(newnode)
		}
	} else {
		lfu.RemoveNode(node)
		if lfu.current == node.frequency && lfu.frequencies[node.frequency].size == 0 {
			lfu.current++
		}
		newnode.frequency = node.frequency + 1
		lfu.AddNode(newnode)

	}
	lfu.cache[key] = newnode
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println("Get: ", cache.Get(1)) // returns 1
	cache.Put(3, 3)
	fmt.Println("Get: ", cache.Get(2)) // returns -1 (not found)
	fmt.Println("Get: ", cache.Get(3)) // returns 3
	cache.Put(4, 4)
	fmt.Println("Get: ", cache.Get(1)) // returns -1 (not found)
	fmt.Println("Get: ", cache.Get(3)) // returns 3
	fmt.Println("Get: ", cache.Get(4)) // returns 4
}
