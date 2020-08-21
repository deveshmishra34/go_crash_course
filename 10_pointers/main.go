package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("a is %d, b is %d \n", a, b)

	// use * to read the value
	fmt.Printf("a is %d, b is %d \n", *&a, *b)

	*b = 10
	fmt.Printf(" after change a is %d, b is %d \n", *&a, *b)
}
