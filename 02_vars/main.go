package main

import "fmt"

func main() {

	// var name = "Devesh"
	// var age = 25
	const bool = true
	// bool = false

	name, age := "Dev", 25

	fmt.Printf("My name is %s and my age is %d and it is %t \n", name, age, bool)
	fmt.Printf("Age is type of %T\n", age)
	fmt.Printf("Name is type of %T\n", name)
	fmt.Printf("Bool is type of %T\n", bool)
}
