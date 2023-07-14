package main

import "fmt"

// By default channels are unbuffered, meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value.
// Buffered channels accepts a limited number of values without a corresponding receiver to those values.
func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
