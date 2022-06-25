package main

import "fmt"

func main() {

	var added int
	added = adder(5, 10)
	fmt.Println(added)
}

// Create function that adds two values
func adder(first, second int) int {
	// define return values at top of func
	var result int
	result = first + second
	return result
}
