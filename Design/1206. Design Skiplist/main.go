package main

import "fmt"

// Skiplist is a data structure that takes O(log(n)) time to add, erase and search.
type Skiplist struct {
	cache []uint16
}

// Constructor creates a Skiplist
func Constructor() Skiplist {
	return Skiplist{make([]uint16, 20001)}
}

// Search returns whether the target exists in the Skiplist or not.
func (sl *Skiplist) Search(target int) bool {
	return sl.cache[target] > 0
}

// Add inserts a value into the SkipList.
func (sl *Skiplist) Add(num int) {
	sl.cache[num]++
}

// Erase removes a value in the Skiplist. If num does not exist in the Skiplist, do nothing and return false. If there exists multiple num values, removing any one of them is fine.
func (sl *Skiplist) Erase(num int) bool {
	if sl.cache[num] > 0 {
		sl.cache[num]--
		return true
	}
	return false
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func main() {
	skiplist := Constructor()
	skiplist.Add(1)
	skiplist.Add(2)
	skiplist.Add(3)
	fmt.Println("Search 0: ", skiplist.Search(0)) // return false.
	skiplist.Add(4)
	fmt.Println("Search 1: ", skiplist.Search(1)) // return true.
	fmt.Println("Erase 0: ", skiplist.Erase(0))   // return false, 0 is not in skiplist.
	fmt.Println("Erase 1: ", skiplist.Erase(1))   // return true.
	fmt.Println("Search 1: ", skiplist.Search(1)) // return false, 1 has already been erased.
	skiplist.Add(20000)
	skiplist.Add(20000)
	fmt.Println("Erase 20000: ", skiplist.Erase(20000))   // return true.
	fmt.Println("Search 20000: ", skiplist.Search(20000)) // return true, duplicate.
}
