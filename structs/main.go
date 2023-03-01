package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	dave := person{
		firstName: "Dave", 
		lastName: "Regan", 
		contactInfo: contactInfo{
			email: "dave@mail.com", 
			zipCode: 80503,
		},
	}

	dave.updateName("New Name")
	dave.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(pointerToPerson).firstName = newFirstName
}