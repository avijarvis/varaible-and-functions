package main

import "fmt"

//defining enums for day of week
type Day int

const (
	Sunday Day = iota //starts with 0 and increase automatically
	Monday
	Tuesday
	Wednesday
	Thrusday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thrusday, Friday, Saturday)

}
