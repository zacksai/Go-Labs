package main

import "fmt"

func main() {

	// Printf formats using specifiers
	fmt.Printf("Printf:\n%s is %d years old and his favorite letter is %q\n", "Zack", 25, 'h')

	// Println prints with a line and spaces arguments
	var a string = "this"
	fmt.Println("Println:\nThe parentheses around (", a, ") are automatically spaced by println")

	// Print just prints
	fmt.Print("Print:\nHello world")

	// Errorf formats using specifiers and returns string giving a desired error
	const name, id = "sanchez", 43
	err := fmt.Errorf("user %q (id #%d) not found", name, id)
	fmt.Println("\nError:\n", err.Error())
}
