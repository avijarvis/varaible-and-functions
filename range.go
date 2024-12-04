package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40}

	// Loop over index and value
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

}
