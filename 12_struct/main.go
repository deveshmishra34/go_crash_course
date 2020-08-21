package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// age       int

	firstName, lastName, city string
	age                       int
}

func (p Person) greetings() string {
	return "hello my name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age) + " year old"
}

func (p *Person) hasBirthDay() {
	p.age++
}

func main() {

	person1 := Person{firstName: "Dev", lastName: "mishra", city: "New Delhi", age: 25}

	fmt.Println(person1.greetings())
	person1.hasBirthDay()
	person1.hasBirthDay()
	person1.hasBirthDay()
	person1.hasBirthDay()
	fmt.Println(person1.greetings())
}
