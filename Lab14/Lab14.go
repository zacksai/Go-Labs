package main

import "fmt"

func main() {

	var n int = 0

	fmt.Println("Counting and printing backwards using defer:")

	for i := 0; i < 100; i++ {
		n = n + i
		defer fmt.Print(i, " ")
	}

	defer fmt.Println(n)

	fmt.Println("Done")

}
