package main

import (
	"fmt"
	"math"

	"github.com/deveshmishra34/go_crash_course/03_packages/strutils"
)

func main() {

	price := 27.024

	fmt.Println(math.Ceil(price))
	fmt.Println(math.Floor(price))
	fmt.Println(math.Sqrt(price))
	fmt.Println(strutils.Reverse("Hello"))
}
