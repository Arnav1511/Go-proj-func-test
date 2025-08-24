package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contactInfo = contactInfo{email: "alex@hotvscode.com", zipCode: 94400}
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName:   "Jim",
		lastName:    "Jimothy",
		contactInfo: contactInfo{email: "Jimothy@hotvscode.com", zipCode: 94400},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
