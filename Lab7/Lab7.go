package main

import "fmt"

func main() {

	var area float32
	area = findArea(23.92)
	fmt.Println(area)
}

func findArea(r float32) float32 {
	var result float32
	result = 3.14 * r * r
	return result
}
