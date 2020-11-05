package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(ch1 chan string, ch2 chan bool) {
	for i := 0; ; i++ {
		select {
		case <-ch2:

			return
		default:
			ch1 <- "ping" + strconv.Itoa(i)
		}

	}
}

func ponger(ch1 chan string, ch2 chan bool) {
	for i := 0; ; i++ {
		select {
		case <-ch2:
			return
		default:
			ch1 <- "pong" + strconv.Itoa(i)
		}
	}
}

func printer(ch1 chan string, ch2 chan bool) {
	for {
		select {
		case <-ch2:
			return
		default:
			msg := <-ch1
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	var input string
	ch1 := make(chan string)
	ch2 := make(chan bool)

	go pinger(ch1, ch2)
	go ponger(ch1, ch2)
	go printer(ch1, ch2)

	fmt.Scanln(&input)

	ch2 <- true

	fmt.Println(input)

}
