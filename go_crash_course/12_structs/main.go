package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName, city, gender string
	age int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++;
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return;
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person { firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25 }
	person2 := Person { "Bob", "Johnson", "New Your", "m", 5 }
	// Alternative
	// person1 := Person { "Samantha", "Smith", "Boston", "f", 25 }

	fmt.Println(person1);
	fmt.Println(person1.firstName);
	
	person1.age++;
	fmt.Println(person1.age);
	
	person1.hasBirthday();
	person1.getMarried("Williams");
	person2.getMarried("Thomson")

	fmt.Println(person1.greet());
	fmt.Println(person2.greet());
}