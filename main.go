package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	antonio := person{firstName: "Antonio", lastName: "Espinoza"}
	fmt.Println(antonio)
}
