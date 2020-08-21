package main

import "fmt"

func main() {
	// Arrays
	// var fruitsArr [2]string
	// fruitsArr[0] = "Apple"
	// fruitsArr[1] = "Orange"
	// fruitsArr[2] = "banana"
	fruitsArr := [2]string{"Apple", "orange"}
	fmt.Println(fruitsArr)

	fruitsSlices := []string{"Apple", "orange", "banana"}
	// fruitsSlices[3] = "Grapes"
	fmt.Println(fruitsSlices)
	fmt.Println(len(fruitsSlices))
	fmt.Println(fruitsSlices[0:2])
}
