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
	s.setFirstName("Bob")
	s.setLastName("Schmitt")
	s.setSchoolID(193035)
	s.setAddress("123 Front St")
	s.setMarks("ABABB")
	s.setPhoneNo(6191234567)
	toString(s)
}

// Setters
func (s *student) setFirstName(firstName string) {
	s.firstName = firstName
}
func (s *student) setLastName(lastName string) {
	s.lastName = lastName
}
func (s *student) setMarks(marks string) {
	s.marks = marks
}
func (s *student) setAddress(address string) {
	s.address = address
}
func (s *student) setPhoneNo(phoneNo int) {
	s.phoneNo = phoneNo
}
func (s *student) setSchoolID(schoolID int) {
	s.schoolID = schoolID
}

// Getters
func (s student) getFirstName() string {
	return s.firstName
}
func (s student) getLastName() string {
	return s.lastName
}
func (s student) getMarks() string {
	return s.marks
}
func (s student) getAddress() string {
	return s.address
}
func (s student) getPhoneNo() int {
	return s.phoneNo
}
func (s student) getSchoolID() int {
	return s.schoolID
}

func toString(s student) {
	fmt.Printf("Student: %s, %s (%d) \nMarks: %s \nPhone: %d \nAddress: %s\n\n", s.getLastName(), s.getFirstName(), s.getSchoolID(), s.getMarks(), s.getPhoneNo(), s.getAddress())
}
