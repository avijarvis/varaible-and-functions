package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a timer that fires after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer...")

	// Wait for the timer to send a value
	<-timer.C
	fmt.Println("Timer fired!")
}
