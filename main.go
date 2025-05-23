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
	// antonio := person{firstName: "Antonio", lastName: "Espinoza"}
	// fmt.Println(antonio)
	// var antonio person
	// antonio.firstName = "Antonio"
	// antonio.lastName = "Espinoza"
	// fmt.Printf("%+v", antonio)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()

}

func (p person) updateName(firstName string) {
	p.firstName = firstName
	fmt.Println("Updated Name:", p.firstName)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
