package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("seven is divisible by 2")

	} else {
		fmt.Println("seven is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")

	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or seven is even")

	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")

	} else if num < 10 {
		fmt.Println(num, "has 1 digit")

	} else {
		fmt.Println(num, "has multiple digits")

	}
}
