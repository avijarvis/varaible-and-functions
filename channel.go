package main

import "fmt"

func main() {
	ch := make(chan string)

	// sender groutine
	go func() {
		ch <- "hello , groutine"

	}()
	//receiver
	message := <-ch
	fmt.Println(message) // output
}
