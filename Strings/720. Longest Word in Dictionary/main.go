package main

import (
	"fmt"
	"strings"
)

// Trie is a struct
type Trie struct {
	next [26]*Trie
	end  int
}

// Constructor initializes Trie.
func Constructor() Trie {
	return Trie{}
}

// Insert adds a word into the Trie.
func (t *Trie) Insert(word string, i int) {
	trie := t
	for _, v := range word {
		v -= 'a'
		if trie.next[v] == nil {
			trie.next[v] = &Trie{}
		}
		trie = trie.next[v]
	}
	trie.end = i
}

// InsertWords adds multiple words into the Trie.
func (t *Trie) InsertWords(words []string) {
	for i, v := range words {
		t.Insert(v, i+1)
	}
}

func (t *Trie) dfs(words []string) string {
	result := ""
	stack := []*Trie{t}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.end > 0 || node == t {
			if node != t {
				temp := words[node.end-1]
				n := len(temp)
				if n > len(result) || (n == len(result) && strings.Compare(result, temp) == 1) {
					result = temp
				}
			}
			for _, val := range node.next {
				if val != nil {
					stack = append(stack, val)
				}
			}
		}
	}
	return result
}

func longestWord(words []string) string {
	trie := Constructor()
	trie.InsertWords(words)
	return trie.dfs(words)
}

func main() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println("Output: ", longestWord(words))
}
