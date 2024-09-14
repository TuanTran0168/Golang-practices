package structs

import "fmt"

type Person struct {
	firstName string // private fields
	lastName  string
	age       int
}

// receiver function
func (person Person) ToString() string {
	return fmt.Sprintf("%v %v %v", person.firstName, person.lastName, person.age)
}

// receiver function
func (person Person) Walk() string {
	return person.firstName + " " + person.lastName + " likes walking"
}

func (person Person) GetFirstName() string {
	return person.firstName
}

func (person Person) GetLastName() string {
	return person.lastName
}

func (person Person) GetAge() int {
	return person.age
}

//Pointer based on receiver functions
func (person *Person) SetFirstName(firstName string) {
	person.firstName = firstName
}

func (person *Person) SetLastName(lastName string) {
	person.lastName = lastName
}

func (person *Person) SetAge(age int) {
	person.age = age
}
