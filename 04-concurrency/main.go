package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!")
	<-done
	fmt.Println("All greetings done!")
}