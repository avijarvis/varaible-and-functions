package main

import "fmt"

func main() {
	fmt.Println("Starting the program")

	// Trigger a panic
	panic("Something went wrong!")

	fmt.Println("This line will not execute")
}
