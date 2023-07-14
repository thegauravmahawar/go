package main

import "fmt"

func main() {

	// Channels are the pipes that connect concurrent goroutines.
	// You can send values into channels from one goroutine and receive those values
	// into another goroutine.
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
