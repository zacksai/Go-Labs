package main

import "fmt"

// Cleint struct
type client struct {
	firstName   string
	lastName    string
	address     string
	phoneNo     int
	accountNo   int
	savingsAmt  int
	checkingAmt int
}

func main() {

	// Create and print a client using getters
	c := client{firstName: "Warren", lastName: "Buffet", address: "1777 Wall St",
		phoneNo: 7601231234, accountNo: 99999999, savingsAmt: 999999999, checkingAmt: 1}
	toString(c)
}

// Getters
func (c *client) getFirstName() string {
	return c.firstName
}
func (c *client) getLastName() string {
	return c.lastName
}
func (c *client) getAddress() string {
	return c.address
}
func (c *client) getPhoneNo() int {
	return c.phoneNo
}
func (c *client) getAccountNo() int {
	return c.accountNo
}
func (c *client) getSavingsAmt() int {
	return c.savingsAmt
}
func (c *client) getCheckingAmt() int {
	return c.checkingAmt
}

func toString(c client) {
	fmt.Printf("Client: %s, %s (Acct %d)\nAddress: %s\nPhone: %d\nChecking:$%d\nSavings:$%d",
		c.lastName, c.firstName, c.accountNo, c.address, c.phoneNo, c.checkingAmt, c.savingsAmt)
}
