package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.

// This is the function we'll run in a goroutine.
// The done channel will be used to notify another goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//Send a value to notify that we're done.
	done <- true
}

func main() {

	// Start a worker goroutine, giving it a channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	// If we removed the <-done line, the program would exit before the worker even started.
	fmt.Println(<-done)
}
