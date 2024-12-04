package main

import "fmt"

//definging interface
type Shape interface {
	Area() float64
}

//defining struct
type Circle struct {
	Radius float64
}

//Area of circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func main() {
	var s Shape // declare vairable of type shape
	s = Circle{Radius: 8}
	fmt.Println("Area of circle:", s.Area()) // call the method
}
