package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello , Goroutine")

}
func main() {
	go sayHello() //starting new goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
