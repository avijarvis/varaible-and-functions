package main

import "fmt"

type Address struct {
	city    string
	zipcode int
}
type Person struct {
	Name    string
	Age     int
	Address // embeded struct

}

func main() {
	P := Person{
		Name: "avinash",
		Age:  27,
		Address: Address{
			city:    "pune",
			zipcode: 410005,
		},
	}
	fmt.Println("Name", P.Name)
	fmt.Println("City", P.city)       // access address city
	fmt.Println("zipcode", P.zipcode) //access addres zipcode
}
