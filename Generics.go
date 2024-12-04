package main

import "fmt"

// Generic function
func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(42)      // Prints: 42
	Print("Hello") // Prints: Hello
	Print(3.14)    // Prints: 3.14
}
