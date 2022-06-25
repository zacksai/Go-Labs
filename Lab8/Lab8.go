package main

import "fmt"

func main() {

	fmt.Println(rectangleArea(3, 5))

}

func rectangleArea(length, breadth int) int {
	var result int
	result = length * breadth
	return result
}
