// package main

// import "fmt"

// func main() {
// 	fmt.Println("Start")

// 	// Deferred function
// 	defer fmt.Println("This is deferred")

// 	fmt.Println("End")
// }

package main

import "fmt"

func main() {
	x := 10
	defer fmt.Println("Deferred value of x:", x) // x is captured as 10
	x = 20
	fmt.Println("Updated value of x:", x)
}
