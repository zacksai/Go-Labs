package main

import "fmt"

func main() {

	var a int = 1000
	var l *int
	l = &a

	fmt.Println(l)
	fmt.Println(*l)

}
