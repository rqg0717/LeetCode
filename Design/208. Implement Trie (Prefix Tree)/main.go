package main

import (
	"fmt"
)

// Trie is a struct
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

// Constructor initializes Trie.
func Constructor() Trie {
	return Trie{}
}

// Insert adds a word into the Trie.
func (t *Trie) Insert(word string) {
	trie := t
	for _, v := range word {
		v -= 'a'
		if trie.next[v] == nil {
			trie.next[v] = &Trie{}
		}
		trie = trie.next[v]
	}
	trie.isEnd = true
}

// Search returns the word in the trie.
func (t *Trie) Search(word string) bool {
	trie := t
	for _, v := range word {
		if trie = trie.next[v-'a']; trie == nil {
			return false
		}
	}
	return trie.isEnd
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	trie := t
	for _, v := range prefix {
		if trie = trie.next[v-'a']; trie == nil {
			return false
		}
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.StartsWith("app")) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // returns true
}
