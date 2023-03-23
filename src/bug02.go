package main

import (
	"fmt"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	main_ch := make(chan bool)
	ch := make(chan int)
	go wait_print(main_ch, ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	<-main_ch
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
}

func wait_print(ch1 chan bool, ch2 <-chan int) {
	Print(ch2)
	ch1 <- true
}
