package main

import "fmt"

func main() {
	// Arrays

	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// fruitArr := [2]string{"Apple", "Orange"}

	//Slice

	fruitArr := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println("Length: ", len(fruitArr))
	fmt.Println("Slice: ", fruitArr[1:2])

}
