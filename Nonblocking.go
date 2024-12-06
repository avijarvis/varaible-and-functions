// Non-Blocking Channel Operations in Go
// Non-blocking channel operations allow you to interact with channels without waiting
// if a channel is not ready for sending or receiving data.
// This is particularly useful for scenarios where you want to avoid stalling your program
// when a channel is busy or empty.

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 42: // Attempt to send data
		fmt.Println("Sent 42 to the channel")
	case val := <-ch: // Attempt to receive data
		fmt.Println("Received:", val)
	default:
		fmt.Println("No channel activity")
	}
}
