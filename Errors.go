package main

import (
	"errors"
	"fmt"
)

func doSomething(condition bool) error {
	if !condition {
		return errors.New("something went wrong")
	}
	return nil
}

func main() {
	err := doSomething(false)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
