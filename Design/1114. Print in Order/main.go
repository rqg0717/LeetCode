package main

import (
	"fmt"
)

// Foo is a struct
type Foo struct {
	startTwo   chan bool
	startThree chan bool
	done       chan bool
}

func (f *Foo) one() {
	fmt.Println("one")
	f.startTwo <- true
}

func (f *Foo) two() {
	<-f.startTwo
	fmt.Println("two")
	f.startThree <- true
}

func (f *Foo) three() {
	<-f.startThree
	fmt.Println("three")
	f.done <- true
}

func main() {
	f := new(Foo)
	f.startTwo = make(chan bool)
	f.startThree = make(chan bool)
	f.done = make(chan bool)
	go f.one()
	go f.two()
	go f.three()
	<-f.done
}
