package main

import "fmt"

// Define a struct type called 'Person'
type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Create an instance of Person
	p1 := Person{Name: "avinash", Age: 27, Email: "avinashk200597@gmail.com"}

	// Access and print struct fields
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println("Email:", p1.Email)

	// Modify a field
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)

	// Create a struct using the field names implicitly
	p2 := Person{"viru", 14, "viru14@gmail.com"}
	fmt.Println("Person 2:", p2)
}
