package main

import "fmt"

func main() {

	// Print name 10 times with for loop
	fmt.Println("Printing first loop:")
	var firstName string = "Zack"
	for i := 0; i < 10; i++ {
		fmt.Print(firstName, " ")
	}

	// Print first name 5 times and last name 8 times
	n := 13
	var lastName string = "Sai"
	fmt.Println("\nPrinting second loop:")

	for n > 0 {
		if n < 9 {
			fmt.Print(lastName, " ")
		} else {
			fmt.Print(firstName, " ")
		}
		n--
	}
}
