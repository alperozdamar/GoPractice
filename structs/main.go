package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{firstName: "Alper", lastName: "Ozdamar"}

	alex.firstName = "Suleyman Alper"
	alex.lastName = "Ozdamar"
	alex.contact.email = "alper@mail.com"
	alex.contact.zipCode = 94122
	fmt.Println(alex)

	var alper person

	fmt.Println(alper)
	fmt.Printf("%+v", alper)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@mail.com",
			zipCode: 94000,
		},
	}
	fmt.Println("")
	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
