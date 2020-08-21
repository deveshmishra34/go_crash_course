package main

import "fmt"

func greetings(name string) string {
	return "hello " + name
}

func sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(greetings("Dev"))
	fmt.Println(sum(7, 5))
}
