//syntax

//func FunctionName[T TypeConstraint](param T) ReturnType {
// Function body
//}

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

//Generics in Go are implemented using type parameters.
//A type parameter is a placeholder for a type that is specified when the function or type is used.
//T is the type parameter.
//any is a built-in type constraint that represents any type.
