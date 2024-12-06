// Tickers in Go, like timers, are part of the time package and are used to repeatedly trigger
// an action at specified intervals. Unlike a Timer, which fires only once, a Ticker continues
//  to send values on its channel at regular intervals until it is stopped.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Stop the ticker after 5 seconds to demonstrate how to stop it
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}()

	// Loop to receive ticks from the ticker's channel
	for range ticker.C {
		fmt.Println("Tick")
	}
}
