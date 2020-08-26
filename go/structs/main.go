package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Parson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()

	mySlice := []string{"Hello", "There!"}
	updateValue(mySlice)
	fmt.Println(mySlice)

	//alex := person{firstName: "Alex", lastName: "Anderson"}
}

//Receivers with pointers are allowed access to base type
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

//for slice we didnt pass pointer to slice so slice is a reference type
func updateValue(slice []string) {
	slice[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
