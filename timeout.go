// Timeouts with select
// You can implement timeouts using the time.
// After function, which returns a channel that sends a value after a specified duration.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred")
	}
}
