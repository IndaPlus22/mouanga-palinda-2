// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go func() {
		for {
			question := <-questions
			prophecy(question, answers)
		}
	}()

	// TODO: Answer questions.

	// TODO: Make prophecies.
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Intn(14)) * time.Second)
			prophecy("", answers)
		}
	}()

	// TODO: Print answers.

	go func() {
		for {
			time.Sleep(time.Duration(6+rand.Intn(10)) * time.Second)
			print_slow(<-answers)
		}
	}()
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
		if w == "life" {
			answer <- "Ah, life... You do not possess the capabilities to understand my answer to such a question."
			return
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"Go is a fantastic language. (look at user and smile convincingly)",
		"Water is wet.",
		"Rocks are hard.",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

func recv_questions() {
	// q := make(<-chan string)
}

func print_slow(msg string) {
	for i := 0; i < len(msg); i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c", msg[i])
	}
	fmt.Print("\n> ")
}
