package main

import (
	"fmt"
)

type student struct {
	schoolID  int
	firstName string
	lastName  string
	address   string
	phoneNo   int
	marks     string
}

func main() {

	s := student{schoolID: 730493, firstName: "Zack", lastName: "Sai", address: "123 S Broadway", phoneNo: 8581234567, marks: "AAAAA"}

	toString(s)

}

func toString(s student) {
	fmt.Printf("Student: %s, %s (%d) \nMarks: %s \nPhone: %d \nAddress: %s", s.lastName, s.firstName, s.schoolID, s.marks, s.phoneNo, s.address)
}
