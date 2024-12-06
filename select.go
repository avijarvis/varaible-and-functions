// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "Message from ch1"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "Message from ch2"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println(msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

//		SELECT WITH DEFAULT

package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("No data received")
	}
}
