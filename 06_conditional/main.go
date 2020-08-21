package main

import "fmt"

func main() {
	x := 14
	y := 10

	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	color := "yellow"

	switch color {
	case "red":
		fmt.Printf("Color is %s\n", color)
		break
	case "blur":
		fmt.Printf("Color is %s\n", color)
		break
	case "green":
		fmt.Printf("Color is %s\n", color)
		break
	default:
		fmt.Printf("No of above color is %s\n", color)
	}
}
