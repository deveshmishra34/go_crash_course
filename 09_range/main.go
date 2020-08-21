package main

import "fmt"

func main() {
	ids := []int{12, 24, 35, 7, 6, 7}

	fmt.Println(ids)

	// print index and value
	for i, id := range ids {
		fmt.Printf("%d : ID - %d\n", i, id)
	}

	// sum of all ids
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Ids sum is %d \n", sum)

	emails := map[string]string{"one": "one@gmail.com", "two": "two@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s - %s \n", k, v)
	}
}
