package main

import (
	"fmt"
	"time"
)

func main() {

	// Demonstrate conditional statements
	var a = time.Now().Weekday()

	isSaturday(a)
	isSaturdaySwitch(a)

}

// Check if it's Saturday using if/else
func isSaturday(a time.Weekday) {

	if a == time.Saturday {
		fmt.Println("TGIS!")
	} else {
		fmt.Println("Today is not Saturday.")
	}
}

// Check if it's Saturday using switch statement
func isSaturdaySwitch(a time.Weekday) {
	switch a {
	case time.Saturday:
		fmt.Println("TGIS!")
	default:
		fmt.Println("Today is not Saturday.")
	}

}
