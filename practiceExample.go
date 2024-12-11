package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)

	fmt.Println("Addition:", num1+num2)
	fmt.Println("Subtraction:", num1-num2)
	fmt.Println("Multiplication:", num1*num2)
	if num1 != 0 && num2 != 0 {
		fmt.Println("Division:", num1/num2)
	}
}
