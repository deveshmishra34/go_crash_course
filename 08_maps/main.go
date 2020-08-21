package main

import "fmt"

func main() {
	// declare
	// emails := make(map[string]string)
	// assign
	// emails["one"] = "one@gmail.com"
	// emails["two"] = "two@gmail.com"

	emails := map[string]string{"one": "one@gmail.com", "two": "two@gmail.com"}
	fmt.Println(emails)
	fmt.Println(emails["two"])
	fmt.Println(len(emails))

	delete(emails, "one")
	fmt.Println(emails)
	fmt.Println(len(emails))
}
