package main

import "fmt"

func main() {
	num := 42

	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0:
		fmt.Println("Positive number")
	}
}
