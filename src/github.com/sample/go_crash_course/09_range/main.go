package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not Using Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with Map

	emails := map[string]string{"Bob": "bob@mail.com", "John": "john@mail.com", "Mike": "mike@mail.com"}
	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

}
