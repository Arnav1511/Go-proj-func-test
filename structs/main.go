package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact = contactInfo{email: "alex@hotvscode.com", zipCode: 94400}
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Jimothy",
		contact:   contactInfo{email: "Jimothy@hotvscode.com", zipCode: 94400},
	}
	fmt.Printf("%+v", jim)
}
