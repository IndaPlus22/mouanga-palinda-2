package main

import (
	"fmt"
)

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}

// Error: the main goroutine is asleep because sending "hello world" through channel blocks the main goroutine, preventing the next line from receiving anything through the channel
// Fix: make the channel buffered, so that it doesn't block (until it reaches > 1 string)

// Original code below

/*

package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}

*/
