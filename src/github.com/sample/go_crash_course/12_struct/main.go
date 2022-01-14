package main

import (
	"fmt"
	"strconv"
)

// Define a Struct

type Person struct {
	firstname, lastname, city, gender string
	age                               int
}

// Greeting method (Value Receiver)

func (p Person) greet() string {
	return "My Name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age)

}

// hasBirthday method ( pointer receiver)
func (p *Person) hasBirthday() {

	p.age++

}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastName
	}

}

func main() {
	// Init a person using struct
	person1 := Person{firstname: "John", lastname: "Doe", city: "London", gender: "f", age: 25}

	// Alternative to  Init a person using struct
	person2 := Person{"John", "Doe", "London", "f", 25}

	fmt.Println(person1)
	fmt.Println(person2)

	// Get a single field from the Struct

	fmt.Println(person2.firstname)
	fmt.Println(person2.age)
	person2.hasBirthday()
	person2.hasBirthday()
	fmt.Println(person2.age)
	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}
