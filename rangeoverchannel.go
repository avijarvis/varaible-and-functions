package main

import (
	"fmt"
)

func main() {
	// Create a channel
	ch := make(chan int)

	// Start a goroutine to send values
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Send values to the channel
		}
		close(ch) // Close the channel when done
	}()

	// Range over the channel
	for val := range ch {
		fmt.Println(val) // Process each value received from the channel
	}
}
