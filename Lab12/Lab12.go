package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float32 = 31.4159
	var area float32

	area = math.Pi * radius * radius

	fmt.Printf("Area of circle with radius %f is %f", radius, area)
}
