package main

import "fmt"

func adder() func(int) int {

	Sum := 0
	return func(x int) int {
		Sum += x
		return Sum
	}

}

func main() {

	Sum := adder()
	for i := 0; i < 10; i++ {

		fmt.Println(Sum(i))

	}

}
