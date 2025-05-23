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
	// antonio := person{firstName: "Antonio", lastName: "Espinoza"}
	// fmt.Println(antonio)
	// var antonio person
	// antonio.firstName = "Antonio"
	// antonio.lastName = "Espinoza"
	// fmt.Printf("%+v", antonio)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jim)

}
