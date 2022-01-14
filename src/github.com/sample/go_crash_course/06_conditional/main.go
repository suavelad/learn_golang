package main

import "fmt"

func main() {

	x := 4
	y := 5

	//If statement
	if x <= y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "green"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color  is NOT blue nor red")
	}

	// Switch statement

	switch color {
	case "red":
		fmt.Println("color is red")

	case "blue":
		fmt.Println("color is blue")

	default:
		fmt.Println("color is NOT red now blue")
	}

}
